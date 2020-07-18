package main

import (
	"cloud/config"
	"cloud/system"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	// defer handle error
	defer handleShutdown()

	// init gin engine
	engine := gin.Default()

	// bootstrap
	system.BootStrap(engine)

	// listen and run server
	err := engine.Run(config.Config().GetServerListenAddr())
	if err != nil {
		fmt.Println("server up error: " + err.Error())
	}
}

func handleShutdown(){
	if e := recover(); e != nil{
		fmt.Println("run error ", e)
	}
}
