package main

import "teamworkers/micro/crud"

func main() {
	Start()
}
func Start() {
	crud.InitRouter("basket", createH, readH, updateH, deleteH) //not really crudable
}
