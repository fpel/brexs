package http

import (
	"brexs-test/domain"
	"encoding/json"
	"io/ioutil"

	"github.com/sirupsen/logrus"

	"net/http"

	"gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
)

func (s *Server) registerRoutes(r *mux.Router) {
	r.HandleFunc("/", s.handleGetBestRoute).Methods("GET")
	r.HandleFunc("/", s.handleSaveNewRoute).Methods("POST")
}

func (s *Server) handleGetBestRoute(w http.ResponseWriter, r *http.Request) {
	// response := "GRU - BRC - SCL - ORL - CDG ao custo de $40"
	var input domain.RouteSchema
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bodyBytes, &input)

	s.routesBusiness.ReadFile()
	// res := s.routesBusiness.Content
	response := s.routesBusiness.FindBestRoute(input)

	// response := input

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logrus.WithError(err).Error("an error happened while encoding health check response")
	}
}

func (s *Server) handleSaveNewRoute(w http.ResponseWriter, r *http.Request) {
	var input domain.RouteSchema
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bodyBytes, &input)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	s.routesBusiness.SaveFile(input)

}
