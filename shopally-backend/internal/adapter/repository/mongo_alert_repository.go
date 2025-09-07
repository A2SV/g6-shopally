package repository

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/shopally-ai/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoAlertRepository implements domain.AlertRepository using MongoDB.
type MongoAlertRepository struct {
	coll *mongo.Collection
}

// NewMongoAlertRepository creates a new MongoAlertRepository with the provided collection.
func NewMongoAlertRepository(coll *mongo.Collection) *MongoAlertRepository {
	return &MongoAlertRepository{coll: coll}
}

func (r *MongoAlertRepository) CreateAlert(alert *domain.Alert) error {
	if alert.ID == "" {
		alert.ID = uuid.New().String()
	}
	// Ensure default active on create if not set
	if !alert.IsActive {
		alert.IsActive = true
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := r.coll.InsertOne(ctx, bson.M{
		"ID":           alert.ID,
		"DeviceID":     alert.DeviceID,
		"ProductID":    alert.ProductID,
		"CurrentPrice": alert.CurrentPrice,
		"IsActive":     alert.IsActive,
	})
	return err
}

func (r *MongoAlertRepository) GetAlert(alertID string) (*domain.Alert, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var doc domain.Alert
	err := r.coll.FindOne(ctx, bson.M{"ID": alertID}).Decode(&doc)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}
		return nil, err
	}
	return &doc, nil
}

func (r *MongoAlertRepository) DeleteAlert(alertID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := r.coll.UpdateOne(ctx, bson.M{"ID": alertID}, bson.M{"$set": bson.M{"IsActive": false}})
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
