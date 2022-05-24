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
	"leastMall_gin/modules"
)

var EsClient *elastic.Client

func init() {
	cfg := modules.Conf.ES
	EsClient, err = elastic.NewClient(elastic.SetURL(cfg.URL), elastic.SetSniff(cfg.Sniff))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("elasticsearch connect success")
}
