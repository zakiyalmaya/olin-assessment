package golangassessment

import (
	"net/http"
	"strconv"
	"strings"

	golangSvc "github.com/zakiyalmaya/olin-assessment/application/golang_assessment"
	"github.com/zakiyalmaya/olin-assessment/transport/response"
)

type Controller struct {
	golangAssessmentSvc golangSvc.Service
}

func NewGolangAssessmentController(golangAssessmentSvc golangSvc.Service) *Controller {
	return &Controller{
		golangAssessmentSvc: golangAssessmentSvc,
	}
}

func (c *Controller) Question1(w http.ResponseWriter, r *http.Request) {
	numsStr := r.URL.Query().Get("nums")
	targetStr := r.URL.Query().Get("target")

	numsArrStr := strings.Split(numsStr, ",")
	nums := make([]int, len(numsArrStr))
	for i, numStr := range numsArrStr {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		nums[i] = num
	}

	target, err := strconv.Atoi(targetStr)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	result := c.golangAssessmentSvc.Question1(nums, target)
	response.RespondWithJSON(w, http.StatusOK, result)
}
