// test.go
package base

import (
	"errors"
	"fmt"
)

func Base() {
	// 编译器确保变量总是被初始化为0，避免出现意外情况
	var i int32
	var s = "hello world"
	println(i, s)

	// 简写类型推导
	simpleInt := 100
	println(simpleInt)

	// 表达式
	/* if */
	x := 100
	if x > 0 {
		println("x > 0")
	} else if x < 0 {
		println("x < 0")
	} else {
		println("0")
	}

	/* switch */
	l := 100
	switch {
	case l > 0:
		println("l > 0")
		break
	case l < 0:
		println("l < 0")
		break
	default:
		println("0")
		break
	}

	/* for */
	for i := 0; i < 5; i++ {
		println(i)
	}

	for i := 4; i >= 0; i-- {
		println(i)
	}
	/* for range */

	q := []int{100, 101, 102}
	for i, n := range q {
		println(i, " ", n)
	}
	/* 函数 */
	a, b := 10, 2
	c, error := div(a, b)
	fmt.Println(c, error)
	// return func() && defer
	y := 100
	f := test(y)
	f()
}

func test(x int) func() {
	defer println("hi\n")
	println("x 的生命周期:")
	return func() {
		println(x)
	}
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}
