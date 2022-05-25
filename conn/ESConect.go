//Package conn
/*
   @author:xie
   @date:2022/5/24
   @note:
*/
package conn

import (
	"github.com/olivere/elastic/v7"
	"leastMall_gin/models"
	"log"
)

var EsClient *elastic.Client

func init() {
	cfg := models.Conf.ES
	EsClient, err = elastic.NewClient(elastic.SetURL(cfg.Host+":"+cfg.Port), elastic.SetSniff(cfg.Sniff))
	if err != nil {
		log.Println(err)
	}
	log.Println("elasticsearch connect success")
}
