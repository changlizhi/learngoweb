package httpserver

import (
	"fmt"
	"net/http"
	"strings"
)

type Router struct {
	Route map[string]map[string]http.HandlerFunc
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := rt.Route[r.Method][r.URL.String()]; ok {
		h(w, r)
	}
}

func (rt *Router) HandleFunc(method, path string, f http.HandlerFunc) {
	method = strings.ToUpper(method)
	if rt.Route == nil {
		rt.Route = make(map[string]map[string]http.HandlerFunc)
	}
	if rt.Route[method] == nil {
		rt.Route[method] = make(map[string]http.HandlerFunc)
	}
	rt.Route[method][path] = f
}

func NewRouter() *Router {
	return new(Router)
}
func My_Listen() {
	r := NewRouter()
	r.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "http GET!")
	})
	r.HandleFunc("POST", "/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "http POST!")
	})
	http.ListenAndServe(":9090", r)
}
