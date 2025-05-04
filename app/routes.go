package app

import (
	"github.com/gorilla/mux"
	"github.com/yohan20-coder/go-toko-online/app/controllers"
)

func (server *Server) InitializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
	server.Router.HandleFunc("/test", controllers.Test).Methods("GET")
}
