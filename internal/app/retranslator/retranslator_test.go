package retranslator

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

type RetranslatorTestSuite struct {
	suite.Suite
	mockCtrl     *gomock.Controller
	repo         *mocks.MockEventRepo
	sender       *mocks.MockEventSender
	cfg          Config
	retranslator Retranslator
	ctx          context.Context
}

func (suite *RetranslatorTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.repo = mocks.NewMockEventRepo(suite.mockCtrl)
	suite.sender = mocks.NewMockEventSender(suite.mockCtrl)
	suite.ctx = context.Background()
	suite.cfg = Config{
		ChannelSize:    512,
		ConsumerCount:  1,
		ConsumeSize:    10,
		ConsumeTimeout: time.Millisecond * 500,
		ProducerCount:  1,
		WorkerCount:    1,
		Repo:           suite.repo,
		Sender:         suite.sender,
	}
	suite.retranslator = NewRetranslator(suite.ctx, suite.cfg)
}

func (suite *RetranslatorTestSuite) TearDownTest() {
	suite.retranslator.Close()
	suite.mockCtrl.Finish()
}

func (suite *RetranslatorTestSuite) TestEventsPipeline() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	lockEvents := []model.AirportEvent{
		{
			ID:     1,
			Type:   model.Created,
			Status: model.Deferred,
			Entity: nil,
		},
	}

	suite.repo.EXPECT().Lock(gomock.Any(), suite.cfg.ConsumeSize).DoAndReturn(func(uint64) ([]model.AirportEvent, error) {
		return lockEvents, nil
	}).Times(1)

	suite.repo.EXPECT().Remove(gomock.Any(), []uint64{1}).Times(1)
	suite.sender.EXPECT().Send(&lockEvents[0]).Do(func(*model.AirportEvent) error {
		wg.Done()
		return nil
	}).Times(1)

	suite.retranslator.Start()
	wg.Wait()
}

func TestRetranslatorTestSuite(t *testing.T) {
	suite.Run(t, new(RetranslatorTestSuite))
}
