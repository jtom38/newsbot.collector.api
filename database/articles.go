package database

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jtom38/newsbot/collector/domain/model"
)

// Generate this struct fr
type ArticlesClient struct {
	rootUri string
}

func (ac *ArticlesClient) List() ([]model.Articles, error) {
	var items []model.Articles
	url := fmt.Sprintf("%v/api/v1/articles", ac.rootUri)
	resp := getContent(url)

	err := json.Unmarshal(resp, &items)
	if err != nil { return []model.Articles{}, err }
	
	return items, nil
}

func (ac *ArticlesClient) FindByID(ID uint) (model.Articles, error) {
	var items model.Articles
	url := fmt.Sprintf("%v/api/v1/articles/%v", ac.rootUri, ID)
	resp := getContent(url)

	err := json.Unmarshal(resp, &items)
	if err != nil { return items, err }

	return items, nil
}

func (ac *ArticlesClient) FindByUrl(url string) (model.Articles, error) {
	var item model.Articles
	get := fmt.Sprintf("%v/api/v1/articles/url/%v", ac.rootUri, url)
	resp := getContent(get)

	if resp.string() == "404 page not found\n" {
		
	}
	err := json.Unmarshal(resp, &item)
	if err != nil { return item, err }

	return item, nil
}

func (ac *ArticlesClient) Delete(id int32) error {
	return errors.New("not implemented")
}

func (ac *ArticlesClient) Add() error {
	return errors.New("not implemented")
}
