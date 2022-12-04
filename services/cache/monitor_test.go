package cache_test

import (
	"testing"

	"github.com/jtom38/newsbot/collector/services/cache"
)

func TestCacheTaintItem(t *testing.T) {
	cc := cache.NewCacheClient("Testing")
	cc.Insert("UnitTesting01", "test")

}
