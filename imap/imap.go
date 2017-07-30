package imap

import (
	"fmt"
)

func MapFunc() {
	// 将字典类型内置， 可以直接从运行层面获得性能优化
	m := make(map[string]int)
	m["a"] = 1
	x, ok := m["b"]
	fmt.Println(x, ok)
	delete(m, "a")
}
