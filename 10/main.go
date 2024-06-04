package main

import (
	"fmt"
	"sort"
)

type SupSet struct {
	nums []float64
}

func main() {
	nums := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	// fmt.Println(nums)
	res := []SupSet{}
	start := 0
	for i := (int(nums[0]) / 10) * 10; i < int((nums[len(nums)-1]/10+1)*10); i += 10 {
		if nums[start] < float64(i+10) && nums[start] > float64(i-10) {
			sup := SupSet{nums: []float64{}}
			for k := start; k < len(nums) && nums[k] < float64(i+10); k++ {
				sup.nums = append(sup.nums, nums[k])
				start = k
			}
			// fmt.Println(sup.nums)
			res = append(res, sup)
			if start < len(nums)-1 {
				start++
			} else {
				break
			}
		}
	}

	fmt.Println(res)

}
