//Package factory
/*
   @author:xie
   @date:2022/6/4
   @note:
*/
package factory

import "leastMall_gin/models"

var GoodAbstractFactory = GoodFactory{}
var ImageAbstractFactory = IFactory{}

type AbstractProduct interface {
	SaveProduct()
}

// AbstractFactory 抽象工厂类--声明一个创建抽象产品对象的操作接口
type AbstractFactory interface {
	CreateAbstractProduct() AbstractProduct
}

type IFactory struct {
}

func (image IFactory) CreateAbstractProduct() AbstractProduct {
	return models.Image{}
}

type GFactory struct {
}

func (good GFactory) CreateAbstractProduct() AbstractProduct {
	return models.Product{}
}
