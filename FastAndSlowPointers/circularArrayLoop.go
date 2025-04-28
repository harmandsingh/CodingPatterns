package main

import (
	"fmt"
)

func circularArrayLoop(nums []int) bool {

	fast, slow := 0, 0
	for slow < len(nums) {
		cycleLength := 0
		fast = slow
		for true {
			cycleLength++
			prev := nums[fast]
			if prev < 0 {
				fast = (fast - nums[fast]) % len(nums)
				if nums[fast] > 0 {
					break
				}
			} else {
				fast = (fast + nums[fast]) % len(nums)
				if nums[fast] < 0 {
					break
				}
			}
			if slow == fast {
				if cycleLength > 1 {
					return true
				} else if cycleLength == 1 {
					break
				}
			}
		}
		slow++
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
}
