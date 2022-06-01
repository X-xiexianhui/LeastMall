//Package models
/*
   @author:xie
   @date:2022/6/1
   @note:
*/
package models

type Product struct {
	id       int
	name     string
	price    float32
	pictures []string
}

func (p *Product) Create() {
}
