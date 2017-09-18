package main

import (
	"time"
)

func main() {
	for {
		select {
		case <-time.After(time.Second * 12):
			go getCron("1")
			go getCron("7")
			go getCron("9")
		}
	}

}
