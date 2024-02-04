package es

import (
	"errors"
	"github.com/olivere/elastic/v7"
	"net/url"
)

type Config struct {
	ESAddr     string
	ESUsername string
	ESPassword string
}

type Searcher struct {
	MustQuery    []elastic.Query
	MustNotQuery []elastic.Query
	ShouldQuery  []elastic.Query
	Filters      []elastic.Query
	Sorters      []elastic.Sorter
	From         int //分页
	Size         int
	MatchAll     bool
}

func (s *Searcher) ToQuery() elastic.Query {
	if s.MatchAll {
		return elastic.NewMatchAllQuery()
	}

	boolQuery := elastic.NewBoolQuery()
	boolQuery.Must(s.MustQuery...)
	boolQuery.MustNot(s.MustNotQuery...)
	boolQuery.Should(s.ShouldQuery...)
	boolQuery.Filter(s.Filters...)

	// 当should不为空时，保证至少匹配should中的一项
	if len(s.ShouldQuery) > 0 {
		boolQuery.MinimumShouldMatch("1")
	}
	return boolQuery
}

func New(cfg Config) (*elastic.Client, error) {
	if cfg.ESAddr == "" {
		return nil, errors.New("must specify elasticsearch addr")
	}
	u, err := url.Parse(cfg.ESAddr)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "" {
		u.Scheme = "http"
	}
	return elastic.NewClient(elastic.SetURL(u.String()), elastic.SetBasicAuth(cfg.ESUsername, cfg.ESPassword), elastic.SetSniff(false))
}
