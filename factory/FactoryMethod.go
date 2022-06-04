//Package factory
/*
   @author:xie
   @date:2022/6/4
   @note:
*/
package factory

import "leastMall_gin/models"

var ImageFactoryObject = ImageFactory{}
var GoodFactoryObject = GoodFactory{}
var ResponseFactoryObject = ResponseFactory{}

type TotalFactory interface {
	CreateProduct() Product
}
type Product interface {
	Echo()
}
type ImageFactory struct {
}

func (i ImageFactory) CreateProduct() Product {
	return models.Image{}
}

type GoodFactory struct {
}

func (g *GoodFactory) CreateProduct() Product {
	return models.Product{}
}

type ResponseFactory struct {
}

func (r *ResponseFactory) CreateProduct() Product {
	return models.Response{}
}
