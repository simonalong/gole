package redis

import (
	goredis "github.com/go-redis/redis/v8"
	"github.com/simonalong/gole/bean"
	"github.com/simonalong/gole/config"
	"github.com/simonalong/gole/constants"
	"github.com/simonalong/gole/logger"
	goleTime "github.com/simonalong/gole/time"
	//"github.com/simonalong/gole/tracing"
	//"github.com/simonalong/gole/tracing"
	"time"
)

var RedisHooks []goredis.Hook

type ConfigError struct {
	ErrMsg string
}

func (error *ConfigError) Error() string {
	return error.ErrMsg
}

func init() {
	config.LoadConfig()

	if config.ExistConfigFile() && config.GetValueBoolDefault("gole.redis.enable", false) {
		err := config.GetValueObject("gole.redis", &config.RedisCfg)
		if err != nil {
			logger.Warn("读取redis配置异常")
			return
		}
	}
	RedisHooks = []goredis.Hook{}
}

func NewClient() (goredis.UniversalClient, error) {
	var rdbClient goredis.UniversalClient
	if config.RedisCfg.Sentinel.Master != "" {
		rdbClient = goredis.NewFailoverClient(getSentinelConfig())
	} else if len(config.RedisCfg.Cluster.Addrs) != 0 {
		rdbClient = goredis.NewClusterClient(getClusterConfig())
	} else {
		rdbClient = goredis.NewClient(getStandaloneConfig())
	}

	for _, hook := range RedisHooks {
		rdbClient.AddHook(hook)
	}
	bean.AddBean(constants.BeanNameRedisPre, &rdbClient)
	return rdbClient, nil
}

func AddRedisHook(hook goredis.Hook) {
	RedisHooks = append(RedisHooks, hook)
}

func getStandaloneConfig() *goredis.Options {
	addr := "127.0.0.1:6379"
	if config.RedisCfg.Standalone.Addr != "" {
		addr = config.RedisCfg.Standalone.Addr
	}

	redisConfig := &goredis.Options{
		Addr: addr,

		DB:       config.RedisCfg.Standalone.Database,
		Network:  config.RedisCfg.Standalone.Network,
		Username: config.RedisCfg.Username,
		Password: config.RedisCfg.Password,

		MaxRetries:      config.RedisCfg.MaxRetries,
		MinRetryBackoff: goleTime.NumToTimeDuration(config.RedisCfg.MinRetryBackoff, time.Millisecond),
		MaxRetryBackoff: goleTime.NumToTimeDuration(config.RedisCfg.MaxRetryBackoff, time.Millisecond),

		DialTimeout:  goleTime.NumToTimeDuration(config.RedisCfg.DialTimeout, time.Millisecond),
		ReadTimeout:  goleTime.NumToTimeDuration(config.RedisCfg.ReadTimeout, time.Millisecond),
		WriteTimeout: goleTime.NumToTimeDuration(config.RedisCfg.WriteTimeout, time.Millisecond),

		PoolFIFO:           config.RedisCfg.PoolFIFO,
		PoolSize:           config.RedisCfg.PoolSize,
		MinIdleConns:       config.RedisCfg.MinIdleConns,
		MaxConnAge:         goleTime.NumToTimeDuration(config.RedisCfg.MaxConnAge, time.Millisecond),
		PoolTimeout:        goleTime.NumToTimeDuration(config.RedisCfg.PoolTimeout, time.Millisecond),
		IdleTimeout:        goleTime.NumToTimeDuration(config.RedisCfg.IdleTimeout, time.Millisecond),
		IdleCheckFrequency: goleTime.NumToTimeDuration(config.RedisCfg.IdleCheckFrequency, time.Millisecond),
	}

	// -------- 命令执行失败配置 --------
	if config.GetValueString("gole.redis.max-retries") == "" {
		// # 命令执行失败时候，最大重试次数，默认3次，-1（不是0）则不重试
		redisConfig.MaxRetries = 3
	}

	if config.GetValueString("gole.redis.min-retry-backoff") == "" {
		// #（单位毫秒） 命令执行失败时候，每次重试的最小回退时间，默认8毫秒，-1则禁止回退
		redisConfig.MinRetryBackoff = 8 * time.Millisecond
	}

	if config.GetValueString("gole.redis.max-retry-backoff") == "" {
		// #（单位毫秒） 命令执行失败时候，每次重试的最小回退时间，默认8毫秒，-1则禁止回退
		redisConfig.MinRetryBackoff = 512 * time.Millisecond
	}

	// -------- 超时配置 --------
	if config.GetValueString("gole.redis.dial-timeout") == "" {
		// # （单位毫秒）超时：创建新链接的拨号超时时间，默认15秒
		redisConfig.DialTimeout = 15 * time.Second
	}

	if config.GetValueString("gole.redis.read-timeout") == "" {
		// # （单位毫秒）超时：读超时，默认3秒，使用-1，使用-1则表示无超时，0的话是表示默认3秒
		redisConfig.ReadTimeout = 3 * time.Second
	}

	if config.GetValueString("gole.redis.write-timeout") == "" {
		// # （单位毫秒）超时：写超时，默认是读超时3秒，使用-1，使用-1则表示无超时，0的话是表示默认3秒
		redisConfig.WriteTimeout = 3 * time.Second
	}

	// -------- 连接池相关配置 --------
	if config.GetValueString("gole.redis.pool-fifo") == "" {
		// # 连接池类型：fifo：true;lifo：false;和lifo相比，fifo开销更高
		redisConfig.PoolFIFO = false
	}

	if config.GetValueString("gole.redis.pool-size") == "" {
		// # 最大连接池大小：默认每个cpu核是10个连接，cpu核数可以根据函数runtime.GOMAXPROCS来配置，默认是runtime.NumCpu
		redisConfig.PoolSize = 10
	}

	if config.GetValueString("gole.redis.min-idle-conns") == "" {
		// # 最小空闲连接数
		redisConfig.MinIdleConns = 10
	}

	if config.GetValueString("gole.redis.max-conn-age") == "" {
		// #（单位毫秒） 连接存活时长，默认不关闭
		redisConfig.MaxConnAge = 12 * 30 * 24 * time.Hour
	}

	if config.GetValueString("gole.redis.pool-timeout") == "" {
		// #（单位毫秒）获取链接池中的链接都在忙，则等待对应的时间，默认读超时+1秒
		redisConfig.PoolTimeout = time.Second
	}

	if config.GetValueString("gole.redis.idle-timeout") == "" {
		// #（单位毫秒）空闲链接时间，超时则关闭，注意：该时间要小于服务端的超时时间，否则会出现拿到的链接失效问题，默认5分钟，-1表示禁用超时检查
		redisConfig.IdleTimeout = 5 * time.Minute
	}

	if config.GetValueString("gole.redis.idle-check-frequency") == "" {
		// #（单位毫秒）空闲链接核查频率，默认1分钟。-1禁止空闲链接核查，即使配置了IdleTime也不行
		redisConfig.IdleCheckFrequency = time.Minute
	}
	return redisConfig
}

func getSentinelConfig() *goredis.FailoverOptions {
	redisConfig := &goredis.FailoverOptions{
		SentinelAddrs: config.RedisCfg.Sentinel.Addrs,
		MasterName:    config.RedisCfg.Sentinel.Master,

		DB:               config.RedisCfg.Sentinel.Database,
		Username:         config.RedisCfg.Username,
		Password:         config.RedisCfg.Password,
		SentinelUsername: config.RedisCfg.Sentinel.SentinelUser,
		SentinelPassword: config.RedisCfg.Sentinel.SentinelPassword,

		MaxRetries:      config.RedisCfg.MaxRetries,
		MinRetryBackoff: goleTime.NumToTimeDuration(config.RedisCfg.MinRetryBackoff, time.Millisecond),
		MaxRetryBackoff: goleTime.NumToTimeDuration(config.RedisCfg.MaxRetryBackoff, time.Millisecond),

		DialTimeout:  goleTime.NumToTimeDuration(config.RedisCfg.DialTimeout, time.Millisecond),
		ReadTimeout:  goleTime.NumToTimeDuration(config.RedisCfg.ReadTimeout, time.Millisecond),
		WriteTimeout: goleTime.NumToTimeDuration(config.RedisCfg.WriteTimeout, time.Millisecond),

		PoolFIFO:           config.RedisCfg.PoolFIFO,
		PoolSize:           config.RedisCfg.PoolSize,
		MinIdleConns:       config.RedisCfg.MinIdleConns,
		MaxConnAge:         goleTime.NumToTimeDuration(config.RedisCfg.MaxConnAge, time.Millisecond),
		PoolTimeout:        goleTime.NumToTimeDuration(config.RedisCfg.PoolTimeout, time.Millisecond),
		IdleTimeout:        goleTime.NumToTimeDuration(config.RedisCfg.IdleTimeout, time.Millisecond),
		IdleCheckFrequency: goleTime.NumToTimeDuration(config.RedisCfg.IdleCheckFrequency, time.Millisecond),
	}

	// -------- 命令执行失败配置 --------
	if config.GetValueString("gole.redis.max-retries") == "" {
		// # 命令执行失败时候，最大重试次数，默认3次，-1（不是0）则不重试
		redisConfig.MaxRetries = 3
	}

	if config.GetValueString("gole.redis.min-retry-backoff") == "" {
		// #（单位毫秒） 命令执行失败时候，每次重试的最小回退时间，默认8毫秒，-1则禁止回退
		redisConfig.MinRetryBackoff = 8 * time.Millisecond
	}

	if config.GetValueString("gole.redis.max-retry-backoff") == "" {
		// #（单位毫秒） 命令执行失败时候，每次重试的最小回退时间，默认8毫秒，-1则禁止回退
		redisConfig.MinRetryBackoff = 512 * time.Millisecond
	}

	// -------- 超时配置 --------
	if config.GetValueString("gole.redis.dial-timeout") == "" {
		// # （单位毫秒）超时：创建新链接的拨号超时时间，默认15秒
		redisConfig.DialTimeout = 15 * time.Second
	}

	if config.GetValueString("gole.redis.read-timeout") == "" {
		// # （单位毫秒）超时：读超时，默认3秒，使用-1，使用-1则表示无超时，0的话是表示默认3秒
		redisConfig.ReadTimeout = 3 * time.Second
	}

	if config.GetValueString("gole.redis.write-timeout") == "" {
		// # （单位毫秒）超时：写超时，默认是读超时3秒，使用-1，使用-1则表示无超时，0的话是表示默认3秒
		redisConfig.WriteTimeout = 3 * time.Second
	}

	// -------- 连接池相关配置 --------
	if config.GetValueString("gole.redis.pool-fifo") == "" {
		// # 连接池类型：fifo：true;lifo：false;和lifo相比，fifo开销更高
		redisConfig.PoolFIFO = false
	}

	if config.GetValueString("gole.redis.pool-size") == "" {
		// # 最大连接池大小：默认每个cpu核是10个连接，cpu核数可以根据函数runtime.GOMAXPROCS来配置，默认是runtime.NumCpu
		redisConfig.PoolSize = 10
	}

	if config.GetValueString("gole.redis.min-idle-conns") == "" {
		// # 最小空闲连接数
		redisConfig.MinIdleConns = 10
	}

	if config.GetValueString("gole.redis.max-conn-age") == "" {
		// #（单位毫秒） 连接存活时长，默认不关闭
		redisConfig.MaxConnAge = 12 * 30 * 24 * time.Hour
	}

	if config.GetValueString("gole.redis.pool-timeout") == "" {
		// #（单位毫秒）获取链接池中的链接都在忙，则等待对应的时间，默认读超时+1秒
		redisConfig.PoolTimeout = time.Second
	}

	if config.GetValueString("gole.redis.idle-timeout") == "" {
		// #（单位毫秒）空闲链接时间，超时则关闭，注意：该时间要小于服务端的超时时间，否则会出现拿到的链接失效问题，默认5分钟，-1表示禁用超时检查
		redisConfig.IdleTimeout = 5 * time.Minute
	}

	if config.GetValueString("gole.redis.idle-check-frequency") == "" {
		// #（单位毫秒）空闲链接核查频率，默认1分钟。-1禁止空闲链接核查，即使配置了IdleTime也不行
		redisConfig.IdleCheckFrequency = time.Minute
	}
	return redisConfig
}

func getClusterConfig() *goredis.ClusterOptions {
	if len(config.RedisCfg.Cluster.Addrs) == 0 {
		config.RedisCfg.Cluster.Addrs = []string{"127.0.0.1:6379"}
	}

	redisConfig := &goredis.ClusterOptions{
		Addrs: config.RedisCfg.Cluster.Addrs,

		Username: config.RedisCfg.Username,
		Password: config.RedisCfg.Password,

		MaxRedirects:   config.RedisCfg.Cluster.MaxRedirects,
		ReadOnly:       config.RedisCfg.Cluster.ReadOnly,
		RouteByLatency: config.RedisCfg.Cluster.RouteByLatency,
		RouteRandomly:  config.RedisCfg.Cluster.RouteRandomly,

		MaxRetries:      config.RedisCfg.MaxRetries,
		MinRetryBackoff: goleTime.NumToTimeDuration(config.RedisCfg.MinRetryBackoff, time.Millisecond),
		MaxRetryBackoff: goleTime.NumToTimeDuration(config.RedisCfg.MaxRetryBackoff, time.Millisecond),

		DialTimeout:  goleTime.NumToTimeDuration(config.RedisCfg.DialTimeout, time.Millisecond),
		ReadTimeout:  goleTime.NumToTimeDuration(config.RedisCfg.ReadTimeout, time.Millisecond),
		WriteTimeout: goleTime.NumToTimeDuration(config.RedisCfg.WriteTimeout, time.Millisecond),
		PoolFIFO:     config.RedisCfg.PoolFIFO,
		PoolSize:     config.RedisCfg.PoolSize,
		MinIdleConns: config.RedisCfg.MinIdleConns,

		MaxConnAge:         goleTime.NumToTimeDuration(config.RedisCfg.MaxConnAge, time.Millisecond),
		PoolTimeout:        goleTime.NumToTimeDuration(config.RedisCfg.PoolTimeout, time.Millisecond),
		IdleTimeout:        goleTime.NumToTimeDuration(config.RedisCfg.IdleTimeout, time.Millisecond),
		IdleCheckFrequency: goleTime.NumToTimeDuration(config.RedisCfg.IdleCheckFrequency, time.Millisecond),
	}

	// -------- 命令执行失败配置 --------
	if config.GetValueString("gole.redis.max-retries") == "" {
		// # 命令执行失败时候，最大重试次数，默认3次，-1（不是0）则不重试
		redisConfig.MaxRetries = 3
	}

	if config.GetValueString("gole.redis.min-retry-backoff") == "" {
		// #（单位毫秒） 命令执行失败时候，每次重试的最小回退时间，默认8毫秒，-1则禁止回退
		redisConfig.MinRetryBackoff = 8 * time.Millisecond
	}

	if config.GetValueString("gole.redis.max-retry-backoff") == "" {
		// #（单位毫秒） 命令执行失败时候，每次重试的最小回退时间，默认8毫秒，-1则禁止回退
		redisConfig.MinRetryBackoff = 512 * time.Millisecond
	}

	// -------- 超时配置 --------
	if config.GetValueString("gole.redis.dial-timeout") == "" {
		// # （单位毫秒）超时：创建新链接的拨号超时时间，默认15秒
		redisConfig.DialTimeout = 15 * time.Second
	}

	if config.GetValueString("gole.redis.read-timeout") == "" {
		// # （单位毫秒）超时：读超时，默认3秒，使用-1，使用-1则表示无超时，0的话是表示默认3秒
		redisConfig.ReadTimeout = 3 * time.Second
	}

	if config.GetValueString("gole.redis.write-timeout") == "" {
		// # （单位毫秒）超时：写超时，默认是读超时3秒，使用-1，使用-1则表示无超时，0的话是表示默认3秒
		redisConfig.WriteTimeout = 3 * time.Second
	}

	// -------- 连接池相关配置 --------
	if config.GetValueString("gole.redis.pool-fifo") == "" {
		// # 连接池类型：fifo：true;lifo：false;和lifo相比，fifo开销更高
		redisConfig.PoolFIFO = false
	}

	if config.GetValueString("gole.redis.pool-size") == "" {
		// # 最大连接池大小：默认每个cpu核是10个连接，cpu核数可以根据函数runtime.GOMAXPROCS来配置，默认是runtime.NumCpu
		redisConfig.PoolSize = 10
	}

	if config.GetValueString("gole.redis.min-idle-conns") == "" {
		// # 最小空闲连接数
		redisConfig.MinIdleConns = 10
	}

	if config.GetValueString("gole.redis.max-conn-age") == "" {
		// #（单位毫秒） 连接存活时长，默认不关闭
		redisConfig.MaxConnAge = 12 * 30 * 24 * time.Hour
	}

	if config.GetValueString("gole.redis.pool-timeout") == "" {
		// #（单位毫秒）获取链接池中的链接都在忙，则等待对应的时间，默认读超时+1秒
		redisConfig.PoolTimeout = time.Second
	}

	if config.GetValueString("gole.redis.idle-timeout") == "" {
		// #（单位毫秒）空闲链接时间，超时则关闭，注意：该时间要小于服务端的超时时间，否则会出现拿到的链接失效问题，默认5分钟，-1表示禁用超时检查
		redisConfig.IdleTimeout = 5 * time.Minute
	}

	if config.GetValueString("gole.redis.idle-check-frequency") == "" {
		// #（单位毫秒）空闲链接核查频率，默认1分钟。-1禁止空闲链接核查，即使配置了IdleTime也不行
		redisConfig.IdleCheckFrequency = time.Minute
	}
	return redisConfig
}
