package main

import (
	"lockfree/cas"
	"time"
)

func main() {
	cas.DoWork()
	for {
		time.Sleep(1 * time.Second)
	}
}
