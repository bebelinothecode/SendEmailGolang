package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int);

	for i:=0; i < len(nums); i++ {
		another := target - nums[i];

		if _, ok := m[another]; ok {
			return []int{m[another],i}
		}
		m[nums[i]] = i;
	}
	return nil
}

func main() {
	lot := [...]int{7,2,13,11}
	fmt.Print(twoSum(lot[:],9))
}


