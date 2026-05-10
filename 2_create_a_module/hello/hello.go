package main

import (
	"fmt"

	"github.com/jackson-perry-1/greetings"
)

func main() {
	message := greetings.Hello("Jackson")
	fmt.Println(message)
}
