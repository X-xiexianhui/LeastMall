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
	"log"
)

var EsClient *elastic.Client

// init函数实饿汉式单例模式
func init() {
	cfg := Conf.ES
	url := fmt.Sprintf("http://%s:%s", cfg.Host, cfg.Port)
	EsClient, err = elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false))
	if err != nil {
		log.Panicln(err)
	}
	log.Println("elasticsearch connect success")
}
