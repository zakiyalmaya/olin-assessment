package golangassessment

type serviceImpl struct {}

func NewGolangAssessmentService() Service {
	return &serviceImpl{}
}

func (s *serviceImpl) Question1(nums []int, target int) []int {
	numIndexMap := make(map[int]int)
	for i, num := range nums {
		diff := target - num

		if j, ok := numIndexMap[diff]; ok {
			return []int{j, i}
		}

		numIndexMap[num] = i
	}

	return []int{}
}

func (s *serviceImpl) Question2(nums []int) [][]int {

	return [][]int{}
}

func (s *serviceImpl) Question3(str string, words []string) []int {
	
	return []int{}
}