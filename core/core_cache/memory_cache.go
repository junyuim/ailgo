package core_cache

import (
	"container/list"
	"encoding/json"
	"github.com/coocood/freecache"
	"strings"
	"unsafe"
)

type MemoryCacheConfig struct {
	//缓存大小(MB)
	Size int `json:"size" yaml:"size"`
}

type MemoryCacheHandler struct {
	cache *freecache.Cache
}

func NewMemoryCacheHandler(config *MemoryCacheConfig) *MemoryCacheHandler {
	return &MemoryCacheHandler{
		cache: freecache.NewCache(config.Size * 1024 * 1024),
	}
}

func (handler *MemoryCacheHandler) Keys(prefix string) ([]string, error) {
	l := list.New()

	iter := handler.cache.NewIterator()

	for i := iter.Next(); i != nil; i = iter.Next() {
		key := string(i.Key)
		if strings.HasPrefix(key, prefix) {
			l.PushBack(key)
		}
	}

	a := make([]string, l.Len())
	ai := 0
	for e := l.Front(); e != nil; e = e.Next() {
		a[ai] = *((*string)(unsafe.Pointer(&e.Value)))
	}

	return a, nil
}

func (handler *MemoryCacheHandler) Has(key string) (bool, error) {
	_, err := handler.cache.Get([]byte(key))

	return err == nil, err
}

//
//func (handler *MemoryCacheHandler) Add(key string, value any, expire int64) (bool, error) {
//	has, err := handler.Has(key)
//
//	if has {
//		return false, err
//	}
//
//	data, err := json.Marshal(value)
//
//	if err != nil {
//		return false, err
//	}
//
//	err = handler.cache.Set([]byte(key), data, int(expire))
//
//	if err != nil {
//		return false, err
//	}
//
//	return true, nil
//}

func (handler *MemoryCacheHandler) Get(key string, out any) (bool, error) {
	res, err := handler.cache.Get([]byte(key))

	if err != nil {
		return false, err
	}

	if len(res) < 1 {
		return false, nil
	}

	return true, json.Unmarshal(res, &out)
}

func (handler *MemoryCacheHandler) Set(key string, value any, expire int64) error {
	data, err := json.Marshal(value)

	if err != nil {
		return err
	}

	return handler.cache.Set([]byte(key), data, int(expire))
}

func (handler *MemoryCacheHandler) Del(key string) (bool, error) {
	return handler.cache.Del([]byte(key)), nil
}
