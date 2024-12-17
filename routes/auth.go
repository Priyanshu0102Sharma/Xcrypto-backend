package routes

import (
	"server/controllers"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	r.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	// r.HandleFunc("/login", controllers.Login).Methods("POST")
	// r.HandleFunc("/forgot-password", controllers.ForgotPassword).Methods("POST")
}
