package producer

import (
	"context"
	"github.com/ozonmp/trv-airport-api/internal/app/repo"
	"github.com/ozonmp/trv-airport-api/internal/app/sender"
	"github.com/ozonmp/trv-airport-api/internal/model"
	"github.com/rs/zerolog/log"
	"sync"

	"github.com/gammazero/workerpool"
)

// Producer creates events for Kafka
type Producer interface {
	Start()
	Close()
}

type producer struct {
	ctx         context.Context
	sender      sender.EventSender
	events      <-chan model.AirportEvent
	repo        repo.EventRepo
	workerCount uint64
	workerPool  *workerpool.WorkerPool
	wg          *sync.WaitGroup
	done        chan bool
}

// NewKafkaProducer creates produces
func NewKafkaProducer(
	ctx context.Context,
	workerCount uint64,
	sender sender.EventSender,
	events <-chan model.AirportEvent,
	workerPool *workerpool.WorkerPool,
	repo repo.EventRepo,
) Producer {

	wg := &sync.WaitGroup{}
	done := make(chan bool)

	return &producer{
		ctx:         ctx,
		workerCount: workerCount,
		sender:      sender,
		events:      events,
		workerPool:  workerPool,
		wg:          wg,
		done:        done,
		repo:        repo,
	}
}

func (p *producer) Start() {
	for i := uint64(0); i < p.workerCount; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for {
				select {
				case event := <-p.events:
					p.handleEvent(event)
				case <-p.done:
					return
				}
			}
		}()
	}
}

func (p *producer) handleEvent(event model.AirportEvent) {
	if err := p.sender.Send(&event); err != nil {
		p.workerPool.Submit(func() {
			err = p.repo.Unlock(p.ctx, []uint64{event.ID})
			if err != nil {
				log.Fatal().Err(err)
			}
		})
	} else {
		p.workerPool.Submit(func() {
			err = p.repo.Remove(p.ctx, []uint64{event.ID})
			if err != nil {
				log.Fatal().Err(err)
			}

		})
	}
}

func (p *producer) Close() {
	close(p.done)
	p.wg.Wait()
}
