package slice

import (
	"fmt"
)

func Slice() {
	// 创建容量为5的切片
	x := make([]int, 0, 5)
	for i := 0; i < 8; i++ {
		// slice 类似动态数组 当大于容量限制时 自动分配更大的存储空间
		x = append(x, i)
	}
	fmt.Println(x)
}
