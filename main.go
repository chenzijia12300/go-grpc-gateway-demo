package main

import (
	"grpc-demo/cmd"
	"grpc-demo/core"
)

func main() {
	core.LoadConfig()
	core.InitDB(true)
	cmd.Execute()
}
