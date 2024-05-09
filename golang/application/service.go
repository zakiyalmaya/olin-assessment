package application

type Service interface {
	Question1(nums []int, target int) []int
	Question2(nums []int) struct {
		Result         [][]int
		TimeComplexity string
	}
	Question3(str string, words []string) []int
}
