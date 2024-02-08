package line

import (
	model "GraduateThesis/biz/model/line"
	"GraduateThesis/platform/es"
	"github.com/olivere/elastic/v7"
)

const (
	StrongDependence = "strong"
	WeakDependence   = "weak"
)

const (
	indexMapping = `{
    "settings": {
        "index": {
            "max_result_window": "1000000000"
        }
    },
    "mappings": {
        "properties": {
            "dependence": {
                "type": "keyword"
            },
            "source": {
                "type": "keyword"
            },
			"sourceIsCore": {
                "type": "boolean"
            },
			"sourceScene": {
                "type": "keyword"
            },
            "target": {
                "type": "keyword"
            },
			"targetIsCore": {
                "type": "boolean"
            },
			"targetScene": {
                "type": "keyword"
            },
            "firstFoundTime": {
              "type": "date"
            },
            "lastFoundTime": {
                "type": "date"
            }
        }
    }
}`
)

type PassStatus struct {
	Status string `json:"status,omitempty"`
}

type Filter struct {
	Sources      []string `json:"sources,omitempty"`
	Targets      []string `json:"targets,omitempty"`
	Dependencies []string `json:"dependencies,omitempty"`
	MatchAll     bool     `json:"match_all,omitempty"`
	PageSize     int64    `json:"page_size,omitempty"`
	From         int64    `json:"from,omitempty"`
	IgnoreIds    []string `json:"ignore_ids,omitempty"`
	IsCore       bool     `json:"is_core,omitempty"`
	Scene        string   `json:"scene,omitempty"`
}

func ParseReqToFilter(req model.LineReq) (filter *Filter, err error) {
	filter = new(Filter)
	filter.Sources = req.Sources
	filter.Targets = req.Targets
	filter.Dependencies = req.Dependencies
	filter.MatchAll = req.GetMatchAll()
	filter.IgnoreIds = req.GetIgnoreIds()
	filter.IsCore = req.GetIsCore()
	filter.Scene = req.GetScene()
	if req.PageSize != 0 {
		filter.PageSize = req.PageSize
		if filter.PageSize <= 0 {
			filter.PageSize = 15
		}
		if req.PageNum != 0 {
			pageNum := req.PageNum
			if pageNum <= 0 {
				pageNum = 1
			}
			filter.From = (pageNum - 1) * filter.PageSize
		}
	}
	return
}

func (f *Filter) ParseToSearcher() (searcher *es.Searcher) {
	searcher = new(es.Searcher)
	searcher.From = int(f.From)
	searcher.Size = int(f.PageSize)

	if f.MatchAll {
		searcher.MatchAll = true
		return
	}

	if len(f.Sources) > 0 {
		switch len(f.Sources) {
		case 1:
			searcher.Filters = append(searcher.Filters, elastic.NewTermsQuery("source", f.Sources[0]))
		default:
			sourceBoolQuery := elastic.NewBoolQuery()
			for _, source := range f.Sources {
				sourceBoolQuery.Should(elastic.NewTermsQuery("source", source))
			}
			sourceBoolQuery.MinimumShouldMatch("1")
			searcher.Filters = append(searcher.Filters, sourceBoolQuery)
		}
	}

	if len(f.Targets) > 0 {
		switch len(f.Targets) {
		case 1:
			searcher.Filters = append(searcher.Filters, elastic.NewTermsQuery("target", f.Targets[0]))
		default:
			sourceBoolQuery := elastic.NewBoolQuery()
			for _, target := range f.Targets {
				sourceBoolQuery.Should(elastic.NewTermsQuery("target", target))
			}
			sourceBoolQuery.MinimumShouldMatch("1")
			searcher.Filters = append(searcher.Filters, sourceBoolQuery)
		}
	}

	if len(f.Dependencies) > 0 {
		switch len(f.Dependencies) {
		case 1:
			searcher.Filters = append(searcher.Filters, elastic.NewTermsQuery("dependence", f.Dependencies[0]))
		default:
			dependencyBoolQuery := elastic.NewBoolQuery()
			for _, dependence := range f.Dependencies {
				dependencyBoolQuery.Should(elastic.NewTermsQuery("dependence", dependence))
			}
			dependencyBoolQuery.MinimumShouldMatch("1")
			searcher.Filters = append(searcher.Filters, dependencyBoolQuery)
		}
	}

	if f.IgnoreIds != nil {
		searcher.Filters = append(searcher.Filters, elastic.NewBoolQuery().MustNot(elastic.NewIdsQuery().Ids(f.IgnoreIds...)))
	}

	if f.IsCore {
		searcher.Filters = append(searcher.Filters, elastic.NewBoolQuery().
			Should(elastic.NewTermsQuery("sourceIsCore", true)).
			Should(elastic.NewTermsQuery("targetIsCore", true)))
	}

	return
}
