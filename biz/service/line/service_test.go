package line

import (
	lineModel "GraduateThesis/biz/model/line"
	"GraduateThesis/biz/service/base"
	"GraduateThesis/conf"
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	jsoniter "github.com/json-iterator/go"
	"log"
	"math/rand"
	"sync"
	"testing"
)

const lineListJson = `[
    {
        "_source": {
            "source": "edith-gateway-nd",
            "target": "aceflow-notefeed-default",
            "dependence": "strong",
            "prejudgment": {},
            "verify": {},
            "labels": [
                {
                    "name": "type",
                    "value": "direct"
                },
                {
                    "name": "scenario",
                    "value": "neiliu"
                },
                {
                    "name": "sourceType",
                    "value": "Service"
                },
                {
                    "name": "targetType",
                    "value": "Service"
                }
            ]
        }
    },
    {
        "_source": {
            "source": "aceflow-notefeed-default",
            "target": "notefeed-gateway-default",
            "dependence": "strong",
            "prejudgment": {},
            "verify": {},
            "labels": [
                {
                    "name": "type",
                    "value": "direct"
                },
                {
                    "name": "scenario",
                    "value": "neiliu"
                },
                {
                    "name": "sourceType",
                    "value": "Service"
                },
                {
                    "name": "targetType",
                    "value": "Service"
                }
            ]
        }
    },
    {
        "_source": {
            "source": "aceflow-notefeed-default",
            "target": "userprofile-service-default",
            "dependence": "strong",
            "prejudgment": {},
            "verify": {},
            "labels": [
                {
                    "name": "type",
                    "value": "direct"
                },
                {
                    "name": "scenario",
                    "value": "neiliu"
                },
                {
                    "name": "sourceType",
                    "value": "Service"
                },
                {
                    "name": "targetType",
                    "value": "Service"
                }
            ]
        }
    },
    {
        "_source": {
            "source": "notefeed-gateway-default",
            "target": "reddedupservice-service-videofeed",
            "dependence": "weak",
            "prejudgment": {},
            "verify": {},
            "labels": [
                {
                    "name": "type",
                    "value": "direct"
                },
                {
                    "name": "scenario",
                    "value": "neiliu"
                },
                {
                    "name": "sourceType",
                    "value": "Service"
                },
                {
                    "name": "targetType",
                    "value": "Service"
                }
            ]
        }
    },
    {
        "_source": {
            "source": "notefeed-gateway-default",
            "target": "userprofile-service-default",
            "dependence": "strong",
            "prejudgment": {},
            "verify": {},
            "labels": [
                {
                    "name": "type",
                    "value": "direct"
                },
                {
                    "name": "scenario",
                    "value": "neiliu"
                },
                {
                    "name": "sourceType",
                    "value": "Service"
                },
                {
                    "name": "targetType",
                    "value": "Service"
                }
            ]
        }
    },
    {
        "_source": {
            "source": "notefeed-gateway-default",
            "target": "zprofile-service-default",
            "dependence": "strong",
            "prejudgment": {},
            "verify": {},
            "labels": [
                {
                    "name": "type",
                    "value": "direct"
                },
                {
                    "name": "scenario",
                    "value": "neiliu"
                },
                {
                    "name": "sourceType",
                    "value": "Service"
                },
                {
                    "name": "targetType",
                    "value": "Service"
                }
            ]
        }
    },
    {
        "_source": {
            "source": "notefeed-gateway-default",
            "target": "noteprofile-service-default",
            "dependence": "strong",
            "prejudgment": {},
            "verify": {},
            "labels": [
                {
                    "name": "type",
                    "value": "direct"
                },
                {
                    "name": "scenario",
                    "value": "neiliu"
                },
                {
                    "name": "sourceType",
                    "value": "Service"
                },
                {
                    "name": "targetType",
                    "value": "Service"
                }
            ]
        }
    },
    {
        "_source": {
            "source": "notefeed-gateway-default",
            "target": "Redis",
            "dependence": "strong",
            "prejudgment": {},
            "verify": {},
            "labels": [
                {
                    "name": "type",
                    "value": "direct"
                },
                {
                    "name": "scenario",
                    "value": "neiliu"
                },
                {
                    "name": "sourceType",
                    "value": "Service"
                },
                {
                    "name": "targetType",
                    "value": "Service"
                }
            ]
        }
    },
    {
        "_source": {
            "source": "notefeed-gateway-default",
            "target": "RedKV",
            "dependence": "strong",
            "prejudgment": {},
            "verify": {},
            "labels": [
                {
                    "name": "type",
                    "value": "direct"
                },
                {
                    "name": "scenario",
                    "value": "neiliu"
                },
                {
                    "name": "sourceType",
                    "value": "Service"
                },
                {
                    "name": "targetType",
                    "value": "Service"
                }
            ]
        }
    },
    {
        "_source": {
            "source": "notefeed-gateway-default",
            "target": "annservice-service-dssmvideo",
            "dependence": "weak",
            "prejudgment": {},
            "verify": {},
            "labels": [
                {
                    "name": "type",
                    "value": "direct"
                },
                {
                    "name": "scenario",
                    "value": "neiliu"
                },
                {
                    "name": "sourceType",
                    "value": "Service"
                },
                {
                    "name": "targetType",
                    "value": "Service"
                }
            ]
        }
    },
    {
        "_source": {
            "source": "notefeed-gateway-default",
            "target": "annservice-shardvideofeed-default",
            "dependence": "weak",
            "prejudgment": {},
            "verify": {},
            "labels": [
                {
                    "name": "type",
                    "value": "direct"
                },
                {
                    "name": "scenario",
                    "value": "neiliu"
                },
                {
                    "name": "sourceType",
                    "value": "Service"
                },
                {
                    "name": "targetType",
                    "value": "Service"
                }
            ]
        }
    },
    {
        "_source": {
            "source": "notefeed-gateway-default",
            "target": "aceoctopus-merger-notefeed",
            "dependence": "weak",
            "prejudgment": {},
            "verify": {},
            "labels": [
                {
                    "name": "type",
                    "value": "direct"
                },
                {
                    "name": "scenario",
                    "value": "neiliu"
                },
                {
                    "name": "sourceType",
                    "value": "Service"
                },
                {
                    "name": "targetType",
                    "value": "Service"
                }
            ]
        }
    },
    {
        "_source": {
            "source": "aceoctopus-merger-notefeed",
            "target": "aceoctopus-recaller-notefeed",
            "dependence": "weak",
            "prejudgment": {},
            "verify": {},
            "labels": [
                {
                    "name": "type",
                    "value": "direct"
                },
                {
                    "name": "scenario",
                    "value": "neiliu"
                },
                {
                    "name": "sourceType",
                    "value": "Service"
                },
                {
                    "name": "targetType",
                    "value": "Service"
                }
            ]
        }
    },
    {
        "_source": {
            "source": "notefeed-gateway-default",
            "target": "LRE",
            "dependence": "weak",
            "prejudgment": {},
            "verify": {},
            "labels": [
                {
                    "name": "type",
                    "value": "direct"
                },
                {
                    "name": "scenario",
                    "value": "neiliu"
                },
                {
                    "name": "sourceType",
                    "value": "Service"
                },
                {
                    "name": "targetType",
                    "value": "Service"
                }
            ]
        }
    },
    {
        "_source": {
            "source": "notefeed-gateway-default",
            "target": "searchqueryrank-service-default",
            "dependence": "weak",
            "prejudgment": {},
            "verify": {},
            "labels": [
                {
                    "name": "type",
                    "value": "direct"
                },
                {
                    "name": "scenario",
                    "value": "neiliu"
                },
                {
                    "name": "sourceType",
                    "value": "Service"
                },
                {
                    "name": "targetType",
                    "value": "Service"
                }
            ]
        }
    },
    {
        "_source": {
            "source": "notefeed-gateway-default",
            "target": "reclambdaservice-service-notefeed-firstrank",
            "dependence": "strong",
            "prejudgment": {},
            "verify": {},
            "labels": [
                {
                    "name": "type",
                    "value": "direct"
                },
                {
                    "name": "scenario",
                    "value": "neiliu"
                },
                {
                    "name": "sourceType",
                    "value": "Service"
                },
                {
                    "name": "targetType",
                    "value": "Service"
                }
            ]
        }
    },
    {
        "_source": {
            "source": "notefeed-gateway-default",
            "target": "reclambdaservice-service-notefeed-finalrank",
            "dependence": "strong",
            "prejudgment": {},
            "verify": {},
            "labels": [
                {
                    "name": "type",
                    "value": "direct"
                },
                {
                    "name": "scenario",
                    "value": "neiliu"
                },
                {
                    "name": "sourceType",
                    "value": "Service"
                },
                {
                    "name": "targetType",
                    "value": "Service"
                }
            ]
        }
    },
    {
        "_source": {
            "source": "notefeed-gateway-default",
            "target": "reclambdaservice-service-notefeed-feature",
            "dependence": "strong",
            "prejudgment": {},
            "verify": {},
            "labels": [
                {
                    "name": "type",
                    "value": "direct"
                },
                {
                    "name": "scenario",
                    "value": "neiliu"
                },
                {
                    "name": "sourceType",
                    "value": "Service"
                },
                {
                    "name": "targetType",
                    "value": "Service"
                }
            ]
        }
    },
    {
        "_source": {
            "source": "notefeed-gateway-default",
            "target": "reclambdaservice-service-notefeed-recall",
            "dependence": "strong",
            "prejudgment": {},
            "verify": {},
            "labels": [
                {
                    "name": "type",
                    "value": "direct"
                },
                {
                    "name": "scenario",
                    "value": "neiliu"
                },
                {
                    "name": "sourceType",
                    "value": "Service"
                },
                {
                    "name": "targetType",
                    "value": "Service"
                }
            ]
        }
    },
    {
        "_source": {
            "source": "notefeed-gateway-default",
            "target": "reclambdaservice-service-notefeed-recall-filter",
            "dependence": "strong",
            "prejudgment": {},
            "verify": {},
            "labels": [
                {
                    "name": "type",
                    "value": "direct"
                },
                {
                    "name": "scenario",
                    "value": "neiliu"
                },
                {
                    "name": "sourceType",
                    "value": "Service"
                },
                {
                    "name": "targetType",
                    "value": "Service"
                }
            ]
        }
    },
    {
        "_source": {
            "source": "notefeed-gateway-default",
            "target": "userrelation-service-default",
            "dependence": "strong",
            "prejudgment": {},
            "verify": {},
            "labels": [
                {
                    "name": "type",
                    "value": "direct"
                },
                {
                    "name": "scenario",
                    "value": "neiliu"
                },
                {
                    "name": "sourceType",
                    "value": "Service"
                },
                {
                    "name": "targetType",
                    "value": "Service"
                }
            ]
        }
    },
    {
        "_source": {
            "source": "notefeed-gateway-default",
            "target": "实验平台",
            "dependence": "strong",
            "prejudgment": {},
            "verify": {},
            "labels": [
                {
                    "name": "type",
                    "value": "direct"
                },
                {
                    "name": "scenario",
                    "value": "neiliu"
                },
                {
                    "name": "sourceType",
                    "value": "Service"
                },
                {
                    "name": "targetType",
                    "value": "Service"
                }
            ]
        }
    },
    {
        "_source": {
            "source": "notefeed-gateway-default",
            "target": "Apollo",
            "dependence": "strong",
            "prejudgment": {},
            "verify": {},
            "labels": [
                {
                    "name": "type",
                    "value": "direct"
                },
                {
                    "name": "scenario",
                    "value": "neiliu"
                },
                {
                    "name": "sourceType",
                    "value": "Service"
                },
                {
                    "name": "targetType",
                    "value": "Service"
                }
            ]
        }
    },
    {
        "_source": {
            "source": "aceflow-notefeed-default",
            "target": "Redis",
            "dependence": "strong",
            "prejudgment": {},
            "verify": {},
            "labels": [
                {
                    "name": "type",
                    "value": "direct"
                },
                {
                    "name": "scenario",
                    "value": "neiliu"
                },
                {
                    "name": "test",
                    "value": "true"
                },
                {
                    "name": "sourceType",
                    "value": "Service"
                },
                {
                    "name": "targetType",
                    "value": "Service"
                }
            ]
        }
    },
    {
        "_source": {
            "source": "aceflow-notefeed-default",
            "target": "reddedupservice-service-videofeed",
            "dependence": "weak",
            "prejudgment": {},
            "verify": {},
            "labels": [
                {
                    "name": "type",
                    "value": "direct"
                },
                {
                    "name": "scenario",
                    "value": "neiliu"
                },
                {
                    "name": "test",
                    "value": "false"
                },
                {
                    "name": "sourceType",
                    "value": "Service"
                },
                {
                    "name": "targetType",
                    "value": "Service"
                }
            ]
        }
    }
]`

var (
	b    base.Base
	line Line
	once sync.Once
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
		line = *New(context.Background(), b)
	})
}

func Trans(jsonString string) (realLineList []*lineModel.Line) {
	lineList := make([]struct {
		Source lineModel.Line `json:"_source"`
	}, 0)
	err := json.Unmarshal([]byte(jsonString), &lineList)
	if err != nil {
		log.Printf("Bulk() error = %v", err)
	}
	realLineList = make([]*lineModel.Line, len(lineList))
	for i, line1 := range lineList {
		line1.Source.SourceScene = "inflow"
		line1.Source.TargetScene = "inflow"
		line1.Source.SourceIsCore = false
		line1.Source.TargetIsCore = false
		line1.Source.VisitCount = 0
		realLineList[i] = &lineModel.Line{
			Source:         line1.Source.Source,
			SourceIsCore:   false,
			SourceScene:    "inflow",
			Target:         line1.Source.Target,
			TargetIsCore:   false,
			TargetScene:    "inflow",
			Dependence:     line1.Source.Dependence,
			LastFoundTime:  0,
			FirstFoundTime: 0,
			VisitCount:     rand.Int63n(50),
		}
	}
	return
}

func TestLine_Trans(t *testing.T) {
	res := Trans(lineListJson)
	t.Log(res)
}

func TestLine_Bulk(t *testing.T) {
	_init()
	lines := []*lineModel.Line{
		{
			Source:         "test1",
			SourceIsCore:   false,
			SourceScene:    "inflow",
			Target:         "test2",
			TargetIsCore:   false,
			TargetScene:    "inflow",
			Dependence:     "weak",
			LastFoundTime:  0,
			FirstFoundTime: 0,
			VisitCount:     66,
		},
	}
	//res := Trans(lineListJson)
	t.Run("Bulk", func(t *testing.T) {
		if err := line.Bulk(context.TODO(), lines); err != nil {
			t.Errorf("Bulk() error = %v", err)
		}
	})
}

func TestLine_BulkDeleteById(t *testing.T) {
	_init()
	line1 := "s*test1*t*test2"
	t.Run("BulkDeleteById", func(t *testing.T) {
		if ok, err := line.DeleteById(context.TODO(), "inflow", line1); err != nil {
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
	t.Run("TopologyAnalyse", func(t *testing.T) {
		ok, indicator := line.TopologyAnalyse(context.TODO(), &req)
		t.Log(ok)
		if ok {
			data, _ := jsoniter.MarshalIndent(indicator, "", " ")
			t.Log(string(data))
		}
	})
}

func TestLine_InitializeConsumers(t *testing.T) {
	_init()
	t.Run("InitializeConsumers", func(t *testing.T) {
		line.InitializeConsumers()
	})
}

func TestLine_SparkQuery(t *testing.T) {
	_init()
	t.Run("SparkQuery", func(t *testing.T) {
		a, b, c := line.SparkQuery(context.TODO(), "spark_result")
		if c != nil {
			t.Log(c)
		} else {
			t.Log(a)
			t.Log(b)
		}
	})
}

func TestLine_GetRank(t *testing.T) {
	_init()
	t.Run("GetRank", func(t *testing.T) {
		a, b := line.GetRank(context.TODO(), "inflow")
		if b != nil {
			t.Log(b)
		} else {
			t.Log(a)
		}
	})
}
