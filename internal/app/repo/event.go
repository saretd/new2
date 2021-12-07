package repo

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/trv-airport-api/internal/model"
)

//EventRepo storage of the events
type EventRepo interface {
	Lock(ctx context.Context, n uint64) ([]model.AirportEvent, error)
	Unlock(ctx context.Context, eventIDs []uint64) error
	Add(ctx context.Context, event []model.AirportEvent) error
	Remove(ctx context.Context, eventIDs []uint64) error
	Close() error
}

type repo struct {
	db *sqlx.DB
}

//NewEventRepo constructor
func NewEventRepo(db *sqlx.DB) EventRepo {
	return &repo{db: db}
}

func (r *repo) Close() error {
	return r.db.Close()
}

func (r *repo) Lock(ctx context.Context, n uint64) ([]model.AirportEvent, error) {
	updateQuery := sq.Update("airport_events").
		Set("is_locked", true).
		Where(sq.Eq{"is_locked": false, "status": model.Deferred}).
		PlaceholderFormat(sq.Dollar).
		Limit(n).
		Suffix("RETURNING id")
	s, args, err := updateQuery.ToSql()
	if err != nil {
		return nil, err
	}
	eventIds := make([]uint64, 0, n)
	err = r.db.SelectContext(ctx, &eventIds, s, args...)
	if err != nil {
		return nil, err
	}
	selectQuery := sq.Select("airport_events.id as event_id", "airport_events.type as event_type", "airport_events.status as event_status", "airports.id as airport_id", "airports.name as airport_name", "airports.location as airport_location").
		Join("airports on airports.id = airport_events.airport_id").
		From("airport_events").
		Where(sq.Eq{"airport_events.id": eventIds})
	s, args, err = selectQuery.ToSql()
	if err != nil {
		return nil, err

	}
	var eventsData []struct {
		EventID         uint64            `db:"event_id"`
		EventType       model.EventType   `db:"event_type"`
		EventStatus     model.EventStatus `db:"event_status"`
		AirportID       uint64            `db:"airport_id"`
		AirportName     string            `db:"airport_name"`
		AirportLocation string            `db:"airport_location"`
	}
	err = r.db.SelectContext(ctx, &eventsData, s, args...)
	events := make([]model.AirportEvent, 0, n)
	for _, ed := range eventsData {
		events = append(events, model.AirportEvent{
			ID:     ed.EventID,
			Type:   ed.EventType,
			Status: ed.EventStatus,
			Entity: &model.Airport{
				ID:       ed.AirportID,
				Name:     ed.AirportName,
				Location: ed.AirportLocation,
			},
		})
	}
	return events, err
}

func (r *repo) Unlock(ctx context.Context, eventIDs []uint64) error {
	query := sq.Update("airport_events").
		Set("is_locked", false).
		Where(sq.Eq{"id": eventIDs}).
		PlaceholderFormat(sq.Dollar)
	s, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, s, args...)
	return err
}
func (r *repo) Remove(ctx context.Context, eventIDs []uint64) error {
	query := sq.Update("airport_events").
		Set("status", model.Processed).
		Where(sq.Eq{"id": eventIDs}).
		PlaceholderFormat(sq.Dollar)
	s, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, s, args...)
	return err
}
func (r *repo) Add(ctx context.Context, events []model.AirportEvent) error {
	return nil
}
