package main

import (
	"log"
	"takoyaki/web"
)

func main() {
	server := web.NewServer(":8081", 100)
	err := server.Run()
	if err != nil {
		log.Printf("server.Run err:%v", err)
	}
}
