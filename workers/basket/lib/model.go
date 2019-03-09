package basket

import (
	"teamworkers/workers/basket/lib/proto"
	"teamworkers/workers/crud"
)

type item basketitem.BasketItem
type Basket []item

var basket Basket //which mean we're not really stateless on this demo !!! dont try running multiple instances

func GetByIdInterface(id string) interface{} {
	return GetById(id)
}

func GetById(id string) *item {
	for _, b := range basket {
		if b.ID == id {
			return &b
		}
	}
	return nil
}

func Add(a basketitem.AddBasketItemRequest) bool {
	if a.ProductId == "" || a.Quantity == 0 {
		return false
	} else if p, err := GetProductFromCatalog(a.ProductId); err != nil {
		//fmt.Println("get product err ", err.Error())
		return false
	} else {
		basket = append(basket, item{Product: *p, ID: crud.GetUUID(), Qauntity: a.Quantity})
		return true
	}

}

func Delete(id string) bool {
	for i, p := range basket {
		if p.ID == id {
			basket = append(basket[:i], basket[i+1:]...) //TODO : lock
			return true
		}
	}
	return false
}

func Update(updated item) bool {
	for i, p := range basket {
		if p.ID == updated.ID {
			basket[i] = updated //TODO : lock
			return true
		}
	}
	return false
}
