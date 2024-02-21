package service

import (
	"GraduateThesis/biz/service/base"
	"GraduateThesis/biz/service/line"
	"GraduateThesis/biz/service/neoNode"
	"GraduateThesis/conf"
	"context"
	"sync"
)

var (
	b       base.Base
	Line    line.Line
	NeoNode neoNode.NeoNode
)

func Init(ctx context.Context, leaderCond *sync.Cond, slaveCond *sync.Cond) {
	b.Init(conf.GConfig)
	Line = *line.New(ctx, b)
	NeoNode = *neoNode.New(ctx, b, Line)
}
