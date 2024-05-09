package controller

import (
	golangSvc "github.com/zakiyalmaya/olin-assessment/application/golang_assessment"
	golangCtrl "github.com/zakiyalmaya/olin-assessment/transport/controller/golang_assessment"
)

type Controller struct {
	GolangAssessmentCtrl *golangCtrl.Controller
}

func NewController(golangAssessmentSvc golangSvc.Service) *Controller {
	return &Controller{
		GolangAssessmentCtrl: golangCtrl.NewGolangAssessmentController(golangAssessmentSvc),
	}
}