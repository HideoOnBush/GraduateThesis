package platform

import (
	"GraduateThesis/conf"
	"GraduateThesis/platform/es"
	"GraduateThesis/platform/neo"
	"GraduateThesis/platform/rabbitmq"
	"GraduateThesis/platform/spark"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
)

var (
	P Platform
)

type Platform struct {
	ES       *elastic.Client
	NEO4J    *neo4j.DriverWithContext
	RABBITMQ *rabbitmq.Connection
	SPARK    *spark.SparkClient
}

func (p *Platform) Init(c *conf.Config) {
	esClient, err := es.New(es.Config{
		ESAddr:     c.ESAddr,
		ESUsername: c.ESUsername,
		ESPassword: c.ESPassword,
	})
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	} else {
		p.ES = esClient
	}
	NEO4JClient, err := neo.New(neo.Config{
		Neo4jAddr:     c.Neo4jAddr,
		Neo4jUsername: c.Neo4jUsername,
		Neo4jPassword: c.Neo4jPassword,
	})
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	} else {
		p.NEO4J = NEO4JClient
	}
	RABBITMQClient, err := rabbitmq.Init(rabbitmq.Conf{
		Addr: c.RabbitMqAddr,
		Port: c.RabbitMqPort,
		User: c.RabbitMqUser,
		Pwd:  c.RabbitMqPwd,
	})
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	} else {
		p.RABBITMQ = RABBITMQClient
	}
	p.SPARK = spark.New(spark.Config{
		Host: c.SparkHost,
	})
}

func Init(c *conf.Config) {
	P.Init(c)
}
