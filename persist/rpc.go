package persist

import (
	"crawler/engine"
	"crawler/persist"
	"gopkg.in/olivere/elastic.v5"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}


func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Index, s.Client, item)
	if err != nil {
		return err
	}
	*result = "ok"
	return nil
}