package api

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ozonmp/trv-airport-api/pkg/trv-airport-api"
)

func (o *airportAPI) ListAirportsV1(
	ctx context.Context,
	req *pb.ListAirportsV1Request,
) (*pb.ListAirportsV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ListAirportV1")
	defer span.Finish()

	switchLogLevel(ctx)

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("ListAirportsV1 - invalid argument")
		// right now it's redundant, but we'll add filter/pagination args in future
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	airports, err := o.repo.List(ctx, req.Cursor, req.Limit)
	if err != nil {
		log.Error().Err(err).Msg("ListAirportsV1 -- failed")
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Msg("DEBUG LEVEL: ListAirportsV1 - success")
	responseAirports := make([]*pb.Airport, 0, len(airports))
	for _, airport := range airports {
		responseAirports = append(responseAirports, &pb.Airport{
			Id:       airport.ID,
			Name:     airport.Name,
			Location: airport.Location,
		})
	}
	log.Info().Msg("INFO LEVEL:ListAirportsV1 - success")
	span.LogKV("airportCount", len(airports))

	return &pb.ListAirportsV1Response{
		Airports: responseAirports,
	}, nil
}
