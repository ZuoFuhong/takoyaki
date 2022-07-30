package main

import (
	"log"
	"takoyaki/web"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
	s := web.NewServer()
	s.Register()
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
