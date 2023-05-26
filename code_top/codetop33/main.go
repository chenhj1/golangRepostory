package main

import "fmt"

// https://leetcode.cn/problems/search-in-rotated-sorted-array/solution/sou-suo-xuan-zhuan-pai-xu-shu-zu-by-leetcode-solut/
// 有序旋转数组的二分查找
/*
关键点：
1. 画图理解。
2. 每次二分得到middle值，要判断middle位于左区间还是右区间。判断方式：如果大于arr[0]，则位于左区间
*/

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	ret := find(arr, 1)
	fmt.Println(ret)

}

func find(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	if len(arr) == 1 {
		if arr[0] == target {
			return 1
		} else {
			return -1
		}
	}

	l := 0
	r := len(arr) - 1

	for l < r {
		middle := (l + r) / 2
		if arr[middle] == target {
			return middle

		}

		if arr[middle] > arr[0] { // 左区间
			if arr[middle] > target && arr[0] <= target { // 左区间的最左边是单调的，可以放心二分。
				r = middle - 1
			} else {
				l = middle + 1
			}
		} else { // 右区间
			if arr[middle] < target && arr[len(arr)-1] >= target { // 右区间的最右边是单调的，可以放心二分。
				l = middle + 1
			} else {
				r = middle - 1
			}
		}
	}
	return -1
}
