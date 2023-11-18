package runeTest

import "fmt"

func Test() {
	s := "pprof.cn博客"
	//runeS := []byte(s)
	runeS := []rune(s)
	for i := 0; i < len(runeS); i++ { //byte
		println(string(runeS[i]))
	}

	//runeS[0] = '狗'
	fmt.Println(string(runeS[8]))
	fmt.Println(len(runeS))
	fmt.Println(string(runeS))
}
