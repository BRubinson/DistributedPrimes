package main

import (
	"fmt"
	"time"
)

func MasterDisplay() {
	go func() {
		for {
			CallClear()
			fmt.Println("Master Node:")
			fmt.Println("current num:", NextNumber)
			fmt.Println("Primes Found:", Primes)
			fmt.Println("Recent API calls(resets every 5 seconds):")
			time.Sleep(5 * time.Second)
		}
	}()
}
