package router

import (
	"net/http"

	"github.com/adrian83/go-mvc-library/library/web/handler"
	"github.com/gorilla/mux"
)

// CreateRouter creates new Router.
func CreateRouter(controllers ...handler.Controller) *mux.Router {
	mux := mux.NewRouter()
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../static/"))))

	for _, controller := range controllers {
		for _, route := range controller.Routes() {
			mux.Handle(route.Path, route.Handler).Methods(route.Method)
		}
	}

	return mux
}