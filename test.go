package main

import (
	"fmt"
	"log"

	uptimerobot "github.com/bitfield/uptimerobot/pkg"
)

func main() {
	fmt.Println("hello world")

	client := uptimerobot.New("u1857235-ec7c1137633fd0592bd3c445")

	monitors, err := client.AllMonitors()
	if err != nil {
		log.Fatal(err)
	}
	for _, m := range monitors {
		fmt.Println(m)
		fmt.Println()
	}

}
