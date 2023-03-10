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
	log.SetPrefix(">>> ")

	serverAddr := fmt.Sprintf(":%s", config.SERVER_PORT)
	r.Run(serverAddr)
}
