package models

import (
	"fmt"
	"github.com/olivere/elastic/v7"
)

var EsClient *elastic.Client

func init() {
	EsClient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"), elastic.SetSniff(false))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("elasticsearch connect success")
}
