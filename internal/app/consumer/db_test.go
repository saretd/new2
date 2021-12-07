package consumer

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/trv-airport-api/internal/mocks"
	"github.com/ozonmp/trv-airport-api/internal/model"
	"github.com/stretchr/testify/suite"
	"sync"
	"testing"
	"time"
)

type ConsumerTestSuite struct {
	suite.Suite
	repo          *mocks.MockEventRepo
	mockCtrl      *gomock.Controller
	events        chan model.AirportEvent
	consumerCount uint64
	batchSize     uint64
	timeout       time.Duration
	consumer      Consumer
	ctx           context.Context
}

func (suite *ConsumerTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.repo = mocks.NewMockEventRepo(suite.mockCtrl)
	suite.events = make(chan model.AirportEvent)
	suite.consumerCount = uint64(1)
	suite.batchSize = uint64(10)
	suite.timeout = time.Millisecond
	suite.ctx = context.Background()
	suite.consumer = NewDbConsumer(
		suite.ctx,
		suite.consumerCount,
		suite.batchSize,
		suite.timeout,
		suite.repo,
		suite.events,
	)
}

func (suite *ConsumerTestSuite) TearDownTest() {
	suite.consumer.Close()
	close(suite.events)
	suite.mockCtrl.Finish()
}

func (suite *ConsumerTestSuite) TestEventCreatedEvent() {
	lockEvents := []model.AirportEvent{
		{
			ID:     1,
			Type:   model.Created,
			Status: model.Deferred,
			Entity: nil,
		},
	}
	suite.repo.EXPECT().Lock(suite.ctx, suite.batchSize).Return(lockEvents, nil).Times(1)
	suite.consumer.Start()

	e := <-suite.events
	suite.Require().Equal(lockEvents[0], e, "received event == sent event")
}
func (suite *ConsumerTestSuite) TestUpdatedEvent() {
	lockEvents := []model.AirportEvent{
		{
			ID:     1,
			Type:   model.Updated,
			Status: model.Deferred,
			Entity: nil,
		},
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	suite.repo.EXPECT().Lock(suite.ctx, uint64(10)).Do(func(uint64) ([]model.AirportEvent, error) {
		wg.Done()
		return lockEvents, nil
	}).Times(1)

	suite.consumer.Start()

	var producedEvent model.AirportEvent
	select {
	case e := <-suite.events:
		producedEvent = e
	default:
		break
	}

	wg.Wait()
	suite.Require().Equal(model.AirportEvent{}, producedEvent)
}

func TestConsumerTestSuite(t *testing.T) {
	suite.Run(t, new(ConsumerTestSuite))
}
