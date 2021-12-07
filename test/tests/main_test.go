//go:build integration
// +build integration

package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	airport "github.com/ozonmp/trv-airport-api/pkg/trv-airport-api"
	"gitlab.ozon.dev/qa/teachers/qa-route-256/item-api/test/internal/config"
	"gitlab.ozon.dev/qa/teachers/qa-route-256/item-api/test/internal/interceptor"
	"google.golang.org/grpc"
)

var (
	ctx        context.Context
	TrvAirportApiServiceClient airport.TrvAirportApiServiceClient
)

func TestMain(m *testing.M) {
	cfg := config.ProcessConfig()
	ctx = context.Background()

	conn, err := grpc.Dial(
		cfg.Host,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptor.Timeout(time.Second)),
		grpc.WithUnaryInterceptor(interceptor.Duration()),
	)
	if err != nil {
		panic(fmt.Errorf("grpc.Dial() err: %v", err))
	}
	defer conn.Close()

	itemClient = airport.NewItemClient(conn)

	m.Run()
}
