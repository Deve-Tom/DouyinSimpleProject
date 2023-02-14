package main

import (
	"DouyinSimpleProject/config"
	"DouyinSimpleProject/router"
	"fmt"
	"log"
)

func main() {
	config.Setup()

	r := router.InitRouter()

	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	serverAddr := fmt.Sprintf("%s:%s", config.SERVER_HOST, config.SERVER_PORT)
	r.Run(serverAddr)
}
