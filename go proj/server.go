package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

var rds redis.Conn

func RedisPollInit() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     5, //最大空闲数
		MaxActive:   0, //最大连接数，0不设上
		Wait:        true,
		IdleTimeout: time.Duration(1) * time.Second, //空闲等待时间
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379") //redis IP地址
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			redis.DialDatabase(0)
			return c, err
		},
	}
}

func RedisInit() {
	rds = RedisPollInit().Get()
}

func RedisClose() {
	_ = rds.Close()
}

// 私人信息组

// 公共信息组
// public/*
// public/food
func create_food(food map[string]interface{}, ch chan string) {
	// 将json转化为string存储到redis,key为id
	str, cov_err := json.Marshal(food)
	if cov_err != nil {
		fmt.Println(cov_err)
		ch <- cov_err.Error()
		return
	}
	_, set_err := rds.Do("set", food["id"], str)
	if set_err != nil {
		fmt.Println(set_err)
		ch <- cov_err.Error()
		return
	}
	ch <- "ok"
}

// 信息格式:html文件
// 每个信息json确保有title的项，作为标题
// 信息若有img的项则选择第一个作为缩略图显示
// 信息里的img项都要是base64格式或者是在线链接，便于显示
// 每个信息本服务器都会分配一个独有id和过期时间
// 独有id用来定位信息
// 过期时间用以定期清除信息
func main() {
	// 初始化redis
	RedisInit()
	defer RedisClose()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	public := router.Group("/public")
	public.POST("/food", func(c *gin.Context) {
		food := make(map[string]interface{}) //注意该结构接受的内容
		// 转换为json
		c.BindJSON(&food)
		// 进行任务，传回消息
		ch := make(chan string)
		go create_food(food, ch)
		msg := <-ch

		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  msg,
		})
	})
	router.Run(":8080")
}
