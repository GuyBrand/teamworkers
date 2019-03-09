package main

import (
	"teamworkers/workers/basket/lib"
	"teamworkers/workers/product/lib"
)

func main() {
	product.Start(false)
	basket.Start(true)
}
