package main

import "fmt"

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

/* 要点：
1.滑动窗口，用start、end标记窗口库边界。
2.使用map标记窗口内元素是否重复

*/

func main() {
	ret := find("abcabcbb")
	fmt.Println(ret)
}

func find(str string) int {
	start := 0
	end := -1
	set := map[byte]bool{}

	ret := 0

	for ; start < len(str); start++ {
		if start != 0 {
			delete(set, str[start-1])
		}

		// 这里使用end+1，先"试探"右边界，判断右边界合法性。如果合法，再真正扩充右边界，即end++。否则，表示右窗口试探到重复元素了，可以不再继续试探右窗口了，后续提高左窗口，继续试探右窗口。
		for end+1 < len(str) && !set[str[end+1]] {
			set[str[end+1]] = true
			end++
		}

		if end-start+1 > ret {
			ret = end - start + 1
		}
	}

	return ret
}
