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

func tjssc_icaile_com(gameType int, period *string) {
	u := fmt.Sprintf("http://tjssc.icaile.com/?op=jb3&num=30")
	resp, err := surfer.Download(&surfer.Request{
		Url: u,
		//DownloaderID: 1,
	})
	if err != nil {
		fmt.Println("tjssc_icaile_com 21:", err)
		return
	}

	defer resp.Body.Close()
	defer surfer.DestroyJsFiles()
	var b []byte
	b, err = ioutil.ReadAll(resp.Body)
	re := regexp.MustCompile(`(20)([0-9]{9})(</td>)([\S\s]+?)(<td>)([,0-9]{9})(</td>)`).FindAllStringSubmatch(string(b), -1)

	c := len(re) - 1
	if len(re) == 0 || len(re) < c || len(re[c]) != 8 {
		return
	}

	if *period != re[c][2] { //期号
		return
	}
	strData := strings.Replace(re[c][6], ",", " ", -1)
	var issue int
	issue, err = strconv.Atoi(re[c][2])
	if err != nil {
		return
	}

	dt := data{Type: gameType, Time: time.Now(), Data: strData, Issue: issue}
	dt.dataIn("tjssc_icaile_com", strData)
}
