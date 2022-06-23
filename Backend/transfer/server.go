package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
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
func convert_to_string_arr(ifs interface{}) []string {
	arr := make([]string, 0)
	switch reflect.TypeOf(ifs).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(ifs)
		for i := 0; i < s.Len(); i++ {
			arr = append(arr, fmt.Sprintf("%s", s.Index(i)))
		}
	}
	return arr
}
func create_info(c *gin.Context) {
	id := generate_unique_id()
	// var info Infomation
	info := make(map[string]interface{})
	bind_err := c.ShouldBind(&info)
	if bind_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": bind_err.Error()})
		return
	}
	//
	switch reflect.TypeOf(info["tags"]).Kind() {
	case reflect.Slice:
		tags_array := reflect.ValueOf(info["tags"])
		for i := 0; i < tags_array.Len(); i++ {
			// tag 代表集合
			// 向集合里添加成员
			rds.Do("SADD", tags_array.Index(i), id)
			// 同时也记录下该信息的tags
			rds.Do("SADD", id+"/tags", tags_array.Index(i))
		}
	}
	//
	data, _ := json.Marshal(info["data"])
	rds.Do("SET", id, data)
	//
	if info["expire"] != nil {
		rds.Do("EXPIRE", id, info["expire"])
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}
func update_command() {
	url := "http://127.0.0.1:5000/get" //请求地址
	http.Get(url)
}

// 私人信息组

// 公共信息组

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

	public := router.Group("/public")
	private := router.Group("/private")

	// 全局连接服务器
	go connServer.run()
	router.GET("/init", func(c *gin.Context) {
		update_command()
		c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	})
	private.POST("/", func(c *gin.Context) {
		create_info(c)
	})
	private.GET("id/:id", func(c *gin.Context) {
		id := c.Param("id")
		data, _ := rds.Do("GET", id)
		if data == nil {
			c.JSON(http.StatusNotFound, gin.H{"msg": "id not exists"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("%s", data)})
	})
	private.GET("tag/:tag", func(c *gin.Context) {
		tag := c.Param("tag")
		ids, _ := rds.Do("SMEMBERS", tag)
		id_arr := convert_to_string_arr(ids)

		c.JSON(http.StatusOK, gin.H{"data": id_arr})
	})
	private.GET("tag/union", func(c *gin.Context) {
		tags := c.QueryArray("tag")
		// 无奈之举,golang真丑陋
		tag_arr := make([]interface{}, 0)
		for i := 0; i < len(tags); i++ {
			tag_arr = append(tag_arr, tags[i])
		}
		ids, _ := rds.Do("SUNION", tag_arr...)
		id_arr := convert_to_string_arr(ids)

		c.JSON(http.StatusOK, gin.H{"data": id_arr})
	})
	private.GET("tag/inter", func(c *gin.Context) {
		tags := c.QueryArray("tag")
		// 无奈之举,golang真丑陋
		tag_arr := make([]interface{}, 0)
		for i := 0; i < len(tags); i++ {
			tag_arr = append(tag_arr, tags[i])
		}
		ids, _ := rds.Do("SINTER", tag_arr...)
		id_arr := convert_to_string_arr(ids)
		c.JSON(http.StatusOK, gin.H{"data": id_arr})
	})
	private.POST("login", func(c *gin.Context) {

	})
	private.GET("/websocket", func(c *gin.Context) {
		websocketServer(c.Writer, c.Request)
	})
	public.POST("/", func(c *gin.Context) {
		create_info(c)
	})
	public.GET("id/:id", func(c *gin.Context) {
		id := c.Param("id")
		data, _ := rds.Do("GET", id)
		if data == nil {
			c.JSON(http.StatusNotFound, gin.H{"msg": "id not exists"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("%s", data)})
	})
	public.GET("tag/:tag", func(c *gin.Context) {
		tag := c.Param("tag")
		ids, _ := rds.Do("SMEMBERS", tag)
		id_arr := convert_to_string_arr(ids)

		c.JSON(http.StatusOK, gin.H{"data": id_arr})
	})
	public.GET("tag/union", func(c *gin.Context) {
		tags := c.QueryArray("tag")
		// 无奈之举,golang真丑陋
		tag_arr := make([]interface{}, 0)
		for i := 0; i < len(tags); i++ {
			tag_arr = append(tag_arr, tags[i])
		}
		ids, _ := rds.Do("SUNION", tag_arr...)
		id_arr := convert_to_string_arr(ids)

		c.JSON(http.StatusOK, gin.H{"data": id_arr})
	})
	public.GET("tag/inter", func(c *gin.Context) {
		tags := c.QueryArray("tag")
		// 无奈之举,golang真丑陋
		tag_arr := make([]interface{}, 0)
		for i := 0; i < len(tags); i++ {
			tag_arr = append(tag_arr, tags[i])
		}
		ids, _ := rds.Do("SINTER", tag_arr...)
		id_arr := convert_to_string_arr(ids)
		c.JSON(http.StatusOK, gin.H{"data": id_arr})
	})
	router.Run(":8080")
}

func generate_unique_id() string {
	prefix := time.Now().Unix()
	rand.Seed(time.Now().UnixMicro()) //随机种子
	max := 0xf00000
	min := 0x100000
	body := rand.Intn(max-min) + min
	return strconv.FormatInt(int64(body), 16) + strconv.FormatInt(prefix, 16)
}
