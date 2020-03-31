package main

import (
	"crawler/engine"
	common "crawler/model"
	"crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	go serveRpc(host, "test1")

	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	item := engine.Item{
		Url:     "https://album.zhenai.com/u/1682657696",
		Id:      "1682657696",
		Type:    "zhenai",
		Payload: common.Profile{
			Name:       "心瘾",
			Marriage:   "离异",
			Age:        "34岁",
			Gender:     "女",
			Height:     "160",
			Weight:     "50kg",
			Income:     "12000",
			Education:  "本科",
			Occupation: "",
			Hukou:      "",
			Xingzuo:    "",
			House:      "",
			Car:        "",
		},
	}

	result := ""
	err = client.Call("ItemSaverService.Save", item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s;err: %s", result, err)
	}
}
