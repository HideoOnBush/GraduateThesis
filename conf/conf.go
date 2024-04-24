package conf

import (
	"flag"
	"os"
	"time"
)

var (
	GConfig *Config
)

const TimeFormat = "2006-01-02 15:04:05"

type ServiceConfig struct {
	ServiceName   string
	Addr          string
	ConnTimeout   time.Duration
	SocketTimeout time.Duration
	Framed        bool
}

type MysqlServiceConfig struct {
	Host     string
	Port     int
	Db       string
	User     string
	Password string
}

type Config struct {
	SparkHost            string
	RedHaLeader          string
	RabbitConCurrencyNum string
	RabbitMqAddr         string
	RabbitMqPort         string
	RabbitMqUser         string
	RabbitMqPwd          string
	Neo4jAddr            string
	Neo4jUsername        string
	Neo4jPassword        string
	ESAddr               string
	ESUsername           string
	ESPassword           string
	RocketMqServer       string
	RedisAddr            string
	RedisPassword        string
	RedisDB              int
	Mysql                MysqlServiceConfig
}

func init() {
	GConfig = new(Config)
}

func Init() {
	flag.StringVar(&GConfig.SparkHost, "spark.host", "http://127.0.0.1:4040", "spark host")
	flag.StringVar(&GConfig.RedHaLeader, "redha.leader", "redha-leader", "redha leader")
	flag.StringVar(&GConfig.RabbitConCurrencyNum, "rabbitmq.concurrency", "1", "rabbitmq.concurrency")
	flag.StringVar(&GConfig.RabbitMqAddr, "rabbitmq.addr", "127.0.0.1", "rabbitmq.addr")
	flag.StringVar(&GConfig.RabbitMqPort, "rabbitmq.port", "5672", "rabbitmq.port")
	flag.StringVar(&GConfig.RabbitMqUser, "rabbitmq.user", "kwq", "rabbitmq.user")
	flag.StringVar(&GConfig.RabbitMqPwd, "rabbitmq.pwd", "123456", "rabbitmq.pwd")
	flag.StringVar(&GConfig.Neo4jAddr, "neo.host", "neo4j://127.0.0.1", "neo host")
	flag.StringVar(&GConfig.Neo4jUsername, "neo.username", "neo4j", "neo username")
	flag.StringVar(&GConfig.Neo4jPassword, "neo.password", "12345678", "neo password")
	flag.StringVar(&GConfig.ESAddr, "es.host", "http://127.0.0.1:9200", "es host")
	flag.StringVar(&GConfig.ESUsername, "es.username", "", "es username")
	flag.StringVar(&GConfig.ESPassword, "es.password", "", "es password")
	flag.StringVar(&GConfig.RedisAddr, "redis.addr", "127.0.0.1:6379", "redis addr")
	flag.StringVar(&GConfig.RedisPassword, "redis.password", "", "redis password")
	flag.IntVar(&GConfig.RedisDB, "redis.db", 1, "redis db")
	flag.Parse()
}

func getEnvOrDefault(key string, defaultValue string) string {
	result := os.Getenv(key)
	if result == "" {
		return defaultValue
	}
	return result
}
