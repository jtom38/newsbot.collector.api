package services

import (
	"errors"

	"github.com/jtom38/newsbot/collector/domain/model"
)

type CacheClient struct{}

var (
	cacheStorage []*model.CacheItem

	ErrCacheRecordMissing = errors.New("unable to find the requested record.")
)


func NewCacheClient() CacheClient {
	return CacheClient{}
}

func (cc *CacheClient) Insert(item *model.CacheItem) {
	//_, err := cc.Find(item.Key, item.Group)
	//if err != nil { }
	cacheStorage = append(cacheStorage, item)
}

func (cc *CacheClient) Find(key string, group string) (*model.CacheItem, error) {
	//go cc.FindExpiredEntries()

	for _, item := range cacheStorage {
		if item.Group != group { continue }

		if item.Key != key { continue }

		return item, nil
	}

	return &model.CacheItem{}, ErrCacheRecordMissing
}