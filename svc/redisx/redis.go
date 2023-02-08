package redisx

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	red "github.com/go-redis/redis/v8"
)

type Client struct {
	conn *red.Client
}

type Options = redis.Options

func New(options *Options) *Client {
	return &Client{
		conn: red.NewClient(options),
	}
}

var (
	_ Cmdable = (*Client)(nil)
)

func (client *Client) Command() (map[string]*CommandInfo, error) {
	return client.CommandCtx(context.Background())
}

func (client *Client) CommandCtx(ctx context.Context) (map[string]*CommandInfo, error) {
	return client.conn.Command(ctx).Result()
}

func (client *Client) ClientGetName() (string, error) {
	return client.ClientGetNameCtx(context.Background())
}

func (client *Client) ClientGetNameCtx(ctx context.Context) (string, error) {
	return client.conn.ClientGetName(ctx).Result()
}

func (client *Client) Echo(message interface{}) (string, error) {
	return client.EchoCtx(context.Background(), message)
}

func (client *Client) EchoCtx(ctx context.Context, message interface{}) (string, error) {
	return client.conn.Echo(ctx, message).Result()
}

func (client *Client) Ping() (string, error) {
	return client.PingCtx(context.Background())
}

func (client *Client) PingCtx(ctx context.Context) (string, error) {
	return client.conn.Ping(ctx).Result()
}

func (client *Client) Del(keys ...string) (int64, error) {
	return client.DelCtx(context.Background(), keys...)
}

func (client *Client) DelCtx(ctx context.Context, keys ...string) (int64, error) {
	return client.conn.Del(ctx, keys...).Result()
}

func (client *Client) Unlink(keys ...string) (int64, error) {
	return client.UnlinkCtx(context.Background(), keys...)
}

func (client *Client) UnlinkCtx(ctx context.Context, keys ...string) (int64, error) {
	return client.conn.Unlink(ctx, keys...).Result()
}

func (client *Client) Dump(key string) (string, error) {
	return client.DumpCtx(context.Background(), key)
}

func (client *Client) DumpCtx(ctx context.Context, key string) (string, error) {
	return client.conn.Dump(ctx, key).Result()
}

func (client *Client) Exists(keys ...string) (int64, error) {
	return client.ExistsCtx(context.Background(), keys...)
}

func (client *Client) ExistsCtx(ctx context.Context, keys ...string) (int64, error) {
	return client.conn.Exists(ctx, keys...).Result()
}

func (client *Client) Expire(key string, expiration time.Duration) (bool, error) {
	return client.ExpireCtx(context.Background(), key, expiration)
}

func (client *Client) ExpireCtx(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return client.conn.Expire(ctx, key, expiration).Result()
}

func (client *Client) ExpireAt(key string, tm time.Time) (bool, error) {
	return client.ExpireAtCtx(context.Background(), key, tm)
}

func (client *Client) ExpireAtCtx(ctx context.Context, key string, tm time.Time) (bool, error) {
	return client.conn.ExpireAt(ctx, key, tm).Result()
}

func (client *Client) ExpireNX(key string, expiration time.Duration) (bool, error) {
	return client.ExpireNXCtx(context.Background(), key, expiration)
}

func (client *Client) ExpireNXCtx(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return client.conn.ExpireNX(ctx, key, expiration).Result()
}

func (client *Client) ExpireXX(key string, expiration time.Duration) (bool, error) {
	return client.ExpireXXCtx(context.Background(), key, expiration)
}

func (client *Client) ExpireXXCtx(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return client.conn.ExpireXX(ctx, key, expiration).Result()
}

func (client *Client) ExpireGT(key string, expiration time.Duration) (bool, error) {
	return client.ExpireGTCtx(context.Background(), key, expiration)
}

func (client *Client) ExpireGTCtx(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return client.conn.ExpireGT(ctx, key, expiration).Result()
}

func (client *Client) ExpireLT(key string, expiration time.Duration) (bool, error) {
	return client.ExpireLTCtx(context.Background(), key, expiration)
}

func (client *Client) ExpireLTCtx(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return client.conn.ExpireLT(ctx, key, expiration).Result()
}

func (client *Client) Keys(pattern string) ([]string, error) {
	return client.KeysCtx(context.Background(), pattern)
}

func (client *Client) KeysCtx(ctx context.Context, pattern string) ([]string, error) {
	return client.conn.Keys(ctx, pattern).Result()
}

func (client *Client) Migrate(host, port, key string, db int, timeout time.Duration) (string, error) {
	return client.MigrateCtx(context.Background(), host, port, key, db, timeout)
}

func (client *Client) MigrateCtx(ctx context.Context, host, port, key string, db int, timeout time.Duration) (string, error) {
	return client.conn.Migrate(ctx, host, port, key, db, timeout).Result()
}

func (client *Client) Move(key string, db int) (bool, error) {
	return client.MoveCtx(context.Background(), key, db)
}

func (client *Client) MoveCtx(ctx context.Context, key string, db int) (bool, error) {
	return client.conn.Move(ctx, key, db).Result()
}

func (client *Client) ObjectRefCount(key string) (int64, error) {
	return client.ObjectRefCountCtx(context.Background(), key)
}

func (client *Client) ObjectRefCountCtx(ctx context.Context, key string) (int64, error) {
	return client.conn.ObjectRefCount(ctx, key).Result()
}

func (client *Client) ObjectEncoding(key string) (string, error) {
	return client.ObjectEncodingCtx(context.Background(), key)
}

func (client *Client) ObjectEncodingCtx(ctx context.Context, key string) (string, error) {
	return client.conn.ObjectEncoding(ctx, key).Result()
}

func (client *Client) ObjectIdleTime(key string) (time.Duration, error) {
	return client.ObjectIdleTimeCtx(context.Background(), key)
}

func (client *Client) ObjectIdleTimeCtx(ctx context.Context, key string) (time.Duration, error) {
	return client.conn.ObjectIdleTime(ctx, key).Result()
}

func (client *Client) Persist(key string) (bool, error) {
	return client.PersistCtx(context.Background(), key)
}

func (client *Client) PersistCtx(ctx context.Context, key string) (bool, error) {
	return client.conn.Persist(ctx, key).Result()
}

func (client *Client) PExpire(key string, expiration time.Duration) (bool, error) {
	return client.PExpireCtx(context.Background(), key, expiration)
}

func (client *Client) PExpireCtx(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return client.conn.PExpire(ctx, key, expiration).Result()
}

func (client *Client) PExpireAt(key string, tm time.Time) (bool, error) {
	return client.PExpireAtCtx(context.Background(), key, tm)
}

func (client *Client) PExpireAtCtx(ctx context.Context, key string, tm time.Time) (bool, error) {
	return client.conn.PExpireAt(ctx, key, tm).Result()
}

func (client *Client) PTTL(key string) (time.Duration, error) {
	return client.PTTLCtx(context.Background(), key)
}

func (client *Client) PTTLCtx(ctx context.Context, key string) (time.Duration, error) {
	return client.conn.PTTL(ctx, key).Result()
}

func (client *Client) RandomKey() (string, error) {
	return client.RandomKeyCtx(context.Background())
}

func (client *Client) RandomKeyCtx(ctx context.Context) (string, error) {
	return client.conn.RandomKey(ctx).Result()
}

func (client *Client) Rename(key, newkey string) (string, error) {
	return client.RenameCtx(context.Background(), key, newkey)
}

func (client *Client) RenameCtx(ctx context.Context, key, newkey string) (string, error) {
	return client.conn.Rename(ctx, key, newkey).Result()
}

func (client *Client) RenameNX(key, newkey string) (bool, error) {
	return client.RenameNXCtx(context.Background(), key, newkey)
}

func (client *Client) RenameNXCtx(ctx context.Context, key, newkey string) (bool, error) {
	return client.conn.RenameNX(ctx, key, newkey).Result()
}

func (client *Client) Restore(key string, ttl time.Duration, value string) (string, error) {
	return client.RestoreCtx(context.Background(), key, ttl, value)
}

func (client *Client) RestoreCtx(ctx context.Context, key string, ttl time.Duration, value string) (string, error) {
	return client.conn.Restore(ctx, key, ttl, value).Result()
}

func (client *Client) RestoreReplace(key string, ttl time.Duration, value string) (string, error) {
	return client.RestoreReplaceCtx(context.Background(), key, ttl, value)
}

func (client *Client) RestoreReplaceCtx(ctx context.Context, key string, ttl time.Duration, value string) (string, error) {
	return client.conn.RestoreReplace(ctx, key, ttl, value).Result()
}

func (client *Client) Sort(key string, sort *Sort) ([]string, error) {
	return client.SortCtx(context.Background(), key, sort)
}

func (client *Client) SortCtx(ctx context.Context, key string, sort *Sort) ([]string, error) {
	return client.conn.Sort(ctx, key, sort).Result()
}

func (client *Client) SortStore(key, store string, sort *Sort) (int64, error) {
	return client.SortStoreCtx(context.Background(), key, store, sort)
}

func (client *Client) SortStoreCtx(ctx context.Context, key, store string, sort *Sort) (int64, error) {
	return client.conn.SortStore(ctx, key, store, sort).Result()
}

func (client *Client) SortInterfaces(key string, sort *Sort) ([]interface{}, error) {
	return client.SortInterfacesCtx(context.Background(), key, sort)
}

func (client *Client) SortInterfacesCtx(ctx context.Context, key string, sort *Sort) ([]interface{}, error) {
	return client.conn.SortInterfaces(ctx, key, sort).Result()
}

func (client *Client) Touch(keys ...string) (int64, error) {
	return client.TouchCtx(context.Background(), keys...)
}

func (client *Client) TouchCtx(ctx context.Context, keys ...string) (int64, error) {
	return client.conn.Touch(ctx, keys...).Result()
}

func (client *Client) TTL(key string) (time.Duration, error) {
	return client.TTLCtx(context.Background(), key)
}

func (client *Client) TTLCtx(ctx context.Context, key string) (time.Duration, error) {
	return client.conn.TTL(ctx, key).Result()
}

func (client *Client) Type(key string) (string, error) {
	return client.TypeCtx(context.Background(), key)
}

func (client *Client) TypeCtx(ctx context.Context, key string) (string, error) {
	return client.conn.Type(ctx, key).Result()
}

func (client *Client) Append(key, value string) (int64, error) {
	return client.AppendCtx(context.Background(), key, value)
}

func (client *Client) AppendCtx(ctx context.Context, key, value string) (int64, error) {
	return client.conn.Append(ctx, key, value).Result()
}

func (client *Client) Decr(key string) (int64, error) {
	return client.DecrCtx(context.Background(), key)
}

func (client *Client) DecrCtx(ctx context.Context, key string) (int64, error) {
	return client.conn.Decr(ctx, key).Result()
}

func (client *Client) DecrBy(key string, decrement int64) (int64, error) {
	return client.DecrByCtx(context.Background(), key, decrement)
}

func (client *Client) DecrByCtx(ctx context.Context, key string, decrement int64) (int64, error) {
	return client.conn.DecrBy(ctx, key, decrement).Result()
}

func (client *Client) Get(key string) (string, error) {
	return client.GetCtx(context.Background(), key)
}

func (client *Client) GetCtx(ctx context.Context, key string) (string, error) {
	return client.conn.Get(ctx, key).Result()
}

func (client *Client) GetRange(key string, start, end int64) (string, error) {
	return client.GetRangeCtx(context.Background(), key, start, end)
}

func (client *Client) GetRangeCtx(ctx context.Context, key string, start, end int64) (string, error) {
	return client.conn.GetRange(ctx, key, start, end).Result()
}

func (client *Client) GetSet(key string, value interface{}) (string, error) {
	return client.GetSetCtx(context.Background(), key, value)
}

func (client *Client) GetSetCtx(ctx context.Context, key string, value interface{}) (string, error) {
	return client.conn.GetSet(ctx, key, value).Result()
}

func (client *Client) GetEx(key string, expiration time.Duration) (string, error) {
	return client.GetExCtx(context.Background(), key, expiration)
}

func (client *Client) GetExCtx(ctx context.Context, key string, expiration time.Duration) (string, error) {
	return client.conn.GetEx(ctx, key, expiration).Result()
}

func (client *Client) GetDel(key string) (string, error) {
	return client.GetDelCtx(context.Background(), key)
}

func (client *Client) GetDelCtx(ctx context.Context, key string) (string, error) {
	return client.conn.GetDel(ctx, key).Result()
}

func (client *Client) Incr(key string) (int64, error) {
	return client.IncrCtx(context.Background(), key)
}

func (client *Client) IncrCtx(ctx context.Context, key string) (int64, error) {
	return client.conn.Incr(ctx, key).Result()
}

func (client *Client) IncrBy(key string, value int64) (int64, error) {
	return client.IncrByCtx(context.Background(), key, value)
}

func (client *Client) IncrByCtx(ctx context.Context, key string, value int64) (int64, error) {
	return client.conn.IncrBy(ctx, key, value).Result()
}

func (client *Client) IncrByFloat(key string, value float64) (float64, error) {
	return client.IncrByFloatCtx(context.Background(), key, value)
}

func (client *Client) IncrByFloatCtx(ctx context.Context, key string, value float64) (float64, error) {
	return client.conn.IncrByFloat(ctx, key, value).Result()
}

func (client *Client) MGet(keys ...string) ([]interface{}, error) {
	return client.MGetCtx(context.Background(), keys...)
}

func (client *Client) MGetCtx(ctx context.Context, keys ...string) ([]interface{}, error) {
	return client.conn.MGet(ctx, keys...).Result()
}

func (client *Client) MSet(values ...interface{}) (string, error) {
	return client.MSetCtx(context.Background(), values)
}

func (client *Client) MSetCtx(ctx context.Context, values ...interface{}) (string, error) {
	return client.conn.MSet(ctx, values).Result()
}

func (client *Client) MSetNX(values ...interface{}) (bool, error) {
	return client.MSetNXCtx(context.Background(), values)
}

func (client *Client) MSetNXCtx(ctx context.Context, values ...interface{}) (bool, error) {
	return client.conn.MSetNX(ctx, values).Result()
}

func (client *Client) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	return client.SetCtx(context.Background(), key, value, expiration)
}

func (client *Client) SetCtx(ctx context.Context, key string, value interface{}, expiration time.Duration) (string, error) {
	return client.conn.Set(ctx, key, value, expiration).Result()
}

func (client *Client) SetArgs(key string, value interface{}, a SetArgs) (string, error) {
	return client.SetArgsCtx(context.Background(), key, value, a)
}

func (client *Client) SetArgsCtx(ctx context.Context, key string, value interface{}, a SetArgs) (string, error) {
	return client.conn.SetArgs(ctx, key, value, a).Result()
}

func (client *Client) SetEX(key string, value interface{}, expiration time.Duration) (string, error) {
	return client.SetEXCtx(context.Background(), key, value, expiration)
}

func (client *Client) SetEXCtx(ctx context.Context, key string, value interface{}, expiration time.Duration) (string, error) {
	return client.conn.SetEX(ctx, key, value, expiration).Result()
}

func (client *Client) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return client.SetNXCtx(context.Background(), key, value, expiration)
}

func (client *Client) SetNXCtx(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return client.conn.SetNX(ctx, key, value, expiration).Result()
}

func (client *Client) SetXX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return client.SetXXCtx(context.Background(), key, value, expiration)
}

func (client *Client) SetXXCtx(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return client.conn.SetXX(ctx, key, value, expiration).Result()
}

func (client *Client) SetRange(key string, offset int64, value string) (int64, error) {
	return client.SetRangeCtx(context.Background(), key, offset, value)
}

func (client *Client) SetRangeCtx(ctx context.Context, key string, offset int64, value string) (int64, error) {
	return client.conn.SetRange(ctx, key, offset, value).Result()
}

func (client *Client) StrLen(key string) (int64, error) {
	return client.StrLenCtx(context.Background(), key)
}

func (client *Client) StrLenCtx(ctx context.Context, key string) (int64, error) {
	return client.conn.StrLen(ctx, key).Result()
}

func (client *Client) Copy(sourceKey string, destKey string, db int, replace bool) (int64, error) {
	return client.CopyCtx(context.Background(), sourceKey, destKey, db, replace)
}

func (client *Client) CopyCtx(ctx context.Context, sourceKey string, destKey string, db int, replace bool) (int64, error) {
	return client.conn.Copy(ctx, sourceKey, destKey, db, replace).Result()
}

func (client *Client) GetBit(key string, offset int64) (int64, error) {
	return client.GetBitCtx(context.Background(), key, offset)
}

func (client *Client) GetBitCtx(ctx context.Context, key string, offset int64) (int64, error) {
	return client.conn.GetBit(ctx, key, offset).Result()
}

func (client *Client) SetBit(key string, offset int64, value int) (int64, error) {
	return client.SetBitCtx(context.Background(), key, offset, value)
}

func (client *Client) SetBitCtx(ctx context.Context, key string, offset int64, value int) (int64, error) {
	return client.conn.SetBit(ctx, key, offset, value).Result()
}

func (client *Client) BitCount(key string, bitCount *BitCount) (int64, error) {
	return client.BitCountCtx(context.Background(), key, bitCount)
}

func (client *Client) BitCountCtx(ctx context.Context, key string, bitCount *BitCount) (int64, error) {
	return client.conn.BitCount(ctx, key, bitCount).Result()
}

func (client *Client) BitOpAnd(destKey string, keys ...string) (int64, error) {
	return client.BitOpAndCtx(context.Background(), destKey, keys...)
}

func (client *Client) BitOpAndCtx(ctx context.Context, destKey string, keys ...string) (int64, error) {
	return client.conn.BitOpAnd(ctx, destKey, keys...).Result()
}

func (client *Client) BitOpOr(destKey string, keys ...string) (int64, error) {
	return client.BitOpOrCtx(context.Background(), destKey, keys...)
}

func (client *Client) BitOpOrCtx(ctx context.Context, destKey string, keys ...string) (int64, error) {
	return client.conn.BitOpOr(ctx, destKey, keys...).Result()
}

func (client *Client) BitOpXor(destKey string, keys ...string) (int64, error) {
	return client.BitOpXorCtx(context.Background(), destKey, keys...)
}

func (client *Client) BitOpXorCtx(ctx context.Context, destKey string, keys ...string) (int64, error) {
	return client.conn.BitOpXor(ctx, destKey, keys...).Result()
}

func (client *Client) BitOpNot(destKey string, key string) (int64, error) {
	return client.BitOpNotCtx(context.Background(), destKey, key)
}

func (client *Client) BitOpNotCtx(ctx context.Context, destKey string, key string) (int64, error) {
	return client.conn.BitOpNot(ctx, destKey, key).Result()
}

func (client *Client) BitPos(key string, bit int64, pos ...int64) (int64, error) {
	return client.BitPosCtx(context.Background(), key, bit, pos...)
}

func (client *Client) BitPosCtx(ctx context.Context, key string, bit int64, pos ...int64) (int64, error) {
	return client.conn.BitPos(ctx, key, bit, pos...).Result()
}

func (client *Client) BitField(key string, args ...interface{}) ([]int64, error) {
	return client.BitFieldCtx(context.Background(), key, args)
}

func (client *Client) BitFieldCtx(ctx context.Context, key string, args ...interface{}) ([]int64, error) {
	return client.conn.BitField(ctx, key, args).Result()
}

func (client *Client) Scan(cursor uint64, match string, count int64) (keys []string, cur uint64, err error) {
	return client.ScanCtx(context.Background(), cursor, match, count)
}

func (client *Client) ScanCtx(ctx context.Context, cursor uint64, match string, count int64) (keys []string, cur uint64, err error) {
	return client.conn.Scan(ctx, cursor, match, count).Result()
}

func (client *Client) ScanType(cursor uint64, match string, count int64, keyType string) (keys []string, cur uint64, err error) {
	return client.ScanTypeCtx(context.Background(), cursor, match, count, keyType)
}

func (client *Client) ScanTypeCtx(ctx context.Context, cursor uint64, match string, count int64, keyType string) (keys []string, cur uint64, err error) {
	return client.conn.ScanType(ctx, cursor, match, count, keyType).Result()
}

func (client *Client) SScan(key string, cursor uint64, match string, count int64) (keys []string, cur uint64, err error) {
	return client.SScanCtx(context.Background(), key, cursor, match, count)
}

func (client *Client) SScanCtx(ctx context.Context, key string, cursor uint64, match string, count int64) (keys []string, cur uint64, err error) {
	return client.conn.SScan(ctx, key, cursor, match, count).Result()
}

func (client *Client) HScan(key string, cursor uint64, match string, count int64) (keys []string, cur uint64, err error) {
	return client.HScanCtx(context.Background(), key, cursor, match, count)
}

func (client *Client) HScanCtx(ctx context.Context, key string, cursor uint64, match string, count int64) (keys []string, cur uint64, err error) {
	return client.conn.HScan(ctx, key, cursor, match, count).Result()
}

func (client *Client) ZScan(key string, cursor uint64, match string, count int64) (keys []string, cur uint64, err error) {
	return client.ZScanCtx(context.Background(), key, cursor, match, count)
}

func (client *Client) ZScanCtx(ctx context.Context, key string, cursor uint64, match string, count int64) (keys []string, cur uint64, err error) {
	return client.conn.ZScan(ctx, key, cursor, match, count).Result()
}

func (client *Client) HDel(key string, fields ...string) (int64, error) {
	return client.HDelCtx(context.Background(), key, fields...)
}

func (client *Client) HDelCtx(ctx context.Context, key string, fields ...string) (int64, error) {
	return client.conn.HDel(ctx, key, fields...).Result()
}

func (client *Client) HExists(key, field string) (bool, error) {
	return client.HExistsCtx(context.Background(), key, field)
}

func (client *Client) HExistsCtx(ctx context.Context, key, field string) (bool, error) {
	return client.conn.HExists(ctx, key, field).Result()
}

func (client *Client) HGet(key, field string) (string, error) {
	return client.HGetCtx(context.Background(), key, field)
}

func (client *Client) HGetCtx(ctx context.Context, key, field string) (string, error) {
	return client.conn.HGet(ctx, key, field).Result()
}

func (client *Client) HGetAll(key string) (map[string]string, error) {
	return client.HGetAllCtx(context.Background(), key)
}

func (client *Client) HGetAllCtx(ctx context.Context, key string) (map[string]string, error) {
	return client.conn.HGetAll(ctx, key).Result()
}

func (client *Client) HIncrBy(key, field string, incr int64) (int64, error) {
	return client.HIncrByCtx(context.Background(), key, field, incr)
}

func (client *Client) HIncrByCtx(ctx context.Context, key, field string, incr int64) (int64, error) {
	return client.conn.HIncrBy(ctx, key, field, incr).Result()
}

func (client *Client) HIncrByFloat(key, field string, incr float64) (float64, error) {
	return client.HIncrByFloatCtx(context.Background(), key, field, incr)
}

func (client *Client) HIncrByFloatCtx(ctx context.Context, key, field string, incr float64) (float64, error) {
	return client.conn.HIncrByFloat(ctx, key, field, incr).Result()
}

func (client *Client) HKeys(key string) ([]string, error) {
	return client.HKeysCtx(context.Background(), key)
}

func (client *Client) HKeysCtx(ctx context.Context, key string) ([]string, error) {
	return client.conn.HKeys(ctx, key).Result()
}

func (client *Client) HLen(key string) (int64, error) {
	return client.HLenCtx(context.Background(), key)
}

func (client *Client) HLenCtx(ctx context.Context, key string) (int64, error) {
	return client.conn.HLen(ctx, key).Result()
}

func (client *Client) HMGet(key string, fields ...string) ([]interface{}, error) {
	return client.HMGetCtx(context.Background(), key, fields...)
}

func (client *Client) HMGetCtx(ctx context.Context, key string, fields ...string) ([]interface{}, error) {
	return client.conn.HMGet(ctx, key, fields...).Result()
}

func (client *Client) HSet(key string, values ...interface{}) (int64, error) {
	return client.HSetCtx(context.Background(), key, values)
}

func (client *Client) HSetCtx(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return client.conn.HSet(ctx, key, values).Result()
}

func (client *Client) HMSet(key string, values ...interface{}) (bool, error) {
	return client.HMSetCtx(context.Background(), key, values)
}

func (client *Client) HMSetCtx(ctx context.Context, key string, values ...interface{}) (bool, error) {
	return client.conn.HMSet(ctx, key, values).Result()
}

func (client *Client) HSetNX(key, field string, value interface{}) (bool, error) {
	return client.HSetNXCtx(context.Background(), key, field, value)
}

func (client *Client) HSetNXCtx(ctx context.Context, key, field string, value interface{}) (bool, error) {
	return client.conn.HSetNX(ctx, key, field, value).Result()
}

func (client *Client) HVals(key string) ([]string, error) {
	return client.HValsCtx(context.Background(), key)
}

func (client *Client) HValsCtx(ctx context.Context, key string) ([]string, error) {
	return client.conn.HVals(ctx, key).Result()
}

func (client *Client) HRandField(key string, count int, withValues bool) ([]string, error) {
	return client.HRandFieldCtx(context.Background(), key, count, withValues)
}

func (client *Client) HRandFieldCtx(ctx context.Context, key string, count int, withValues bool) ([]string, error) {
	return client.conn.HRandField(ctx, key, count, withValues).Result()
}

func (client *Client) BLPop(timeout time.Duration, keys ...string) ([]string, error) {
	return client.BLPopCtx(context.Background(), timeout, keys...)
}

func (client *Client) BLPopCtx(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
	return client.conn.BLPop(ctx, timeout, keys...).Result()
}

func (client *Client) BRPop(timeout time.Duration, keys ...string) ([]string, error) {
	return client.BRPopCtx(context.Background(), timeout, keys...)
}

func (client *Client) BRPopCtx(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
	return client.conn.BRPop(ctx, timeout, keys...).Result()
}

func (client *Client) BRPopLPush(source, destination string, timeout time.Duration) (string, error) {
	return client.BRPopLPushCtx(context.Background(), source, destination, timeout)
}

func (client *Client) BRPopLPushCtx(ctx context.Context, source, destination string, timeout time.Duration) (string, error) {
	return client.conn.BRPopLPush(ctx, source, destination, timeout).Result()
}

func (client *Client) LIndex(key string, index int64) (string, error) {
	return client.LIndexCtx(context.Background(), key, index)
}

func (client *Client) LIndexCtx(ctx context.Context, key string, index int64) (string, error) {
	return client.conn.LIndex(ctx, key, index).Result()
}

func (client *Client) LInsert(key, op string, pivot, value interface{}) (int64, error) {
	return client.LInsertCtx(context.Background(), key, op, pivot, value)
}

func (client *Client) LInsertCtx(ctx context.Context, key, op string, pivot, value interface{}) (int64, error) {
	return client.conn.LInsert(ctx, key, op, pivot, value).Result()
}

func (client *Client) LInsertBefore(key string, pivot, value interface{}) (int64, error) {
	return client.LInsertBeforeCtx(context.Background(), key, pivot, value)
}

func (client *Client) LInsertBeforeCtx(ctx context.Context, key string, pivot, value interface{}) (int64, error) {
	return client.conn.LInsertBefore(ctx, key, pivot, value).Result()
}

func (client *Client) LInsertAfter(key string, pivot, value interface{}) (int64, error) {
	return client.LInsertAfterCtx(context.Background(), key, pivot, value)
}

func (client *Client) LInsertAfterCtx(ctx context.Context, key string, pivot, value interface{}) (int64, error) {
	return client.conn.LInsertAfter(ctx, key, pivot, value).Result()
}

func (client *Client) LLen(key string) (int64, error) {
	return client.LLenCtx(context.Background(), key)
}

func (client *Client) LLenCtx(ctx context.Context, key string) (int64, error) {
	return client.conn.LLen(ctx, key).Result()
}

func (client *Client) LPop(key string) (string, error) {
	return client.LPopCtx(context.Background(), key)
}

func (client *Client) LPopCtx(ctx context.Context, key string) (string, error) {
	return client.conn.LPop(ctx, key).Result()
}

func (client *Client) LPopCount(key string, count int) ([]string, error) {
	return client.LPopCountCtx(context.Background(), key, count)
}

func (client *Client) LPopCountCtx(ctx context.Context, key string, count int) ([]string, error) {
	return client.conn.LPopCount(ctx, key, count).Result()
}

func (client *Client) LPos(key string, value string, args LPosArgs) (int64, error) {
	return client.LPosCtx(context.Background(), key, value, args)
}

func (client *Client) LPosCtx(ctx context.Context, key string, value string, args LPosArgs) (int64, error) {
	return client.conn.LPos(ctx, key, value, args).Result()
}

func (client *Client) LPosCount(key string, value string, count int64, args LPosArgs) ([]int64, error) {
	return client.LPosCountCtx(context.Background(), key, value, count, args)
}

func (client *Client) LPosCountCtx(ctx context.Context, key string, value string, count int64, args LPosArgs) ([]int64, error) {
	return client.conn.LPosCount(ctx, key, value, count, args).Result()
}

func (client *Client) LPush(key string, values ...interface{}) (int64, error) {
	return client.LPushCtx(context.Background(), key, values)
}

func (client *Client) LPushCtx(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return client.conn.LPush(ctx, key, values).Result()
}

func (client *Client) LPushX(key string, values ...interface{}) (int64, error) {
	return client.LPushXCtx(context.Background(), key, values)
}

func (client *Client) LPushXCtx(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return client.conn.LPushX(ctx, key, values).Result()
}

func (client *Client) LRange(key string, start, stop int64) ([]string, error) {
	return client.LRangeCtx(context.Background(), key, start, stop)
}

func (client *Client) LRangeCtx(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return client.conn.LRange(ctx, key, start, stop).Result()
}

func (client *Client) LRem(key string, count int64, value interface{}) (int64, error) {
	return client.LRemCtx(context.Background(), key, count, value)
}

func (client *Client) LRemCtx(ctx context.Context, key string, count int64, value interface{}) (int64, error) {
	return client.conn.LRem(ctx, key, count, value).Result()
}

func (client *Client) LSet(key string, index int64, value interface{}) (string, error) {
	return client.LSetCtx(context.Background(), key, index, value)
}

func (client *Client) LSetCtx(ctx context.Context, key string, index int64, value interface{}) (string, error) {
	return client.conn.LSet(ctx, key, index, value).Result()
}

func (client *Client) LTrim(key string, start, stop int64) (string, error) {
	return client.LTrimCtx(context.Background(), key, start, stop)
}

func (client *Client) LTrimCtx(ctx context.Context, key string, start, stop int64) (string, error) {
	return client.conn.LTrim(ctx, key, start, stop).Result()
}

func (client *Client) RPop(key string) (string, error) {
	return client.RPopCtx(context.Background(), key)
}

func (client *Client) RPopCtx(ctx context.Context, key string) (string, error) {
	return client.conn.RPop(ctx, key).Result()
}

func (client *Client) RPopCount(key string, count int) ([]string, error) {
	return client.RPopCountCtx(context.Background(), key, count)
}

func (client *Client) RPopCountCtx(ctx context.Context, key string, count int) ([]string, error) {
	return client.conn.RPopCount(ctx, key, count).Result()
}

func (client *Client) RPopLPush(source, destination string) (string, error) {
	return client.RPopLPushCtx(context.Background(), source, destination)
}

func (client *Client) RPopLPushCtx(ctx context.Context, source, destination string) (string, error) {
	return client.conn.RPopLPush(ctx, source, destination).Result()
}

func (client *Client) RPush(key string, values ...interface{}) (int64, error) {
	return client.RPushCtx(context.Background(), key, values)
}

func (client *Client) RPushCtx(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return client.conn.RPush(ctx, key, values).Result()
}

func (client *Client) RPushX(key string, values ...interface{}) (int64, error) {
	return client.RPushXCtx(context.Background(), key, values)
}

func (client *Client) RPushXCtx(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return client.conn.RPushX(ctx, key, values).Result()
}

func (client *Client) LMove(source, destination, srcpos, destpos string) (string, error) {
	return client.LMoveCtx(context.Background(), source, destination, srcpos, destpos)
}

func (client *Client) LMoveCtx(ctx context.Context, source, destination, srcpos, destpos string) (string, error) {
	return client.conn.LMove(ctx, source, destination, srcpos, destpos).Result()
}

func (client *Client) BLMove(source, destination, srcpos, destpos string, timeout time.Duration) (string, error) {
	return client.BLMoveCtx(context.Background(), source, destination, srcpos, destpos, timeout)
}

func (client *Client) BLMoveCtx(ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration) (string, error) {
	return client.conn.BLMove(ctx, source, destination, srcpos, destpos, timeout).Result()
}

func (client *Client) SAdd(key string, members ...interface{}) (int64, error) {
	return client.SAddCtx(context.Background(), key, members)
}

func (client *Client) SAddCtx(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return client.conn.SAdd(ctx, key, members).Result()
}

func (client *Client) SCard(key string) (int64, error) {
	return client.SCardCtx(context.Background(), key)
}

func (client *Client) SCardCtx(ctx context.Context, key string) (int64, error) {
	return client.conn.SCard(ctx, key).Result()
}

func (client *Client) SDiff(keys ...string) ([]string, error) {
	return client.SDiffCtx(context.Background(), keys...)
}

func (client *Client) SDiffCtx(ctx context.Context, keys ...string) ([]string, error) {
	return client.conn.SDiff(ctx, keys...).Result()
}

func (client *Client) SDiffStore(destination string, keys ...string) (int64, error) {
	return client.SDiffStoreCtx(context.Background(), destination, keys...)
}

func (client *Client) SDiffStoreCtx(ctx context.Context, destination string, keys ...string) (int64, error) {
	return client.conn.SDiffStore(ctx, destination, keys...).Result()
}

func (client *Client) SInter(keys ...string) ([]string, error) {
	return client.SInterCtx(context.Background(), keys...)
}

func (client *Client) SInterCtx(ctx context.Context, keys ...string) ([]string, error) {
	return client.conn.SInter(ctx, keys...).Result()
}

func (client *Client) SInterStore(destination string, keys ...string) (int64, error) {
	return client.SInterStoreCtx(context.Background(), destination, keys...)
}

func (client *Client) SInterStoreCtx(ctx context.Context, destination string, keys ...string) (int64, error) {
	return client.conn.SInterStore(ctx, destination, keys...).Result()
}

func (client *Client) SIsMember(key string, member interface{}) (bool, error) {
	return client.SIsMemberCtx(context.Background(), key, member)
}

func (client *Client) SIsMemberCtx(ctx context.Context, key string, member interface{}) (bool, error) {
	return client.conn.SIsMember(ctx, key, member).Result()
}

func (client *Client) SMIsMember(key string, members ...interface{}) ([]bool, error) {
	return client.SMIsMemberCtx(context.Background(), key, members)
}

func (client *Client) SMIsMemberCtx(ctx context.Context, key string, members ...interface{}) ([]bool, error) {
	return client.conn.SMIsMember(ctx, key, members).Result()
}

func (client *Client) SMembers(key string) ([]string, error) {
	return client.SMembersCtx(context.Background(), key)
}

func (client *Client) SMembersCtx(ctx context.Context, key string) ([]string, error) {
	return client.conn.SMembers(ctx, key).Result()
}

func (client *Client) SMembersMap(key string) (map[string]struct{}, error) {
	return client.SMembersMapCtx(context.Background(), key)
}

func (client *Client) SMembersMapCtx(ctx context.Context, key string) (map[string]struct{}, error) {
	return client.conn.SMembersMap(ctx, key).Result()
}

func (client *Client) SMove(source, destination string, member interface{}) (bool, error) {
	return client.SMoveCtx(context.Background(), source, destination, member)
}

func (client *Client) SMoveCtx(ctx context.Context, source, destination string, member interface{}) (bool, error) {
	return client.conn.SMove(ctx, source, destination, member).Result()
}

func (client *Client) SPop(key string) (string, error) {
	return client.SPopCtx(context.Background(), key)
}

func (client *Client) SPopCtx(ctx context.Context, key string) (string, error) {
	return client.conn.SPop(ctx, key).Result()
}

func (client *Client) SPopN(key string, count int64) ([]string, error) {
	return client.SPopNCtx(context.Background(), key, count)
}

func (client *Client) SPopNCtx(ctx context.Context, key string, count int64) ([]string, error) {
	return client.conn.SPopN(ctx, key, count).Result()
}

func (client *Client) SRandMember(key string) (string, error) {
	return client.SRandMemberCtx(context.Background(), key)
}

func (client *Client) SRandMemberCtx(ctx context.Context, key string) (string, error) {
	return client.conn.SRandMember(ctx, key).Result()
}

func (client *Client) SRandMemberN(key string, count int64) ([]string, error) {
	return client.SRandMemberNCtx(context.Background(), key, count)
}

func (client *Client) SRandMemberNCtx(ctx context.Context, key string, count int64) ([]string, error) {
	return client.conn.SRandMemberN(ctx, key, count).Result()
}

func (client *Client) SRem(key string, members ...interface{}) (int64, error) {
	return client.SRemCtx(context.Background(), key, members)
}

func (client *Client) SRemCtx(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return client.conn.SRem(ctx, key, members).Result()
}

func (client *Client) SUnion(keys ...string) ([]string, error) {
	return client.SUnionCtx(context.Background(), keys...)
}

func (client *Client) SUnionCtx(ctx context.Context, keys ...string) ([]string, error) {
	return client.conn.SUnion(ctx, keys...).Result()
}

func (client *Client) SUnionStore(destination string, keys ...string) (int64, error) {
	return client.SUnionStoreCtx(context.Background(), destination, keys...)
}

func (client *Client) SUnionStoreCtx(ctx context.Context, destination string, keys ...string) (int64, error) {
	return client.conn.SUnionStore(ctx, destination, keys...).Result()
}

func (client *Client) ZAdd(key string, members ...*Z) (int64, error) {
	return client.ZAddCtx(context.Background(), key, members...)
}

func (client *Client) ZAddCtx(ctx context.Context, key string, members ...*Z) (int64, error) {
	return client.conn.ZAdd(ctx, key, members...).Result()
}

func (client *Client) ZAddNX(key string, members ...*Z) (int64, error) {
	return client.ZAddNXCtx(context.Background(), key, members...)
}

func (client *Client) ZAddNXCtx(ctx context.Context, key string, members ...*Z) (int64, error) {
	return client.conn.ZAddNX(ctx, key, members...).Result()
}

func (client *Client) ZAddXX(key string, members ...*Z) (int64, error) {
	return client.ZAddXXCtx(context.Background(), key, members...)
}

func (client *Client) ZAddXXCtx(ctx context.Context, key string, members ...*Z) (int64, error) {
	return client.conn.ZAddXX(ctx, key, members...).Result()
}

func (client *Client) ZAddCh(key string, members ...*Z) (int64, error) {
	return client.ZAddChCtx(context.Background(), key, members...)
}

func (client *Client) ZAddChCtx(ctx context.Context, key string, members ...*Z) (int64, error) {
	return client.conn.ZAddCh(ctx, key, members...).Result()
}

func (client *Client) ZAddNXCh(key string, members ...*Z) (int64, error) {
	return client.ZAddNXChCtx(context.Background(), key, members...)
}

func (client *Client) ZAddNXChCtx(ctx context.Context, key string, members ...*Z) (int64, error) {
	return client.conn.ZAddNXCh(ctx, key, members...).Result()
}

func (client *Client) ZAddXXCh(key string, members ...*Z) (int64, error) {
	return client.ZAddXXChCtx(context.Background(), key, members...)
}

func (client *Client) ZAddXXChCtx(ctx context.Context, key string, members ...*Z) (int64, error) {
	return client.conn.ZAddXXCh(ctx, key, members...).Result()
}

func (client *Client) ZAddArgs(key string, args ZAddArgs) (int64, error) {
	return client.ZAddArgsCtx(context.Background(), key, args)
}

func (client *Client) ZAddArgsCtx(ctx context.Context, key string, args ZAddArgs) (int64, error) {
	return client.conn.ZAddArgs(ctx, key, args).Result()
}

func (client *Client) ZAddArgsIncr(key string, args ZAddArgs) (float64, error) {
	return client.ZAddArgsIncrCtx(context.Background(), key, args)
}

func (client *Client) ZAddArgsIncrCtx(ctx context.Context, key string, args ZAddArgs) (float64, error) {
	return client.conn.ZAddArgsIncr(ctx, key, args).Result()
}

func (client *Client) ZIncr(key string, member *Z) (float64, error) {
	return client.ZIncrCtx(context.Background(), key, member)
}

func (client *Client) ZIncrCtx(ctx context.Context, key string, member *Z) (float64, error) {
	return client.conn.ZIncr(ctx, key, member).Result()
}

func (client *Client) ZIncrNX(key string, member *Z) (float64, error) {
	return client.ZIncrNXCtx(context.Background(), key, member)
}

func (client *Client) ZIncrNXCtx(ctx context.Context, key string, member *Z) (float64, error) {
	return client.conn.ZIncrNX(ctx, key, member).Result()
}

func (client *Client) ZIncrXX(key string, member *Z) (float64, error) {
	return client.ZIncrXXCtx(context.Background(), key, member)
}

func (client *Client) ZIncrXXCtx(ctx context.Context, key string, member *Z) (float64, error) {
	return client.conn.ZIncrXX(ctx, key, member).Result()
}

func (client *Client) ZCard(key string) (int64, error) {
	return client.ZCardCtx(context.Background(), key)
}

func (client *Client) ZCardCtx(ctx context.Context, key string) (int64, error) {
	return client.conn.ZCard(ctx, key).Result()
}

func (client *Client) ZCount(key, min, max string) (int64, error) {
	return client.ZCountCtx(context.Background(), key, min, max)
}

func (client *Client) ZCountCtx(ctx context.Context, key, min, max string) (int64, error) {
	return client.conn.ZCount(ctx, key, min, max).Result()
}

func (client *Client) ZLexCount(key, min, max string) (int64, error) {
	return client.ZLexCountCtx(context.Background(), key, min, max)
}

func (client *Client) ZLexCountCtx(ctx context.Context, key, min, max string) (int64, error) {
	return client.conn.ZLexCount(ctx, key, min, max).Result()
}

func (client *Client) ZIncrBy(key string, increment float64, member string) (float64, error) {
	return client.ZIncrByCtx(context.Background(), key, increment, member)
}

func (client *Client) ZIncrByCtx(ctx context.Context, key string, increment float64, member string) (float64, error) {
	return client.conn.ZIncrBy(ctx, key, increment, member).Result()
}

func (client *Client) ZInter(store *ZStore) ([]string, error) {
	return client.ZInterCtx(context.Background(), store)
}

func (client *Client) ZInterCtx(ctx context.Context, store *ZStore) ([]string, error) {
	return client.conn.ZInter(ctx, store).Result()
}

func (client *Client) ZInterWithScores(store *ZStore) ([]Z, error) {
	return client.ZInterWithScoresCtx(context.Background(), store)
}

func (client *Client) ZInterWithScoresCtx(ctx context.Context, store *ZStore) ([]Z, error) {
	return client.conn.ZInterWithScores(ctx, store).Result()
}

func (client *Client) ZInterStore(destination string, store *ZStore) (int64, error) {
	return client.ZInterStoreCtx(context.Background(), destination, store)
}

func (client *Client) ZInterStoreCtx(ctx context.Context, destination string, store *ZStore) (int64, error) {
	return client.conn.ZInterStore(ctx, destination, store).Result()
}

func (client *Client) ZMScore(key string, members ...string) ([]float64, error) {
	return client.ZMScoreCtx(context.Background(), key, members...)
}

func (client *Client) ZMScoreCtx(ctx context.Context, key string, members ...string) ([]float64, error) {
	return client.conn.ZMScore(ctx, key, members...).Result()
}

func (client *Client) ZPopMax(key string, count ...int64) ([]Z, error) {
	return client.ZPopMaxCtx(context.Background(), key, count...)
}

func (client *Client) ZPopMaxCtx(ctx context.Context, key string, count ...int64) ([]Z, error) {
	return client.conn.ZPopMax(ctx, key, count...).Result()
}

func (client *Client) ZPopMin(key string, count ...int64) ([]Z, error) {
	return client.ZPopMinCtx(context.Background(), key, count...)
}

func (client *Client) ZPopMinCtx(ctx context.Context, key string, count ...int64) ([]Z, error) {
	return client.conn.ZPopMin(ctx, key, count...).Result()
}

func (client *Client) ZRange(key string, start, stop int64) ([]string, error) {
	return client.ZRangeCtx(context.Background(), key, start, stop)
}

func (client *Client) ZRangeCtx(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return client.conn.ZRange(ctx, key, start, stop).Result()
}

func (client *Client) ZRangeWithScores(key string, start, stop int64) ([]Z, error) {
	return client.ZRangeWithScoresCtx(context.Background(), key, start, stop)
}

func (client *Client) ZRangeWithScoresCtx(ctx context.Context, key string, start, stop int64) ([]Z, error) {
	return client.conn.ZRangeWithScores(ctx, key, start, stop).Result()
}

func (client *Client) ZRangeByScore(key string, opt *ZRangeBy) ([]string, error) {
	return client.ZRangeByScoreCtx(context.Background(), key, opt)
}

func (client *Client) ZRangeByScoreCtx(ctx context.Context, key string, opt *ZRangeBy) ([]string, error) {
	return client.conn.ZRangeByScore(ctx, key, opt).Result()
}

func (client *Client) ZRangeByLex(key string, opt *ZRangeBy) ([]string, error) {
	return client.ZRangeByLexCtx(context.Background(), key, opt)
}

func (client *Client) ZRangeByLexCtx(ctx context.Context, key string, opt *ZRangeBy) ([]string, error) {
	return client.conn.ZRangeByLex(ctx, key, opt).Result()
}

func (client *Client) ZRangeByScoreWithScores(key string, opt *ZRangeBy) ([]Z, error) {
	return client.ZRangeByScoreWithScoresCtx(context.Background(), key, opt)
}

func (client *Client) ZRangeByScoreWithScoresCtx(ctx context.Context, key string, opt *ZRangeBy) ([]Z, error) {
	return client.conn.ZRangeByScoreWithScores(ctx, key, opt).Result()
}

func (client *Client) ZRangeArgs(z ZRangeArgs) ([]string, error) {
	return client.ZRangeArgsCtx(context.Background(), z)
}

func (client *Client) ZRangeArgsCtx(ctx context.Context, z ZRangeArgs) ([]string, error) {
	return client.conn.ZRangeArgs(ctx, z).Result()
}

func (client *Client) ZRangeArgsWithScores(z ZRangeArgs) ([]Z, error) {
	return client.ZRangeArgsWithScoresCtx(context.Background(), z)
}

func (client *Client) ZRangeArgsWithScoresCtx(ctx context.Context, z ZRangeArgs) ([]Z, error) {
	return client.conn.ZRangeArgsWithScores(ctx, z).Result()
}

func (client *Client) ZRangeStore(dst string, z ZRangeArgs) (int64, error) {
	return client.ZRangeStoreCtx(context.Background(), dst, z)
}

func (client *Client) ZRangeStoreCtx(ctx context.Context, dst string, z ZRangeArgs) (int64, error) {
	return client.conn.ZRangeStore(ctx, dst, z).Result()
}

func (client *Client) ZRank(key, member string) (int64, error) {
	return client.ZRankCtx(context.Background(), key, member)
}

func (client *Client) ZRankCtx(ctx context.Context, key, member string) (int64, error) {
	return client.conn.ZRank(ctx, key, member).Result()
}

func (client *Client) ZRem(key string, members ...interface{}) (int64, error) {
	return client.ZRemCtx(context.Background(), key, members)
}

func (client *Client) ZRemCtx(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return client.conn.ZRem(ctx, key, members).Result()
}

func (client *Client) ZRemRangeByRank(key string, start, stop int64) (int64, error) {
	return client.ZRemRangeByRankCtx(context.Background(), key, start, stop)
}

func (client *Client) ZRemRangeByRankCtx(ctx context.Context, key string, start, stop int64) (int64, error) {
	return client.conn.ZRemRangeByRank(ctx, key, start, stop).Result()
}

func (client *Client) ZRemRangeByScore(key, min, max string) (int64, error) {
	return client.ZRemRangeByScoreCtx(context.Background(), key, min, max)
}

func (client *Client) ZRemRangeByScoreCtx(ctx context.Context, key, min, max string) (int64, error) {
	return client.conn.ZRemRangeByScore(ctx, key, min, max).Result()
}

func (client *Client) ZRemRangeByLex(key, min, max string) (int64, error) {
	return client.ZRemRangeByLexCtx(context.Background(), key, min, max)
}

func (client *Client) ZRemRangeByLexCtx(ctx context.Context, key, min, max string) (int64, error) {
	return client.conn.ZRemRangeByLex(ctx, key, min, max).Result()
}

func (client *Client) ZRevRange(key string, start, stop int64) ([]string, error) {
	return client.ZRevRangeCtx(context.Background(), key, start, stop)
}

func (client *Client) ZRevRangeCtx(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return client.conn.ZRevRange(ctx, key, start, stop).Result()
}

func (client *Client) ZRevRangeWithScores(key string, start, stop int64) ([]Z, error) {
	return client.ZRevRangeWithScoresCtx(context.Background(), key, start, stop)
}

func (client *Client) ZRevRangeWithScoresCtx(ctx context.Context, key string, start, stop int64) ([]Z, error) {
	return client.conn.ZRevRangeWithScores(ctx, key, start, stop).Result()
}

func (client *Client) ZRevRangeByScore(key string, opt *ZRangeBy) ([]string, error) {
	return client.ZRevRangeByScoreCtx(context.Background(), key, opt)
}

func (client *Client) ZRevRangeByScoreCtx(ctx context.Context, key string, opt *ZRangeBy) ([]string, error) {
	return client.conn.ZRevRangeByScore(ctx, key, opt).Result()
}

func (client *Client) ZRevRangeByLex(key string, opt *ZRangeBy) ([]string, error) {
	return client.ZRevRangeByLexCtx(context.Background(), key, opt)
}

func (client *Client) ZRevRangeByLexCtx(ctx context.Context, key string, opt *ZRangeBy) ([]string, error) {
	return client.conn.ZRevRangeByLex(ctx, key, opt).Result()
}

func (client *Client) ZRevRangeByScoreWithScores(key string, opt *ZRangeBy) ([]Z, error) {
	return client.ZRevRangeByScoreWithScoresCtx(context.Background(), key, opt)
}

func (client *Client) ZRevRangeByScoreWithScoresCtx(ctx context.Context, key string, opt *ZRangeBy) ([]Z, error) {
	return client.conn.ZRevRangeByScoreWithScores(ctx, key, opt).Result()
}

func (client *Client) ZRevRank(key, member string) (int64, error) {
	return client.ZRevRankCtx(context.Background(), key, member)
}

func (client *Client) ZRevRankCtx(ctx context.Context, key, member string) (int64, error) {
	return client.conn.ZRevRank(ctx, key, member).Result()
}

func (client *Client) ZScore(key, member string) (float64, error) {
	return client.ZScoreCtx(context.Background(), key, member)
}

func (client *Client) ZScoreCtx(ctx context.Context, key, member string) (float64, error) {
	return client.conn.ZScore(ctx, key, member).Result()
}

func (client *Client) ZUnionStore(dest string, store *ZStore) (int64, error) {
	return client.ZUnionStoreCtx(context.Background(), dest, store)
}

func (client *Client) ZUnionStoreCtx(ctx context.Context, dest string, store *ZStore) (int64, error) {
	return client.conn.ZUnionStore(ctx, dest, store).Result()
}

func (client *Client) ZUnion(store ZStore) ([]string, error) {
	return client.ZUnionCtx(context.Background(), store)
}

func (client *Client) ZUnionCtx(ctx context.Context, store ZStore) ([]string, error) {
	return client.conn.ZUnion(ctx, store).Result()
}

func (client *Client) ZUnionWithScores(store ZStore) ([]Z, error) {
	return client.ZUnionWithScoresCtx(context.Background(), store)
}

func (client *Client) ZUnionWithScoresCtx(ctx context.Context, store ZStore) ([]Z, error) {
	return client.conn.ZUnionWithScores(ctx, store).Result()
}

func (client *Client) ZRandMember(key string, count int, withScores bool) ([]string, error) {
	return client.ZRandMemberCtx(context.Background(), key, count, withScores)
}

func (client *Client) ZRandMemberCtx(ctx context.Context, key string, count int, withScores bool) ([]string, error) {
	return client.conn.ZRandMember(ctx, key, count, withScores).Result()
}

func (client *Client) ZDiff(keys ...string) ([]string, error) {
	return client.ZDiffCtx(context.Background(), keys...)
}

func (client *Client) ZDiffCtx(ctx context.Context, keys ...string) ([]string, error) {
	return client.conn.ZDiff(ctx, keys...).Result()
}

func (client *Client) ZDiffWithScores(keys ...string) ([]Z, error) {
	return client.ZDiffWithScoresCtx(context.Background(), keys...)
}

func (client *Client) ZDiffWithScoresCtx(ctx context.Context, keys ...string) ([]Z, error) {
	return client.conn.ZDiffWithScores(ctx, keys...).Result()
}

func (client *Client) ZDiffStore(destination string, keys ...string) (int64, error) {
	return client.ZDiffStoreCtx(context.Background(), destination, keys...)
}

func (client *Client) ZDiffStoreCtx(ctx context.Context, destination string, keys ...string) (int64, error) {
	return client.conn.ZDiffStore(ctx, destination, keys...).Result()
}

func (client *Client) FlushAll() (string, error) {
	return client.FlushAllCtx(context.Background())
}

func (client *Client) FlushAllCtx(ctx context.Context) (string, error) {
	return client.conn.FlushAll(ctx).Result()
}

func (client *Client) FlushAllAsync() (string, error) {
	return client.FlushAllAsyncCtx(context.Background())
}

func (client *Client) FlushAllAsyncCtx(ctx context.Context) (string, error) {
	return client.conn.FlushAllAsync(ctx).Result()
}

func (client *Client) FlushDB() (string, error) {
	return client.FlushDBCtx(context.Background())
}

func (client *Client) FlushDBCtx(ctx context.Context) (string, error) {
	return client.conn.FlushDB(ctx).Result()
}

func (client *Client) FlushDBAsync() (string, error) {
	return client.FlushDBAsyncCtx(context.Background())
}

func (client *Client) FlushDBAsyncCtx(ctx context.Context) (string, error) {
	return client.conn.FlushDBAsync(ctx).Result()
}

func (client *Client) Info(section ...string) (string, error) {
	return client.InfoCtx(context.Background(), section...)
}

func (client *Client) InfoCtx(ctx context.Context, section ...string) (string, error) {
	return client.conn.Info(ctx, section...).Result()
}

func (client *Client) Save() (string, error) {
	return client.SaveCtx(context.Background())
}

func (client *Client) SaveCtx(ctx context.Context) (string, error) {
	return client.conn.Save(ctx).Result()
}

func (client *Client) Eval(script string, keys []string, args ...interface{}) (interface{}, error) {
	return client.EvalCtx(context.Background(), script, keys, args)
}

func (client *Client) EvalCtx(ctx context.Context, script string, keys []string, args ...interface{}) (interface{}, error) {
	return client.conn.Eval(ctx, script, keys, args).Result()
}

func (client *Client) EvalSha(sha1 string, keys []string, args ...interface{}) (interface{}, error) {
	return client.EvalShaCtx(context.Background(), sha1, keys, args)
}

func (client *Client) EvalShaCtx(ctx context.Context, sha1 string, keys []string, args ...interface{}) (interface{}, error) {
	return client.conn.EvalSha(ctx, sha1, keys, args).Result()
}

func (client *Client) ScriptExists(hashes ...string) ([]bool, error) {
	return client.ScriptExistsCtx(context.Background(), hashes...)
}

func (client *Client) ScriptExistsCtx(ctx context.Context, hashes ...string) ([]bool, error) {
	return client.conn.ScriptExists(ctx, hashes...).Result()
}
