//Package backendFactory
/*
   @author:xie
   @date:2022/6/1
   @note:
*/
package factory

import "leastMall_gin/models"

func SimpleFactory(object string) models.ProductObject {
	if object == "product" {
		return &models.Product{}
	} else if object == "image" {
		return &models.Image{}
	}
	return nil
}
