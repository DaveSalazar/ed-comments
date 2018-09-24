package routes

import (
	"github.com/DaveSalazar/ed-comments/controllers"
	"github.com/gorilla/mux"
)

//SetLoginRouter router para login
func SetLoginRouter(router *mux.Router) {
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
}
