package api

import(
	"github.com/lionelritchie29/staem-backend/middleware"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router{
	r:= mux.NewRouter()
	r.Use(middleware.LogMiddleware)
	r.Use(middleware.Auth)

	return r
}
