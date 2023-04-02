package router

import (
	"fmt"
	"net/http"
)

type Router struct {
	mux *http.ServeMux
}

func NewRouter() *Router {
	r := &Router{
		mux: http.NewServeMux(),
	}
	return r
}

func (r *Router) Serve() {
	http.ListenAndServe(fmt.Sprintf(":%s", "8080"),r.mux)
}
