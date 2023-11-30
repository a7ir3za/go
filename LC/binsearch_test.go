package lc

import (
	"log"
	"testing"
)

func init() {
	log.Print("> BinSearch")
}

func TestBSLeftMost(t *testing.T) {
	bSearch := func(nums []int, k int) int {
		l, r := 0, len(nums)
		for l < r {
			m := l + (r-l)>>1
			if nums[m] >= k {
				r = m
			} else {
				l = m + 1
			}
		}
		return l
	}

	log.Print("2 ?= ", bSearch([]int{1, 2, 3, 4, 5}, 3))
	log.Print("0 ?= ", bSearch([]int{1, 2, 3}, 1))
	log.Print("0 ?= ", bSearch([]int{1, 2, 3, 4}, 0))
	log.Print("2 ?= ", bSearch([]int{1, 2, 5, 7}, 3))
	log.Print("3 ?= ", bSearch([]int{1, 2, 5, 7}, 7))
	log.Print("4 ?= ", bSearch([]int{1, 2, 3, 4}, 5))
}

// 33m Search in Rotated Sorted Array
func Test33(t *testing.T) {
	search := func(nums []int, target int) int {
		l, r := 0, len(nums)-1
		for l < r {
			m := l + (r-l)>>1

			log.Printf("%d   %d>%d | %d>%d | %d>%d ", target, l, nums[l], m, nums[m], r, nums[r])

			if nums[l] <= nums[m] {
				if nums[l] <= target && target <= nums[m] {
					r = m
				} else {
					l = m + 1
				}

			} else {
				if nums[m] <= target && target <= nums[r] {
					l = m
				} else {
					r = m - 1
				}
			}

		}

		if l < len(nums) && nums[l] == target {
			return l
		}
		return -1
	}

	log.Print("4 ?= ", search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	log.Print("-1 ?= ", search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	log.Print("-1 ?= ", search([]int{5, 5, 6, 7, 1, 1, 2}, 0))
}

// 153m Find Minimum in Rotated Sorted Array
func Test153(t *testing.T) {
	findMin := func(nums []int) int {
		l, r := 0, len(nums)-1

		for l < r {
			m := l + (r-l)>>1
			log.Printf("%d>%d | %d>%d | %d>%d", l, nums[l], m, nums[m], r, nums[r])

			if nums[m] < nums[r] {
				r = m
			} else {
				l = m + 1
			}
		}

		return nums[l]
	}

	log.Print("1 ?= ", findMin([]int{5, 1, 2, 3, 4}))
	log.Print("1 ?= ", findMin([]int{3, 4, 5, 1, 2}))
	log.Print("0 ?= ", findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	log.Print("0 ?= ", findMin([]int{5, 6, 7, 0, 1, 2, 3, 4}))
	log.Print("0 ?= ", findMin([]int{0, 1, 2, 3, 4, 5, 6, 7}))
	log.Print("0 ?= ", findMin([]int{1, 2, 3, 4, 5, 6, 7, 0}))
}
