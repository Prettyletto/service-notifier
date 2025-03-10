package main

import "github.com/Prettyletto/service-notifier/cmd/server"

func main() {

	s := server.New(":8080")
	s.Start()
	defer s.Stop()

}
