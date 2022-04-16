package services

import (
	"errors"
	"time"

	"github.com/jtom38/newsbot/collector/domain/model"
)

type CacheClient struct{}



var cacheStorage []*model.CacheItem

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

	return &model.CacheItem{}, errors.New("unable to find the requested record.")
}

// This will be fired off each time an cache a
func (cc *CacheClient) FindExpiredEntries() {
	now := time.Now()
	for index, item := range cacheStorage {
		res := now.After(item.Expires)
		if res {
			cc.removeExpiredEntries(index)
		}
	}
}

// This will create a new slice and add the valid items to it and ignore the one to be removed.
// The existing cacheStorage will be replaced.
func (cc *CacheClient) removeExpiredEntries(arrayEntry int) {
	var temp []*model.CacheItem
	for index, item := range cacheStorage {
		if index == arrayEntry { continue }
		temp = append(temp, item)
	}
	cacheStorage = temp
}