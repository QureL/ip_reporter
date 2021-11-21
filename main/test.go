package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func showMode(code int) {
	fmt.Println(os.FileMode(code).String())
}
func test01() {
	showMode(0777)
	showMode(0766)
	showMode(0764)
}

func main() {
	data := []byte("hello goo\n")
	i := 0
	for {
		// 覆盖式写入
		time.Sleep(5 * time.Second)
		ioutil.WriteFile("test.txt", data, 0664)
		data = []byte(fmt.Sprintf("%d\n", i))
		log.Println(string(data))
		i++
	}
}
