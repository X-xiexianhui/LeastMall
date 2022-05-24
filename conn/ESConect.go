//Package conn
/*
   @author:xie
   @date:2022/5/24
   @note:
*/
package conn

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
)

var EsClient *elasticsearch.Client

func init() {
	EsClient, err = elasticsearch.NewDefaultClient()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("elasticsearch connect success")
}
