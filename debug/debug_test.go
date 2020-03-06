package main

import (
	"fmt"
	"testing"
)

func Test_main(t *testing.T) {

	nums := []int{3, 2, 2, 3}
	d := 3

	a := removeElement(nums, d)
	fmt.Println("ssss", a)
	fmt.Print(nums)
}

func removeElement(nums []int, val int) int {

	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i

}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {

		fmt.Println("i:", i, "j:", j)
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
		fmt.Println("iI:", i)

	}
	return i

}
