package arrayTest

import "fmt"

func ArrayTest() {
	var arr1 [5]int
	printArr(&arr1)
	fmt.Println(arr1)
	arr2 := [...]int{2, 4, 6, 8, 10}
	printArr(&arr2)
	fmt.Println(arr2)
}

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func Addtest() {
	myTest([5]int{1, 3, 5, 8, 7}, 8)
}

func myTest(arr [5]int, num int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == num {
				fmt.Printf("%d,%d\n", i, j)
			}
		}
	}
}
