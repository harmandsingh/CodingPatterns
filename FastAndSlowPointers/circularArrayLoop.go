package main

import (
	"fmt"
)

func circularArrayLoop(nums []int) bool {

	for i := range nums {
		// Initilize slow and fast pointers
		slow, fast := i, i

		// Set direction of pointer
		forward := nums[i] > 0

		for {
			// Move slow pointer
			slow = (slow + nums[slow]) % len(nums)
			if slow < 0 {
				slow += len(nums)
			}
			// Check if cycle does not exist
			if forward != (nums[slow] > 0) || nums[slow]%len(nums) == 0 {
				break
			}

			// Move fast pointer
			fast = (fast + nums[fast]) % len(nums)
			if fast < 0 {
				fast += len(nums)
			}
			// Check if cycle does not exist
			if forward != (nums[fast] > 0) || nums[fast]%len(nums) == 0 {
				break
			}

			// Move fast pointer again
			fast = (fast + nums[fast]) % len(nums)
			if fast < 0 {
				fast += len(nums)
			}
			// Check if cycle does not exist
			if forward != (nums[fast] > 0) || nums[fast]%len(nums) == 0 {
				break
			}

			// Check if slow and fast pointer values are equal
			if slow == fast {
				return true
			}
		}
	}
	return false
}

func main() {
	nums := []int{1, 3, -2, -4, 1}
	fmt.Println(circularArrayLoop(nums))
	nums2 := []int{2, 1, -1, -2}
	fmt.Println(circularArrayLoop(nums2))
	nums3 := []int{5, 4, -2, -1, 3}
	fmt.Println(circularArrayLoop(nums3))
	nums4 := []int{1, 2, -3, 3, 4, 7, 1}
	fmt.Println(circularArrayLoop(nums4))
	nums5 := []int{3, 3, 1, -1, 2}
	fmt.Println(circularArrayLoop(nums5))
}
