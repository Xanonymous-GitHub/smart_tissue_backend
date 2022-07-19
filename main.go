package main

import (
	"backend/model"
	"backend/server"
)

func main() {
	model.Setup()
	server.SetupRouter()
}
