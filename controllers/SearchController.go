//Package controllers
/*
   @author:xie
   @date:2022/6/3
   @note:
*/
package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"leastMall_gin/conn"
	"leastMall_gin/models"
	"log"
	"strconv"
)

func init() {
	exists, err := conn.EsClient.IndexExists("product").Do(context.Background())
	if err != nil {
		log.Panic(err)
	}
	if !exists {
		// Create a new index.
		fmt.Println("Create a new index.")
		mapping := `
		{
			"settings": {
			  "number_of_shards": 1,
			  "number_of_replicas": 0
			},
			"mappings": {
				"properties": {
					"product_name": {
						"type": "text",
				  		"analyzer": "ik_max_word",
				 	 	"search_analyzer": "ik_max_word"
					},
					price: {
						"type": float
					}
				}
			}
		  }
		`
		_, err := conn.EsClient.CreateIndex("product").Body(mapping).Do(context.Background())
		if err != nil {
			// Handle error
			log.Panic(err)
		}
		var product []models.Product
		conn.Db.Find(&product)
		fmt.Println(product)
		for i := 0; i < len(product); i++ {
			Add(product[i])
		}
	}
}

// Add 增加商品数据
func Add(product models.Product) {
	_, err := conn.EsClient.Index().
		Index("product").
		Id(strconv.Itoa(int(product.Id))).
		BodyJson(product).
		Do(context.Background())
	if err != nil {
		// Handle error
		fmt.Println("添加商品数据失败")
		return
	}
}

// Delete 删除
func Delete(id string) {
	res, err := conn.EsClient.Delete().
		Index("product").
		Type("_doc").
		Id(id).
		Do(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("delete result %s\n", res.Result)
}

// Update 更新数据
func Update(product models.Product) {
	res, err := conn.EsClient.Update().
		Index("product").
		Type("_doc").
		Id(strconv.FormatInt(product.Id, 10)).
		Doc(product).
		Do(context.Background())
	if err != nil {
		fmt.Println("修改数据失败")
		return
	}
	fmt.Printf("update %s\n", res.Result)
}

// Query 搜索
func Query(c *gin.Context) {
	keyWord := strconv.Quote(c.Query("keyWord"))
	matchQuery := elastic.NewMatchQuery("product_name", keyWord)
	res, err := conn.EsClient.Search().
		Index("product").
		Query(matchQuery).
		Sort("productName", true).
		Do(context.Background())
	fmt.Println(res)
	if err != nil {
		c.JSON(500, models.NewResponse(false, "搜索失败", "出错了，请稍后再试"))
		return
	}
	c.JSON(200, models.NewResponse(true, res, "搜索成功"))
}
