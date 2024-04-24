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
			SparkHost:            "http://127.0.0.1:4040",
			Neo4jAddr:            "neo4j://127.0.0.1",
			Neo4jUsername:        "neo4j",
			Neo4jPassword:        "12345678",
			RabbitConCurrencyNum: "1",
			RabbitMqAddr:         "127.0.0.1",
			RabbitMqPort:         "5672",
			RabbitMqUser:         "kwq",
			RabbitMqPwd:          "123456",
			RedisAddr:            "127.0.0.1:6379",
			RedisPassword:        "",
			RedisDB:              1,
			ESAddr:               "http://127.0.0.1:9200",
			ESUsername:           "",
			ESPassword:           "",
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
		Scene:        "inflow",
	}
	t.Run("EsToNeo", func(t *testing.T) {
		if err := n.EsToNeo(context.TODO(), &req); err != nil {
			t.Errorf("EsToNeo() error = %v", err)
		}
	})
}

func TestNeoNode_DeleteAllNodes(t *testing.T) {
	_init()
	t.Run("deleteAllNode", func(t *testing.T) {
		err := n.DeleteAllNodes(context.TODO())
		if err != nil {
			t.Errorf("deleteAllNode() error = %v", err)
		}
	})
}
