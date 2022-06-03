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
				"content": {
				  "type": "text",
				  "analyzer": "ik_max_word",
				  "search_analyzer": "ik_max_word"
				},
				"title": {
				  "type": "text",
				  "analyzer": "ik_max_word",
				  "search_analyzer": "ik_max_word"
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

//查找
func gets(id string) {
	//通过id查找
	get1, err := conn.EsClient.Get().Index("product").Type("_doc").Id(id).Do(context.Background())
	if err != nil {
		fmt.Println("查找失败：" + err.Error())
	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}
}
