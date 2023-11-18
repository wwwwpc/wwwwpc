package httptest

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetTest() {
	resp, err := http.Get("https://api.digifinextest.com/need_timestamp")
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read from resp.Body failed,err:", err)
		return
	}
	fmt.Print(1232131321)
	fmt.Print(body)
}
