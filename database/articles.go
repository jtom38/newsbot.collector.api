package database

import (
	"errors"

	"github.com/jtom38/newsbot/collector/domain/model"
)

// Generate this struct fr
type ArticlesClient struct {
	rootUri string
}

func (ac *ArticlesClient) List() []model.Articles {
	var items []model.Articles
	return items
}

func (ac *ArticlesClient) Find() []model.Articles {
	var items []model.Articles
	return items
}

func (ac *ArticlesClient) FindByUrl(url string) model.Articles {
	return model.Articles{}
}

func (ac *ArticlesClient) Delete(id int32) error {
	return errors.New("not implemented")
}

func (ac *ArticlesClient) Add() error {
	return errors.New("not implemented")
}
