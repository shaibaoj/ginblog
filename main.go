package main

import (
	"fmt"
	"os"

	"github.com/thhy/ginblog/conf"
	"github.com/thhy/ginblog/router"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	//change path to word dir
	os.Chdir(conf.SRCPATH)

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	os.Chdir(conf.SRCPATH)
	router.InitializeRoutes(server)

	server.Run(conf.BINDADDRESS)
}
