package pachong

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	// w代表大小写字母+数字+下划线
	reEmail = `\w+@\w+\.\w+`
	// s?有或者没有s
	// +代表出1次或多次
	//\s\S各种字符
	// +?代表贪婪模式
	reLinke  = `href="(https?://[\s\S]+?)"`
	rePhone  = `1[3456789]\d\s?\d{4}\s?\d{4}`
	reIdcard = `[123456789]\d{5}((19\d{2})|(20[01]\d))((0[1-9])|(1[012]))((0[1-9])|([12]\d)|(3[01]))\d{3}[\dXx]`
	reImg    = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

// 处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

func GetEmail2(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reEmail)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result)
	}
}

// 抽取根据url获取内容
func GetPageStr(url string) (pageStr string) {
	resp, err := http.Get(url)
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	// 字节转字符串
	pageStr = string(pageBytes)
	return pageStr
}

func Test1() {
	// 2.抽取的爬邮箱
	// GetEmail2("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	// 3.爬链接
	GetLink("https://www.bilibili.com/video/BV1CV4y1u7yu/?spm_id_from=333.1007.tianma.4-1-12.click&vd_source=ab27aee0d4d818c12a20a42576a26b6e")
	// 4.爬手机号
	//GetPhone("https://www.zhaohaowang.com/")
	// 5.爬身份证号
	//GetIdCard("https://www.zhaohaowang.com/")
	// 6.爬图片
	GetImg("https://www.figma.com/file/oxIbFK8l9zH917WZ3Z8BKD/v1.2.0%E4%BC%98%E6%83%A0%E5%88%B8?node-id=0%3A1")
}
func GetIdCard(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reLinke)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result)
	}
}

// 爬链接
func GetLink(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reLinke)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result[1])
	}
}

//爬手机号
func GetPhone(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(rePhone)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result)
	}
}

func GetImg(url string) {
	print("start")
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	print(results)
	for _, result := range results {
		fmt.Println(result[0])
	}
}
