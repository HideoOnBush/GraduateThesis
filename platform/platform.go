package platform

import (
	"GraduateThesis/conf"
	"GraduateThesis/platform/es"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
)

var (
	P Platform
)

type Platform struct {
	ES *elastic.Client
}

func (p *Platform) Init(c *conf.Config) {
	client, err := es.New(es.Config{
		ESAddr:     c.ESAddr,
		ESUsername: c.ESUsername,
		ESPassword: c.ESPassword,
	})
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	} else {
		p.ES = client
	}
}

func Init(c *conf.Config) {
	P.Init(c)
}
