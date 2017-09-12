package main

import (
	"fmt"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-redis/redis"
)

func redisInit() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:        "localhost:6379",
		Password:    "", // no password set
		DB:          1,  // use default DB
		PoolSize:    999,
		DialTimeout: time.Second * 3,
	})
	pong, err := redisClient.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

var (
	Db *pg.DB
)

func Database() *pg.DB {
	return Db
}

func createSchema(db *pg.DB) error {
	return nil
}

func dbInit() {
	Db = pg.Connect(&pg.Options{
		Network:            "tcp",
		Addr:               fmt.Sprintf("%s:%s", "127.0.0.1", "5432"),
		User:               "root",
		Password:           "root",
		Database:           "cp33",
		DialTimeout:        3 * time.Second,
		ReadTimeout:        3 * time.Second,
		WriteTimeout:       3 * time.Second,
		PoolSize:           20, //postgresql默认是100的,多了无意义对于插入来说，同时看开了多少个采集端毕竟pg默认连接数才100
		PoolTimeout:        time.Second * 3,
		IdleTimeout:        time.Second * 10,
		IdleCheckFrequency: time.Second,
	})

	err := createSchema(Db)
	if err != nil {
		fmt.Println(err.Error())
		t, _ := time.ParseDuration("5s")
		time.Sleep(t)
		dbInit()
	}
}

func init() {
	redisInit()
	dbInit()
	go loopDoneNotice()
	//go getCqsscAll()
}
