package greetings_test

import (
	"fmt"
	httpUtil "github.com/kirahuang/test-go/src/http"
)

//import (
//	"fmt"
//	httpUtil "github.com/kirahuang/test-go/src/http"
//	greetings "github.com/kirahuang/test-go/src/test"
//)

func main_1() {

	// Get a greeting message and print it.
	//message := greetings.Hello("Gladys")
	message := Hello("Gladys")
	fmt.Println(message)



	bytes:=httpUtil.HttpRequest("https://www.baidu.com")
	str := httpUtil.String(bytes)
	//bytes[0] = 'i'//注意这一行，bytes在这里修改了数据，但是str打印出来的依然没变化，
	fmt.Println(str)
}