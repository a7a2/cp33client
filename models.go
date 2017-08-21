package main

import (
	"net/url"
	"time"

	"github.com/go-redis/redis"
)

var (
	cqsscId     = 1
	redisClient *redis.Client
	ChanArray   = make([]chan int, 100)
)

type Result struct { //全站通用json返回结果
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type OpenInfo struct { //输出当前期号、开奖信息等
	Type                  int    `json:"type"`        //上期期号
	Last_period           int    `json:"last_period"` //上期期号
	Last_open             string `json:"last_open"`   //上期开奖号码
	Current_period        int    `json:"current_period"`
	Current_period_status string `json:"current_period_status"`
	Timeleft              int64  `json:"timeleft"`
}

type data struct {
	Id    int
	Type  int
	Time  time.Time
	Data  string
	Issue int
}

type dataTime struct {
	Id         int
	Type       int
	ActionNo   int
	ActionTime string
	StopTime   string
}

type httpReq struct {
	method          string
	url             url.URL
	referer         string
	httpProxyStatus bool
	httpProxy       httpProxy
}

type httpProxy struct {
	user     string
	password string
	host     string
	port     string
}
