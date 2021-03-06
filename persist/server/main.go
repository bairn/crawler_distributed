package main

import (
	"github.com/bairn/crawler_distributed/config"
	"github.com/bairn/crawler_distributed/persist"
	"github.com/bairn/crawler_distributed/rpcsupport"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func main() {
	log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))
}


func serveRpc(host string, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}