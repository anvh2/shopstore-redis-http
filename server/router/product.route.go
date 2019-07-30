package router

import (
	"net/http"
	"shopstore/src/server/controller"
	constroller "shopstore/src/server/controller"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

//NewProductRouter -
func NewProductRouter() *mux.Router {
	router.HandleFunc("/products", controller.GetAllProducts).Methods("GET")
	router.HandleFunc("/products", constroller.InsertProduct).Methods("POST")

	staticFileDirectory := http.Dir("./resources/app/static/html")
	staticFileHandler := http.StripPrefix("/", http.FileServer(staticFileDirectory))

	router.PathPrefix("/").Handler(staticFileHandler).Methods("GET")

	return router
}
