package main

// 快排
func quickSort(arr []int, low int, high int) {
	if low < high {
		middle := getMiddle(arr, low, high) //将list数组进行一分为二
		quickSort(arr, low, middle-1)       //对低字表进行递归排序
		quickSort(arr, middle+1, high)      //对高字表进行递归排序
	}
}

func getMiddle(arr []int, low int, high int) int {
	middleValue := arr[low] // 数组的第一个作为中轴，同时middleValue也起到了暂存的作用，最后要恢复。
	for low < high {
		for low < high && arr[high] > middleValue { // 先看高位。找到第一个不合法的高位，放到低位。
			high--
		}
		arr[low] = arr[high]

		for low < high && arr[low] < middleValue { // 再看低位。找到第一个不合法的低位，放到高位
			low++
		}
		arr[high] = arr[low]
	}

	arr[low] = middleValue // 处理完毕，把低位的值恢复为middleValue。

	return low
}
