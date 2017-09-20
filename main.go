package main

import (
	"time"
)

func main() {
	for {
		select {
		case <-time.After(time.Second * 12):
			go getCron("1") //重庆时时彩
			go getCron("7") //西藏时时彩
			go getCron("9") //北京pk10
			go getCron("4") //天津时时彩
		}
	}

}
