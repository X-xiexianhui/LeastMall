//Package models
/*
   @author:xie
   @date:2022/6/1
   @note:
*/
package models

type Product struct {
	ProductId    int
	ProductName  string
	Price        float32
	Descriptions string
	Pictures     []string
}

func (p *Product) Create() {
}
