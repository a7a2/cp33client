package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
	//	"reflect"
	"regexp"

	"github.com/henrylee2cn/surfer"
)

func pk10_bwlc_net(period *string) {
	u := fmt.Sprintf("http://www.bwlc.net/")
	resp, err := surfer.Download(&surfer.Request{
		Url: u,
		//DownloaderID: 1,
	})
	if err != nil {
		fmt.Println("pk10_bwlc_net 22:", err)
		return
	}

	defer resp.Body.Close()
	defer surfer.DestroyJsFiles()
	var b []byte
	b, err = ioutil.ReadAll(resp.Body)
	if len(b) < 2200 {
		return
	}
	re := regexp.MustCompile(`(<span class="ml10 b red fa f14">)([0-9]{6})(</span>)([.\S\s]+?)(<ul class="dib">)([.\S\s]*?)(</ul>)`).FindAllStringSubmatch(string(b[21000:]), -1)

	if len(re) != 1 && len(re[0]) != 7 {
		return
	}

	if *period != re[0][2] { //期号
		return
	}
	reData := regexp.MustCompile(`[0-9]{2}`).FindAllString(re[0][6], -1)
	if len(reData) != 10 {
		return
	}
	var strData string //开奖号码
	for i := 0; i < len(reData); i++ {
		strData = fmt.Sprintf("%s %s", strData, reData[i])
	}
	strData = strData[1:] //去掉前面的空格
	var issue int
	issue, err = strconv.Atoi(*period)
	if err != nil {
		return
	}
	dt := data{Type: 9, Time: time.Now(), Data: strData, Issue: issue}
	dt.dataIn("pk10_bwlc_net", strData)
}
