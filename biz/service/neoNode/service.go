package neoNode

import (
	lineModel "GraduateThesis/biz/model/line"
	"GraduateThesis/biz/service/base"
	"GraduateThesis/biz/service/line"
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
)

var mergeQueryString = "MERGE (n:NeoNode {name: $name,isCore: $isCore,scene: $scene}) RETURN n"
var mergeRelationString = "MATCH (n:NeoNode {name: $name}) UNWIND $targets AS target_name MATCH (target:NeoNode {name: target_name}) MERGE (n)-[:Relation]->(target)"

type NeoNode struct {
	base.Base
	line.Line
}

func New(ctx context.Context, base base.Base, line line.Line) *NeoNode {
	neoNode := &NeoNode{Base: base, Line: line}
	return neoNode
}

func (n *NeoNode) MergeNode(ctx context.Context, nodes []map[string]any) bool {
	for _, node := range nodes {
		_, err := neo4j.ExecuteQuery(ctx, *n.NEO4J, mergeQueryString, node, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
		if err != nil {
			log.Printf("neo4j MergeNode failed,err=%s", err.Error())
			return false
		}
	}
	for _, node := range nodes {
		if node["targets"] != "" {
			_, err := neo4j.ExecuteQuery(ctx, *n.NEO4J, mergeRelationString, node, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
			if err != nil {
				log.Printf("neo4j MergeNode failed,err=%s", err.Error())
				return false
			}
		}
	}
	return true
}

func (n *NeoNode) DeleteAllNodes(ctx context.Context) error {
	deleteQueryString := "MATCH (n) DETACH DELETE n"
	_, err := neo4j.ExecuteQuery(ctx, *n.NEO4J, deleteQueryString, nil, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		log.Printf("Failed to delete all nodes from Neo4j, error: %s", err.Error())
		return err
	}
	return nil
}

func (n *NeoNode) EsToNeo(ctx context.Context, req *lineModel.LineReq) error {
	c := &app.RequestContext{}
	lines, total := n.Line.Query(ctx, c, req)
	if c.Errors != nil {
		log.Printf("query lines in EsToNeo failed")
		return c.Errors.Last()
	}
	if total == 0 {
		return nil
	}
	relations := make(map[string][]string)
	otherInformations := make(map[string]map[string]any)
	for _, line1 := range lines {
		if _, ex := otherInformations[line1.GetSource()]; !ex {
			otherInformations[line1.GetSource()] = make(map[string]any)
		}
		if _, ex := otherInformations[line1.GetTarget()]; !ex {
			otherInformations[line1.GetTarget()] = make(map[string]any)
		}
		otherInformations[line1.GetSource()]["name"] = line1.GetSource()
		otherInformations[line1.GetTarget()]["name"] = line1.GetTarget()
		otherInformations[line1.GetSource()]["isCore"] = line1.GetSourceIsCore()
		otherInformations[line1.GetTarget()]["isCore"] = line1.GetTargetIsCore()
		otherInformations[line1.GetSource()]["scene"] = line1.GetSourceScene()
		otherInformations[line1.GetTarget()]["scene"] = line1.GetTargetScene()
		relations[line1.GetSource()] = append(relations[line1.GetSource()], line1.GetTarget())
	}
	nodes := make([]map[string]any, 0, len(otherInformations))
	for source, otherInformation := range otherInformations {
		otherInformation["targets"] = relations[source]
		nodes = append(nodes, otherInformation)
	}
	ok := n.MergeNode(ctx, nodes)
	if !ok {
		fmt.Println("mergeNode in EsToNeo failed")
		return errors.New("mergeNode in EsToNeo failed")
	}
	return nil
}
