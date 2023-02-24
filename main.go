package main

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack/v3"
	"github.com/wow-unbelievable/blog/global"
	"github.com/wow-unbelievable/blog/internal/model"
	"github.com/wow-unbelievable/blog/internal/routers"
	"github.com/wow-unbelievable/blog/pkg/logger"
	"github.com/wow-unbelievable/blog/pkg/setting"
	"github.com/wow-unbelievable/blog/pkg/tracer"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
}

// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService https://github.com/go-programming-tour-book
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
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
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Tracer", &global.TracerSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	global.AppSetting.Timeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	l, _ := lumberjack.NewRoller(global.AppSetting.LogSavePath+"/"+global.AppSetting.LogFileName+global.AppSetting.LogFileExt, 600*1024*1024, &lumberjack.Options{MaxAge: 10, LocalTime: true})
	global.Logger = logger.NewLogger(l, "", log.LstdFlags).WithCaller(2)
	return nil
}

func setupTracer() error {

	trace, _, err := tracer.NewJaegerTracer(
		global.TracerSetting.ServiceName,
		global.TracerSetting.AgentUrl,
	)
	if err != nil {
		return err
	}
	global.Tracer = trace
	return nil
}
