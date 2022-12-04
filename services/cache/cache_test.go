package cache_test

import (
	"testing"

	"github.com/jtom38/newsbot/collector/services/cache"
)

func TestNewCacheClient(t *testing.T) {
	_ = cache.NewCacheClient("placeholder")
}

func TestInsert(t *testing.T) {
	cache := cache.NewCacheClient("Testing")
	cache.Insert("UnitTesting", "Something, or nothing")
}

func TestFindGroupMissing(t *testing.T) {
	cache := cache.NewCacheClient("faker")
	_, err := cache.FindByKey("UnitTesting")
	if err == nil {
		panic("Nothing was appended with the requested group.")
	}
}

func TestFindGroupExists(t *testing.T) {
	cache := cache.NewCacheClient("Testing")
	cache.Insert("UnitTesting", "Something")
	_, err := cache.FindByKey("UnitTesting")
	if err != nil {
		panic("")
	}
}

func TestCacheStorage(t *testing.T) {
	cc := cache.NewCacheClient("Testing")
	cc.Insert("UnitTesting01", "test")
	cc.Insert("UnitTesting02", "Test")

	cache := cache.NewCacheClient("Testing")
	_, err := cache.FindByKey("UnitTesting02")
	if err != nil {
		panic("expected to find the value")
	}
}
