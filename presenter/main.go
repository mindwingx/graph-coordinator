package main

import "github.com/mindwingx/graph-coordinator/bootstrap"

func main() {
	service := bootstrap.NewApp()
	service.Init()
	service.Start()
}
