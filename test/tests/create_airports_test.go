//go:build integration
// +build integration

package tests

import (
	"context"
	airport "github.com/ozonmp/trv-airport-api/pkg/trv-airport-api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	ctx                        context.Context
	trvAirportApiServiceClient airport.TrvAirportApiServiceClient
)

func TestCreateAirport(t *testing.T) {

	testCases := []struct {
		title     string
		nameq     string
		locationq string
	}{
		{
			"Test 1",
			"AAA",
			"Location of AAA",
		},
		{
			"Test 2",
			"BBB",
			"Location of BBB",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			// t.Parallel()
			// errTest := errors.New("test announcement err")

			req := airport.CreateAirportV1Request{Airport: &airport.Airport{
				Name:     tc.nameq,
				Location: tc.locationq,
			}}

			res, err := airport.CreateAirportV1(ctx, &req)

			require.NoError(t, err)
			assert.NotZero(t, res.Id)
			t.Logf("id: %v", res.Id)
		})
	}
}
