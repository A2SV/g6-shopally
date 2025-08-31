package gateway

import (
	"context"
	"errors"
	"testing"

	"firebase.google.com/go/v4/messaging"
	"github.com/shopally-ai/internal/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type FCMGatewaySuite struct {
	suite.Suite
	ctx context.Context
	mc  *mocks.FCMClient
	gw  *FCMGateway
}

func (s *FCMGatewaySuite) SetupTest() {
	s.ctx = context.Background()
	s.mc = mocks.NewFCMClient(s.T())
	s.gw = NewFCMGatewayWithClient(s.mc)
}

func (s *FCMGatewaySuite) TestSend_Success() {
	token := "test-token"
	title := "Hello"
	body := "World"
	data := map[string]string{"a": "1"}

	// we assert that the message built has our fields set
	s.mc.On("Send", s.ctx, mock.MatchedBy(func(msg *messaging.Message) bool {
		return msg != nil && msg.Token == token && msg.Notification != nil &&
			msg.Notification.Title == title && msg.Notification.Body == body &&
			msg.Data["a"] == "1"
	})).Return("id-123", nil).Once()

	id, err := s.gw.Send(s.ctx, token, title, body, data)
	s.Require().NoError(err)
	s.Equal("id-123", id)
}

func (s *FCMGatewaySuite) TestSend_Error() {
	s.mc.On("Send", s.ctx, mock.Anything).Return("", errors.New("boom")).Once()
	id, err := s.gw.Send(s.ctx, "t", "a", "b", nil)
	s.Error(err)
	s.Equal("", id)
}

func TestFCMGatewaySuite(t *testing.T) { suite.Run(t, new(FCMGatewaySuite)) }
