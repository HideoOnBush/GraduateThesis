package neo

import (
	"errors"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Config struct {
	Neo4jAddr     string
	Neo4jUsername string
	Neo4jPassword string
}

type NeoClient struct {
	client *neo4j.DriverWithContext
}

func New(cfg Config) (*neo4j.DriverWithContext, error) {
	if cfg.Neo4jAddr == "" {
		return nil, errors.New("must specify neo addr")
	}
	dbUri := cfg.Neo4jAddr // scheme://host(:port) (default port is 7687)
	driver, err := neo4j.NewDriverWithContext(dbUri, neo4j.NoAuth())
	if err != nil {
		return nil, errors.New("initialize neo failed")
	}
	return &driver, nil
}
