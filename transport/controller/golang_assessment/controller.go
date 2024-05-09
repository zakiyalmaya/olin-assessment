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

	if numsStr == "" || targetStr == "" {
		response.RespondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}

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

func (c *Controller) Question2(w http.ResponseWriter, r *http.Request) {
	numsStr := r.URL.Query().Get("nums")

	if numsStr == "" {
		response.RespondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}

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

	result := c.golangAssessmentSvc.Question2(nums)
	response.RespondWithJSON(w, http.StatusOK, result)
}

func (c *Controller) Question3(w http.ResponseWriter, r *http.Request) {
	sentence := r.URL.Query().Get("sentence")
	words := r.URL.Query().Get("words")
	wordsArr := strings.Split(words, ",")

	if sentence == "" || words == "" {
		response.RespondWithError(w, http.StatusBadRequest, "invalid request")
		return
	}

	result := c.golangAssessmentSvc.Question3(sentence, wordsArr)
	response.RespondWithJSON(w, http.StatusOK, result)
}
