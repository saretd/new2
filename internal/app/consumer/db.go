package consumer

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"sync"
	"time"

	"github.com/ozonmp/trv-airport-api/internal/app/repo"
	"github.com/ozonmp/trv-airport-api/internal/model"
)

var (
	totalAirportEventProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "trv_airport_api_airport_event_processed_total",
		Help: "Total number of airport events that were processed",
	})
)

//Consumer sends events to Kafka Producer
type Consumer interface {
	Start()
	Close()
}

type consumer struct {
	workerCount uint64
	events      chan<- model.AirportEvent
	repo        repo.EventRepo
	batchSize   uint64
	timeout     time.Duration
	done        chan bool
	wg          *sync.WaitGroup
	ctx         context.Context
}

//NewDbConsumer constructor
func NewDbConsumer(
	ctx context.Context,
	workerCount uint64,
	batchSize uint64,
	consumeTimeout time.Duration,
	repo repo.EventRepo,
	events chan<- model.AirportEvent,
) Consumer {

	wg := &sync.WaitGroup{}
	done := make(chan bool)

	return &consumer{
		workerCount: workerCount,
		batchSize:   batchSize,
		timeout:     consumeTimeout,
		repo:        repo,
		events:      events,
		wg:          wg,
		done:        done,
		ctx:         ctx,
	}
}

func (c *consumer) Start() {
	for i := uint64(0); i < c.workerCount; i++ {
		c.wg.Add(1)

		go func() {
			defer c.wg.Done()
			ticker := time.NewTicker(c.timeout)
			for {
				select {
				case <-ticker.C:
					c.handleEventsBatch()
				case <-c.done:
					return
				}
			}
		}()
	}
}

func (c *consumer) handleEventsBatch() {
	events, err := c.repo.Lock(c.ctx, c.batchSize)
	if err != nil {
		return
	}
	for _, event := range events {
		switch event.Type {
		case model.Created:
			c.events <- event
			totalAirportEventProcessed.Inc()
		}
	}
}

func (c *consumer) Close() {
	close(c.done)
	c.wg.Wait()
}
