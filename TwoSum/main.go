package main

func main() {
	//Call one of this method and tell me what do you think
}

//TwoSum ... Return indexs of the two values that sum the target value
//Runtime: 4 ms, faster than 98.68% of Go online submissions for Two Sum.
//Memory Usage: 3.7 MB, less than 45.10% of Go online submissions for Two Sum.
func TwoSum(nums []int, target int) []int {
	substractions := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		temp := target - nums[i]
		if index, ok := substractions[nums[i]]; ok {
			return []int{index, i} //Index of the previous substraction and the next
		}
		substractions[temp] = i
	}
	return nil
}

//TwoSumVersion1 ... The same than TwoSum, but it is ohter version
//Runtime: 36 ms, faster than 36.37% of Go online submissions for Two Sum.
//Memory Usage: 3 MB, less than 81.78% of Go online submissions for Two Sum.
func TwoSumVersion1(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

//TwoSumVersion2 ... The same than TwoSum, but it is ohter version
//Runtime: 32 ms, faster than 38.80% of Go online submissions for Two Sum.
//Memory Usage: 3 MB, less than 72.41% of Go online submissions for Two Sum.
func TwoSumVersion2(nums []int, target int) []int {
	for i, ii := range nums {
		for j, jj := range nums[i+1:] {
			if jj+ii == target {
				return []int{i, j + i + 1}
			}
		}
	}
	return nil
}
