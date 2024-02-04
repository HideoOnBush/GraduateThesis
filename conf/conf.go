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
	ESAddr             string
	ESUsername         string
	ESPassword         string
	ChaosAddr          string
	ChaosDebug         bool
	RocketMqServer     string
	RedisAddr          string
	RedisPassword      string
	RedisDB            int
	RedHaLeader        string
	TreeHost           string
	TreeApp            string
	TreeSecret         string
	TreeDebug          bool
	OnesHost           string
	OnesApp            string
	OnesSecret         string
	OnesDebug          bool
	AccountService     ServiceConfig
	EdsHosts           string
	EnableAuth         bool
	Mysql              MysqlServiceConfig
	NocHost            string
	NocDebug           bool
	NocToken           string
	XrayHost           string
	XrayDebug          bool
	XrayApp            string
	XraySecret         string
	XrayToken          string
	RedcloudHost       string
	RedcloudDebug      bool
	RcaHost            string
	RcaDebug           bool
	ForbidVerifyPeriod string
	CrossRoadHost      string
	CrossRoadDebug     bool
	CrossRoadAgentId   int
	CrossRoadToken     string
	RcmHost            string
	RcmDebug           bool
}

func init() {
	GConfig = new(Config)
}

func Init() {
	flag.StringVar(&GConfig.ESAddr, "es.host", "http://127.0.0.1:9200", "es host")
	flag.StringVar(&GConfig.ESUsername, "es.username", "", "es username")
	flag.StringVar(&GConfig.ESPassword, "es.password", "", "es password")
	flag.StringVar(&GConfig.ChaosAddr, "chaos.host", "http://srechaos.devops.xiaohongshu.com", "chaos host")
	flag.BoolVar(&GConfig.ChaosDebug, "chaos.debug", true, "chaos debug")
	flag.StringVar(&GConfig.RocketMqServer, "rocketmq.server", "10.4.44.168:9876", "rocketmq server")
	flag.StringVar(&GConfig.RedisAddr, "redis.addr", "127.0.0.1:6379", "redis addr")
	flag.StringVar(&GConfig.RedisPassword, "redis.password", "", "redis password")
	flag.IntVar(&GConfig.RedisDB, "redis.db", 1, "redis db")
	flag.StringVar(&GConfig.TreeHost, "tree.host", "http://tree.devops.xiaohongshu.com", "服务树 host")
	flag.StringVar(&GConfig.TreeApp, "tree.app", "ones", "服务树平台账号")
	flag.StringVar(&GConfig.TreeSecret, "tree.secret", "", "服务树平台账号密钥")
	flag.BoolVar(&GConfig.TreeDebug, "tree.debug", false, "调试服务树")
	flag.StringVar(&GConfig.OnesHost, "ones.host", "http://ones.devops.xiaohongshu.com", "ones host")
	flag.StringVar(&GConfig.OnesApp, "ones.app", "ones", "ones 平台帐号")
	flag.StringVar(&GConfig.OnesSecret, "ones.secret", "", "ones 平台账号密钥")
	flag.BoolVar(&GConfig.OnesDebug, "ones.debug", false, "调试ones")
	flag.StringVar(&GConfig.AccountService.ServiceName, "account.servicename", "com.xiaohongshu.fls.rpc.skywalker.account.service.AccountService", "account service")
	flag.StringVar(&GConfig.AccountService.Addr, "account.addr", "", "account addr")
	flag.DurationVar(&GConfig.AccountService.ConnTimeout, "account.conn-timeout", 300, "account conn-timeout")
	flag.DurationVar(&GConfig.AccountService.SocketTimeout, "account.socket-timeout", 1000, "account socket-timeout")
	flag.BoolVar(&GConfig.AccountService.Framed, "account.framed", true, "account framed")
	flag.StringVar(&GConfig.EdsHosts, "eds-hosts", "10.4.28.45:80", "eds hosts")
	flag.BoolVar(&GConfig.EnableAuth, "enable-auth", true, "enable auth")
	flag.StringVar(&GConfig.Mysql.User, "mysql.user", "redha_rw", "mysql user")
	flag.StringVar(&GConfig.Mysql.Password, "mysql.password", "kXfkW0Mc5fWU6J*V", "mysql password")
	flag.StringVar(&GConfig.Mysql.Host, "mysql.host", "10.4.41.80", "mysql host")
	flag.IntVar(&GConfig.Mysql.Port, "mysql.port", 33071, "mysql port")
	flag.StringVar(&GConfig.Mysql.Db, "mysql.db", "redha", "mysql db")
	flag.StringVar(&GConfig.NocHost, "noc.host", "http://noc.devops.xiaohongshu.com", "noc host")
	flag.StringVar(&GConfig.NocToken, "noc.token", "e54cad4a73c3428299c953e16eacc26f", "noc token")
	flag.BoolVar(&GConfig.NocDebug, "noc.debug", false, "debug noc")
	flag.StringVar(&GConfig.XrayHost, "xray.host", "http://xray.int.xiaohongshu.com", "xray host")
	flag.BoolVar(&GConfig.XrayDebug, "xray.debug", false, "debug xray")
	flag.StringVar(&GConfig.XrayApp, "xray.app", "ones", "xray 平台帐号")
	flag.StringVar(&GConfig.XraySecret, "xray.secret", "", "xray 平台账号密钥")
	flag.StringVar(&GConfig.XrayToken, "xray.token", "6868046b-06cb-4d13-a2a2-3359bafada88", "xray 平台的token")
	flag.StringVar(&GConfig.RedHaLeader, "redha.leader", "redha-leader", "redha leader")
	flag.StringVar(&GConfig.RedcloudHost, "redcloud.host", "http://redcloud.int.xiaohongshu.com", "redcloud host")
	flag.BoolVar(&GConfig.RedcloudDebug, "redcloud.debug", false, "debug redcloud")
	flag.StringVar(&GConfig.RcaHost, "rca.host", "https://rca.devops.sit.xiaohongshu.com", "rca host")
	flag.BoolVar(&GConfig.RcaDebug, "rca.debug", false, "debug rca")
	flag.StringVar(&GConfig.CrossRoadHost, "crossroad.host", "http://crossroad.int.xiaohongshu.com", "crossroad host")
	flag.BoolVar(&GConfig.CrossRoadDebug, "crossroad.debug", false, "debug crossroad")
	flag.StringVar(&GConfig.CrossRoadToken, "crossroad.token", "fd81d87b00d94f72bfa914ce1700e8aa", "crossroad token")
	flag.IntVar(&GConfig.CrossRoadAgentId, "crossroad.agentid", 1000162, "crossroad agentId")
	flag.StringVar(&GConfig.ForbidVerifyPeriod, "forbid-verify-period", "[[\"00:00:00\",\"10:00:00\"],[\"11:30:00\",\"13:30:00\"],[\"20:00:00\",\"23:59:59\"]]", "禁止验证的时间")
	flag.StringVar(&GConfig.RcmHost, "rcm.host", "http://rcm.devops.xiaohongshu.com", "rcm host")
	flag.BoolVar(&GConfig.RcmDebug, "rcm.debug", false, "rcm debug")
	flag.Parse()
}

func getEnvOrDefault(key string, defaultValue string) string {
	result := os.Getenv(key)
	if result == "" {
		return defaultValue
	}
	return result
}
