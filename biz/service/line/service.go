package line

import (
	model "GraduateThesis/biz/model/line"
	"GraduateThesis/biz/service/base"
	"GraduateThesis/platform/es"
	"bytes"
	"container/list"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	jsoniter "github.com/json-iterator/go"
	"github.com/olivere/elastic/v7"
	"log"
	"time"
)

type Line struct {
	base.Base
}

func New(ctx context.Context, base base.Base) *Line {
	line := &Line{Base: base}
	go line.InitializeConsumers()
	return line
}

func (l *Line) GetIndex() []string {
	query := elastic.NewCatIndicesService(l.ES).
		Columns("index").
		Pretty(true)

	result, err := query.Do(context.Background())
	if err != nil {
		panic(err)
	}

	var ans []string
	for _, index := range result {
		ans = append(ans, index.Index)
	}
	return ans
}

func (l *Line) searchWithSearcher(ctx context.Context, searchFilter *es.Searcher, index string) (lines []*model.Line, totalCount int, err error) {
	searcher := l.ES.Search().Version(true).SeqNoAndPrimaryTerm(true).Index(index).Query(searchFilter.ToQuery()).SortBy(searchFilter.Sorters...).From(searchFilter.From).Size(searchFilter.Size).TrackTotalHits(true)
	searchResp, err := searcher.Do(ctx)
	if err != nil {
		e, _ := jsoniter.MarshalToString(err)
		log.Println(e)
		return
	}
	if searchResp.Hits == nil || searchResp.Hits.Hits == nil || len(searchResp.Hits.Hits) == 0 {
		return
	}
	totalCount = int(searchResp.Hits.TotalHits.Value)
	lines = make([]*model.Line, 0, len(searchResp.Hits.Hits))
	for _, hit := range searchResp.Hits.Hits {
		id := hit.Id
		line := &model.Line{}
		if hit.Source == nil {
			continue
		}
		log.Println(string(hit.Source))
		if err = json.Unmarshal(hit.Source, line); err == nil {
			line.DocId = id
			line.DocVersion = *hit.Version
			line.DocSeqNo = *hit.SeqNo
			line.DocPrimaryTerm = *hit.PrimaryTerm
			lines = append(lines, line)
		} else {
			log.Println(err.Error())
			return
		}
	}
	return
}

func (l *Line) search(ctx context.Context, filter *Filter, index string) (lines []*model.Line, totalCount int, err error) {
	lines = make([]*model.Line, 0)
	searchFilter := filter.ParseToSearcher()
	searchFilter.Sorters = append(searchFilter.Sorters, elastic.NewFieldSort("source").Asc(), elastic.NewFieldSort("target").Asc())

	lines, totalCount, err = l.searchWithSearcher(ctx, searchFilter, index)
	return
}

func (l *Line) Query(ctx context.Context, app *app.RequestContext, req *model.LineReq) (lines []*model.Line, totalCount int) {
	index := l.SceneToIndex(req.GetScene())
	exist, err := l.ExistsIndex(ctx, index)
	if err != nil {
		_ = app.Error(err)
		return
	}
	if !exist {
		_ = app.Error(errors.New("查询index错误或者不存在"))
		return
	}
	if req.PageSize == 0 {
		var size int64 = 15
		req.PageSize = size
	}
	if req.PageNum == 0 {
		var num int64 = 1
		req.PageNum = num
	}

	filter, err := ParseReqToFilter(*req)
	if err != nil {
		_ = app.Error(err)
		return
	}

	lines, totalCount, err = l.search(ctx, filter, index)
	if err != nil {
		_ = app.Error(err)
		return
	}
	return
}

func (l *Line) CreateIndex(ctx context.Context, sceneName string) (string, error) {
	index := l.SceneToIndex(sceneName)
	exist, err := l.ExistsIndex(ctx, index)
	if err != nil {
		return index, err
	}
	if !exist {
		createIndex, err := l.ES.CreateIndex(index).BodyString(indexMapping).Do(ctx)
		if err != nil || !createIndex.Acknowledged {
			return index, errors.New("索引创建失败" + err.Error())
		}
	}
	return index, nil
}

func (l *Line) ExistsIndex(ctx context.Context, index string) (bool, error) {
	exist, err := l.ES.IndexExists(index).Do(ctx)
	if err != nil {
		return true, err
	}
	if !exist {
		return false, nil
	} else {
		return true, nil
	}
}

func (l *Line) SceneToIndex(scene string) string {
	var buffer bytes.Buffer
	buffer.WriteString(scene)
	buffer.WriteString("_")
	buffer.WriteString("lines")
	index := buffer.String()
	return index
}

func (l *Line) ExistsScene(ctx context.Context, sceneName string) (bool, error) {
	index := l.SceneToIndex(sceneName)
	exist, err := l.ExistsIndex(ctx, index)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, nil
	}
	return true, nil
}

func (l *Line) Bulk(ctx context.Context, lines []*model.Line) (err error) {
	bulkRequest := l.ES.Bulk().Refresh("true")
	now := time.Now().Unix()
	for _, line := range lines {
		GenerateLineDocId(line)
		if line.GetFirstFoundTime() == 0 {
			line.FirstFoundTime = now
		}
		if line.GetLastFoundTime() == 0 {
			line.LastFoundTime = now
		}
		sceneMap := map[string]struct{}{}
		sceneMap[line.GetSourceScene()] = struct{}{}
		sceneMap[line.GetTargetScene()] = struct{}{}
		for sceneName, _ := range sceneMap {
			var index string
			ex, err := l.ExistsScene(ctx, sceneName)
			if err != nil {
				log.Printf("查看是否存在scene(%s)index错误\n", sceneName)
				return err
			}
			if !ex {
				index, err = l.CreateIndex(ctx, sceneName)
				if err != nil {
					return err
				}
				log.Printf("默认创建了index %s\n", index)
			} else {
				index = l.SceneToIndex(sceneName)
			}
			doc := elastic.NewBulkUpdateRequest().
				Index(index).
				Id(line.GetDocId()).
				Doc(struct {
					Source         string `thrift:"source,5" json:"source" form:"source" query:"source"`
					SourceIsCore   bool   `thrift:"sourceIsCore,6" json:"sourceIsCore" form:"sourceIsCore" query:"sourceIsCore"`
					SourceScene    string `thrift:"sourceScene,7" json:"sourceScene" form:"sourceScene" query:"sourceScene"`
					Target         string `thrift:"target,8" json:"target" form:"target" query:"target"`
					TargetIsCore   bool   `thrift:"targetIsCore,9" json:"targetIsCore" form:"targetIsCore" query:"targetIsCore"`
					TargetScene    string `thrift:"targetScene,10" json:"targetScene" form:"targetScene" query:"targetScene"`
					Dependence     string `thrift:"dependence,11" json:"dependence" form:"dependence" query:"dependence"`
					LastFoundTime  int64  `thrift:"lastFoundTime,12" json:"lastFoundTime" form:"lastFoundTime" query:"lastFoundTime"`
					FirstFoundTime int64  `thrift:"firstFoundTime,13" json:"firstFoundTime" form:"firstFoundTime" query:"firstFoundTime"`
				}{
					Source:         line.GetSource(),
					SourceIsCore:   line.GetSourceIsCore(),
					SourceScene:    line.GetSourceScene(),
					Target:         line.GetTarget(),
					TargetIsCore:   line.GetTargetIsCore(),
					TargetScene:    line.GetTargetScene(),
					Dependence:     line.GetDependence(),
					LastFoundTime:  line.GetLastFoundTime(),
					FirstFoundTime: line.GetFirstFoundTime(),
				}).
				DocAsUpsert(true)
			bulkRequest.Add(doc)
		}
	}
	bulkResponse, err := bulkRequest.Do(ctx)
	if err != nil {
		return
	}

	bad := bulkResponse.Failed()
	if len(bad) > 0 {
		s, _ := jsoniter.MarshalToString(bad)
		err = errors.New("部分记录更新失败 " + s)
	}

	return
}

func (l *Line) BulkUpdateDependence(ctx context.Context, lines map[string]string, index string) (err error) {
	if ex, err1 := l.ExistsIndex(ctx, index); !ex {
		err = err1
		return
	}
	bulkRequest := l.ES.Bulk().Refresh("true")
	for docId, dependence := range lines {
		doc := elastic.NewBulkUpdateRequest().
			Index(index).
			Id(docId).
			Doc(struct {
				Dependence string `thrift:"dependence,11" json:"dependence" form:"dependence" query:"dependence"`
			}{
				Dependence: dependence,
			})
		bulkRequest.Add(doc)
	}
	bulkResponse, err := bulkRequest.Do(ctx)
	if err != nil {
		return
	}

	bad := bulkResponse.Failed()
	if len(bad) > 0 {
		s, _ := jsoniter.MarshalToString(bad)
		err = errors.New("部分记录更新失败 " + s)
	}

	return
}

func (l *Line) BulkCreate(ctx context.Context, lines []*model.Line, index string) (bad []*elastic.BulkResponseItem, err error, ids []string) {
	bulkRequest := l.ES.Bulk().Refresh("true")
	ids = make([]string, 0, len(lines))
	for _, line := range lines {
		GenerateLineDocId(line)
		doc := elastic.NewBulkCreateRequest().
			Index(index).
			Type("_doc").
			Id(line.GetDocId()).
			Doc(struct {
				Source         string `thrift:"source,5" json:"source" form:"source" query:"source"`
				SourceIsCore   bool   `thrift:"sourceIsCore,6" json:"sourceIsCore" form:"sourceIsCore" query:"sourceIsCore"`
				SourceScene    string `thrift:"sourceScene,7" json:"sourceScene" form:"sourceScene" query:"sourceScene"`
				Target         string `thrift:"target,8" json:"target" form:"target" query:"target"`
				TargetIsCore   bool   `thrift:"targetIsCore,9" json:"targetIsCore" form:"targetIsCore" query:"targetIsCore"`
				TargetScene    string `thrift:"targetScene,10" json:"targetScene" form:"targetScene" query:"targetScene"`
				Dependence     string `thrift:"dependence,11" json:"dependence" form:"dependence" query:"dependence"`
				LastFoundTime  int64  `thrift:"lastFoundTime,12" json:"lastFoundTime" form:"lastFoundTime" query:"lastFoundTime"`
				FirstFoundTime int64  `thrift:"firstFoundTime,13" json:"firstFoundTime" form:"firstFoundTime" query:"firstFoundTime"`
			}{
				Source:         line.GetSource(),
				SourceIsCore:   line.GetSourceIsCore(),
				SourceScene:    line.GetSourceScene(),
				Target:         line.GetTarget(),
				TargetIsCore:   line.GetTargetIsCore(),
				TargetScene:    line.GetTargetScene(),
				Dependence:     line.GetDependence(),
				LastFoundTime:  line.GetLastFoundTime(),
				FirstFoundTime: line.GetFirstFoundTime(),
			})
		bulkRequest.Add(doc)
		ids = append(ids, line.DocId)
	}

	bulkResponse, err := bulkRequest.Do(ctx)
	if err != nil {
		return
	}

	bad = bulkResponse.Failed()
	if len(bad) > 0 {
		s, _ := jsoniter.MarshalToString(bad)
		err = errors.New("部分记录更新失败 " + s)
	}

	return
}

func (l *Line) BulkCreateFromChan(ctx context.Context, lines chan *model.Line, index string) (bad []*elastic.BulkResponseItem, err error, ids []interface{}) {
	if ex, err1 := l.ExistsIndex(ctx, index); !ex {
		err = err1
		return
	}
	bulkRequest := l.ES.Bulk().Refresh("true")
	ids = make([]interface{}, 0, len(lines))
	for {
		chanClose := false
		select {
		case line, ok := <-lines:
			if !ok {
				chanClose = true
				break
			}
			GenerateLineDocId(line)
			doc := elastic.NewBulkUpdateRequest().
				Index(index).
				Id(line.GetDocId()).Doc(struct {
				LastFoundTime int64 `thrift:"lastFoundTime,10" json:"lastFoundTime" form:"lastFoundTime" query:"lastFoundTime"`
			}{
				LastFoundTime: line.GetLastFoundTime(),
			}).
				Upsert(struct {
					Source         string `thrift:"source,5" json:"source" form:"source" query:"source"`
					SourceIsCore   bool   `thrift:"sourceIsCore,6" json:"sourceIsCore" form:"sourceIsCore" query:"sourceIsCore"`
					SourceScene    string `thrift:"sourceScene,7" json:"sourceScene" form:"sourceScene" query:"sourceScene"`
					Target         string `thrift:"target,8" json:"target" form:"target" query:"target"`
					TargetIsCore   bool   `thrift:"targetIsCore,9" json:"targetIsCore" form:"targetIsCore" query:"targetIsCore"`
					TargetScene    string `thrift:"targetScene,10" json:"targetScene" form:"targetScene" query:"targetScene"`
					Dependence     string `thrift:"dependence,11" json:"dependence" form:"dependence" query:"dependence"`
					LastFoundTime  int64  `thrift:"lastFoundTime,12" json:"lastFoundTime" form:"lastFoundTime" query:"lastFoundTime"`
					FirstFoundTime int64  `thrift:"firstFoundTime,13" json:"firstFoundTime" form:"firstFoundTime" query:"firstFoundTime"`
				}{
					Source:         line.GetSource(),
					SourceIsCore:   line.GetSourceIsCore(),
					SourceScene:    line.GetSourceScene(),
					Target:         line.GetTarget(),
					TargetIsCore:   line.GetTargetIsCore(),
					TargetScene:    line.GetTargetScene(),
					Dependence:     line.GetDependence(),
					LastFoundTime:  line.GetLastFoundTime(),
					FirstFoundTime: line.GetFirstFoundTime(),
				})
			bulkRequest.Add(doc)
			ids = append(ids, line.DocId)
		default:
		}
		if chanClose {
			break
		}
	}

	bulkResponse, err := bulkRequest.Do(ctx)
	if err != nil {
		return
	}

	bad = bulkResponse.Failed()
	if len(bad) > 0 {
		err = errors.New("部分记录更新失败 ")
	}

	return
}

func (l *Line) BulkUpdateFoundTime(ctx context.Context, index string, ids ...interface{}) (bool, error) {
	if ex, err1 := l.ExistsIndex(ctx, index); !ex {
		return false, err1
	}
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Filter(elastic.NewTermsQuery("_id", ids...))
	script := elastic.NewScript(`
        ctx._source.lastFoundTime = params.date
    `).Params(map[string]interface{}{
		"date": time.Now().Unix(),
	})
	exec := l.ES.UpdateByQuery().Index(index).Query(boolQuery).Script(script).Refresh("true")
	_, err := exec.Do(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (l *Line) DeleteById(ctx context.Context, scene string, id string) (bool, error) {
	index := l.SceneToIndex(scene)
	var err error
	if ex, err := l.ExistsIndex(ctx, index); !ex {
		return false, err
	}
	_, err = l.ES.Delete().Index(index).Id(id).Do(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (l *Line) Get(ctx context.Context, app *app.RequestContext, id string, index string) (line *model.Line) {
	getter := l.ES.Get().Index(index).Id(id)
	getResp, err := getter.Do(ctx)
	if err != nil {
		_ = app.Error(err)
		return
	}

	if getResp.Source == nil {
		_ = app.Error(errors.New("source is nil"))
		return
	}

	if !getResp.Found {
		return
	}

	line = &model.Line{}
	if err = json.Unmarshal(getResp.Source, line); err == nil {
		line.DocId = getResp.Id
		line.DocVersion = *getResp.Version
		line.DocSeqNo = *getResp.SeqNo
		line.DocPrimaryTerm = *getResp.PrimaryTerm
	} else {
		_ = app.Error(err)
	}
	return
}

func GenerateLineDocId(line *model.Line) {
	if line.GetDocId() != "" {
		return
	}
	id := fmt.Sprintf("s*%s*t*%s", line.GetSource(), line.GetTarget())
	line.DocId = id
}

func (l *Line) UpsertById(ctx context.Context, app *app.RequestContext, line *model.Line, index string) bool {
	if ex, _ := l.ExistsIndex(ctx, index); !ex {
		return false
	}
	if line == nil {
		_ = app.Error(errors.New("line id为空，禁止更新"))
		return false
	}
	if line.GetSource() == "" || line.GetTarget() == "" {
		_ = app.Error(errors.New("源或者目标不能为空"))
		return false
	}

	if line.GetDocId() == "" {
		GenerateLineDocId(line)
	}

	upserter := l.ES.Index().Index(index).Id(line.GetDocId()).BodyJson(struct {
		Source         string `thrift:"source,5" json:"source" form:"source" query:"source"`
		SourceIsCore   bool   `thrift:"sourceIsCore,6" json:"sourceIsCore" form:"sourceIsCore" query:"sourceIsCore"`
		SourceScene    string `thrift:"sourceScene,7" json:"sourceScene" form:"sourceScene" query:"sourceScene"`
		Target         string `thrift:"target,8" json:"target" form:"target" query:"target"`
		TargetIsCore   bool   `thrift:"targetIsCore,9" json:"targetIsCore" form:"targetIsCore" query:"targetIsCore"`
		TargetScene    string `thrift:"targetScene,10" json:"targetScene" form:"targetScene" query:"targetScene"`
		Dependence     string `thrift:"dependence,11" json:"dependence" form:"dependence" query:"dependence"`
		LastFoundTime  int64  `thrift:"lastFoundTime,12" json:"lastFoundTime" form:"lastFoundTime" query:"lastFoundTime"`
		FirstFoundTime int64  `thrift:"firstFoundTime,13" json:"firstFoundTime" form:"firstFoundTime" query:"firstFoundTime"`
	}{
		Source:         line.GetSource(),
		SourceIsCore:   line.GetSourceIsCore(),
		SourceScene:    line.GetSourceScene(),
		Target:         line.GetTarget(),
		TargetIsCore:   line.GetTargetIsCore(),
		TargetScene:    line.GetTargetScene(),
		Dependence:     line.GetDependence(),
		LastFoundTime:  line.GetLastFoundTime(),
		FirstFoundTime: line.GetFirstFoundTime(),
	}).Refresh("true")

	if line.GetDocSeqNo() > 0 && line.GetDocPrimaryTerm() > 0 {
		upserter.IfSeqNo(line.GetDocSeqNo()).IfPrimaryTerm(line.GetDocPrimaryTerm())
	}

	_, err := upserter.Do(ctx)
	if err != nil {
		_ = app.Error(err)
		return false
	}
	return true
}

func (l *Line) UpMergeById(ctx context.Context, app *app.RequestContext, id string, data interface{}, seqNo, primaryTerm int64, index string) bool {
	if id == "" {
		_ = app.Error(errors.New("id为空，禁止更新"))
		return false
	}

	upMerger := l.ES.Update().Index(index).Id(id).Doc(data).Refresh("true")
	if seqNo > 0 && primaryTerm > 0 {
		upMerger.IfSeqNo(seqNo).IfPrimaryTerm(primaryTerm)
	}
	_, err := upMerger.Do(ctx)
	if err != nil {
		_ = app.Error(err)
		return false
	}
	return true
}

func (l *Line) updateByQuery(ctx context.Context, app *app.RequestContext, index string, filter *Filter, script *elastic.Script) bool {
	searchFilter := filter.ParseToSearcher()
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Must(searchFilter.MustQuery...)
	boolQuery.MustNot(searchFilter.MustNotQuery...)
	boolQuery.Should(searchFilter.ShouldQuery...)
	boolQuery.Filter(searchFilter.Filters...)
	if len(searchFilter.ShouldQuery) > 0 {
		boolQuery.MinimumShouldMatch("1")
	}
	var err error
	_, err = l.ES.UpdateByQuery().Index(index).Query(boolQuery).Script(script).Refresh("true").Do(ctx)
	if err != nil {
		_ = app.Error(err)
		return false
	}
	return true
}

func (l *Line) ChangeDependenceById(ctx context.Context, app *app.RequestContext, id, dependence, startUpDependence, index string, seqNo, primaryTerm int64) bool {
	if ex, _ := l.ExistsIndex(ctx, index); !ex {
		return false
	}
	if dependence != StrongDependence && dependence != WeakDependence && dependence != "" {
		_ = app.Error(errors.New(fmt.Sprintf("不支持的依赖类型（%s, %s", StrongDependence, WeakDependence)))
		return false
	}

	patch := map[string]string{}

	if dependence != "" {
		patch["dependence"] = dependence
	}

	return l.UpMergeById(ctx, app, id, patch, seqNo, primaryTerm, index)
}

func (l *Line) ChangeDependenceByRelation(ctx context.Context, c *app.RequestContext, req model.ChangeDependenceWithRelationReq) bool {
	if req.Source == "" || req.Target == "" {
		_ = c.Error(errors.New("source 和 target 不能为空"))
		return false
	}

	if req.GetRelationType() != StrongDependence && req.GetRelationType() != WeakDependence && req.GetRelationType() != "" {
		_ = c.Error(errors.New(fmt.Sprintf("不支持的依赖类型（%s, %s", StrongDependence, WeakDependence)))
		return false
	}
	index := l.SceneToIndex(req.GetScene())
	exist, err := l.ExistsIndex(ctx, index)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	if !exist {
		log.Println("不可以使用不存在的index")
		return false
	}
	newLine := &model.Line{
		Source:     req.GetSource(),
		Target:     req.GetTarget(),
		Dependence: req.GetRelationType(),
	}
	_, _, _ = l.BulkCreate(ctx, []*model.Line{newLine}, index)
	filter := Filter{
		Sources: []string{req.GetSource()},
		Targets: []string{req.GetTarget()},
	}

	script := elastic.NewScript(`
ctx._source.dependence=params.dependence;
    `).Params(map[string]interface{}{
		"dependence": req.GetRelationType(),
	})

	resp := l.updateByQuery(ctx, c, index, &filter, script)
	if resp {
		log.Printf("%s中，服务%s->服务%s，依赖被修改为%s", req.GetScene(), req.GetSource(), req.GetTarget(), req.GetRelationType())
	}

	return resp
}

func (l *Line) QueryToChan(ctx context.Context, lineChan chan *model.Line, searchFilter *es.Searcher, index string) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			lines, _, err := l.searchWithSearcher(ctx, searchFilter, index)
			if err != nil {
				return err
			}
			if len(lines) == 0 {
				return nil
			}
			searchFilter.From = searchFilter.From + len(lines)
			for _, line := range lines {
				lineChan <- line
			}
		}
	}
}

func (l *Line) TopologyAnalyse(ctx context.Context, req *model.TopologyIndicatorReq) (bool, model.Indicator) {
	resp := model.Indicator{}
	c := &app.RequestContext{}
	lines, count := l.Query(ctx, c, &model.LineReq{
		Sources:      nil,
		Targets:      nil,
		Dependencies: nil,
		MatchAll:     false,
		PageSize:     0,
		PageNum:      0,
		IgnoreIds:    nil,
		IsCore:       false,
		Scene:        req.GetScene(),
	})
	if c.Errors != nil {
		log.Printf("TopologyAnalyse中Query失败,err=%s", c.Errors.String())
		return false, resp
	}
	if count == 0 {
		return true, resp
	}

	stack := list.New()
	relations := make(map[string][]string)
	inDegree := make(map[string]int)
	//initialize the status
	for _, line := range lines {
		inDegree[line.GetSource()] = 0
		inDegree[line.GetTarget()] = 0
	}
	for _, line := range lines {
		inDegree[line.GetTarget()]++
		relations[line.GetSource()] = append(relations[line.GetSource()], line.GetTarget())
	}
	for name, degree := range inDegree {
		if degree == 0 {
			stack.PushBack(name)
		}
	}
	for stack.Len() != 0 {
		line := stack.Back()
		stack.Remove(line)
		sourceName := line.Value.(string)
		for _, targetName := range relations[sourceName] {
			inDegree[targetName]--
			if inDegree[targetName] == 0 {
				stack.PushBack(targetName)
			}
		}
		delete(relations, sourceName)
	}
	if len(relations) == 0 {
		resp.IsCycle = false
	} else {
		resp.IsCycle = true
	}
	return true, resp
}
