package retranslator

import (
	"context"
	"github.com/ozonmp/trv-airport-api/internal/app/consumer"
	"github.com/ozonmp/trv-airport-api/internal/app/producer"
	"github.com/ozonmp/trv-airport-api/internal/app/repo"
	"github.com/ozonmp/trv-airport-api/internal/app/sender"
	"github.com/ozonmp/trv-airport-api/internal/model"
	"time"

	"github.com/gammazero/workerpool"
)

//Retranslator interface
type Retranslator interface {
	Start()
	Close()
}

// Config for the app
type Config struct {
	ChannelSize    uint64
	ConsumerCount  uint64
	ConsumeSize    uint64
	ConsumeTimeout time.Duration
	ProducerCount  uint64
	WorkerCount    int
	Repo           repo.EventRepo
	Sender         sender.EventSender
}

type retranslator struct {
	events     chan model.AirportEvent
	consumer   consumer.Consumer
	producer   producer.Producer
	workerPool *workerpool.WorkerPool
}

// NewRetranslator creates retranslator
func NewRetranslator(ctx context.Context, cfg Config) Retranslator {
	events := make(chan model.AirportEvent, cfg.ChannelSize)
	workerPool := workerpool.New(cfg.WorkerCount)

	consumer := consumer.NewDbConsumer(
		ctx,
		cfg.ConsumerCount,
		cfg.ConsumeSize,
		cfg.ConsumeTimeout,
		cfg.Repo,
		events,
	)
	producer := producer.NewKafkaProducer(
		ctx,
		cfg.ProducerCount,
		cfg.Sender,
		events,
		workerPool,
		cfg.Repo,
	)

	return &retranslator{
		events:     events,
		consumer:   consumer,
		producer:   producer,
		workerPool: workerPool,
	}
}

func (r *retranslator) Start() {
	r.producer.Start()
	r.consumer.Start()
}

func (r *retranslator) Close() {
	r.consumer.Close()
	r.producer.Close()
	r.workerPool.StopWait()
}
