package main

import (
	"reflect"
	"unsafe"
	//	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	//	"strings"
	"time"

	"github.com/henrylee2cn/surfer"
)

//func getXqsscAll() { //重庆时时彩2009-12-13至今的数据用作趋势分析
//	go func() {
//		select {
//		case d := <-fangLouChan:
//			cqsscAll(d)
//		}
//	}()
//	var dStr string
//	tt, _ := time.ParseInLocation("2006-01-02", "2009-12-12", time.Local)
//	for {
//		dStr = tt.Add(time.Hour * 24).Format("2006-01-02")
//		tt, _ = time.ParseInLocation("2006-01-02", dStr, time.Local)
//		go cqsscAll(dStr)
//		if dStr == "2017-08-21" {
//			break
//		}
//		time.Sleep(time.Microsecond * 300) //任务完成速度跟不上会吃光内存
//	}
//}

//func xzsscAll(d string) {
//	u := fmt.Sprintf("http://chart.cp.360.cn/kaijiang/kaijiang?lotId=255401&spanType=2&span=%s_%s", d, d)
//	h := map[string][]string{
//		"Accept-Language":  []string{"zh-CN,zh;q=0.8,en;q=0.6,th;q=0.4"},
//		"Accept":           []string{"application/json, text/javascript, */*; q=0.01"},
//		"Connection":       []string{"close"},
//		"Content-Type":     []string{"application/x-www-form-urlencoded; charset=UTF-8"},
//		"Host":             []string{"127.0.0.1:8080"},
//		"Origin":           []string{"http://127.0.0.1:8080"},
//		"User-Agent":       []string{"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Mobile Safari/537.36"},
//		"X-Requested-With": []string{"XMLHttpRequest"},
//	}
//	resp, err := surfer.Download(&surfer.Request{
//		Header: h,
//		Url:    u,
//		//DownloaderID: 1,
//	})

//	if err != nil {
//		//fmt.Println("94:", err, "	", u)
//		time.Sleep(time.Second)
//		fangLouChan <- d
//		return
//	}

//	defer resp.Body.Close()
//	defer surfer.DestroyJsFiles()
//	var b []byte
//	b, err = ioutil.ReadAll(resp.Body)
//	re := regexp.MustCompile(`(<tr><td class='gray'>)([0-9]{3})(</td><td class='red big'>)([0-9]{5})(</td><td class='gray'>)`)
//	a := re.FindAllStringSubmatch(string(b), -1)

//	for i := 0; i < len(a) && len(a[i]) != 5; i++ {
//		tt, _ := time.ParseInLocation("2006-01-02", d, time.Local)
//		dd := a[i][4][0:1] + " " + a[i][4][1:2] + " " + a[i][4][2:3] + " " + a[i][4][3:4] + " " + a[i][4][4:5]
//		strIssue := tt.Format("060102") + a[i][2]
//		issue, _ := strconv.Atoi(strIssue)
//		dt := data{Type: 1, Time: time.Now(), Data: dd, Issue: issue}
//		dt.dataIn("", strIssue)
//	}

//}
func Slice(s string) (b []byte) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pbytes.Data = pstring.Data
	pbytes.Len = pstring.Len
	pbytes.Cap = pstring.Len
	return
}

func xzssc_xjflcp_com(period *string) {
	var strPeriod string
	slicePeriod := string(Slice(*period))
	if len(slicePeriod) == 9 && slicePeriod[6:7] == "0" {
		strPeriod = fmt.Sprintf("%s%s", slicePeriod[0:6], slicePeriod[7:9])
	} else {
		strPeriod = slicePeriod
	}

	u := fmt.Sprintf("http://www.xjflcp.com/game/sscOpenDetail?gameId=7&lotteryIssue=20%s", strPeriod)
	resp, err := surfer.Download(&surfer.Request{
		Url: u,
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
	//fmt.Println(string(b))
	//log.Println(string(b), err)
	re := regexp.MustCompile(`(20)([0-9]{8})([\S\s]*)(<td colspan=.?3.?>)([0-9\ ]{9})(</td>)`).FindAllStringSubmatch(string(b), -1)
	if len(re) != 1 || len(re[0]) != 7 {
		return
	}
	fmt.Println(*period, "	", len(re), " ", re[0][2], " ", re[0][5])
	var issue int
	issue, err = strconv.Atoi(*period)
	if err != nil {
		return
	}
	dt := data{Type: 7, Time: time.Now(), Data: re[0][5], Issue: issue}
	dt.dataIn("xzssc_xjflcp_com", re[0][5])
}
