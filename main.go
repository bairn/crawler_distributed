package main

import (
	"github.com/bairn/crawler/engine"
	"github.com/bairn/crawler/scheduler"
	"github.com/bairn/crawler/zhenai/parser"
	"github.com/bairn/crawler_distributed/config"
	client2 "github.com/bairn/crawler_distributed/persist/client"
	"github.com/bairn/crawler_distributed/rpcsupport"
	"github.com/bairn/crawler_distributed/worker/client"
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"strings"
)

var (
	workerHosts = flag.String("worker_hosts", "", "worker hosts (comma separated")
)

func main() {
	flag.Parse()

	itemChan , err := client2.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}



	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor, err := client.CreateProcessor(pool)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/beijing",
		Parser: engine.NewFuncParser(parser.ParseCity, "ParseCity"),
	})
}


func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client

	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
		} else {
			log.Printf("error connetct to %s:%v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()

	return out
}
