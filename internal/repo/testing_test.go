package repo_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/trv-airport-api/internal/repo"

	"github.com/stretchr/testify/suite"
	"testing"
)

type AirportRepoTestSuite struct {
	suite.Suite
	repo repo.AirportRepo
	mock sqlmock.Sqlmock
	ctx  context.Context
}

const batchSize = 5

func (suite *AirportRepoTestSuite) SetupTest() {
	mockDB, mock, _ := sqlmock.New()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	suite.mock = mock
	suite.repo = repo.NewAirportRepo(sqlxDB, batchSize)
	suite.ctx = context.Background()
}

func (suite *AirportRepoTestSuite) TearDownTest() {
	suite.repo.Close()
}

func TestAirportRepoTestSuite(t *testing.T) {
	suite.Run(t, new(AirportRepoTestSuite))
}
