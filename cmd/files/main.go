package main

import (
	// "github.com/mijgona/wallet/pkg/wallet"
	//"github.com/mijgona/wallet/pkg/types"
	"log"
	"time"
)


func main() {
	log.Print("main started")

	go func ()  {
		log.Print("goroutines")
	}()

	log.Print("main finished")

	time.Sleep(time.Second *5)
}