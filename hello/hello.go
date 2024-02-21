package main

import (
	"fmt"
	"log"

	"github.com/charlieolinsky/go-project/greetings"
)

func main() {
	/* Set properties of the predefined Logger */
	log.SetPrefix("greetings: ") //Add a suffix to the logger
	log.SetFlags(0)              //Disable default printing of time, source file, and line number.

	// Request a greeting message.
	message, err := greetings.Hello("Charlie")

	// If an error was returned, print it to the console and exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message to the console.
	fmt.Println(message)
}
