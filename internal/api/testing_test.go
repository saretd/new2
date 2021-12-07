package api_test

import (
	"context"
	"github.com/golang/mock/gomock"
	api_ "github.com/ozonmp/trv-airport-api/internal/api"
	"github.com/ozonmp/trv-airport-api/internal/mocks"
	"github.com/ozonmp/trv-airport-api/internal/repo"
	"github.com/ozonmp/trv-airport-api/pkg/trv-airport-api"
	pb "github.com/ozonmp/trv-airport-api/pkg/trv-airport-api"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"testing"
)

func dialer(repo repo.AirportRepo) func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()
	api := api_.NewAirportAPI(repo)
	trv_airport_api.RegisterTrvAirportApiServiceServer(server, api)
	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal().Err(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

type GrpcV1TestSuite struct {
	suite.Suite
	repo     *mocks.MockAirportRepo
	mockCtrl *gomock.Controller
	ctx      context.Context
	conn     *grpc.ClientConn
	client   pb.TrvAirportApiServiceClient
}

func (suite *GrpcV1TestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.repo = mocks.NewMockAirportRepo(suite.mockCtrl)
	suite.ctx = context.Background()
	conn, err := grpc.DialContext(suite.ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer(suite.repo)))
	suite.conn = conn
	if err != nil {
		log.Fatal().Err(err)
	}

	suite.client = pb.NewTrvAirportApiServiceClient(conn)
}

func (suite *GrpcV1TestSuite) TearDownTest() {
	suite.conn.Close()
	suite.mockCtrl.Finish()
}

func TestGrpcV1TestSuite(t *testing.T) {
	suite.Run(t, new(GrpcV1TestSuite))
}
