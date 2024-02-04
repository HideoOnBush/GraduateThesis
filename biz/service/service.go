package service

import (
	"context"
	"sync"

	"GraduateThesis/biz/service/base"
	"GraduateThesis/conf"
)

var (
	b base.Base
)

func Init(ctx context.Context, leaderCond *sync.Cond, slaveCond *sync.Cond) {
	b.Init(conf.GConfig)
}
