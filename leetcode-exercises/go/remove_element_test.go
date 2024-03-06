package main

// removeElement removes `val` entries in nums slice and return the number of elements not removed.
// there are no tests because it worked first time, moving to the next one :)
func removeElement(nums []int, val int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			count++
			i--
		}
	}
	return len(nums)
}
