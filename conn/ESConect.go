//Package conn
/*
   @author:xie
   @date:2022/5/24
   @note:
*/
package conn

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
