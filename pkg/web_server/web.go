package web_server

import (
	"context"
	"log"
	"net/http"
	"opencodes/pkg/config"
	"opencodes/router"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

// 初始化web 服务器
func InitWebServer() {
	//gin.SetMode(gin.ReleaseMode)
	r := router.InitRouter()
	r.Use(gin.Recovery())
	//r.Use()
	srv := &http.Server{
		Addr:    getAddr(),
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Println("Server Run http://127.0.0.1:" + config.ApplicationConfig.Port + "/")
	log.Println("Enter Control + C Shutdown Server")

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}

//获取Addr
func getAddr() string {
	return config.ApplicationConfig.Host + ":" + config.ApplicationConfig.Port
}
