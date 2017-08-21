package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/henrylee2cn/surfer"
)

//redisClient.SetNX(s,ife,)

func cqssc_163_com() { //163的比较慢 质量较差
	resp, err := surfer.Download(&surfer.Request{
		Url: "http://trend.caipiao.163.com/cqssc/jiben-5xing.html?periodNumber=30",
		//DownloaderID: 1,
	})
	if err != nil {
		fmt.Println("67:", err)
		return
	}
	defer resp.Body.Close()
	defer surfer.DestroyJsFiles()
	var b []byte
	b, err = ioutil.ReadAll(resp.Body)
	//log.Println(string(b), err)
	re := regexp.MustCompile(`(<tr data-period=")([0-9]{9})(" data-award=\")([0-9 ]{9})(">)`).FindAllStringSubmatch(string(b), 200)

	if len(re) != 30 || len(re[29]) <= 4 {
		fmt.Println("33 regexp规则错误，或数据错误！")
		fmt.Println(string(b))
		return
	}

	var issue int
	issue, err = strconv.Atoi(re[29][2])
	if err != nil {
		return
	}
	dt := data{Type: 1, Time: time.Now(), Data: re[29][4], Issue: issue}
	dt.dataIn()

}

func cqssc_cqcp_net() {
	resp, err := surfer.Download(&surfer.Request{
		Url: "http://buy.cqcp.net/Game/GetNum.aspx?iType=3&time=Mon%20Aug%2007%202017%2001:49:30%20GMT+0800%20(HKT)",
		//DownloaderID: 1,
	})
	if err != nil {
		fmt.Println("95:", err)
		return
	}

	defer resp.Body.Close()
	defer surfer.DestroyJsFiles()
	var b []byte
	b, err = ioutil.ReadAll(resp.Body)

	//log.Println(string(b), err)
	re := regexp.MustCompile(`([\d]{9})(</li><li class=.?openli2.?>)([,0-9]{9})(</li>)`).FindAllStringSubmatch(string(b), -1)
	//fmt.Println("105:", re[0][3])
	//fmt.Println("106:", re[0])
	if len(re) < 1 || len(re[0]) != 5 {
		fmt.Println("67 regexp规则错误，或数据错误！")
		return
	}

	var issue int
	issue, err = strconv.Atoi(re[0][1])
	if err != nil {
		return
	}
	dt := data{Type: 1, Time: time.Now(), Data: strings.Replace(re[0][3], ",", " ", -1), Issue: issue}
	dt.dataIn()
}
