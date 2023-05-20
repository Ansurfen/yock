package main

import (
	"fmt"

	"github.com/ansurfen/yock/daemon/interface/client"
)

func main() {
	client.Gopt.Parse()
	client := client.New(client.Gopt)
	defer client.Close()
	if err := client.Ping(); err != nil {
		panic(err)
	}
	info, err := client.Info()
	if err != nil {
		panic(err)
	}
	fmt.Println(info)
}
