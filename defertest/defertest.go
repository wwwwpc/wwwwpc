package defertest

import "fmt"

func TestDefer() {
	a := 1
	b := 2
	defer func() {
		println(2) // 将 interface{} 转型为具体类型。
	}()
	defer func() {
		println(1) // 将 interface{} 转型为具体类型。

	}()
	defer func() {
		println(0) // 将 interface{} 转型为具体类型。

	}()
	//panic("err!")
	var br = make(chan int, 10)
	testChan(a, br)
	var y = <-br
	println("y", y)
	println(br)
	br <- 1
	println(br)
	x := <-br
	println(x)
	println(a, "+", b, a+b)
}

func testChan(a int, br chan int) {
	br <- a
}

func DeferTest1() {
	fmt.Println(test1())
	fmt.Println(test2())
	fmt.Println(test3())
	fmt.Println(test4())

	return
}

func test1() (v int) {
	defer fmt.Println(v)
	return v
}

func test2() (v int) {
	defer func() {
		fmt.Println(v)
	}()
	return 3
}

func test3() (v int) {
	defer fmt.Println(v)
	v = 3
	return 4
}

func test4() (v int) {
	defer func(n int) {
		fmt.Println(n)
	}(v)
	return 5
}
