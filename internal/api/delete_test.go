package api_test

import (
	"github.com/golang/mock/gomock"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	pb "github.com/ozonmp/trv-airport-api/pkg/trv-airport-api"
)

func (suite *GrpcV1TestSuite) TestDeleteAirportSuccess() {
	airportID := uint64(1)
	suite.repo.EXPECT().Delete(gomock.Any(), airportID).Return(airportID, nil).Times(1)

	request := &pb.DeleteAirportV1Request{AirportId: airportID}
	response, err := suite.client.DeleteAirportV1(suite.ctx, request)
	suite.Require().Equal(airportID, response.GetAirportId())
	suite.Require().Nil(err)
}
