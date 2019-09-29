package main

import (
	"fmt"
	"strconv"
)

//----------------------------------------------------------------------------------------------------------------------
// 验证：使用for range时，如果每个元素是指针，对每个元素的修改是可以直接生效的。
// for-range 的坑：如果每个元素是指针，把每个元素添加到其他数组，会发现全都是最后一个元素。
//----------------------------------------------------------------------------------------------------------------------

type Student struct {
	Id   int
	Name string
}

func main() {
	studentArr := []*Student{
		{
			Id:   0,
			Name: "first",
		},
		{
			Id:   1,
			Name: "second",
		},
		{
			Id:   2,
			Name: "third",
		},
	}

	for i, s := range studentArr {
		s.Name = s.Name + strconv.Itoa(i) // 修改Name
	}

	// 输出：
	//&{0 first0}
	//&{1 second1}
	//&{2 third2}
	for _, s := range studentArr {
		fmt.Println(s)
	}

}
