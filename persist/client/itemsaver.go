package client

import (
	"crawler/engine"
	"crawler/model"
	"crawler_distributed/config"
	"crawler_distributed/rpcsupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error)  {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out

			if item.Payload == nil {
				continue
			}

			profile := item.Payload.(model.Profile)
			if profile.Age == "" {
				continue
			}

			log.Printf("Item Saver:got item "+"#%d: %v", itemCount, item)
			itemCount ++

			result := ""
			err = client.Call(config.ItemSaverPrc, item, &result)

			if err != nil || result != "ok" {
				log.Printf("Item Saver: error " + "saving item %v:%v", item, err)
			}
		}
	}()

	return out, nil
}
