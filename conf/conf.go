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
	Neo4jAddr      string
	Neo4jUsername  string
	Neo4jPassword  string
	ESAddr         string
	ESUsername     string
	ESPassword     string
	RocketMqServer string
	RedisAddr      string
	RedisPassword  string
	RedisDB        int
	Mysql          MysqlServiceConfig
}

func init() {
	GConfig = new(Config)
}

func Init() {
	flag.StringVar(&GConfig.Neo4jAddr, "neo.host", "neo4j://127.0.0.1", "neo host")
	flag.StringVar(&GConfig.Neo4jUsername, "neo.username", "neo4j", "neo username")
	flag.StringVar(&GConfig.Neo4jPassword, "neo.password", "12345678", "neo password")
	flag.StringVar(&GConfig.ESAddr, "es.host", "http://127.0.0.1:9200", "es host")
	flag.StringVar(&GConfig.ESUsername, "es.username", "", "es username")
	flag.StringVar(&GConfig.ESPassword, "es.password", "", "es password")
	flag.StringVar(&GConfig.RocketMqServer, "rocketmq.server", "10.4.44.168:9876", "rocketmq server")
	flag.StringVar(&GConfig.RedisAddr, "redis.addr", "127.0.0.1:6379", "redis addr")
	flag.StringVar(&GConfig.RedisPassword, "redis.password", "", "redis password")
	flag.IntVar(&GConfig.RedisDB, "redis.db", 1, "redis db")
	flag.StringVar(&GConfig.Mysql.User, "mysql.user", "redha_rw", "mysql user")
	flag.StringVar(&GConfig.Mysql.Password, "mysql.password", "kXfkW0Mc5fWU6J*V", "mysql password")
	flag.StringVar(&GConfig.Mysql.Host, "mysql.host", "10.4.41.80", "mysql host")
	flag.IntVar(&GConfig.Mysql.Port, "mysql.port", 33071, "mysql port")
	flag.StringVar(&GConfig.Mysql.Db, "mysql.db", "redha", "mysql db")
	flag.Parse()
}

func getEnvOrDefault(key string, defaultValue string) string {
	result := os.Getenv(key)
	if result == "" {
		return defaultValue
	}
	return result
}
