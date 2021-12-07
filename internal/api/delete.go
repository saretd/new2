package api

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ozonmp/trv-airport-api/pkg/trv-airport-api"
)

func (o *airportAPI) DeleteAirportV1(
	ctx context.Context,
	req *pb.DeleteAirportV1Request,
) (*pb.DeleteAirportV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "DeleteAirportV1")
	defer span.Finish()

	switchLogLevel(ctx)

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DeleteAirportV1 - invalid argument")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	airportID, err := o.repo.Delete(ctx, req.AirportId)
	if err != nil {
		log.Error().Err(err).Msg("DeleteAirportV1 -- failed")
		return nil, status.Error(codes.Internal, err.Error())
	}
	if airportID == 0 {
		log.Error().Err(err).Uint64("airportId", req.AirportId).Msg("airport not found")
		totalAirportNotFound.Inc()
		return nil, status.Error(codes.NotFound, "airport not found")
	}
	span.LogKV("airportID", airportID)
	log.Debug().Msg("DeleteAirportV1 - success")
	totalAirportDeleted.Inc()

	return &pb.DeleteAirportV1Response{
		AirportId: airportID,
	}, nil
}
