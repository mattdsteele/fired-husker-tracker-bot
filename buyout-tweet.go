package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Hello %s", "world!\n")
	now := time.Now()
	fmt.Println(now.Hour())
	fmt.Printf(now.String())
}
