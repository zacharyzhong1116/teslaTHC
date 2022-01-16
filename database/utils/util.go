package utils

import (
	"net/http"
	"teslaTHC/database/persistence"
)

type RouterImpl struct {
	Router mux.Router
	Db     persistence.DbImpl
}

func (ri *Router) NewRouter() {
	ri.Router = mux.NewRouter()
}

//Register  Regist the handler, currently we have the post only and we only get the integer and save it into data
func (ri *Router) Register(path string, f func(w http.ResponseWriter, r *http.Request), verb string) {
	ri.Router.HandleFunc(path, f).Methods(verb)
}
