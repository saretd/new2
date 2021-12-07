package producer

import (
	"context"
	"errors"
	"github.com/gammazero/workerpool"
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/trv-airport-api/internal/mocks"
	"github.com/ozonmp/trv-airport-api/internal/model"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ProducerTestSuite struct {
	suite.Suite
	ctx            context.Context
	mockCtrl       *gomock.Controller
	sender         *mocks.MockEventSender
	repo           *mocks.MockEventRepo
	producerCount  uint64
	events         chan model.AirportEvent
	workerPoolSize int
	workerPool     *workerpool.WorkerPool
	producer       Producer
	testEvent      model.AirportEvent
}

func (suite *ProducerTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.sender = mocks.NewMockEventSender(suite.mockCtrl)
	suite.repo = mocks.NewMockEventRepo(suite.mockCtrl)
	suite.producerCount = uint64(1)
	suite.events = make(chan model.AirportEvent)
	suite.workerPoolSize = 1
	suite.workerPool = workerpool.New(suite.workerPoolSize)
	suite.producer = NewKafkaProducer(
		suite.ctx,
		suite.producerCount,
		suite.sender,
		suite.events,
		suite.workerPool,
		suite.repo,
	)
	suite.testEvent = model.AirportEvent{
		ID:     1,
		Type:   model.Created,
		Status: model.Deferred,
		Entity: nil}
}

func (suite *ProducerTestSuite) TearDownTest() {
	suite.producer.Close()
	suite.workerPool.StopWait()
	close(suite.events)
	suite.mockCtrl.Finish()
}

func (suite *ProducerTestSuite) TestStart() {
	suite.producer.Start()
}

func (suite *ProducerTestSuite) TestEventSendError() {
	suite.producer.Start()

	suite.sender.EXPECT().Send(&suite.testEvent).Return(errors.New("test error")).Times(1)
	suite.repo.EXPECT().Unlock(gomock.Any(), []uint64{suite.testEvent.ID})
	suite.events <- suite.testEvent

}
func (suite *ProducerTestSuite) TestEventSendOk() {
	suite.producer.Start()

	suite.sender.EXPECT().Send(&suite.testEvent).Return(nil).Times(1)
	suite.repo.EXPECT().Remove(gomock.Any(), []uint64{suite.testEvent.ID}).Times(1)
	suite.events <- suite.testEvent
}
func TestProducerTestSuite(t *testing.T) {
	suite.Run(t, new(ProducerTestSuite))
}
