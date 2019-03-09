package product

import pb "teamworkers/workers/product/lib/proto"

var products productsDb //which mean we're not really stateless on this demo !!! dont try running multiple instances

type product pb.Product
type productsDb []product

func init() {
	products = productsDb{product{ID: "1A", Description: "Burger", Price: 12.34},
		product{ID: "2B", Description: "Fries", Price: 5.67}}
}

func GetByIdInterface(id string) interface{} {
	return (*pb.Product)(GetById(id))
}

func GetById(id string) *product {
	for _, p := range products {
		if p.ID == id {
			return &p
		}
	}
	return nil
}

func Add(p product) {
	products = append(products, p)
}

func Delete(id string) bool {
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...) //TODO : lock
			return true
		}
	}
	return false
}

func Update(updated product) bool {
	for i, p := range products {
		if p.ID == updated.ID {
			products[i] = updated //TODO : lock
			return true
		}
	}
	return false
}
