package main

import (
	"fmt"
	"time"
)

var ch = make(chan string, 10)

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second) // 模拟耗时操作
	ch <- url
}

func main()  {
	var loop int = 5
	for i:= 0; i < loop; i++ {
		go download("https://baidu.com/"+string(i+'0'))
	}
	for i:=0;i<loop;i++{
		msg := <- ch
		fmt.Println("finish",msg)
	}
	fmt.Println("Done!")
}