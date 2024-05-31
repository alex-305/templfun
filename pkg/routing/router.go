package routing

import "net/http"

type Router struct {
	*http.ServeMux
	Routes map[string]http.Handler
}

func NewRouter() Router {
	return Router{
		ServeMux: http.NewServeMux(),
		Routes:   make(map[string]http.Handler),
	}
}
