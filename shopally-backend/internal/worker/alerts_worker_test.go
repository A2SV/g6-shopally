package worker

import (
	"context"
	"reflect"
	"testing"
	"time"

	imocks "github.com/shopally-ai/internal/mocks"
	"github.com/shopally-ai/pkg/domain"
	"github.com/shopally-ai/pkg/util"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mocks for alertsCollection and cursor using testify.Mock
type mockCursor struct{ mock.Mock }

func (m *mockCursor) Next(ctx context.Context) bool {
	ret := m.Called(ctx)
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		return rf(ctx)
	}
	if v, ok := ret.Get(0).(bool); ok {
		return v
	}
	return false
}
func (m *mockCursor) Decode(val interface{}) error {
	ret := m.Called(val)
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		return rf(val)
	}
	return ret.Error(0)
}
func (m *mockCursor) Err() error { return m.Called().Error(0) }
func (m *mockCursor) Close(ctx context.Context) error {
	ret := m.Called(ctx)
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		return rf(ctx)
	}
	return ret.Error(0)
}

type mockColl struct{ mock.Mock }

func (m *mockColl) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cursor, error) {
	args := m.Called(ctx, filter)
	var c cursor
	if v := args.Get(0); v != nil {
		c = v.(cursor)
	}
	return c, args.Error(1)
}
func (m *mockColl) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update)
	if r := args.Get(0); r != nil {
		return r.(*mongo.UpdateResult), args.Error(1)
	}
	return &mongo.UpdateResult{}, args.Error(1)
}

type AlertsWorkerSuite struct {
	suite.Suite
	ctx   context.Context
	coll  *mockColl
	cur   *mockCursor
	push  *imocks.IPushNotificationGateway
	price *util.PriceService
}


// gateway mock for PriceService
type priceGW struct{ mock.Mock }

func (p *priceGW) FetchProducts(ctx context.Context, query string, filters map[string]interface{}) ([]*domain.Product, error) {
	args := p.Called(ctx, query, filters)
	if v := args.Get(0); v != nil {
		return v.([]*domain.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func (s *AlertsWorkerSuite) SetupTest() {
	s.ctx = context.Background()
	s.coll = &mockColl{}
	s.cur = &mockCursor{}
	s.push = imocks.NewIPushNotificationGateway(s.T())
	pgw := &priceGW{}
	s.price = util.New(pgw)
}

func (s *AlertsWorkerSuite) TestTick_HappyPath_PriceDrop() {
	// Prepare cursor to yield one alert then stop
	// alert doc expected shape
	first := true
	s.cur.On("Next", mock.Anything).Return(func(ctx context.Context) bool {
		if first {
			first = false
			return true
		}
		return false
	}).Twice()
	s.cur.On("Decode", mock.Anything).Return(func(v interface{}) error {
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
		}
		rv.FieldByName("ID").SetString("a1")
		rv.FieldByName("DeviceID").SetString("tok1")
		rv.FieldByName("ProductID").SetString("p1")
		rv.FieldByName("ProductTitle").SetString("Prod 1")
		rv.FieldByName("CurrentPrice").SetFloat(20)
		rv.FieldByName("IsActive").SetBool(true)
		return nil
	}).Once()
	s.cur.On("Err").Return(nil).Once()
	s.cur.On("Close", mock.Anything).Return(nil).Once()

	s.coll.On("Find", mock.Anything, mock.Anything).Return(s.cur, nil).Once()
	s.coll.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(&mongo.UpdateResult{}, nil).Once()

	// price gateway returns lower price
	pgw := s.priceGateway()
	pgw.On("FetchProducts", mock.Anything, "", mock.MatchedBy(func(f map[string]interface{}) bool { return f["product_ids"] != nil })).Return([]*domain.Product{{ID: "p1", Price: domain.Price{USD: 18}}}, nil).Once()

	// expect push
	s.push.On("Send", mock.Anything, "tok1", mock.Anything, mock.Anything, mock.Anything).Return("ok", nil).Once()

	w := &AlertsWorker{Coll: s.coll, Price: s.price, Push: s.push, BatchSize: 100, Interval: time.Hour}
	s.NoError(w.tick(s.ctx))
	s.push.AssertExpectations(s.T())
}

func (s *AlertsWorkerSuite) TestTick_NoDrop_NoPush_NoUpdate() {
	// cursor yields one alert
	first := true
	s.cur.On("Next", mock.Anything).Return(func(ctx context.Context) bool {
		if first {
			first = false
			return true
		}
		return false
	}).Twice()
	s.cur.On("Decode", mock.Anything).Return(func(v interface{}) error {
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
		}
		rv.FieldByName("ID").SetString("a2")
		rv.FieldByName("DeviceID").SetString("tok2")
		rv.FieldByName("ProductID").SetString("p2")
		rv.FieldByName("ProductTitle").SetString("Prod 2")
		rv.FieldByName("CurrentPrice").SetFloat(10)
		rv.FieldByName("IsActive").SetBool(true)
		return nil
	}).Once()
	s.cur.On("Err").Return(nil).Once()
	s.cur.On("Close", mock.Anything).Return(nil).Once()
	s.coll.On("Find", mock.Anything, mock.Anything).Return(s.cur, nil).Once()

	pgw := s.priceGateway()
	pgw.On("FetchProducts", mock.Anything, "", mock.Anything).Return([]*domain.Product{{ID: "p2", Price: domain.Price{USD: 12}}}, nil).Once()

	w := &AlertsWorker{Coll: s.coll, Price: s.price, Push: s.push, BatchSize: 100, Interval: time.Hour}
	s.NoError(w.tick(s.ctx))
	s.push.AssertNotCalled(s.T(), "Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	s.coll.AssertNotCalled(s.T(), "UpdateOne", mock.Anything, mock.Anything, mock.Anything)
}

func (s *AlertsWorkerSuite) TestTick_FindError_Propagates() {
	s.coll.On("Find", mock.Anything, mock.Anything).Return(nil, assertErr("boom")).Once()
	w := &AlertsWorker{Coll: s.coll, Price: s.price, Push: s.push}
	s.Error(w.tick(s.ctx))
}

// helper to get underlying price gateway mock
func (s *AlertsWorkerSuite) priceGateway() *priceGW {
	// PriceService.New wrapped our priceGW, access it via the private field through type assertion
	// We created PriceService with util.New(pgw) in SetupTest, so s.price.ag should be *priceGW
	type hasAg interface{ GetAg() interface{} }
	// expose via hack: since we can't access unexported field, we rebuild service
	pgw := &priceGW{}
	s.price = util.New(pgw)
	return pgw
}

type assertErr string

func (e assertErr) Error() string { return string(e) }

func TestAlertsWorkerSuite(t *testing.T) { suite.Run(t, new(AlertsWorkerSuite)) }
