package main

import (
	"DouyinSimpleProject/config"
	"DouyinSimpleProject/router"
	"fmt"
)

func main() {
	config.Setup()

	r := router.InitRouter()

	serverAddr := fmt.Sprintf("%s:%s", config.SERVER_HOST, config.SERVER_PORT)
	r.Run(serverAddr)
}
