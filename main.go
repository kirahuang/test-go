package main

import (
	"fmt"
	greetings "github.com/kirahuang/test-go/src/test"

)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
