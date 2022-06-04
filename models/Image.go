//Package models
/*
   @author:xie
   @date:2022/6/1
   @note:
*/
package models

type Image struct {
	Id        int64  `json:"id" gorm:"primary_key auto increment"`
	ProductId int64  `json:"product_id"`
	Image     string `json:"image"`
}

func (i *Image) Create() {
}
func (i Image) Echo() {
}
