package api_test

import (
	"github.com/golang/mock/gomock"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	"github.com/ozonmp/trv-airport-api/internal/model"

	pb "github.com/ozonmp/trv-airport-api/pkg/trv-airport-api"
)

func (suite *GrpcV1TestSuite) TestListAirport() {
	airports := []model.Airport{
		{ID: 1, Name: "OVB", Location: "Novosibirsk"},
		{ID: 2, Name: "LAX", Location: "Los-Angeles"},
	}
	cursor := uint64(5)
	limit := uint64(100)
	suite.repo.EXPECT().List(gomock.Any(), cursor, limit).Return(airports, nil)

	request := &pb.ListAirportsV1Request{Cursor: cursor, Limit: limit}
	response, err := suite.client.ListAirportsV1(suite.ctx, request)
	suite.Require().Nil(err)
	fetchedAirports := response.GetAirports()
	suite.Require().NotNil(fetchedAirports)
	suite.Require().Len(fetchedAirports, 2)
	for i, fetchedAirport := range fetchedAirports {
		suite.Require().Equal(airports[i].ID, fetchedAirport.GetId())
		suite.Require().Equal(airports[i].Name, fetchedAirport.GetName())
		suite.Require().Equal(airports[i].Location, fetchedAirport.GetLocation())
	}
}
