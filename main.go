package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"PDSgroupon/config"
	"PDSgroupon/model"
	"PDSgroupon/router"
)

var (
	cfg = pflag.StringP("config", "c", "", "groupon config file path.")
)

func main() {
	// 解析命令行参数
	pflag.Parse()

	// 初始化配置
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// init db
	model.DB.Init()
	defer model.DB.Close()

	// init redis
	err := model.RC.Init()
	if err != nil {
		panic(err)
	}
	defer model.RC.Close()

	// viper热更新配置文件
	//for {
	//	fmt.Println(viper.GetString("runmode"))
	//	time.Sleep(4*time.Second)
	//}

	// 设置gin模式
	gin.SetMode(viper.GetString("runmode"))

	// 创建引擎
	g := gin.New()

	// gin 中间件
	middlewares := []gin.HandlerFunc{}

	// 路由加载
	router.Load(
		// gin 核心引擎
		g,
		// 中间件列表加载
		middlewares...,
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
