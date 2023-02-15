package main

import (
	"DouyinSimpleProject/config"
	"DouyinSimpleProject/router"
	"log"
)

func main() {
	config.Setup()

	r := router.InitRouter()

	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	r.Run()
}
