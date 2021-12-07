package api

import (
	"context"
	"github.com/ozonmp/trv-airport-api/internal/repo"
	"github.com/ozonmp/trv-airport-api/pkg/trv-airport-api"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/metadata"
	"strings"
)

//NewAirportAPI constructs API
func NewAirportAPI(r repo.AirportRepo) trv_airport_api.TrvAirportApiServiceServer {
	return &airportAPI{repo: r}
}

var (
	totalAirportNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "trv_airport_api_airport_not_found_total",
		Help: "Total number of airports that were not found",
	})
	totalAirportCreated = promauto.NewCounter(prometheus.CounterOpts{
		Name: "trv_airport_api_airport_created_total",
		Help: "Total number of airports that were created",
	})
	totalAirportDeleted = promauto.NewCounter(prometheus.CounterOpts{
		Name: "trv_airport_api_airport_deleted_total",
		Help: "Total number of airports that were deleted",
	})
)

type airportAPI struct {
	trv_airport_api.UnimplementedTrvAirportApiServiceServer
	repo repo.AirportRepo
}

func parseLevel(str string) (zerolog.Level, bool) {
	switch strings.ToLower(str) {
	case "debug":
		return zerolog.DebugLevel, true
	case "info":
		return zerolog.InfoLevel, true
	case "warn":
		return zerolog.WarnLevel, true
	case "error":
		return zerolog.ErrorLevel, true
	default:
		return zerolog.DebugLevel, false
	}
}

func switchLogLevel(ctx context.Context) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		levels := md.Get("log-level")
		log.Info().Strs("levels", levels).Msg("got log levels")
		if len(levels) > 0 {
			if parsedLevel, ok := parseLevel(levels[0]); ok {
				zerolog.SetGlobalLevel(parsedLevel)

			}
		}
	}
}
