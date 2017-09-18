package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	//	"time"

	"github.com/henrylee2cn/surfer"
)

func getCron(gameId string) {
	var err error
	var resp *http.Response
	resp, err = surfer.Download(&surfer.Request{
		Url: "http://127.0.0.1:8080/apiMyself/" + gameId,
		//DownloaderID: 1,
	})
	if err != nil {
		fmt.Println("getCron()43:", err.Error())
		return
	}
	defer resp.Body.Close()
	var b []byte
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("getCron()50:", err.Error())
		return
	}

	var openInfo OpenInfo
	err = json.Unmarshal(b, &openInfo)
	if err != nil {
		fmt.Println("getCron()61:", err.Error())
		return
	}
	if openInfo.Last_open != "" {
		return
	}
	//fmt.Println("getCron()64:", openInfo.Last_period, "	", diffTime)
	openInfo.checkIsGot(strconv.Itoa(openInfo.Last_period))
}

func (self *OpenInfo) checkIsGot(period string) {
	if redisClient.Exists("Client_"+strconv.Itoa(self.Type)+"_"+strconv.Itoa(self.Last_period)).Val() == 0 { //不存在
		switch self.Type {
		case 1:
			//fmt.Println("checkIsGot():1		", time.Now(), "		period=", period)
			go cqssc_163_com(period)
			go cqssc_cqcp_net()
		case 7:
			go xzssc_xjflcp_com(&period)
		case 9:
			go pk10_bwlc_net(&period)
		default:
			fmt.Println("checkIsGot():default")
		}
	}
}
