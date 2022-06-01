//Package models
/*
   @author:xie
   @date:2022/6/1
   @note:
*/
package models

type Banner struct {
	Id        int
	ProductId int
	picture   string
}

func (b *Banner) Create() {
}
