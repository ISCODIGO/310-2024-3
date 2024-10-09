package main

import (
	"fmt"
	"time"
)

func main() {
	for range time.Tick(time.Second * 3) {
		fmt.Println(time.Now())
	}
}