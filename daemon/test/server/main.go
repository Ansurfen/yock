package main

import "github.com/ansurfen/yock/daemon/interface/server"

func main() {
	server.Gopt.Parse()
	s := server.New(server.Gopt)
	defer s.Close()
	s.Run()
}
