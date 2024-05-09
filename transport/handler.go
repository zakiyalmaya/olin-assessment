package transport

import (
	"github.com/gorilla/mux"
	golangSvc "github.com/zakiyalmaya/olin-assessment/application/golang_assessment"
	"github.com/zakiyalmaya/olin-assessment/transport/controller"
)

func Handler(golangAssessmentSvc golangSvc.Service, r *mux.Router) {
	ctrl := controller.NewController(golangAssessmentSvc)

	r.HandleFunc("/golang/question1", ctrl.GolangAssessmentCtrl.Question1).Methods("GET")
}