package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

var route *mux.Router

func SetRoute(r *mux.Router) {
	route = r
}

func Name2URL(routeName string, pairs ...string) string {

	url, err := route.Get(routeName).URL(pairs...)
	if err != nil {
		return ""
	} else {
		return url.String()
	}
}

func GetRouteVariable(key string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[key]
}
