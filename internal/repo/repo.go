package repo

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/ozonmp/trv-airport-api/internal/model"
)

// AirportRepo is DAO for Airport
type AirportRepo interface {
	Get(ctx context.Context, airportID uint64) (*model.Airport, error)
	Add(ctx context.Context, name string, location string) (uint64, error)
	List(ctx context.Context, cursor, limit uint64) ([]model.Airport, error)
	Delete(ctx context.Context, airportID uint64) (uint64, error)
	Close() error
}

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewAirportRepo returns AirportRepo interface
func NewAirportRepo(db *sqlx.DB, batchSize uint) AirportRepo {
	return &repo{db: db, batchSize: batchSize}
}

func (r *repo) Close() error {
	return r.db.Close()
}

func (r *repo) Get(ctx context.Context, airportID uint64) (*model.Airport, error) {
	query := sq.Select("id", "name", "location").
		PlaceholderFormat(sq.Dollar).
		From("airports").
		Where(sq.Eq{"id": airportID})
	s, args, err := query.ToSql()
	if err != nil {
		return nil, err

	}
	airports := []model.Airport{}
	err = r.db.SelectContext(ctx, &airports, s, args...)
	if len(airports) == 0 {
		return nil, nil
	}
	return &airports[0], err
}

func (r *repo) Add(ctx context.Context, name string, location string) (airportID uint64, err error) {
	query := sq.Insert("airports").
		PlaceholderFormat(sq.Dollar).
		Columns("name", "location").
		Values(name, location).
		Suffix("RETURNING id").RunWith(r.db)
	rows, err := query.QueryContext(ctx)
	if err != nil {
		return 0, err
	}
	if rows.Next() {
		err = rows.Scan(&airportID)
	}
	return
}

func (r *repo) List(ctx context.Context, cursor, limit uint64) ([]model.Airport, error) {
	query := sq.Select("id", "name", "location").
		PlaceholderFormat(sq.Dollar).
		From("airports").
		Where(sq.Gt{"id": cursor}).
		Limit(limit)
	s, args, err := query.ToSql()
	if err != nil {
		return nil, err

	}
	airports := make([]model.Airport, 0, limit)
	err = r.db.SelectContext(ctx, &airports, s, args...)
	return airports, err
}
func (r *repo) Delete(ctx context.Context, airportID uint64) (uint64, error) {
	query := sq.Delete("airports").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": airportID})
	s, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}
	_, err = r.db.ExecContext(ctx, s, args...)
	return airportID, err
}
