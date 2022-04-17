package services

import (
	"time"

	"github.com/jtom38/newsbot/collector/domain/model"
)

type CacheMonitor struct {}

func NewCacheMonitorClient() CacheMonitor {
	return CacheMonitor{}
}

func (cm *CacheMonitor) Enable() {
	
}

// This will be fired off each time an cache a
func (cm *CacheMonitor) FindExpiredEntries() {
	now := time.Now()
	for index, item := range cacheStorage {
		res := now.After(item.Expires)
		if res {
			cm.removeExpiredEntries(index)
		}
	}
}

// This will create a new slice and add the valid items to it and ignore the one to be removed.
// The existing cacheStorage will be replaced.
func (cc *CacheMonitor) removeExpiredEntries(arrayEntry int) {
	var temp []*model.CacheItem
	for index, item := range cacheStorage {
		if index == arrayEntry { continue }
		temp = append(temp, item)
	}
	cacheStorage = temp
} 