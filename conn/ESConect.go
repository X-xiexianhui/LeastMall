//Package conn
/*
   @author:xie
   @date:2022/5/24
   @note:
*/
package conn

import (
	"github.com/olivere/elastic/v7"
	"log"
)

var EsClient *elastic.Client

// init函数实饿汉式单例模式
func init() {
	cfg := Conf.ES
	EsClient, err = elastic.NewClient(elastic.SetURL(cfg.Host+":"+cfg.Port), elastic.SetSniff(cfg.Sniff))
	if err != nil {
		log.Println(err)
	}
	log.Println("elasticsearch connect success")
}
