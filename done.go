package main

import (
	"fmt"
	"time"

	//"regexp"
	"strconv"

	"github.com/henrylee2cn/surfer"
)

func (self *data) done() {
	resp, err := surfer.Download(&surfer.Request{
		Url: "http://127.0.0.1:8080/dataInNotice/" + strconv.Itoa(self.Type) + "/" + strconv.Itoa(self.Issue),
		//DownloaderID: 1,
	})
	if err != nil {
		fmt.Println("done() 22:", err)
		return
	}

	defer resp.Body.Close()
	defer surfer.DestroyJsFiles()
	redisClient.Set("Client_"+strconv.Itoa(self.Type)+"_"+strconv.Itoa(self.Issue), self.Data, time.Minute*30)
	fmt.Println(time.Now(), " 成功入库：", "	", self.Type, self.Issue, "		", self.Data)
}
