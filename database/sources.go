package database

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jtom38/newsbot/collector/domain/model"
)

type SourcesClient struct {
	rootUri string
}

func (sb *SourcesClient) List() ([]model.Sources, error) {
	var items []model.Sources
	url := fmt.Sprintf("%v/api/v1/sources", sb.rootUri)
	resp := getContent(url)

	err := json.Unmarshal(resp, &items)
	if err != nil { return []model.Sources{}, err }
	
	return items, nil
}

func (sb *SourcesClient) FindBySource(SourceType string) ([]model.Sources, error) {
	items, err := sb.List()
	if err != nil { log.Panicln(err) }

	var res []model.Sources
	for _, item := range(items) {
		if item.Source == SourceType {
			res = append(res, item)
		}
	}
	return res, nil
}