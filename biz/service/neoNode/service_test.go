package neoNode

import (
	"GraduateThesis/biz/model/line"
	"GraduateThesis/biz/service/base"
	Line "GraduateThesis/biz/service/line"
	"GraduateThesis/conf"
	"context"
	"sync"
	"testing"
)

var (
	b    base.Base
	l    Line.Line
	once sync.Once
	n    NeoNode
)

func _init() {
	once.Do(func() {
		conf.GConfig = &conf.Config{
			RedisAddr:     "127.0.0.1:6379",
			RedisPassword: "",
			RedisDB:       1,
			ESAddr:        "http://127.0.0.1:9200",
			ESUsername:    "",
			ESPassword:    "",
			Neo4jAddr:     "neo4j://127.0.0.1",
			Neo4jUsername: "neo4j",
			Neo4jPassword: "12345678",
		}
		b.Init(conf.GConfig)
		l = *Line.New(context.Background(), b)
		n = *New(context.Background(), b, l)
	})
}

func TestNeoNode_EsToNeo(t *testing.T) {
	_init()
	req := line.LineReq{
		Sources:      nil,
		Targets:      nil,
		Dependencies: nil,
		MatchAll:     false,
		PageSize:     0,
		PageNum:      0,
		IgnoreIds:    nil,
		IsCore:       false,
		Scene:        "test",
	}
	t.Run("EsToNeo", func(t *testing.T) {
		if err := n.EsToNeo(context.TODO(), &req); err != nil {
			t.Errorf("EsToNeo() error = %v", err)
		}
	})
}
