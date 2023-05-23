package main

import "fmt"

// https://leetcode.cn/problems/maximum-subarray/solution/zui-da-zi-xu-he-by-leetcode-solution/
// 数组的最大子序列和

/*
用f(i) 代表以第 i 个数结尾的"连续子数组的最大和"

只需要求出每个位置的f(i)，然后返回 f 数组中的最大值即可。那么我们如何求f(i) 呢？
可以考虑 nums[i] 单独成为一段还是加入 f(i−1) 对应的那一段，这取决于
nums[i] 和 f(i−1)+nums[i] 的大小，我们希望获得一个比较大的
*/

func main() {
	arr := []int{1, 2, 3, 20}
	fmt.Println(maxSubArray(arr))
}

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		// 核心代码。nums[i-1]的值已经被覆盖为此位置的连续子数组最大值。
		// 则判断i的位置时，判断加上i之后，会不会得到更大值的子数组，如果得到更大子数组，则子数组拼上i。如果得不到更大子数组，则从i重新开启新的子数组。
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}

		if nums[i] > max {
			max = nums[i]
		}
		//fmt.Println(nums)
	}
	return max
}
