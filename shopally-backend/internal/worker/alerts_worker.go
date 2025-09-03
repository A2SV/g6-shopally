package worker

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/shopally-ai/pkg/domain"
	"github.com/shopally-ai/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AlertsWorker periodically checks saved alerts and sends notifications on price drops.
type AlertsWorker struct {
	Coll      alertsCollection
	Price     *util.PriceService
	Push      domain.IPushNotificationGateway
	BatchSize int
	Interval  time.Duration
}

// alertRecord is the minimal projection of an alert document needed for processing.
type alertRecord struct {
	ID           string
	DeviceID     string
	ProductID    string
	ProductTitle string
	CurrentPrice float64
}

// alertsCollection abstracts the subset of mongo.Collection used by the worker.
type alertsCollection interface {
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cursor, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
}

// cursor abstracts the subset of mongo.Cursor used by the worker.
type cursor interface {
	Next(context.Context) bool
	Decode(val interface{}) error
	Err() error
	Close(context.Context) error
}

// NewAlertsWorker constructs a worker using a real mongo collection by wrapping it into adapters.
func NewAlertsWorker(coll *mongo.Collection, price *util.PriceService, push domain.IPushNotificationGateway) *AlertsWorker {
	return &AlertsWorker{
		Coll:      &mongoCollectionAdapter{c: coll},
		Price:     price,
		Push:      push,
		BatchSize: 500,
		Interval:  4 * time.Hour,
	}
}

// Run starts the periodic loop. It blocks.
func (w *AlertsWorker) Run(ctx context.Context) {
	ticker := time.NewTicker(w.Interval)
	defer ticker.Stop()
	for {
		if err := w.tick(ctx); err != nil {
			log.Printf("alerts worker tick error: %v", err)
		}
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
	}
}

// tick performs a single scan-and-notify cycle.
func (w *AlertsWorker) tick(ctx context.Context) error {
	// 1) Stream active alerts in pages
	filter := bson.M{"IsActive": true}
	opts := options.Find().SetBatchSize(int32(w.BatchSize))
	cur, err := w.Coll.Find(ctx, filter, opts)
	if err != nil {
		return err
	}
	defer cur.Close(ctx)

	// Aggregate alerts by productID
	var batch []alertRecord

	for cur.Next(ctx) {
		var doc struct {
			ID           string
			DeviceID     string
			ProductID    string
			ProductTitle string
			CurrentPrice float64
			IsActive     bool
		}
		if err := cur.Decode(&doc); err != nil {
			return err
		}
		batch = append(batch, alertRecord{ID: doc.ID, DeviceID: doc.DeviceID, ProductID: doc.ProductID, ProductTitle: doc.ProductTitle, CurrentPrice: doc.CurrentPrice})
		if len(batch) >= w.BatchSize {
			if err := w.processBatch(ctx, batch); err != nil {
				log.Printf("alerts worker batch error: %v", err)
			}
			batch = batch[:0]
		}
	}
	if err := cur.Err(); err != nil {
		return err
	}
	if len(batch) > 0 {
		if err := w.processBatch(ctx, batch); err != nil {
			log.Printf("alerts worker batch error: %v", err)
		}
	}
	return nil
}

func (w *AlertsWorker) processBatch(ctx context.Context, alerts []alertRecord) error {
	// Dedupe by product
	ids := make([]string, 0, len(alerts))
	current := make(map[string]float64, len(alerts))
	seen := map[string]struct{}{}
	for _, a := range alerts {
		if _, ok := seen[a.ProductID]; !ok {
			ids = append(ids, a.ProductID)
			seen[a.ProductID] = struct{}{}
			current[a.ProductID] = a.CurrentPrice
		} else if _, ok := current[a.ProductID]; !ok {
			current[a.ProductID] = a.CurrentPrice
		}
	}

	// Fetch prices
	res, err := w.Price.UpdatePricesIfChangedBatch(ctx, ids, current)
	if err != nil {
		return err
	}

	// Fan-out to alerts and notify on drop
	for _, a := range alerts {
		pc, ok := res[a.ProductID]
		if !ok {
			continue
		}
		dropped := pc.Price < a.CurrentPrice-1e-6
		if dropped { // drop
			// Send push
			if w.Push != nil && a.DeviceID != "" {
				// Treat DeviceID saved on alert as the FCM token for this app instance
				tokens := []string{a.DeviceID}
				title := "Price drop on a saved product"
				pt := a.ProductTitle
				if pt == "" {
					pt = a.ProductID
				}
				body := fmt.Sprintf("There is a price drop on '%s'. Refresh to check the new price.", pt)
				for _, tok := range tokens {
					if _, err := w.Push.Send(ctx, tok, title, body, map[string]string{
						"productId": a.ProductID,
						"oldPrice":  fmt.Sprintf("%.2f", a.CurrentPrice),
						"newPrice":  fmt.Sprintf("%.2f", pc.Price),
					}); err != nil {
						// best-effort: just log; consider removing/marking invalid on known FCM errors later
						log.Printf("push send failed for token %s: %v", tok, err)
					}
				}
			}
		}
		// Update stored price ONLY if drop, to avoid ratcheting up and causing false drop notifications later
		if dropped {
			_, _ = w.Coll.UpdateOne(ctx, bson.M{"ID": a.ID}, bson.M{"$set": bson.M{"CurrentPrice": pc.Price}})
		}
	}
	return nil
}

// mongo adapters to satisfy interfaces
type mongoCollectionAdapter struct{ c *mongo.Collection }

func (m *mongoCollectionAdapter) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cursor, error) {
	cur, err := m.c.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	return &mongoCursorAdapter{cur: cur}, nil
}

func (m *mongoCollectionAdapter) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return m.c.UpdateOne(ctx, filter, update, opts...)
}

type mongoCursorAdapter struct{ cur *mongo.Cursor }

func (mc *mongoCursorAdapter) Next(ctx context.Context) bool   { return mc.cur.Next(ctx) }
func (mc *mongoCursorAdapter) Decode(val interface{}) error    { return mc.cur.Decode(val) }
func (mc *mongoCursorAdapter) Err() error                      { return mc.cur.Err() }
func (mc *mongoCursorAdapter) Close(ctx context.Context) error { return mc.cur.Close(ctx) }
