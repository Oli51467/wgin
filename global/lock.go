package global

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
	"wgin/util"
)

// LockInterface 定义分布式锁需要实现的接口
type LockInterface interface {
	Get() bool
	Block(second int64) bool
	Release() bool
	ForceRelease()
}

type lock struct {
	context context.Context
	name    string // 锁名称
	owner   string // 锁标识
	seconds int64  // 有效期
}

// 释放锁 Lua 脚本，防止任何客户端都能解锁
const releaseLockLuaScript = `
if redis.call("get",KEYS[1]) == ARGV[1] then
    return redis.call("del",KEYS[1])
else
    return 0
end
`

// MakeLock 生成一把锁
func MakeLock(name string, seconds int64) LockInterface {
	return &lock{
		context.Background(),
		name,
		util.RandString(16),
		seconds,
	}
}

// Get 尝试获取一把锁
func (l *lock) Get() bool {
	return App.Redis.SetNX(l.context, l.name, l.owner, time.Duration(l.seconds)*time.Second).Val()
}

// Block 阻塞获取一把锁 如果在second时间内仍然没有获得锁 则退出竞争
func (l *lock) Block(seconds int64) bool {
	startTime := time.Now().Unix()
	for {
		if !l.Get() {
			time.Sleep(time.Duration(1) * time.Second)
			if time.Now().Unix()-seconds >= startTime {
				return false
			}
		} else {
			return true
		}
	}
}

// Release 释放一把锁
func (l *lock) Release() bool {
	releaseLuaScript := redis.NewScript(releaseLockLuaScript)
	result := releaseLuaScript.Run(l.context, App.Redis, []string{l.name}, l.owner).Val().(int64)
	return result != 0
}

func (l *lock) ForceRelease() {
	App.Redis.Del(l.context, l.name)
}
