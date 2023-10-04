package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Container is running...")

		time.Sleep(5 * time.Second) 
	}
}
