//package main

//import (
//	"fmt"
//	"io/ioutil"
//	"net/url"
//	"regexp"
//)

//func main() {
//	h := httpReq{
//		method:          "GET",
//		url:             url.URL{Scheme: "http", Host: "trend.caipiao.163.com", RawPath: "/cqssc/", ForceQuery: false},
//		referer:         "",
//		httpProxyStatus: false,
//		httpProxy:       httpProxy{},
//	}
//	r, err := h.Gethtml()
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	var bytesBody []byte
//	bytesBody, err = ioutil.ReadAll(r.Body)
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	strBody := string(bytesBody)
//	fmt.Println(strBody[200:])
//	re := regexp.MustCompile(`(cpdata)`).FindStringSubmatch(strBody)
//	fmt.Println(re[0])
//}

package main

import (
	"time"
)

func main() {
	//fmt.Println(time.Now().Format("15:04:05"))
	for {
		select {
		case <-time.After(time.Second * 10):
			//go getCron(1)
			go getCron("7")
			//fmt.Println(openInfo.Last_period)
		}
	}

}
