package main

func majorityElement(nums []int) int {
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	maxK := 0
	maxV := m[0]
	for k, v := range m {
		if v > maxV {
			maxK = k
			maxV = v
		}
	}
	return maxK
}
