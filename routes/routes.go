package routes

import (
	"github.com/gorilla/mux"
	"main.go/controllers"
	"main.go/middlewares"
)

func GetRouter() *mux.Router {
	Router := mux.NewRouter()
	Router.HandleFunc("/", middlewares.DBcontext(controllers.Home)).Methods("GET")
	Router.HandleFunc("/createUser", middlewares.DBcontext(controllers.CreateUser)).Methods("POST")
	Router.HandleFunc("/createProduct", middlewares.DBcontext(controllers.CreateProduct)).Methods("POST")
	Router.HandleFunc("/getUsers", middlewares.DBcontext(controllers.GetUsers)).Methods("POST")
	Router.HandleFunc("/getProducts", middlewares.DBcontext(controllers.GetProducts)).Methods("POST")
	Router.HandleFunc("/verify", middlewares.DBcontext(controllers.VerifyOrder)).Methods("POST")

	return Router
}
