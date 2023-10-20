package main

import (
	"fmt"
	"github/KronosSolucionesTIC/greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings")
	log.SetFlags(0)

	names := []string{"Diego", "Alejandro", "Natalia"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
