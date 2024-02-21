package neo

import (
	"context"
	"errors"
	"fmt"
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

func (n *NeoClient) insertItem(ctx context.Context) (*Item, error) {
	result, err := neo4j.ExecuteQuery(ctx, *n.client,
		"CREATE (n:Item { id: $id, name: $name }) RETURN n",
		map[string]any{
			"id":   1,
			"name": "Item 1",
		}, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}
	itemNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "n")
	if err != nil {
		return nil, fmt.Errorf("could not find node n")
	}
	id, err := neo4j.GetProperty[int64](itemNode, "id")
	if err != nil {
		return nil, err
	}
	name, err := neo4j.GetProperty[string](itemNode, "name")
	if err != nil {
		return nil, err
	}
	return &Item{Id: id, Name: name}, nil
}

type Item struct {
	Id   int64
	Name string
}

func (i *Item) String() string {
	return fmt.Sprintf("Item (id: %d, name: %q)", i.Id, i.Name)
}
