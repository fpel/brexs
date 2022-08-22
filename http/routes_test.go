package http

import (
	"brexs-test/domain"
	"brexs-test/services"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleGetBestRoute(t *testing.T) {

	routesBusiness := services.NewRoutesBusiness("../input_test.csv")

	t.Run("it should successfully get the best route", func(t *testing.T) {
		reqSchema := domain.RouteSchema{
			Origin:  "GRU",
			Destiny: "ORL",
		}
		body, err := json.Marshal(reqSchema)
		assert.NoError(t, err)

		req := httptest.NewRequest("GET", "/", bytes.NewReader(body))
		w := httptest.NewRecorder()
		s := NewServer("3000", routesBusiness)
		s.router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code, "The http code does not match the expected")
		assert.Equal(t, w.Body.Bytes(), []byte("GRU-BRC-SCL-ORL > $35"))
	})

	t.Run("it should successfully save the new route", func(t *testing.T) {
		reqSchema := domain.RouteSchema{
			Origin:  "POA",
			Destiny: "ORL",
			Cost:    120,
		}
		body, err := json.Marshal(reqSchema)
		assert.NoError(t, err)

		req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
		w := httptest.NewRecorder()
		s := NewServer("3000", routesBusiness)
		s.router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code, "The http code does not match the expected")

	})

}
