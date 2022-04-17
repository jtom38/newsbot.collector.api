package services_test

import (
	"testing"
	"time"

	"github.com/jtom38/newsbot/collector/domain/model"
	"github.com/jtom38/newsbot/collector/services"
)

func TestNewCacheClient(t *testing.T) {
	_ = services.NewCacheClient()
}

func TestInsert(t *testing.T) {
	cache := services.NewCacheClient()
	var item *model.CacheItem = &model.CacheItem{
		Key: "UnitTesting",
		Value: "Something, or nothing",
		Group: "Testing",
		Expires: time.Now().Add(5 * time.Second),
	}
	cache.Insert(item)
}

func TestFindGroupMissing(t *testing.T) {
	cache := services.NewCacheClient()
	_, err := cache.Find("UnitTesting", "Unknown")
	if err == nil { panic("Nothing was appended with the requested group.") }
}

func TestFindGroupExists(t *testing.T) {
	cache := services.NewCacheClient()
	var item *model.CacheItem = &model.CacheItem{
		Key: "UnitTesting",
		Value: "Something, or nothing",
		Group: "Testing",
		Expires: time.Now().Add(5 * time.Second),
	}
	cache.Insert(item)
	_, err := cache.Find("UnitTesting", "Testing2")
	//t.Log(res)
	if err == nil { panic("") }
}


func TestCacheStorage(t *testing.T) {
	cc := services.NewCacheClient()
	
	item1 := &model.CacheItem {
		Key: "UnitTesting01",
		Value: "",
		Group: "Testing",
		Expires: time.Now().Add(5 * time.Minute),	
	}
	cc.Insert(item1)

	item2 := &model.CacheItem {
		Key: "UnitTesting02",
		Value: "",
		Group: "Testing",
		Expires: time.Now().Add(5 * time.Minute),	
	}
	cc.Insert(item2)

	cache := services.NewCacheClient()
	_, err := cache.Find("UnitTesting02", "Testing")
	if err != nil { panic("expected to find the value")}
}