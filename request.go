package main

import (
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

func (self *httpReq) Gethtml() (*http.Response, error) {
	var client *http.Client
	var request *http.Request
	var err error
	var gCookieJar, _ = cookiejar.New(nil)

	request, err = http.NewRequest("GET", self.url.Scheme+"://"+self.url.Host, strings.NewReader(self.url.Path))
	if err != nil {
		return nil, err
	}
	if strings.ToLower(self.method) == "post" {
		request, err = http.NewRequest(self.method, self.url.Host, strings.NewReader(self.url.Path))
		if err != nil {
			return nil, err
		}
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	request.Header.Set("Accept-Encoding", "*")
	request.Header.Set("Host", "*")
	request.Header.Set("Accept", "*/*")
	request.Header.Set("Connection", "close")
	if self.referer != "" {
		request.Header.Set("Referer", self.referer)
	}
	request.Header.Set("Cache-Control", "no-cache")
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	if self.httpProxyStatus == true {
		proxy, err := url.Parse(self.url.Scheme + "://" + self.httpProxy.user + ":" + self.httpProxy.password + "@" + self.httpProxy.host + ":" + self.httpProxy.port)
		if err != nil {
			return nil, err
		}
		client = &http.Client{
			CheckRedirect: nil,
			Jar:           gCookieJar,
			Transport: &http.Transport{
				Dial: func(netw, addr string) (net.Conn, error) {
					c, err := net.DialTimeout(netw, addr, time.Second*2)
					if err != nil {
						return nil, err
					}
					return c, nil
				},
				Proxy:                 http.ProxyURL(proxy),
				MaxIdleConnsPerHost:   4,
				ResponseHeaderTimeout: time.Second * 3,
			},
		}
	} else {
		client = &http.Client{
			CheckRedirect: nil,
			Jar:           gCookieJar,
			Transport: &http.Transport{
				Dial: func(netw, addr string) (net.Conn, error) {
					c, err := net.DialTimeout(netw, addr, time.Second*2)
					if err != nil {
						return nil, err
					}
					return c, nil
				},
				MaxIdleConnsPerHost:   4,
				ResponseHeaderTimeout: time.Second * 3,
			},
		}
	}
	return client.Do(request)
}
