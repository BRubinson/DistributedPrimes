package main

import (
	"fmt"
	"time"
)

func OverlordDisplay() {
	go func() {
		for {
			CallClear()
			fmt.Println("OVERLORD REEEE")
			fmt.Println("current num:", Leader.CurrentCount)
			fmt.Println("Current Master Node", Leader.MasterNode.Ip)
			fmt.Println("Primes Found:", Leader.Primes)
			fmt.Println("Laster master Ping:", Leader.LastMasterPing)
			fmt.Println("Recent API calls(resets every 5 seconds):")
			time.Sleep(5 * time.Second)
		}
	}()
}
