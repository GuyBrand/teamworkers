package basketitem

import "teamworkers/micro/product/proto"

type BasketItem struct {
	ID       string          `json:"id"`
	Product  product.Product `json:"product"`
	Quantity float64         `json:"quantity"`
	Discount float64         `json:"discount,omitempty"`
}

type AddBasketItemRequest struct {
	ProductId string  `json:"productId"`
	Quantity  float64 `json:"quantity"`
}
