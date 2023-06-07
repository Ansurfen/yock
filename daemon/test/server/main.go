package main

import "github.com/ansurfen/yock/daemon/server"

func main() {
	server.Gopt.Parse()
	s := server.New(server.Gopt)
	defer s.Close()
	s.Run()
}
