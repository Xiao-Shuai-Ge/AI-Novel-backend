package config

import (
	"Ai-Novel/common/gormx"
	"Ai-Novel/common/redisx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql   gormx.Mysql
	Redis   redisx.Redis
	LogConf logx.LogConf
	Auth    struct {
		AccessSecret string
		AccessExpire int64
	}
}
