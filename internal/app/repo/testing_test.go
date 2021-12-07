package repo_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	repo "github.com/ozonmp/trv-airport-api/internal/app/repo"

	"github.com/stretchr/testify/suite"
	"testing"
)

type EventRepoTestSuite struct {
	suite.Suite
	repo repo.EventRepo
	mock sqlmock.Sqlmock
	ctx  context.Context
}

func (suite *EventRepoTestSuite) SetupTest() {
	mockDB, mock, _ := sqlmock.New()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	suite.mock = mock
	suite.repo = repo.NewEventRepo(sqlxDB)
	suite.ctx = context.Background()
}

func (suite *EventRepoTestSuite) TearDownTest() {
	suite.repo.Close()
}

func TestEventRepoTestSuite(t *testing.T) {
	suite.Run(t, new(EventRepoTestSuite))
}
