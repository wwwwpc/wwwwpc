package test1

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 != nil && l2 == nil {
		return l1
	}
	if l1 == nil && l2 != nil {
		return l2
	}

	tempList := &ListNode{0, nil}
	tempList1 := tempList
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tempList1.Next = l1
			l1 = l1.Next
		} else {
			tempList1.Next = l2
			l2 = l2.Next
		}
		tempList1 = tempList1.Next
	}

	if l1 != nil {
		tempList1.Next = l1
	}

	if l2 != nil {
		tempList1.Next = l2
	}
	return tempList.Next
}

func Test2() {
	defer func() {
		recover()
		fmt.Printf("recover:", recover())
	}()
	list3 := ListNode{
		3,
		nil,
	}
	list2 := ListNode{
		2,
		&list3,
	}

	list1 := ListNode{
		1,
		&list2,
	}

	list6 := ListNode{
		6,
		nil,
	}
	list5 := ListNode{
		5,
		&list6,
	}

	list4 := ListNode{
		4,
		&list5,
	}
	result := MergeTwoLists(&list1, &list4)
	for result != nil {
		fmt.Println("val :", result.Val)
		result = result.Next
	}
}

// 可以引⼊的库和版本相关请参考 “环境说明”
// 必须定义⼀个 包名为 `main` 的包

// 本题面试官已设置测试用例
func solution(str string) string {
	// 在这⾥写代码
	ss := strings.Split(str, " ")
	newStr := ""
	for k, v := range ss {
		if len(v) >= 5 {
			ss[k] = changeStr(ss[k])
		}
		if k == 0 {
			newStr += ss[k]
		} else {
			newStr = newStr + " " + ss[k]
		}
	}
	return newStr
}

func changeStr(str string) string {
	length := len(str)
	array := make([]string, length)
	for i, v := range str {
		array[i] = string(v)
	}
	newArray := make([]string, length)
	for i := 0; i < length; i++ {
		newArray[i] = array[length-i-1]
	}
	str1 := ""
	for _, v := range newArray {
		str1 += v
	}
	return str1
}

func Append(s []int){ s = append(s,5) } func Add(s []int) { for i :=range s