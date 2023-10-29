package main

import (
	"fmt"
	"log"

	"github.com/shan-xi/golang_learn/greetings"
)

func main() {
	// fmt.Println(quote.Go())
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// Get a greeting message and print it.
	// message, err := greetings.Hello("Spin")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(message)

	names := []string{"Spin", "SPIN", "spin"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
