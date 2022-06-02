//Package models
/*
   @author:xie
   @date:2022/6/1
   @note:
*/
package models

type Product struct {
	ProductId    int     `json:"product_id"`
	ProductName  string  `json:"product_name"`
	Price        float64 `json:"price"`
	Descriptions string  `json:"descriptions"`
	Cover        string  `json:"cover"`
}

func (p *Product) Create() {
}
