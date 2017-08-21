package main

import (
	"fmt"
	"strconv"
	//"strconv"
	"time"

	//"github.com/henrylee2cn/surfer"
	//"github.com/go-pg/pg"
)

func (self *data) dataIn(src string) {
	strType := strconv.Itoa(self.Type)
	for !(redisClient.HSetNX(strType, "client_lock", true).Val()) {
		fmt.Println("dataIn() client_lock ,false,Type=", self.Type)
		time.Sleep(time.Second / 5)
	}
	redisClient.HDel(strType, "client_lock")
	tx, err := Db.Begin()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//var stmt *pg.Stmt
	var b bool
	b, err = tx.Model(&self).Where("type=? and issue=?", self.Type, self.Issue).SelectOrInsert()
	if err != nil {
		fmt.Println("dataIn()96:" + err.Error())
		return
	}

	defer tx.Rollback()
	err = tx.Commit()
	if err != nil {
		fmt.Println("dataIn()103:" + err.Error())
		return
	}
	if b == true {
		fmt.Println("来源：", src)
		self.done()
		//redisClient.Set("caiji_"self.Type+"_"+self.Issue)
	}

	return
}
