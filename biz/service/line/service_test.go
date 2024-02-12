package line

import (
	lineModel "GraduateThesis/biz/model/line"
	"GraduateThesis/biz/service/base"
	"GraduateThesis/conf"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	jsoniter "github.com/json-iterator/go"
	"sync"
	"testing"
)

var (
	b    base.Base
	line Line
	once sync.Once
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
		}
		b.Init(conf.GConfig)
		line = *New(context.Background(), b)
	})
}

func TestLine_Bulk(t *testing.T) {
	_init()
	lines := []*lineModel.Line{
		//{
		//	Source:         "edith-gateway-nd",
		//	SourceIsCore:   false,
		//	SourceScene:    "test",
		//	Target:         "apex-predator",
		//	TargetIsCore:   false,
		//	TargetScene:    "test",
		//	Dependence:     "strong",
		//	LastFoundTime:  0,
		//	FirstFoundTime: 0,
		//},
		{
			Source:         "notefeed-gateway-default",
			SourceIsCore:   false,
			SourceScene:    "test",
			Target:         "apex-predator",
			TargetIsCore:   false,
			TargetScene:    "test",
			Dependence:     "weak",
			LastFoundTime:  0,
			FirstFoundTime: 0,
		},
	}

	t.Run("Bulk", func(t *testing.T) {
		if err := line.Bulk(context.TODO(), lines); err != nil {
			t.Errorf("Bulk() error = %v", err)
		}
	})
}

func TestLine_BulkDeleteById(t *testing.T) {
	_init()
	line1 := "s*edith-gateway-nd*t*aceflow-notefeed-default"
	t.Run("BulkDeleteById", func(t *testing.T) {
		if ok, err := line.DeleteById(context.TODO(), "test", line1); err != nil {
			t.Log(ok)
			t.Errorf("BulkDeleteById() error = %v", err)
		} else {
			t.Log(ok)
		}
	})
}

func TestLine_ChangeDependenceByRelation(t *testing.T) {
	_init()
	req := lineModel.ChangeDependenceWithRelationReq{
		Source:       "edith-gateway-nd",
		Target:       "aceflow-notefeed-default",
		RelationType: "weak",
		Scene:        "test",
	}
	c := &app.RequestContext{}
	t.Run("ChangeDependenceByRelation", func(t *testing.T) {
		ok := line.ChangeDependenceByRelation(context.TODO(), c, req)
		if c.Errors != nil {
			t.Log(ok)
			t.Errorf("ChangeDependenceByRelation() error=%v", c.Errors.String())
		}
		t.Log(ok)
	})
}

func TestLine_Query(t *testing.T) {
	_init()
	c := &app.RequestContext{}
	req := lineModel.LineReq{
		Sources:      []string{"edith-gateway-nd"},
		Targets:      []string{"aceflow-notefeed-default"},
		Dependencies: nil,
		MatchAll:     false,
		PageSize:     0,
		PageNum:      0,
		IgnoreIds:    nil,
		IsCore:       false,
		Scene:        "test",
	}
	t.Run("Query", func(t *testing.T) {
		lines, total := line.Query(context.TODO(), c, &req)
		if c.Errors != nil {
			t.Errorf("ChangeDependenceByRelation() error=%v", c.Errors.String())
		}
		t.Log(total)
		data, _ := jsoniter.MarshalIndent(lines, "", " ")
		t.Log(string(data))
	})
}

func TestLine_TopologyAnalyse(t *testing.T) {
	_init()
	req := lineModel.TopologyIndicatorReq{
		Scene: "test",
	}
	t.Run("Query", func(t *testing.T) {
		ok, indicator := line.TopologyAnalyse(context.TODO(), &req)
		t.Log(ok)
		if ok {
			data, _ := jsoniter.MarshalIndent(indicator, "", " ")
			t.Log(string(data))
		}
	})
}
