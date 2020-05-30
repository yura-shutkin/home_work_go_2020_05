package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func HelloNow() {
	timeNow := time.Now().Round(time.Second)
	fmt.Printf("current time: %v\n", timeNow)

	exactTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		log.Fatalf("An error occurred. Error is: %v", err)
	}

	fmt.Printf("exact time: %v\n", exactTime.Round(time.Second))
}

func main() {
	HelloNow()
}
