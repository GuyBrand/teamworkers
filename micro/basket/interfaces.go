package main

//mmm...autonomous vs  library coupling ("dont share model")
import (
	"teamworkers/micro/crud"
	pb "teamworkers/micro/product/proto"
)

func GetProductFromCatalog(ProductId string) (*pb.Product, error) {
	p := pb.Product{}
	if err := crud.GetEntityFromService("product", ProductId, &p); err != nil {
		return nil, err
	} else {
		return &p, nil
	}
}
