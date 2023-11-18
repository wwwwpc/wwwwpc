package runtimeTest

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func Main() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for {
	}
}

func GoschedTest() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("hello")
	}
}

func GoroutinePool() {
	// 需要2个管道
	// 1.job管道
	jobChan := make(chan *Job, 128)
	// 2.结果管道
	resultChan := make(chan *Result, 128)
	// 3.创建工作池
	createPool(100000, jobChan, resultChan)
	// 4.开个打印的协程
	go func(resultChan chan *Result) {
		// 遍历结果管道打印
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id,
				result.job.RandNum, result.sum)
		}
	}(resultChan)
	var id int
	// 循环创建job，输入到管道
	for {
		id++
		// 生成随机数
		r_num := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}
		jobChan <- job
	}
}

// 创建工作池
// 参数1：开几个协程
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	// 根据开协程个数，去跑运行
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			// 执行运算
			// 遍历job管道所有数据，进行相加
			for job := range jobChan {
				// 随机数接过来
				r_num := job.RandNum
				// 随机数每一位相加
				// 定义返回值
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				// 想要的结果是Result
				r := &Result{
					job: job,
					sum: sum,
				}
				//运算结果扔到管道
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}

type Job struct {
	// id
	Id int
	// 需要计算的随机数
	RandNum int
}

type Result struct {
	// 这里必须传对象实例
	job *Job
	// 求和
	sum int
}

func Test() {
	var m = map[string]int{
		"one":   2,
		"two":   3,
		"three": 4,
	}
	var v []string
	for k, val := range m {
		fmt.Println(val)
		v = append(v, k)
	}
	sort.Strings(v)
	for _, s := range v {
		fmt.Println(m[s])
	}
	return
}

function(){
if(!SUCCEEDED(Operation1())){
	return OPERATION1FAILED
}
if(!SUCCEEDED(Operation2())){
return OPERATION2FAILED
}
if(!SUCCEEDED(Operation3())){
return OPERATION3FAILED
}
if(!SUCCEEDED(Operation4())){
return OPERATION4FAILED
}
return S_OK
}

func findPair(arr []int, sum int){
	if len(arr) <= 1 {
		fmt.Println("数组最少有两个数！")
		return
	}
	sort.Ints(arr)
	low := 0
	high := len(arr) - 1
	for low <= high{
		if arr[low] + arr[high] == sum{
			fmt.Printf("(%d, %d).\n", arr[low], arr[high])
			low++
			high--
		} else if arr[low] + arr[high] < sum {
			low++
		} else {
			high--
		}
	}
}

func ChangeIp(str string) int {
	strArr := strings.Split(str, ".")
	if len(strArr) != 4 {
		fmt.Printf("输入的ip错误")
		return 0
	}
	//将ip改为255进制
	num1,_ := strconv.ParseFloat(strArr[0],64)
	num2,_ := strconv.ParseFloat(strArr[1],64)
	num3,_ := strconv.ParseFloat(strArr[2],64)
	num4,_ := strconv.ParseFloat(strArr[3],64)
	return int(num1 * math.Pow(255,3) + num2 * math.Pow(255,2) + num3 * math.Pow(255,1) + num4 * math.Pow(255,0))
}