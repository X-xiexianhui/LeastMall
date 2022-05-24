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
	"log"
)

var EsClient *elastic.Client

func init() {
	cfg := modules.Conf.ES
	fmt.Println(cfg.URL)
	EsClient, err = elastic.NewClient(elastic.SetURL(cfg.URL), elastic.SetSniff(cfg.Sniff))
	if err != nil {
		log.Println(err)
	}
	fmt.Println("elasticsearch connect success")
}
