package cache

import (
	"errors"

	"github.com/jtom38/newsbot/collector/domain/model"
)

var (
	cacheStorage []*model.CacheItem

	ErrCacheRecordMissing = errors.New("unable to find the requested record")
)