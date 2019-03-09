package main

import (
	"teamworkers/micro/crud"
)

func main() {
	Start()
}

func Start() {
	crud.InitRouter("product", createH, readH, updateH, deleteH)
}
