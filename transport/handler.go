package transport

import (
	"github.com/gorilla/mux"
	"github.com/zakiyalmaya/olin-assessment/application"
	"github.com/zakiyalmaya/olin-assessment/transport/controller"
)

func Handler(service application.Service, r *mux.Router) {
	ctrl := controller.NewController(service)

	r.HandleFunc("/question1", ctrl.Question1).Methods("GET")
	r.HandleFunc("/question2", ctrl.Question2).Methods("GET")
	r.HandleFunc("/question3", ctrl.Question3).Methods("GET")
}