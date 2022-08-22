package http

import (
	"brexs-test/services"
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
)

type Server struct {
	router         *mux.Router
	server         *http.Server
	routesBusiness *services.RoutesBusiness
}

func NewServer(addr string, r *services.RoutesBusiness) *Server {
	s := newServer(addr, r)
	s.registerRoutes(s.router)

	return s
}

func newServer(addr string, r *services.RoutesBusiness) *Server {
	router := mux.NewRouter(
		mux.WithAnalytics(true),
	)

	s := &Server{
		server:         &http.Server{},
		router:         router,
		routesBusiness: r,
	}

	s.server.Addr = ":" + addr
	s.server.Handler = s.router

	return s
}

func (s *Server) Open() {
	defer logrus.Infof("webserver started on port %s", s.server.Addr)

	go func() {
		_ = s.server.ListenAndServe()
	}()
}

func (s *Server) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return s.server.Shutdown(ctx)
}
