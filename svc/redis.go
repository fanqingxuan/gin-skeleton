package svc

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type AWRedis struct {
	ctx    context.Context
	client *redis.Client
}

type Cmdable interface {
	Command() *redis.CommandsInfoCmd
	ClientGetName() *redis.StringCmd
	Echo(message interface{}) *redis.StringCmd
	Ping() *redis.StatusCmd
	Quit() *redis.StatusCmd
	Del(keys ...string) *redis.IntCmd
	Unlink(keys ...string) *redis.IntCmd
	Dump(key string) *redis.StringCmd
	Exists(keys ...string) *redis.IntCmd
	Expire(key string, expiration time.Duration) *redis.BoolCmd
	ExpireAt(key string, tm time.Time) *redis.BoolCmd
	ExpireNX(key string, expiration time.Duration) *redis.BoolCmd
	ExpireXX(key string, expiration time.Duration) *redis.BoolCmd
	ExpireGT(key string, expiration time.Duration) *redis.BoolCmd
	ExpireLT(key string, expiration time.Duration) *redis.BoolCmd
	Keys(pattern string) *redis.StringSliceCmd
	Migrate(host, port, key string, db int, timeout time.Duration) *redis.StatusCmd
	Move(key string, db int) *redis.BoolCmd
	ObjectRefCount(key string) *redis.IntCmd
	ObjectEncoding(key string) *redis.StringCmd
	ObjectIdleTime(key string) *redis.DurationCmd
	Persist(key string) *redis.BoolCmd
	PExpire(key string, expiration time.Duration) *redis.BoolCmd
	PExpireAt(key string, tm time.Time) *redis.BoolCmd
	PTTL(key string) *redis.DurationCmd
	RandomKey() *redis.StringCmd
	Rename(key, newkey string) *redis.StatusCmd
	RenameNX(key, newkey string) *redis.BoolCmd
	Restore(key string, ttl time.Duration, value string) *redis.StatusCmd
	RestoreReplace(key string, ttl time.Duration, value string) *redis.StatusCmd
	Touch(keys ...string) *redis.IntCmd
	TTL(key string) *redis.DurationCmd
	Type(key string) *redis.StatusCmd
	Append(key, value string) *redis.IntCmd
	Decr(key string) *redis.IntCmd
	DecrBy(key string, decrement int64) *redis.IntCmd
	Get(key string) *redis.StringCmd
	GetRange(key string, start, end int64) *redis.StringCmd
	GetSet(key string, value interface{}) *redis.StringCmd
	GetEx(key string, expiration time.Duration) *redis.StringCmd
	GetDel(key string) *redis.StringCmd
	Incr(key string) *redis.IntCmd
	IncrBy(key string, value int64) *redis.IntCmd
	IncrByFloat(key string, value float64) *redis.FloatCmd
	MGet(keys ...string) *redis.SliceCmd
	MSet(values ...interface{}) *redis.StatusCmd
	MSetNX(values ...interface{}) *redis.BoolCmd
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	SetArgs(key string, value interface{}, a redis.SetArgs) *redis.StatusCmd
	// TODO: rename to SetEx
	SetEX(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	SetNX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd
	SetXX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd
	SetRange(key string, offset int64, value string) *redis.IntCmd
	StrLen(key string) *redis.IntCmd
	Copy(sourceKey string, destKey string, db int, replace bool) *redis.IntCmd

	Scan(cursor uint64, match string, count int64) *redis.ScanCmd
	ScanType(cursor uint64, match string, count int64, keyType string) *redis.ScanCmd
	SScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd
	HScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd
	ZScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd

	HDel(key string, fields ...string) *redis.IntCmd
	HExists(key, field string) *redis.BoolCmd
	HGet(key, field string) *redis.StringCmd
	HGetAll(key string) *redis.StringStringMapCmd
	HIncrBy(key, field string, incr int64) *redis.IntCmd
	HIncrByFloat(key, field string, incr float64) *redis.FloatCmd
	HKeys(key string) *redis.StringSliceCmd
	HLen(key string) *redis.IntCmd
	HMGet(key string, fields ...string) *redis.SliceCmd
	HSet(key string, values ...interface{}) *redis.IntCmd
	HMSet(key string, values ...interface{}) *redis.BoolCmd
	HSetNX(key, field string, value interface{}) *redis.BoolCmd
	HVals(key string) *redis.StringSliceCmd
	HRandField(key string, count int, withValues bool) *redis.StringSliceCmd

	BLPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd
	BRPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd
	BRPopLPush(source, destination string, timeout time.Duration) *redis.StringCmd
	LIndex(key string, index int64) *redis.StringCmd
	LInsert(key, op string, pivot, value interface{}) *redis.IntCmd
	LInsertBefore(key string, pivot, value interface{}) *redis.IntCmd
	LInsertAfter(key string, pivot, value interface{}) *redis.IntCmd
	LLen(key string) *redis.IntCmd
	LPop(key string) *redis.StringCmd
	LPopCount(key string, count int) *redis.StringSliceCmd
	LPos(key string, value string, args redis.LPosArgs) *redis.IntCmd
	LPosCount(key string, value string, count int64, args redis.LPosArgs) *redis.IntSliceCmd
	LPush(key string, values ...interface{}) *redis.IntCmd
	LPushX(key string, values ...interface{}) *redis.IntCmd
	LRange(key string, start, stop int64) *redis.StringSliceCmd
	LRem(key string, count int64, value interface{}) *redis.IntCmd
	LSet(key string, index int64, value interface{}) *redis.StatusCmd
	LTrim(key string, start, stop int64) *redis.StatusCmd
	RPop(key string) *redis.StringCmd
	RPopCount(key string, count int) *redis.StringSliceCmd
	RPopLPush(source, destination string) *redis.StringCmd
	RPush(key string, values ...interface{}) *redis.IntCmd
	RPushX(key string, values ...interface{}) *redis.IntCmd
	LMove(source, destination, srcpos, destpos string) *redis.StringCmd
	BLMove(source, destination, srcpos, destpos string, timeout time.Duration) *redis.StringCmd

	SAdd(key string, members ...interface{}) *redis.IntCmd
	SCard(key string) *redis.IntCmd
	SDiff(keys ...string) *redis.StringSliceCmd
	SDiffStore(destination string, keys ...string) *redis.IntCmd
	SInter(keys ...string) *redis.StringSliceCmd
	SInterStore(destination string, keys ...string) *redis.IntCmd
	SIsMember(key string, member interface{}) *redis.BoolCmd
	SMIsMember(key string, members ...interface{}) *redis.BoolSliceCmd
	SMembers(key string) *redis.StringSliceCmd
	SMembersMap(key string) *redis.StringStructMapCmd
	SMove(source, destination string, member interface{}) *redis.BoolCmd
	SPop(key string) *redis.StringCmd
	SPopN(key string, count int64) *redis.StringSliceCmd
	SRandMember(key string) *redis.StringCmd
	SRandMemberN(key string, count int64) *redis.StringSliceCmd
	SRem(key string, members ...interface{}) *redis.IntCmd
	SUnion(keys ...string) *redis.StringSliceCmd
	SUnionStore(destination string, keys ...string) *redis.IntCmd

	// TODO: remove
	//		ZAddCh
	//		ZIncr
	//		ZAddNXCh
	//		ZAddXXCh
	//		ZIncrNX
	//		ZIncrXX
	// 	in v9.
	// 	use ZAddArgs and ZAddArgsIncr.

	ZAdd(key string, members ...*redis.Z) *redis.IntCmd
	ZAddNX(key string, members ...*redis.Z) *redis.IntCmd
	ZAddXX(key string, members ...*redis.Z) *redis.IntCmd
	ZAddCh(key string, members ...*redis.Z) *redis.IntCmd
	ZAddNXCh(key string, members ...*redis.Z) *redis.IntCmd
	ZAddXXCh(key string, members ...*redis.Z) *redis.IntCmd
	ZAddArgs(key string, args redis.ZAddArgs) *redis.IntCmd
	ZAddArgsIncr(key string, args redis.ZAddArgs) *redis.FloatCmd
	ZIncr(key string, member *redis.Z) *redis.FloatCmd
	ZIncrNX(key string, member *redis.Z) *redis.FloatCmd
	ZIncrXX(key string, member *redis.Z) *redis.FloatCmd
	ZCard(key string) *redis.IntCmd
	ZCount(key, min, max string) *redis.IntCmd
	ZLexCount(key, min, max string) *redis.IntCmd
	ZIncrBy(key string, increment float64, member string) *redis.FloatCmd
	ZInter(store *redis.ZStore) *redis.StringSliceCmd
	ZInterWithScores(store *redis.ZStore) *redis.ZSliceCmd
	ZInterStore(destination string, store *redis.ZStore) *redis.IntCmd
	ZMScore(key string, members ...string) *redis.FloatSliceCmd
	ZPopMax(key string, count ...int64) *redis.ZSliceCmd
	ZPopMin(key string, count ...int64) *redis.ZSliceCmd
	ZRange(key string, start, stop int64) *redis.StringSliceCmd
	ZRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd
	ZRangeByScore(key string, opt *redis.ZRangeBy) *redis.StringSliceCmd
	ZRangeByLex(key string, opt *redis.ZRangeBy) *redis.StringSliceCmd
	ZRangeByScoreWithScores(key string, opt *redis.ZRangeBy) *redis.ZSliceCmd
	ZRangeArgs(z redis.ZRangeArgs) *redis.StringSliceCmd
	ZRangeArgsWithScores(z redis.ZRangeArgs) *redis.ZSliceCmd
	ZRangeStore(dst string, z redis.ZRangeArgs) *redis.IntCmd
	ZRank(key, member string) *redis.IntCmd
	ZRem(key string, members ...interface{}) *redis.IntCmd
	ZRemRangeByRank(key string, start, stop int64) *redis.IntCmd
	ZRemRangeByScore(key, min, max string) *redis.IntCmd
	ZRemRangeByLex(key, min, max string) *redis.IntCmd
	ZRevRange(key string, start, stop int64) *redis.StringSliceCmd
	ZRevRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd
	ZRevRangeByScore(key string, opt *redis.ZRangeBy) *redis.StringSliceCmd
	ZRevRangeByLex(key string, opt *redis.ZRangeBy) *redis.StringSliceCmd
	ZRevRangeByScoreWithScores(key string, opt *redis.ZRangeBy) *redis.ZSliceCmd
	ZRevRank(key, member string) *redis.IntCmd
	ZScore(key, member string) *redis.FloatCmd
	ZUnionStore(dest string, store *redis.ZStore) *redis.IntCmd
	ZUnion(store redis.ZStore) *redis.StringSliceCmd
	ZUnionWithScores(store redis.ZStore) *redis.ZSliceCmd
	ZRandMember(key string, count int, withScores bool) *redis.StringSliceCmd
	ZDiff(keys ...string) *redis.StringSliceCmd
	ZDiffWithScores(keys ...string) *redis.ZSliceCmd
	ZDiffStore(destination string, keys ...string) *redis.IntCmd

	BgRewriteAOF() *redis.StatusCmd
	BgSave() *redis.StatusCmd
	ClientKill(ipPort string) *redis.StatusCmd
	ClientKillByFilter(keys ...string) *redis.IntCmd
	ClientList() *redis.StringCmd
	ClientPause(dur time.Duration) *redis.BoolCmd
	ClientID() *redis.IntCmd
	ConfigGet(parameter string) *redis.SliceCmd
	ConfigResetStat() *redis.StatusCmd
	ConfigSet(parameter, value string) *redis.StatusCmd
	ConfigRewrite() *redis.StatusCmd
	DBSize() *redis.IntCmd
	FlushAll() *redis.StatusCmd
	FlushAllAsync() *redis.StatusCmd
	FlushDB() *redis.StatusCmd
	FlushDBAsync() *redis.StatusCmd
	Info(section ...string) *redis.StringCmd
	LastSave() *redis.IntCmd
	Save() *redis.StatusCmd
	Shutdown() *redis.StatusCmd
	ShutdownSave() *redis.StatusCmd
	ShutdownNoSave() *redis.StatusCmd
	SlaveOf(host, port string) *redis.StatusCmd
	Time() *redis.TimeCmd
	DebugObject(key string) *redis.StringCmd
	ReadOnly() *redis.StatusCmd
	ReadWrite() *redis.StatusCmd
	MemoryUsage(key string, samples ...int) *redis.IntCmd

	Eval(script string, keys []string, args ...interface{}) *redis.Cmd
	EvalSha(sha1 string, keys []string, args ...interface{}) *redis.Cmd
	ScriptExists(hashes ...string) *redis.BoolSliceCmd
	ScriptFlush() *redis.StatusCmd
	ScriptKill() *redis.StatusCmd
	ScriptLoad(script string) *redis.StringCmd
}

var (
	_ Cmdable = (*AWRedis)(nil)
)

func NewRedis(ctx context.Context, client *redis.Client) *AWRedis {
	var redis *AWRedis
	redis = &AWRedis{
		ctx:    ctx,
		client: client,
	}
	client.AddHook(&CheckServerStatusHook{})
	return redis
}

type CheckServerStatusHook struct {
	redis.Hook
}

func (h *CheckServerStatusHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	return ctx, nil
}
func (h *CheckServerStatusHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	err := cmd.Err()
	if err != nil && err != redis.Nil {
		panic(err.Error())

	}
	return err
}

func (h *CheckServerStatusHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	return ctx, nil
}
func (h *CheckServerStatusHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	return nil
}

func (that *AWRedis) Command() *redis.CommandsInfoCmd {
	return that.client.Command(that.ctx)
}

func (that *AWRedis) ClientGetName() *redis.StringCmd {
	return that.client.ClientGetName(that.ctx)
}

func (that *AWRedis) Echo(message interface{}) *redis.StringCmd {
	return that.client.Echo(that.ctx, message)
}

func (that *AWRedis) Ping() *redis.StatusCmd {
	return that.client.Ping(that.ctx)
}

func (that *AWRedis) Quit() *redis.StatusCmd {
	return that.client.Quit(that.ctx)
}

func (that *AWRedis) Del(keys ...string) *redis.IntCmd {
	return that.client.Del(that.ctx, keys...)
}

func (that *AWRedis) Unlink(keys ...string) *redis.IntCmd {
	return that.client.Unlink(that.ctx, keys...)
}

func (that *AWRedis) Dump(key string) *redis.StringCmd {
	return that.client.Dump(that.ctx, key)
}

func (that *AWRedis) Exists(keys ...string) *redis.IntCmd {
	return that.client.Exists(that.ctx, keys...)
}

func (that *AWRedis) Expire(key string, expiration time.Duration) *redis.BoolCmd {
	return that.client.Expire(that.ctx, key, expiration)
}

func (that *AWRedis) ExpireAt(key string, tm time.Time) *redis.BoolCmd {
	return that.client.ExpireAt(that.ctx, key, tm)
}

func (that *AWRedis) ExpireNX(key string, expiration time.Duration) *redis.BoolCmd {
	return that.client.ExpireNX(that.ctx, key, expiration)
}

func (that *AWRedis) ExpireXX(key string, expiration time.Duration) *redis.BoolCmd {
	return that.client.ExpireXX(that.ctx, key, expiration)
}

func (that *AWRedis) ExpireGT(key string, expiration time.Duration) *redis.BoolCmd {
	return that.client.ExpireGT(that.ctx, key, expiration)
}

func (that *AWRedis) ExpireLT(key string, expiration time.Duration) *redis.BoolCmd {
	return that.client.ExpireLT(that.ctx, key, expiration)
}

func (that *AWRedis) Keys(pattern string) *redis.StringSliceCmd {
	return that.client.Keys(that.ctx, pattern)
}

func (that *AWRedis) Migrate(host, port, key string, db int, timeout time.Duration) *redis.StatusCmd {
	return that.client.Migrate(that.ctx, host, port, key, db, timeout)
}

func (that *AWRedis) Move(key string, db int) *redis.BoolCmd {
	return that.client.Move(that.ctx, key, db)
}

func (that *AWRedis) ObjectRefCount(key string) *redis.IntCmd {
	return that.client.ObjectRefCount(that.ctx, key)
}

func (that *AWRedis) ObjectEncoding(key string) *redis.StringCmd {
	return that.client.ObjectEncoding(that.ctx, key)
}

func (that *AWRedis) ObjectIdleTime(key string) *redis.DurationCmd {
	return that.client.ObjectIdleTime(that.ctx, key)
}

func (that *AWRedis) Persist(key string) *redis.BoolCmd {
	return that.client.Persist(that.ctx, key)
}

func (that *AWRedis) PExpire(key string, expiration time.Duration) *redis.BoolCmd {
	return that.client.PExpire(that.ctx, key, expiration)
}

func (that *AWRedis) PExpireAt(key string, tm time.Time) *redis.BoolCmd {
	return that.client.PExpireAt(that.ctx, key, tm)
}

func (that *AWRedis) PTTL(key string) *redis.DurationCmd {
	return that.client.PTTL(that.ctx, key)
}

func (that *AWRedis) RandomKey() *redis.StringCmd {
	return that.client.RandomKey(that.ctx)
}

func (that *AWRedis) Rename(key, newkey string) *redis.StatusCmd {
	return that.client.Rename(that.ctx, key, newkey)
}

func (that *AWRedis) RenameNX(key, newkey string) *redis.BoolCmd {
	return that.client.RenameNX(that.ctx, key, newkey)
}

func (that *AWRedis) Restore(key string, ttl time.Duration, value string) *redis.StatusCmd {
	return that.client.Restore(that.ctx, key, ttl, value)
}

func (that *AWRedis) RestoreReplace(key string, ttl time.Duration, value string) *redis.StatusCmd {
	return that.client.RestoreReplace(that.ctx, key, ttl, value)
}

func (that *AWRedis) Touch(keys ...string) *redis.IntCmd {
	return that.client.Touch(that.ctx, keys...)
}

func (that *AWRedis) TTL(key string) *redis.DurationCmd {
	return that.client.TTL(that.ctx, key)
}

func (that *AWRedis) Type(key string) *redis.StatusCmd {
	return that.client.Type(that.ctx, key)
}

func (that *AWRedis) Append(key, value string) *redis.IntCmd {
	return that.client.Append(that.ctx, key, value)
}

func (that *AWRedis) Decr(key string) *redis.IntCmd {
	return that.client.Decr(that.ctx, key)
}

func (that *AWRedis) DecrBy(key string, decrement int64) *redis.IntCmd {
	return that.client.DecrBy(that.ctx, key, decrement)
}

func (that *AWRedis) Get(key string) *redis.StringCmd {
	return that.client.Get(that.ctx, key)
}

func (that *AWRedis) GetRange(key string, start, end int64) *redis.StringCmd {
	return that.client.GetRange(that.ctx, key, start, end)
}

func (that *AWRedis) GetSet(key string, value interface{}) *redis.StringCmd {
	return that.client.GetSet(that.ctx, key, value)
}

func (that *AWRedis) GetEx(key string, expiration time.Duration) *redis.StringCmd {
	return that.client.GetEx(that.ctx, key, expiration)
}

func (that *AWRedis) GetDel(key string) *redis.StringCmd {
	return that.client.GetDel(that.ctx, key)
}

func (that *AWRedis) Incr(key string) *redis.IntCmd {
	return that.client.Incr(that.ctx, key)
}

func (that *AWRedis) IncrBy(key string, value int64) *redis.IntCmd {
	return that.client.IncrBy(that.ctx, key, value)
}

func (that *AWRedis) IncrByFloat(key string, value float64) *redis.FloatCmd {
	return that.client.IncrByFloat(that.ctx, key, value)
}

func (that *AWRedis) MGet(keys ...string) *redis.SliceCmd {
	return that.client.MGet(that.ctx, keys...)
}

func (that *AWRedis) MSet(values ...interface{}) *redis.StatusCmd {
	return that.client.MSet(that.ctx, values...)
}

func (that *AWRedis) MSetNX(values ...interface{}) *redis.BoolCmd {
	return that.client.MSetNX(that.ctx, values...)
}

func (that *AWRedis) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return that.client.Set(that.ctx, key, value, expiration)
}

func (that *AWRedis) SetArgs(key string, value interface{}, a redis.SetArgs) *redis.StatusCmd {
	return that.client.SetArgs(that.ctx, key, value, a)
}

func (that *AWRedis) SetEX(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return that.client.SetEX(that.ctx, key, value, expiration)
}

func (that *AWRedis) SetNX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	return that.client.SetNX(that.ctx, key, value, expiration)
}

func (that *AWRedis) SetXX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	return that.client.SetXX(that.ctx, key, value, expiration)
}

func (that *AWRedis) SetRange(key string, offset int64, value string) *redis.IntCmd {
	return that.client.SetRange(that.ctx, key, offset, value)
}

func (that *AWRedis) StrLen(key string) *redis.IntCmd {
	return that.client.StrLen(that.ctx, key)
}

func (that *AWRedis) Copy(sourceKey string, destKey string, db int, replace bool) *redis.IntCmd {
	return that.client.Copy(that.ctx, sourceKey, destKey, db, replace)
}

func (that *AWRedis) Scan(cursor uint64, match string, count int64) *redis.ScanCmd {
	return that.client.Scan(that.ctx, cursor, match, count)
}

func (that *AWRedis) ScanType(cursor uint64, match string, count int64, keyType string) *redis.ScanCmd {
	return that.client.ScanType(that.ctx, cursor, match, count, keyType)
}

func (that *AWRedis) SScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return that.client.SScan(that.ctx, key, cursor, match, count)
}

func (that *AWRedis) HScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return that.client.HScan(that.ctx, key, cursor, match, count)
}

func (that *AWRedis) ZScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return that.client.ZScan(that.ctx, key, cursor, match, count)
}

func (that *AWRedis) HDel(key string, fields ...string) *redis.IntCmd {
	return that.client.HDel(that.ctx, key, fields...)
}

func (that *AWRedis) HExists(key, field string) *redis.BoolCmd {
	return that.client.HExists(that.ctx, key, field)
}

func (that *AWRedis) HGet(key, field string) *redis.StringCmd {
	return that.client.HGet(that.ctx, key, field)
}

func (that *AWRedis) HGetAll(key string) *redis.StringStringMapCmd {
	return that.client.HGetAll(that.ctx, key)
}

func (that *AWRedis) HIncrBy(key, field string, incr int64) *redis.IntCmd {
	return that.client.HIncrBy(that.ctx, key, field, incr)
}

func (that *AWRedis) HIncrByFloat(key, field string, incr float64) *redis.FloatCmd {
	return that.client.HIncrByFloat(that.ctx, key, field, incr)
}

func (that *AWRedis) HKeys(key string) *redis.StringSliceCmd {
	return that.client.HKeys(that.ctx, key)
}

func (that *AWRedis) HLen(key string) *redis.IntCmd {
	return that.client.HLen(that.ctx, key)
}

func (that *AWRedis) HMGet(key string, fields ...string) *redis.SliceCmd {
	return that.client.HMGet(that.ctx, key, fields...)
}

func (that *AWRedis) HSet(key string, values ...interface{}) *redis.IntCmd {
	return that.client.HSet(that.ctx, key, values...)
}

func (that *AWRedis) HMSet(key string, values ...interface{}) *redis.BoolCmd {
	return that.client.HMSet(that.ctx, key, values...)
}

func (that *AWRedis) HSetNX(key, field string, value interface{}) *redis.BoolCmd {
	return that.client.HSetNX(that.ctx, key, field, value)
}

func (that *AWRedis) HVals(key string) *redis.StringSliceCmd {
	return that.client.HVals(that.ctx, key)
}

func (that *AWRedis) HRandField(key string, count int, withValues bool) *redis.StringSliceCmd {
	return that.client.HRandField(that.ctx, key, count, withValues)
}

func (that *AWRedis) BLPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	return that.client.BLPop(that.ctx, timeout, keys...)
}

func (that *AWRedis) BRPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	return that.client.BRPop(that.ctx, timeout, keys...)
}

func (that *AWRedis) BRPopLPush(source, destination string, timeout time.Duration) *redis.StringCmd {
	return that.client.BRPopLPush(that.ctx, source, destination, timeout)
}

func (that *AWRedis) LIndex(key string, index int64) *redis.StringCmd {
	return that.client.LIndex(that.ctx, key, index)
}

func (that *AWRedis) LInsert(key, op string, pivot, value interface{}) *redis.IntCmd {
	return that.client.LInsert(that.ctx, key, op, pivot, value)
}

func (that *AWRedis) LInsertBefore(key string, pivot, value interface{}) *redis.IntCmd {
	return that.client.LInsertBefore(that.ctx, key, pivot, value)
}

func (that *AWRedis) LInsertAfter(key string, pivot, value interface{}) *redis.IntCmd {
	return that.client.LInsertAfter(that.ctx, key, pivot, value)
}

func (that *AWRedis) LLen(key string) *redis.IntCmd {
	return that.client.LLen(that.ctx, key)
}

func (that *AWRedis) LPop(key string) *redis.StringCmd {
	return that.client.LPop(that.ctx, key)
}

func (that *AWRedis) LPopCount(key string, count int) *redis.StringSliceCmd {
	return that.client.LPopCount(that.ctx, key, count)
}

func (that *AWRedis) LPos(key string, value string, args redis.LPosArgs) *redis.IntCmd {
	return that.client.LPos(that.ctx, key, value, args)
}

func (that *AWRedis) LPosCount(key string, value string, count int64, args redis.LPosArgs) *redis.IntSliceCmd {
	return that.client.LPosCount(that.ctx, key, value, count, args)
}

func (that *AWRedis) LPush(key string, values ...interface{}) *redis.IntCmd {
	return that.client.LPush(that.ctx, key, values...)
}

func (that *AWRedis) LPushX(key string, values ...interface{}) *redis.IntCmd {
	return that.client.LPushX(that.ctx, key, values...)
}

func (that *AWRedis) LRange(key string, start, stop int64) *redis.StringSliceCmd {
	return that.client.LRange(that.ctx, key, start, stop)
}

func (that *AWRedis) LRem(key string, count int64, value interface{}) *redis.IntCmd {
	return that.client.LRem(that.ctx, key, count, value)
}

func (that *AWRedis) LSet(key string, index int64, value interface{}) *redis.StatusCmd {
	return that.client.LSet(that.ctx, key, index, value)
}

func (that *AWRedis) LTrim(key string, start, stop int64) *redis.StatusCmd {
	return that.client.LTrim(that.ctx, key, start, stop)
}

func (that *AWRedis) RPop(key string) *redis.StringCmd {
	return that.client.RPop(that.ctx, key)
}

func (that *AWRedis) RPopCount(key string, count int) *redis.StringSliceCmd {
	return that.client.RPopCount(that.ctx, key, count)
}

func (that *AWRedis) RPopLPush(source, destination string) *redis.StringCmd {
	return that.client.RPopLPush(that.ctx, source, destination)
}

func (that *AWRedis) RPush(key string, values ...interface{}) *redis.IntCmd {
	return that.client.RPush(that.ctx, key, values...)
}

func (that *AWRedis) RPushX(key string, values ...interface{}) *redis.IntCmd {
	return that.client.RPushX(that.ctx, key, values...)
}

func (that *AWRedis) LMove(source, destination, srcpos, destpos string) *redis.StringCmd {
	return that.client.LMove(that.ctx, source, destination, srcpos, destpos)
}

func (that *AWRedis) BLMove(source, destination, srcpos, destpos string, timeout time.Duration) *redis.StringCmd {
	return that.client.BLMove(that.ctx, source, destination, srcpos, destpos, timeout)
}

func (that *AWRedis) SAdd(key string, members ...interface{}) *redis.IntCmd {
	return that.client.SAdd(that.ctx, key, members...)
}

func (that *AWRedis) SCard(key string) *redis.IntCmd {
	return that.client.SCard(that.ctx, key)
}

func (that *AWRedis) SDiff(keys ...string) *redis.StringSliceCmd {
	return that.client.SDiff(that.ctx, keys...)
}

func (that *AWRedis) SDiffStore(destination string, keys ...string) *redis.IntCmd {
	return that.client.SDiffStore(that.ctx, destination, keys...)
}

func (that *AWRedis) SInter(keys ...string) *redis.StringSliceCmd {
	return that.client.SInter(that.ctx, keys...)
}

func (that *AWRedis) SInterStore(destination string, keys ...string) *redis.IntCmd {
	return that.client.SInterStore(that.ctx, destination, keys...)
}

func (that *AWRedis) SIsMember(key string, member interface{}) *redis.BoolCmd {
	return that.client.SIsMember(that.ctx, key, member)
}

func (that *AWRedis) SMIsMember(key string, members ...interface{}) *redis.BoolSliceCmd {
	return that.client.SMIsMember(that.ctx, key, members...)
}

func (that *AWRedis) SMembers(key string) *redis.StringSliceCmd {
	return that.client.SMembers(that.ctx, key)
}

func (that *AWRedis) SMembersMap(key string) *redis.StringStructMapCmd {
	return that.client.SMembersMap(that.ctx, key)
}

func (that *AWRedis) SMove(source, destination string, member interface{}) *redis.BoolCmd {
	return that.client.SMove(that.ctx, source, destination, member)
}

func (that *AWRedis) SPop(key string) *redis.StringCmd {
	return that.client.SPop(that.ctx, key)
}

func (that *AWRedis) SPopN(key string, count int64) *redis.StringSliceCmd {
	return that.client.SPopN(that.ctx, key, count)
}

func (that *AWRedis) SRandMember(key string) *redis.StringCmd {
	return that.client.SRandMember(that.ctx, key)
}

func (that *AWRedis) SRandMemberN(key string, count int64) *redis.StringSliceCmd {
	return that.client.SRandMemberN(that.ctx, key, count)
}

func (that *AWRedis) SRem(key string, members ...interface{}) *redis.IntCmd {
	return that.client.SRem(that.ctx, key, members...)
}

func (that *AWRedis) SUnion(keys ...string) *redis.StringSliceCmd {
	return that.client.SUnion(that.ctx, keys...)
}

func (that *AWRedis) SUnionStore(destination string, keys ...string) *redis.IntCmd {
	return that.client.SUnionStore(that.ctx, destination, keys...)
}

func (that *AWRedis) ZAdd(key string, members ...*redis.Z) *redis.IntCmd {
	return that.client.ZAdd(that.ctx, key, members...)
}

func (that *AWRedis) ZAddNX(key string, members ...*redis.Z) *redis.IntCmd {
	return that.client.ZAddNX(that.ctx, key, members...)
}

func (that *AWRedis) ZAddXX(key string, members ...*redis.Z) *redis.IntCmd {
	return that.client.ZAddXX(that.ctx, key, members...)
}

func (that *AWRedis) ZAddCh(key string, members ...*redis.Z) *redis.IntCmd {
	return that.client.ZAddCh(that.ctx, key, members...)
}

func (that *AWRedis) ZAddNXCh(key string, members ...*redis.Z) *redis.IntCmd {
	return that.client.ZAddNXCh(that.ctx, key, members...)
}

func (that *AWRedis) ZAddXXCh(key string, members ...*redis.Z) *redis.IntCmd {
	return that.client.ZAddXXCh(that.ctx, key, members...)
}

func (that *AWRedis) ZAddArgs(key string, args redis.ZAddArgs) *redis.IntCmd {
	return that.client.ZAddArgs(that.ctx, key, args)
}

func (that *AWRedis) ZAddArgsIncr(key string, args redis.ZAddArgs) *redis.FloatCmd {
	return that.client.ZAddArgsIncr(that.ctx, key, args)
}

func (that *AWRedis) ZIncr(key string, member *redis.Z) *redis.FloatCmd {
	return that.client.ZIncr(that.ctx, key, member)
}

func (that *AWRedis) ZIncrNX(key string, member *redis.Z) *redis.FloatCmd {
	return that.client.ZIncrNX(that.ctx, key, member)
}

func (that *AWRedis) ZIncrXX(key string, member *redis.Z) *redis.FloatCmd {
	return that.client.ZIncrXX(that.ctx, key, member)
}

func (that *AWRedis) ZCard(key string) *redis.IntCmd {
	return that.client.ZCard(that.ctx, key)
}

func (that *AWRedis) ZCount(key, min, max string) *redis.IntCmd {
	return that.client.ZCount(that.ctx, key, min, max)
}

func (that *AWRedis) ZLexCount(key, min, max string) *redis.IntCmd {
	return that.client.ZLexCount(that.ctx, key, min, max)
}

func (that *AWRedis) ZIncrBy(key string, increment float64, member string) *redis.FloatCmd {
	return that.client.ZIncrBy(that.ctx, key, increment, member)
}

func (that *AWRedis) ZInter(store *redis.ZStore) *redis.StringSliceCmd {
	return that.client.ZInter(that.ctx, store)
}

func (that *AWRedis) ZInterWithScores(store *redis.ZStore) *redis.ZSliceCmd {
	return that.client.ZInterWithScores(that.ctx, store)
}

func (that *AWRedis) ZInterStore(destination string, store *redis.ZStore) *redis.IntCmd {
	return that.client.ZInterStore(that.ctx, destination, store)
}

func (that *AWRedis) ZMScore(key string, members ...string) *redis.FloatSliceCmd {
	return that.client.ZMScore(that.ctx, key, members...)
}

func (that *AWRedis) ZPopMax(key string, count ...int64) *redis.ZSliceCmd {
	return that.client.ZPopMax(that.ctx, key, count...)
}

func (that *AWRedis) ZPopMin(key string, count ...int64) *redis.ZSliceCmd {
	return that.client.ZPopMin(that.ctx, key, count...)
}

func (that *AWRedis) ZRange(key string, start, stop int64) *redis.StringSliceCmd {
	return that.client.ZRange(that.ctx, key, start, stop)
}

func (that *AWRedis) ZRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd {
	return that.client.ZRangeWithScores(that.ctx, key, start, stop)
}

func (that *AWRedis) ZRangeByScore(key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	return that.client.ZRangeByScore(that.ctx, key, opt)
}

func (that *AWRedis) ZRangeByLex(key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	return that.client.ZRangeByLex(that.ctx, key, opt)
}

func (that *AWRedis) ZRangeByScoreWithScores(key string, opt *redis.ZRangeBy) *redis.ZSliceCmd {
	return that.client.ZRangeByScoreWithScores(that.ctx, key, opt)
}

func (that *AWRedis) ZRangeArgs(z redis.ZRangeArgs) *redis.StringSliceCmd {
	return that.client.ZRangeArgs(that.ctx, z)
}

func (that *AWRedis) ZRangeArgsWithScores(z redis.ZRangeArgs) *redis.ZSliceCmd {
	return that.client.ZRangeArgsWithScores(that.ctx, z)
}

func (that *AWRedis) ZRangeStore(dst string, z redis.ZRangeArgs) *redis.IntCmd {
	return that.client.ZRangeStore(that.ctx, dst, z)
}

func (that *AWRedis) ZRank(key, member string) *redis.IntCmd {
	return that.client.ZRank(that.ctx, key, member)
}

func (that *AWRedis) ZRem(key string, members ...interface{}) *redis.IntCmd {
	return that.client.ZRem(that.ctx, key, members...)
}

func (that *AWRedis) ZRemRangeByRank(key string, start, stop int64) *redis.IntCmd {
	return that.client.ZRemRangeByRank(that.ctx, key, start, stop)
}

func (that *AWRedis) ZRemRangeByScore(key, min, max string) *redis.IntCmd {
	return that.client.ZRemRangeByScore(that.ctx, key, min, max)
}

func (that *AWRedis) ZRemRangeByLex(key, min, max string) *redis.IntCmd {
	return that.client.ZRemRangeByLex(that.ctx, key, min, max)
}

func (that *AWRedis) ZRevRange(key string, start, stop int64) *redis.StringSliceCmd {
	return that.client.ZRevRange(that.ctx, key, start, stop)
}

func (that *AWRedis) ZRevRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd {
	return that.client.ZRevRangeWithScores(that.ctx, key, start, stop)
}

func (that *AWRedis) ZRevRangeByScore(key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	return that.client.ZRevRangeByScore(that.ctx, key, opt)
}

func (that *AWRedis) ZRevRangeByLex(key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	return that.client.ZRevRangeByLex(that.ctx, key, opt)
}

func (that *AWRedis) ZRevRangeByScoreWithScores(key string, opt *redis.ZRangeBy) *redis.ZSliceCmd {
	return that.client.ZRevRangeByScoreWithScores(that.ctx, key, opt)
}

func (that *AWRedis) ZRevRank(key, member string) *redis.IntCmd {
	return that.client.ZRevRank(that.ctx, key, member)
}

func (that *AWRedis) ZScore(key, member string) *redis.FloatCmd {
	return that.client.ZScore(that.ctx, key, member)
}

func (that *AWRedis) ZUnionStore(dest string, store *redis.ZStore) *redis.IntCmd {
	return that.client.ZUnionStore(that.ctx, dest, store)
}

func (that *AWRedis) ZUnion(store redis.ZStore) *redis.StringSliceCmd {
	return that.client.ZUnion(that.ctx, store)
}

func (that *AWRedis) ZUnionWithScores(store redis.ZStore) *redis.ZSliceCmd {
	return that.client.ZUnionWithScores(that.ctx, store)
}

func (that *AWRedis) ZRandMember(key string, count int, withScores bool) *redis.StringSliceCmd {
	return that.client.ZRandMember(that.ctx, key, count, withScores)
}

func (that *AWRedis) ZDiff(keys ...string) *redis.StringSliceCmd {
	return that.client.ZDiff(that.ctx, keys...)
}

func (that *AWRedis) ZDiffWithScores(keys ...string) *redis.ZSliceCmd {
	return that.client.ZDiffWithScores(that.ctx, keys...)
}

func (that *AWRedis) ZDiffStore(destination string, keys ...string) *redis.IntCmd {
	return that.client.ZDiffStore(that.ctx, destination, keys...)
}

func (that *AWRedis) BgRewriteAOF() *redis.StatusCmd {
	return that.client.BgRewriteAOF(that.ctx)
}

func (that *AWRedis) BgSave() *redis.StatusCmd {
	return that.client.BgSave(that.ctx)
}

func (that *AWRedis) ClientKill(ipPort string) *redis.StatusCmd {
	return that.client.ClientKill(that.ctx, ipPort)
}

func (that *AWRedis) ClientKillByFilter(keys ...string) *redis.IntCmd {
	return that.client.ClientKillByFilter(that.ctx, keys...)
}

func (that *AWRedis) ClientList() *redis.StringCmd {
	return that.client.ClientList(that.ctx)
}

func (that *AWRedis) ClientPause(dur time.Duration) *redis.BoolCmd {
	return that.client.ClientPause(that.ctx, dur)
}

func (that *AWRedis) ClientID() *redis.IntCmd {
	return that.client.ClientID(that.ctx)
}

func (that *AWRedis) ConfigGet(parameter string) *redis.SliceCmd {
	return that.client.ConfigGet(that.ctx, parameter)
}

func (that *AWRedis) ConfigResetStat() *redis.StatusCmd {
	return that.client.ConfigResetStat(that.ctx)
}

func (that *AWRedis) ConfigSet(parameter, value string) *redis.StatusCmd {
	return that.client.ConfigSet(that.ctx, parameter, value)
}

func (that *AWRedis) ConfigRewrite() *redis.StatusCmd {
	return that.client.ConfigRewrite(that.ctx)
}

func (that *AWRedis) DBSize() *redis.IntCmd {
	return that.client.DBSize(that.ctx)
}

func (that *AWRedis) FlushAll() *redis.StatusCmd {
	return that.client.FlushAll(that.ctx)
}

func (that *AWRedis) FlushAllAsync() *redis.StatusCmd {
	return that.client.FlushAllAsync(that.ctx)
}

func (that *AWRedis) FlushDB() *redis.StatusCmd {
	return that.client.FlushDB(that.ctx)
}

func (that *AWRedis) FlushDBAsync() *redis.StatusCmd {
	return that.client.FlushDBAsync(that.ctx)
}

func (that *AWRedis) Info(section ...string) *redis.StringCmd {
	return that.client.Info(that.ctx, section...)
}

func (that *AWRedis) LastSave() *redis.IntCmd {
	return that.client.LastSave(that.ctx)
}

func (that *AWRedis) Save() *redis.StatusCmd {
	return that.client.Save(that.ctx)
}

func (that *AWRedis) Shutdown() *redis.StatusCmd {
	return that.client.Shutdown(that.ctx)
}

func (that *AWRedis) ShutdownSave() *redis.StatusCmd {
	return that.client.ShutdownSave(that.ctx)
}

func (that *AWRedis) ShutdownNoSave() *redis.StatusCmd {
	return that.client.ShutdownNoSave(that.ctx)
}

func (that *AWRedis) SlaveOf(host, port string) *redis.StatusCmd {
	return that.client.SlaveOf(that.ctx, host, port)
}

func (that *AWRedis) Time() *redis.TimeCmd {
	return that.client.Time(that.ctx)
}

func (that *AWRedis) DebugObject(key string) *redis.StringCmd {
	return that.client.DebugObject(that.ctx, key)
}

func (that *AWRedis) ReadOnly() *redis.StatusCmd {
	return that.client.ReadOnly(that.ctx)
}

func (that *AWRedis) ReadWrite() *redis.StatusCmd {
	return that.client.ReadWrite(that.ctx)
}

func (that *AWRedis) MemoryUsage(key string, samples ...int) *redis.IntCmd {
	return that.client.MemoryUsage(that.ctx, key, samples...)
}

func (that *AWRedis) Eval(script string, keys []string, args ...interface{}) *redis.Cmd {
	return that.client.Eval(that.ctx, script, keys, args...)
}

func (that *AWRedis) EvalSha(sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	return that.client.EvalSha(that.ctx, sha1, keys, args...)
}

func (that *AWRedis) ScriptExists(hashes ...string) *redis.BoolSliceCmd {
	return that.client.ScriptExists(that.ctx, hashes...)
}

func (that *AWRedis) ScriptFlush() *redis.StatusCmd {
	return that.client.ScriptFlush(that.ctx)
}

func (that *AWRedis) ScriptKill() *redis.StatusCmd {
	return that.client.ScriptKill(that.ctx)
}

func (that *AWRedis) ScriptLoad(script string) *redis.StringCmd {
	return that.client.ScriptLoad(that.ctx, script)
}
