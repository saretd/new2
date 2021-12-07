package api

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ozonmp/trv-airport-api/pkg/trv-airport-api"
)

func (o *airportAPI) CreateAirportV1(
	ctx context.Context,
	req *pb.CreateAirportV1Request,
) (*pb.CreateAirportV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateAirportV1")
	defer span.Finish()

	switchLogLevel(ctx)

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("CreateAirportV1 - invalid argument")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	airportID, err := o.repo.Add(ctx, req.Name, req.Location)
	if err != nil {
		log.Error().Err(err).Msg("CreateAirportV1 -- failed")
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Msg("CreateAirportV1 - success")
	span.LogKV("airportID", airportID)
	totalAirportCreated.Inc()

	return &pb.CreateAirportV1Response{
		Id: airportID,
	}, nil
}
