package main

import (
	"crawler_distributed/rpcsupport"
	"crawler_distributed/worker"
	"fmt"
	"testing"
	"time"
)

func TestCrawService(t *testing.T) {
	go rpcsupport.ServeRpc(":9000", worker.CrawService{})
	
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(":9000")
	if err != nil {
		panic(err)
	}
	
	
	req := worker.Request{
		Url:    "https://album.zhenai.com/u/1598901060",
		Parser: worker.SerializedParser{
			Name: "ParseProfile",
			Args : "安静的雪",
		},
	}
	var result worker.ParseResult
	err = client.Call("CrawService.Process", req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
