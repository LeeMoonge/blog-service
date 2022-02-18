package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go/blog-service/global"
	"github.com/go/blog-service/internal/routers"
	"github.com/go/blog-service/pkg/setting"
	"log"
	"net/http"
	"time"
)

// 初始化配置读取：在这里调用了初始化配置的方法，起到把配置文件内容映射到应用配置结构体中的作用
func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

func main() {
	// r := gin.Default() 使用新的路由组代替
	//
	//
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"msg": "pong"})
	//})
	//
	//_ = r.Run(":8080")

	gin.SetMode(global.ServerSetting.RunMode)

	router := routers.NewRouter()

	s := &http.Server{
		Addr:              ":" + global.ServerSetting.HttpPort,
		Handler:           router,
		ReadHeaderTimeout: global.ServerSetting.ReadTimeout,
		WriteTimeout:      global.ServerSetting.WriteTimeout,
		MaxHeaderBytes:    1 << 20,
	}

	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}
