package main

import (
	"fmt"

	"github.com/kirahuang/test-go/src/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
