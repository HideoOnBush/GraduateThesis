package service

import (
	"GraduateThesis/biz/service/line"
	"context"
	"sync"

	"GraduateThesis/biz/service/base"
	"GraduateThesis/conf"
)

var (
	b    base.Base
	Line line.Line
)

func Init(ctx context.Context, leaderCond *sync.Cond, slaveCond *sync.Cond) {
	b.Init(conf.GConfig)
	Line = *line.New(ctx, b)
}
