package cache

import (
	"time"

	"github.com/jtom38/newsbot/collector/domain/models"
)

type CacheClient struct {
	group        string
	DefaultTimer time.Duration
}

func NewCacheClient(group string) CacheClient {
	return CacheClient{
		group:        group,
		DefaultTimer: time.Hour,
	}
}

func (cc *CacheClient) Insert(key string, value string) {
	item := models.CacheItem{
		Key:       key,
		Value:     value,
		Group:     cc.group,
		Expires:   time.Now().Add(1 * time.Hour),
		IsTainted: false,
	}
	cacheStorage = append(cacheStorage, &item)
}

func (cc *CacheClient) FindByKey(key string) (*models.CacheItem, error) {
	for _, item := range cacheStorage {
		if item.Group != cc.group {
			continue
		}
		if item.Key != key {
			continue
		}

		// if it was tainted, renew the timer and remove the taint as this record was still needed
		if item.IsTainted {
			item.IsTainted = false
			item.Expires = time.Now().Add(1 * time.Hour)
		}
		return item, nil
	}

	return &models.CacheItem{}, ErrCacheRecordMissing
}

func (cc *CacheClient) FindByValue(value string) (*models.CacheItem, error) {
	for _, item := range cacheStorage {
		if item.Group != cc.group {
			continue
		}
		if item.Value != value {
			continue
		}

		// if it was tainted, renew the timer and remove the taint as this record was still needed
		if item.IsTainted {
			item.IsTainted = false
			item.Expires = time.Now().Add(1 * time.Hour)
		}
		return item, nil
	}
	return &models.CacheItem{}, ErrCacheRecordMissing
}
