//Package models
/*
   @author:xie
   @date:2022/6/1
   @note:
*/
package models

type Banner struct {
	Id        int64  `json:"id"`
	ProductId int64  `json:"product_id"`
	Picture   string `json:"picture"`
}

func (b *Banner) Create() {
}
