package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-oauth/config"
	db "go-oauth/data"
	"go-oauth/router"
)

func setConfiguration(configPath string) {
	config.Setup(configPath)
	db.SetupDB()
	gin.SetMode(config.GetConfig().Server.Mode)
}

func main() {
	configPath := "./config.yml"
	setConfiguration(configPath)
	conf := config.GetConfig()
	web := router.Setup()
	fmt.Println("Go API REST Running on port " + conf.Server.Port)
	fmt.Println("==================>")
	_ = web.Run(":" + conf.Server.Port)
}
