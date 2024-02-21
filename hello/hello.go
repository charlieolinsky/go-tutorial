package main

import (
	"fmt"

	"github.com/charlieolinsky/go-project/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Charlie")
	fmt.Println(message)
}
