package api_test

import (
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/trv-airport-api/internal/model"
	"google.golang.org/grpc/status"
	"strings"

	pb "github.com/ozonmp/trv-airport-api/pkg/trv-airport-api"
)

func (suite *GrpcV1TestSuite) TestCreateAirportSuccess() {
	airport := model.Airport{ID: 1, Name: "LAX", Location: "Los-Angeles"}
	suite.repo.EXPECT().Add(gomock.Any(), airport.Name, airport.Location).Return(airport.ID, nil).Times(1)

	request := &pb.CreateAirportV1Request{Name: airport.Name, Location: airport.Location}
	response, err := suite.client.CreateAirportV1(suite.ctx, request)
	suite.Require().Nil(err)
	suite.Require().Equal(uint64(1), response.GetId())
}

func (suite *GrpcV1TestSuite) TestCreateAirportInvalidArgs() {

	tt := []struct {
		request *pb.CreateAirportV1Request
		errMsg  string
	}{
		{request: &pb.CreateAirportV1Request{Name: "A", Location: "Moscow"},
			errMsg: "invalid CreateAirportV1Request.Name: value length must be 3 runes",
		},
		{request: &pb.CreateAirportV1Request{Name: "OVB", Location: "M"},
			errMsg: "invalid CreateAirportV1Request.Location: value length must be between 2 and 255 runes, inclusive",
		},

		{request: &pb.CreateAirportV1Request{Name: "OVB", Location: strings.Repeat("OVB", 256)},
			errMsg: "invalid CreateAirportV1Request.Location: value length must be between 2 and 255 runes, inclusive",
		},
	}

	for _, test := range tt {
		response, err := suite.client.CreateAirportV1(suite.ctx, test.request)

		suite.Require().Equal(uint64(0), response.GetId())
		suite.Require().NotNil(err)
		er, ok := status.FromError(err)
		suite.Require().True(ok)
		suite.Require().Equal(test.errMsg, er.Message())
	}
}
