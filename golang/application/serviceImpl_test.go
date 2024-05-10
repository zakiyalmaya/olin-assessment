package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestion1(t *testing.T) {
	service := NewService()

	testCases := []struct {
		nums     []int
		target   int
		response []int
	}{
		{nums: []int{1, 2, 3}, target: 4, response: []int{0, 2}},
		{nums: []int{2, 7, 3, 11, -3}, target: 0, response: []int{2, 4}},
		{nums: []int{2, 7, 3, 11, -3}, target: 20, response: []int{}},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.response, service.Question1(testCase.nums, testCase.target))
	}
}

func TestQuestion2(t *testing.T) {
	service := NewService()

	testCases := []struct {
		nums     []int
		response struct {
			result         [][]int
			timeComplexity string
		}
	}{
		{nums: []int{1, -1, -1, 2, 3},
			response: struct {
				result         [][]int
				timeComplexity string
			}{result: [][]int{{-1, -1, 2}}, timeComplexity: "O(n^2)"},
		},
		{nums: []int{1, -1, -1, 2, -1, -1, 0, 2, -4},
			response: struct {
				result         [][]int
				timeComplexity string
			}{result: [][]int{{-4, 2, 2}, {-1, -1, 2}, {-1, 0, 1}}, timeComplexity: "O(n^2)"},
		},
	}

	for _, testCase := range testCases {
		result := service.Question2(testCase.nums)
		assert.Equal(t, testCase.response.result, result.Result)
	}
}

func TestQuestion3(t *testing.T) {
	service := NewService()

	testCases := []struct {
		str      string
		words []string
		response []int
	}{
		{str: "wordgoodwordgoodbestword", words: []string{"word", "good", "best", "word"}, response: []int{8}},
		{str: "wordgoodgoodgoodbestword", words: []string{"word", "good", "best", "word"}, response: []int{}},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.response, service.Question3(testCase.str, testCase.words))
	}
}
