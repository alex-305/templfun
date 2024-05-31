package routing

import "net/http"

func (r *Router) HandleRoute(path string, handler http.Handler) {
	r.Handle(path, handler)
	r.Routes[path] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, exists := r.Routes[req.URL.Path]
	if exists {
		handler.ServeHTTP(w, req)
	} else {
		r.Serve404(w, req)
	}
}

func (r *Router) Serve404(w http.ResponseWriter, req *http.Request) {
	_, exists := r.Routes["/404"]

	if exists {
		http.Redirect(w, req, "/404", http.StatusSeeOther)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
