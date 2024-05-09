package main

import (
	"net/http"

	"github.com/gorilla/mux"
	golangSvc "github.com/zakiyalmaya/olin-assessment/application/golang_assessment"
	"github.com/zakiyalmaya/olin-assessment/transport"
)

func main() {

	// instatiate service
	golangAssessmentSvc := golangSvc.NewGolangAssessmentService()

	// instatiate router
	r := mux.NewRouter()

	// call handlers
	transport.Handler(golangAssessmentSvc, r)

	http.ListenAndServe(":8080", r)
}