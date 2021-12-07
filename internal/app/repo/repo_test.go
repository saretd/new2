package repo_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ozonmp/trv-airport-api/internal/model"
)

func (suite *EventRepoTestSuite) TestLock() {
	limit := uint64(5)
	var sourceEvents = []model.AirportEvent{
		{ID: 1, Type: model.Created, Status: model.Deferred, Entity: &model.Airport{ID: 1, Name: "LAX", Location: "Los-Angeles"}},
		{ID: 2, Type: model.Created, Status: model.Deferred, Entity: &model.Airport{ID: 2, Name: "OVB", Location: "Novosibirsk"}},
	}
	suite.mock.
		ExpectQuery("UPDATE airport_events SET is_locked = \\$1 WHERE is_locked = \\$2 AND status = \\$3 LIMIT 5 RETURNING id").
		WillReturnRows(sqlmock.
			NewRows([]string{"id"}).
			AddRow(1).
			AddRow(2))
	rows := sqlmock.NewRows([]string{"event_id", "event_type", "event_status", "airport_id", "airport_name", "airport_location"})
	for _, event := range sourceEvents {
		rows.AddRow(event.ID, event.Type, event.Status, event.Entity.ID, event.Entity.Name, event.Entity.Location)
	}
	suite.mock.
		ExpectQuery("SELECT airport_events.id as event_id, airport_events.type as event_type, airport_events.status as event_status, airports.id as airport_id, airports.name as airport_name, airports.location as airport_location FROM airport_events JOIN airports on airports.id = airport_events.airport_id WHERE airport_events.id IN").
		WithArgs(1, 2).
		WillReturnRows(rows)
	events, err := suite.repo.Lock(suite.ctx, limit)
	suite.Require().Nil(err)
	suite.Require().NotNil(events)
	suite.Require().Len(events, 2)

	for i, expectedEvent := range sourceEvents {
		suite.Require().NotNil(events[i])
		suite.Require().Equal(expectedEvent, events[i])
	}

}

func (suite *EventRepoTestSuite) TestUnlock() {
	eventIds := []uint64{1, 2, 3, 4}
	suite.mock.ExpectExec("UPDATE airport_events SET is_locked = \\$1 WHERE id IN \\\\?").
		WithArgs(false, 1, 2, 3, 4).
		WillReturnResult(sqlmock.NewResult(0, 1))
	err := suite.repo.Unlock(suite.ctx, eventIds)
	suite.Require().Nil(err)
}
func (suite *EventRepoTestSuite) TestDelete() {
	eventIds := []uint64{1, 2, 3, 4}
	suite.mock.ExpectExec("UPDATE airport_events SET status = \\$1 WHERE id IN \\\\?").
		WithArgs(model.Processed, 1, 2, 3, 4).
		WillReturnResult(sqlmock.NewResult(0, 1))
	err := suite.repo.Remove(suite.ctx, eventIds)
	suite.Require().Nil(err)
}
