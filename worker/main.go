package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello!")
		time.Sleep(10 * time.Second)
	}
}
