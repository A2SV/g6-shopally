package worker

import (
	"context"
	"reflect"
	"testing"
	"time"

	imocks "github.com/shopally-ai/internal/mocks"
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

// fake fetcher implementing util.PriceFetcher
type fakeFetcher struct{ mock.Mock }

func (f *fakeFetcher) FetchPrices(ctx context.Context, ids []string) (map[string]util.PriceAmounts, error) {
	args := f.Called(ctx, ids)
	if v := args.Get(0); v != nil {
		return v.(map[string]util.PriceAmounts), args.Error(1)
	}
	return nil, args.Error(1)
}

func (s *AlertsWorkerSuite) SetupTest() {
	s.ctx = context.Background()
	s.coll = &mockColl{}
	s.cur = &mockCursor{}
	s.push = imocks.NewIPushNotificationGateway(s.T())
	s.price = util.NewWithFetcher(&fakeFetcher{})
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

	// fetcher returns lower price
	ff := &fakeFetcher{}
	ff.On("FetchPrices", mock.Anything, mock.Anything).Return(map[string]util.PriceAmounts{"p1": {USD: 18}}, nil).Once()
	s.price = util.NewWithFetcher(ff)

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

	ff := &fakeFetcher{}
	ff.On("FetchPrices", mock.Anything, mock.Anything).Return(map[string]util.PriceAmounts{"p2": {USD: 12}}, nil).Once()
	s.price = util.NewWithFetcher(ff)

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

// no need for gateway accessor since we inject the fetcher directly

type assertErr string

func (e assertErr) Error() string { return string(e) }

func TestAlertsWorkerSuite(t *testing.T) { suite.Run(t, new(AlertsWorkerSuite)) }
