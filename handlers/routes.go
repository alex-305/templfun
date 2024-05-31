package handlers

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/alex-305/templfun/pkg/routing"
	"github.com/alex-305/templfun/templates/pages"
)

type APIServer struct {
	ListenAddress string
	router        routing.Router
}

func CreateServer(listenAddress string) APIServer {
	return APIServer{
		ListenAddress: listenAddress,
	}
}

func (s *APIServer) Start() error {
	s.router = routing.NewRouter()
	s.defineRoutes()

	log.Printf("Server is listening on %s", s.ListenAddress)
	return http.ListenAndServe(s.ListenAddress, &s.router)
}

func (s *APIServer) defineRoutes() {
	s.router.HandleRoute("/", templ.Handler(pages.Index()))
	s.router.HandleRoute("/users", templ.Handler(pages.Users([]string{"thing 1", "thing 2", "thing 3", "thing 4"})))
	s.router.HandleRoute("/404", templ.Handler(pages.NotFound()))
}
