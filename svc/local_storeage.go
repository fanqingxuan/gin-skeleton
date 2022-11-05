package svc

import (
	"gin-skeleton/config"
	"time"

	"github.com/patrickmn/go-cache"
)

type LocalStorage struct {
	cache  *cache.Cache
	prefix string
}

func NewLocalStorage(c config.Config) *LocalStorage {

	return &LocalStorage{
		cache:  cache.New(time.Duration(c.LocalStorage.DefaultExpiration)*time.Second, time.Duration(c.LocalStorage.CleanupInterval)*time.Second),
		prefix: c.LocalStorage.KeyPrefix,
	}
}

func (that *LocalStorage) withPrefix(k string) string {
	return that.prefix + k
}

func (that *LocalStorage) Get(k string) (interface{}, bool) {
	return that.cache.Get(that.withPrefix(k))
}

func (that *LocalStorage) GetWithExpiration(k string) (interface{}, time.Time, bool) {
	return that.cache.GetWithExpiration(that.withPrefix(k))
}

func (that *LocalStorage) Set(k string, value interface{}) {
	that.cache.SetDefault(that.withPrefix(k), value)
}

func (that *LocalStorage) SetWithExpire(k string, value interface{}, d time.Duration) {
	that.cache.Set(that.withPrefix(k), value, d)
}

func (that *LocalStorage) Add(k string, value interface{}, d time.Duration) error {
	return that.cache.Add(that.withPrefix(k), value, d)
}

func (that *LocalStorage) Replace(k string, value interface{}, d time.Duration) error {
	return that.cache.Replace(that.withPrefix(k), value, d)
}

func (that *LocalStorage) Decr(k string, n int64) error {
	return that.cache.Decrement(that.withPrefix(k), n)
}

func (that *LocalStorage) Incr(k string, n int64) error {
	return that.cache.Increment(that.withPrefix(k), n)
}

func (that *LocalStorage) Delete(k string) {
	that.cache.Delete(that.withPrefix(k))
}

func (that *LocalStorage) DeleteExpired() {
	that.cache.DeleteExpired()
}

func (that *LocalStorage) Flush() {
	that.cache.Flush()
}

func (that *LocalStorage) Items() map[string]cache.Item {
	return that.cache.Items()
}
