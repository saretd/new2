package api_test

import (
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/trv-airport-api/internal/model"
	"google.golang.org/grpc/status"

	pb "github.com/ozonmp/trv-airport-api/pkg/trv-airport-api"
)

func (suite *GrpcV1TestSuite) TestDescribeAirportSuccess() {
	airport := model.Airport{ID: 1, Name: "LAX", Location: "Los-Angeles"}
	suite.repo.EXPECT().Get(gomock.Any(), airport.ID).Return(&airport, nil).Times(1)

	request := &pb.DescribeAirportV1Request{AirportId: 1}
	response, err := suite.client.DescribeAirportV1(suite.ctx, request)
	suite.Require().Nil(err)
	suite.Require().NotNil(response.GetValue())
	suite.Require().Equal(airport.ID, response.GetValue().GetId())
}
func (suite *GrpcV1TestSuite) TestDescribeAirportNotFound() {
	airportID := uint64(1)
	suite.repo.EXPECT().Get(gomock.Any(), airportID).Return(nil, nil).Times(1)

	request := &pb.DescribeAirportV1Request{AirportId: airportID}
	response, err := suite.client.DescribeAirportV1(suite.ctx, request)
	suite.Require().Nil(response.GetValue())
	suite.Require().NotNil(err)
	er, ok := status.FromError(err)
	suite.Require().True(ok)
	suite.Require().Equal("airport not found", er.Message())
}
