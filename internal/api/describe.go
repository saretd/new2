package api

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ozonmp/trv-airport-api/pkg/trv-airport-api"
)

func (o *airportAPI) DescribeAirportV1(
	ctx context.Context,
	req *pb.DescribeAirportV1Request,
) (*pb.DescribeAirportV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "DescribeAirportV1")
	defer span.Finish()

	switchLogLevel(ctx)

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DescribeAirportV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	airport, err := o.repo.Get(ctx, req.AirportId)
	if err != nil {
		log.Error().Err(err).Msg("DescribeAirportV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if airport == nil {
		log.Error().Err(err).Uint64("airportId", req.AirportId).Msg("airport not found")
		totalAirportNotFound.Inc()
		return nil, status.Error(codes.NotFound, "airport not found")
	}
	span.LogKV("airportID", req.AirportId)
	log.Debug().Msg("DescribeAirportV1 - success")
	return &pb.DescribeAirportV1Response{
		Value: &pb.Airport{
			Id:       airport.ID,
			Name:     airport.Name,
			Location: airport.Location,
		},
	}, nil
}
