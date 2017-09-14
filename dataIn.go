package main

import (
	"fmt"
	"strconv"
	"time"
)

func (self *data) dataIn(src, issue string) {
	var err error
	strType := strconv.Itoa(self.Type)
	for {
		c := redisClient.SetNX("caiji_"+strType+"_"+issue, "", time.Second)
		if c.Err() != nil {
			fmt.Println("dataIn() client_lock ,false, ", strType+"_"+issue, " ", c.Err().Error())
			time.Sleep(time.Second / 5)
			continue
		} else if c.Val() == false {
			fmt.Println("redisClient.HSetNX boo == false, ", strType+"_"+issue)
			time.Sleep(time.Second / 5)
			continue
		} else {
			break
		}
	}
	defer redisClient.Del("caiji_" + strType + "_" + issue)
	tx, err := Db.Begin()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//var stmt *pg.Stmt
	var b bool
	b, err = tx.Model(self).Where("type=? and issue=?", self.Type, self.Issue).SelectOrInsert()
	if err != nil {
		fmt.Println("dataIn()96:" + err.Error())
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("dataIn()103:" + err.Error())
		return
	}
	if b == true {
		if src != "" {
			fmt.Println("来源：", src)
		}
		self.done()
		//redisClient.Set("caiji_"self.Type+"_"+self.Issue)
	}
	return
}
