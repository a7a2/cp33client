package main

import (
	"reflect"
	"unsafe"

	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"

	"time"

	"github.com/henrylee2cn/surfer"
)

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
	re := regexp.MustCompile(`(20)([0-9]{8})([\S\s]*)(<td colspan=.?3.?>)([0-9\ ]{9})(</td>)`).FindAllStringSubmatch(string(b), -1)
	if len(re) != 1 || len(re[0]) != 7 {
		return
	}
	var issue int
	issue, err = strconv.Atoi(*period)
	if err != nil {
		return
	}
	dt := data{Type: 7, Time: time.Now(), Data: re[0][5], Issue: issue}
	dt.dataIn("xzssc_xjflcp_com", re[0][5])
}
