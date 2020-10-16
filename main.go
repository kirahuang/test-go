package main

import (
	"fmt"
	greetings "github.com/kirahuang/test-go/src/test"
	httpUtil "github.com/kirahuang/test-go/src/http"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)


		bytes:=httpRequest("https://www.baidu.com")
		str := String(bytes)
		//bytes[0] = 'i'//注意这一行，bytes在这里修改了数据，但是str打印出来的依然没变化，
		fmt.Println(str)
}
