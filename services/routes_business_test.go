package services

import (
	"brexs-test/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoutesBusiness(t *testing.T) {

	routesBusiness := NewRoutesBusiness("../input_test.csv")
	routesBusiness.ReadFile()

	t.Run("it should successfully get the best route", func(t *testing.T) {
		routeSchema := domain.RouteSchema{
			Origin:  "GRU",
			Destiny: "ORL",
		}
		bestRoute := routesBusiness.FindBestRoute(routeSchema)
		assert.Equal(t, bestRoute, "GRU-BRC-SCL-ORL > $35")
	})

	t.Run("it should successfully save the new route", func(t *testing.T) {
		routeSchema := domain.RouteSchema{
			Origin:  "POA",
			Destiny: "ORL",
			Cost:    120,
		}

		err := routesBusiness.SaveFile(routeSchema)
		assert.NoError(t, err)
	})

}
