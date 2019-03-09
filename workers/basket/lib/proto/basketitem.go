package basketitem

import "teamworkers/workers/product/lib/proto"

type BasketItem struct {
	ID       string          `json:"id"`
	Product  product.Product `json:"product"`
	Qauntity float64         `json:"qauntity"`
	Discount float64         `json:"discount,omitempty"`
}

type AddBasketItemRequest struct {
	ProductId string  `json:"productId"`
	Quantity  float64 `json:"quantity"`
}
