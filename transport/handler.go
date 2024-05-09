package transport

import (
	"github.com/gorilla/mux"
	golangSvc "github.com/zakiyalmaya/olin-assessment/application/golang_assessment"
	"github.com/zakiyalmaya/olin-assessment/transport/controller"
)

func Handler(golangAssessmentSvc golangSvc.Service, r *mux.Router) {
	ctrl := controller.NewController(golangAssessmentSvc)

	r.HandleFunc("/golang/question1", ctrl.GolangAssessmentCtrl.Question1).Methods("GET")
	r.HandleFunc("/golang/question2", ctrl.GolangAssessmentCtrl.Question2).Methods("GET")
	r.HandleFunc("/golang/question3", ctrl.GolangAssessmentCtrl.Question3).Methods("GET")
}