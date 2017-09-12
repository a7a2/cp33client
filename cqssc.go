package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/henrylee2cn/surfer"
)

var fangLouChan = make(chan string, 10000)

func getCqsscAll() { //重庆时时彩2009-12-13至今的数据用作趋势分析
	go func() {
		select {
		case d := <-fangLouChan:
			cqsscAll(d)
		}
	}()
	var dStr string
	tt, _ := time.ParseInLocation("2006-01-02", "2009-12-12", time.Local)
	for {
		dStr = tt.Add(time.Hour * 24).Format("2006-01-02")
		tt, _ = time.ParseInLocation("2006-01-02", dStr, time.Local)
		go cqsscAll(dStr)
		if dStr == "2017-08-21" {
			break
		}
		time.Sleep(time.Microsecond * 300) //任务完成速度跟不上会吃光内存
	}
}

func cqsscAll(d string) {
	u := fmt.Sprintf("http://chart.cp.360.cn/kaijiang/kaijiang?lotId=255401&spanType=2&span=%s_%s", d, d)
	h := map[string][]string{
		"Accept-Language":  []string{"zh-CN,zh;q=0.8,en;q=0.6,th;q=0.4"},
		"Accept":           []string{"application/json, text/javascript, */*; q=0.01"},
		"Connection":       []string{"close"},
		"Content-Type":     []string{"application/x-www-form-urlencoded; charset=UTF-8"},
		"Host":             []string{"127.0.0.1:8080"},
		"Origin":           []string{"http://127.0.0.1:8080"},
		"User-Agent":       []string{"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Mobile Safari/537.36"},
		"X-Requested-With": []string{"XMLHttpRequest"},
	}
	resp, err := surfer.Download(&surfer.Request{
		Header: h,
		Url:    u,
		//DownloaderID: 1,
	})

	if err != nil {
		//fmt.Println("94:", err, "	", u)
		time.Sleep(time.Second)
		fangLouChan <- d
		return
	}

	defer resp.Body.Close()
	defer surfer.DestroyJsFiles()
	var b []byte
	b, err = ioutil.ReadAll(resp.Body)

	re := regexp.MustCompile(`(<tr><td class='gray'>)([0-9]{3})(</td><td class='red big'>)([0-9]{5})(</td><td class='gray'>)`)
	a := re.FindAllStringSubmatch(string(b), -1)
	//fmt.Println(len(a), "	", len(a[0]))

	for i := 0; i < len(a) && len(a[i]) != 5; i++ {
		tt, _ := time.ParseInLocation("2006-01-02", d, time.Local)
		dd := a[i][4][0:1] + " " + a[i][4][1:2] + " " + a[i][4][2:3] + " " + a[i][4][3:4] + " " + a[i][4][4:5]
		strIssue := tt.Format("060102") + a[i][2]
		issue, _ := strconv.Atoi(strIssue)
		dt := data{Type: 1, Time: time.Now(), Data: dd, Issue: issue}
		dt.dataIn("", strIssue)
	}

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
	dt.dataIn("cqssc_cqcp_net", re[0][1])
}

type sJson163 struct {
	AwardNumberInfoList []a163 `json:"awardNumberInfoList"`
	Status              string `json:"status"`
}

type a163 struct {
	DaXiaoBi      string
	FirstXt       string
	GeWei         string
	Hezhi         string
	HouSan        string
	JiOuBi        string
	Period        string
	SecondXt      string
	ShiWei        string
	WinningNumber string
	XingTai       string
}

func cqssc_163_com(period string) {
	u := fmt.Sprintf("http://caipiao.163.com/award/getAwardNumberInfo.html?gameEn=ssc&cache=%v&period=%s", time.Now().UnixNano(), period)
	resp, err := surfer.Download(&surfer.Request{
		Url: u,
		//DownloaderID: 1,
	})

	if err != nil {
		fmt.Println("94:", err)
		return
	}

	defer resp.Body.Close()
	defer surfer.DestroyJsFiles()
	var b []byte
	b, err = ioutil.ReadAll(resp.Body)

	var j sJson163
	err = json.Unmarshal(b, &j)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if len(j.AwardNumberInfoList) > 0 && j.AwardNumberInfoList[0].Period == period {
		issue, _ := strconv.Atoi(j.AwardNumberInfoList[0].Period)
		dt := data{Type: 1, Time: time.Now(), Data: j.AwardNumberInfoList[0].WinningNumber, Issue: issue}
		dt.dataIn("cqssc_163_com", j.AwardNumberInfoList[0].Period)
	}
}
