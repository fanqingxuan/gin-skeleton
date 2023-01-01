package svc

import (
	"gin-skeleton/config"
	"time"

	"github.com/patrickmn/go-cache"
)

type XCache struct {
	cache  *cache.Cache
	prefix string
}

func NewXCache(c config.Config) *XCache {

	return &XCache{
		cache:  cache.New(time.Duration(c.XCache.DefaultExpiration)*time.Second, time.Duration(c.XCache.CleanupInterval)*time.Second),
		prefix: c.XCache.KeyPrefix,
	}
}

func (that *XCache) getKey(k string) string {
	return that.prefix + k
}

func (that *XCache) Get(k string) (interface{}, bool) {
	return that.cache.Get(that.getKey(k))
}

func (that *XCache) GetWithExpiration(k string) (interface{}, time.Time, bool) {
	return that.cache.GetWithExpiration(that.getKey(k))
}

func (that *XCache) Set(k string, value interface{}) {
	that.cache.SetDefault(that.getKey(k), value)
}

func (that *XCache) SetWithExpire(k string, value interface{}, d time.Duration) {
	that.cache.Set(that.getKey(k), value, d)
}

func (that *XCache) Add(k string, value interface{}, d time.Duration) error {
	return that.cache.Add(that.getKey(k), value, d)
}

func (that *XCache) Replace(k string, value interface{}, d time.Duration) error {
	return that.cache.Replace(that.getKey(k), value, d)
}

func (that *XCache) Decr(k string, n int64) error {
	return that.cache.Decrement(that.getKey(k), n)
}

func (that *XCache) Incr(k string, n int64) error {
	return that.cache.Increment(that.getKey(k), n)
}

func (that *XCache) Delete(k string) {
	that.cache.Delete(that.getKey(k))
}

func (that *XCache) DeleteExpired() {
	that.cache.DeleteExpired()
}

func (that *XCache) Flush() {
	that.cache.Flush()
}

func (that *XCache) Items() map[string]cache.Item {
	return that.cache.Items()
}
