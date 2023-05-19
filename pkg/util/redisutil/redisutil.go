package redisutil

import (
	"context"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gmlock"
	"github.com/gogf/gf/v2/util/gconv"
)

const unlockLua = "if redis.call(\"get\",KEYS[1]) == ARGV[1] " +
	"then " +
	"    return redis.call(\"del\",KEYS[1]) " +
	"else " +
	"    return 0 " +
	"end "

func TryLock(ctx context.Context, key string, value interface{}, ttlSeconds int64) bool {
	if !gmlock.TryLock(key) {
		return false
	}

	gv, err := g.Redis().Set(ctx, key, value, gredis.SetOption{
		TTLOption: gredis.TTLOption{
			EX: gconv.PtrInt64(ttlSeconds),
		},
		NX: true,
	})
	if err != nil {
		return false
	}
	return gv.Bool()
}

func Unlock(ctx context.Context, key string, value interface{}) (bool, error) {
	defer gmlock.Unlock(key)
	v, err := g.Redis().Eval(ctx, unlockLua, 1, []string{key}, []interface{}{value})
	if err != nil {
		return false, err
	}
	if v.String() == "0" {
		return false, gerror.New("unlock key=" + key + " failed")
	}

	return true, nil
}
