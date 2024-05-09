package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zakiyalmaya/olin-assessment/application"
	"github.com/zakiyalmaya/olin-assessment/transport"
)

func main() {

	// instatiate service
	golangAssessmentSvc := application.NewService()

	// instatiate router
	r := mux.NewRouter()

	// call handlers
	transport.Handler(golangAssessmentSvc, r)

	http.ListenAndServe(":8080", r)
}