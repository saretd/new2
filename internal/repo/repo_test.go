package repo_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ozonmp/trv-airport-api/internal/model"
)

func (suite *AirportRepoTestSuite) TestAdd() {
	airport := model.Airport{ID: 1, Name: "OVB", Location: "Novosibirsk"}
	rows := sqlmock.NewRows([]string{"id"}).AddRow(airport.ID)
	suite.mock.ExpectQuery("INSERT INTO airports").
		WithArgs(airport.Name, airport.Location).
		WillReturnRows(rows)
	airportID, err := suite.repo.Add(suite.ctx, airport.Name, airport.Location)
	suite.Require().Nil(err)
	suite.Require().NotNil(airport)
	suite.Require().Equal(airport.ID, airportID)
}

func (suite *AirportRepoTestSuite) TestGetSuccess() {
	airportID := uint64(1)
	suite.mock.
		ExpectQuery("SELECT id, name, location FROM airports").
		WithArgs(airportID).
		WillReturnRows(sqlmock.
			NewRows([]string{"id", "name", "location"}).
			AddRow(1, "OVB", "Novosibirsk"))
	airport, err := suite.repo.Get(suite.ctx, airportID)
	suite.Require().Nil(err)
	suite.Require().NotNil(airport)
	suite.Require().Equal(airport.ID, airportID)
}
func (suite *AirportRepoTestSuite) TestGetNotFound() {
	airportID := uint64(1)
	suite.mock.
		ExpectQuery("SELECT id, name, location FROM airports").
		WithArgs(airportID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "location"}))
	airport, err := suite.repo.Get(suite.ctx, airportID)
	suite.Require().Nil(err)
	suite.Require().Nil(airport)
}

func (suite *AirportRepoTestSuite) TestListSuccess() {
	cursor := uint64(5)
	suite.mock.
		ExpectQuery("SELECT id, name, location FROM airports").
		WithArgs(cursor).
		WillReturnRows(sqlmock.
			NewRows([]string{"id", "name", "location"}).
			AddRow(1, "OVB", "Novosibirsk").
			AddRow(2, "OVB", "Novosibirsk"))
	airports, err := suite.repo.List(suite.ctx, cursor, uint64(10))
	suite.Require().Nil(err)
	suite.Require().NotNil(airports)
	suite.Require().Len(airports, 2)
	suite.Require().Equal(airports[0].ID, uint64(1))
	suite.Require().Equal(airports[1].ID, uint64(2))
}
func (suite *AirportRepoTestSuite) TestListEmpty() {
	cursor := uint64(5)
	suite.mock.
		ExpectQuery("SELECT id, name, location FROM airports").
		WithArgs(cursor).
		WillReturnRows(sqlmock.
			NewRows([]string{"id", "name", "location"}))
	airports, err := suite.repo.List(suite.ctx, cursor, uint64(10))
	suite.Require().Nil(err)
	suite.Require().NotNil(airports)
	suite.Require().Len(airports, 0)
}

func (suite *AirportRepoTestSuite) TestADelete() {
	airportID := uint64(10)
	suite.mock.ExpectExec("DELETE FROM airports WHERE id = ").
		WithArgs(airportID).WillReturnResult(sqlmock.NewResult(0, 1))
	deletedAirportID, err := suite.repo.Delete(suite.ctx, airportID)
	suite.Require().Nil(err)
	suite.Require().Equal(airportID, deletedAirportID)
}
