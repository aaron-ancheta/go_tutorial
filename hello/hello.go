package main

import (
	"fmt"
	"log"

	"github.com/aaron-ancheta/go_tutorial/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Jenica")

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(message)
}