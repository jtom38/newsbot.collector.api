package cache

import (
	"time"

	"github.com/jtom38/newsbot/collector/domain/model"
)

// When a record becomes tainted, it needs to be renewed or it will be dropped from the cache.
// If a record is tainted and used again, the taint will be removed and a new Expires value will be set.
// If its not renewed, it will be dropped.
type CacheAgeMonitor struct{}

func NewCacheAgeMonitor() CacheAgeMonitor {
	return CacheAgeMonitor{}
}

// This is an automated job that will review all the objects for age and taint them if needed.
func (cam CacheAgeMonitor) CheckExpiredEntries() {
	now := time.Now()
	for index, item := range cacheStorage {
		if now.After(item.Expires) {

			// the timer expired, and its not tainted, taint it
			if !item.IsTainted {
				item.IsTainted = true
				item.Expires = now.Add(1 * time.Hour)
			}

			// if its tainted and the timer didnt get renewed, delete
			if item.IsTainted {
				cacheStorage = cam.removeEntry(index)
			}
		}
	}
}

// This creates a new slice and skips over the item that needs to be dropped
func (cam CacheAgeMonitor) removeEntry(index int) []*model.CacheItem {
	var temp []*model.CacheItem
	for i, item := range cacheStorage {
		if i != index {
			temp = append(temp, item)
		}
	}
	return temp
}
