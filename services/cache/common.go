package cache

import (
	"errors"

	"github.com/jtom38/newsbot/collector/domain/models"
)

var (
	cacheStorage []*models.CacheItem

	ErrCacheRecordMissing = errors.New("unable to find the requested record")
)
