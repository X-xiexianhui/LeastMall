//Package models
/*
   @author:xie
   @date:2022/6/1
   @note:
*/
package models

type Product struct {
	Id           int64   `json:"id" gorm:"primary_key auto increment"`
	ProductName  string  `json:"product_name"`
	Price        float64 `json:"price"`
	Descriptions string  `json:"descriptions"`
	Cover        string  `json:"cover"`
}

func (p *Product) Create() {
}
func (p Product) Echo() {
}
