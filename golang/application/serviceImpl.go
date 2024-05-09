package application

import "sort"

type serviceImpl struct{}

func NewService() Service {
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

func (s *serviceImpl) Question2(nums []int) struct {
	Result         [][]int
	TimeComplexity string
} {
	result := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			indexLeft, indexRight := i+1, len(nums)-1
			sum := 0 - nums[i]
			isAdded := true

			for indexLeft < indexRight {
				if nums[indexLeft]+nums[indexRight] == sum {
					for _, r := range result {
						if r[0] == nums[i] && r[1] == nums[indexLeft] && r[2] == nums[indexRight] {
							isAdded = false
							break
						}
						isAdded = true
					}

					if isAdded {
						result = append(result, []int{nums[i], nums[indexLeft], nums[indexRight]})
					}

					indexLeft++
					indexRight--
				} else if nums[indexLeft]+nums[indexRight] < sum {
					indexLeft++
				} else {
					indexRight--
				}
			}
		}
	}

	timeComplexity := "The time complexity is O(n^2) because there is a double nested loop. " +
		"The first loop has a complexity of O(n), iterating through n array elements, " +
		"while the second loop, which checks the result array, also has a complexity of O(n)."

	return struct {
		Result         [][]int
		TimeComplexity string
	}{
		Result:         result,
		TimeComplexity: timeComplexity,
	}
}

func (s *serviceImpl) Question3(str string, words []string) []int {
	result := make([]int, 0)
	wordLen := len(words[0])
	totalWords := len(words)
	wordsMap := make(map[string]int)

	for _, word := range words {
		wordsMap[word]++
	}

	for i := 0; i <= len(str)-wordLen*totalWords; i++ {
		countSeenWordMap := make(map[string]int)
		wordCombination := 0
		for wordCombination < totalWords {
			word := str[i+wordCombination*wordLen : i+(wordCombination+1)*wordLen]
			if count, found := wordsMap[word]; found {
				countSeenWordMap[word]++
				if countSeenWordMap[word] > count {
					break
				}
			} else {
				break
			}
			wordCombination++
		}
		if wordCombination == totalWords {
			result = append(result, i)
		}
	}

	return result
}
