package svc

import (
	"gin-skeleton/config"
	"time"

	"github.com/patrickmn/go-cache"
)

type CacheX struct {
	cache  *cache.Cache
	prefix string
}

func NewCacheX(c config.Config) *CacheX {

	return &CacheX{
		cache:  cache.New(time.Duration(c.CacheX.DefaultExpiration)*time.Second, time.Duration(c.CacheX.CleanupInterval)*time.Second),
		prefix: c.CacheX.KeyPrefix,
	}
}

func (that *CacheX) getKey(k string) string {
	return that.prefix + k
}

func (that *CacheX) Get(k string) (interface{}, bool) {
	return that.cache.Get(that.getKey(k))
}

func (that *CacheX) GetWithExpiration(k string) (interface{}, time.Time, bool) {
	return that.cache.GetWithExpiration(that.getKey(k))
}

func (that *CacheX) Set(k string, value interface{}) {
	that.cache.SetDefault(that.getKey(k), value)
}

func (that *CacheX) SetWithExpire(k string, value interface{}, d time.Duration) {
	that.cache.Set(that.getKey(k), value, d)
}

func (that *CacheX) Add(k string, value interface{}, d time.Duration) error {
	return that.cache.Add(that.getKey(k), value, d)
}

func (that *CacheX) Replace(k string, value interface{}, d time.Duration) error {
	return that.cache.Replace(that.getKey(k), value, d)
}

func (that *CacheX) Decr(k string, n int64) error {
	return that.cache.Decrement(that.getKey(k), n)
}

func (that *CacheX) Incr(k string, n int64) error {
	return that.cache.Increment(that.getKey(k), n)
}

func (that *CacheX) Delete(k string) {
	that.cache.Delete(that.getKey(k))
}

func (that *CacheX) DeleteExpired() {
	that.cache.DeleteExpired()
}

func (that *CacheX) Flush() {
	that.cache.Flush()
}

func (that *CacheX) Items() map[string]cache.Item {
	return that.cache.Items()
}
