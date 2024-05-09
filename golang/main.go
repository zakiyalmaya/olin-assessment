package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zakiyalmaya/olin-assessment/golang/application"
	"github.com/zakiyalmaya/olin-assessment/golang/transport"
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