namespace go line

include "base.thrift"

struct Line {
  1:string docId;
  2:i64 docVersion;
  3:i64 docSeqNo;
  4:i64 docPrimaryTerm;
  5:string source;
  6:bool sourceIsCore;
  7:string sourceScene;
  8:string target;
  9:bool targetIsCore;
  10:string targetScene;
  11:string dependence;
  12:base.Timestamp lastFoundTime;
  13:base.Timestamp firstFoundTime;
}

struct LineReq {
  1: list<string> sources (api.query="sources[]");
  2: list<string> targets (api.query="targets[]");
  3: list<string> dependencies (api.query="dependencies[]");
  4: bool matchAll (api.query="matchAll");
  5: i64 pageSize (api.query="pageSize");
  6: i64 pageNum (api.query="pageNum");
  7: list<string> ignoreIds (api.query="ignoreIds[]");
  8: bool isCore (api.query="isCore");
  9: string scene (api.query="scene");
}

struct LineBulkReq {
  1: required list<Line> lines (api.body="lines");
}

struct LineDeleteReq {
  1: string docId (api.body="doc_id");
  2: string scene (api.body="scene")
}

struct LineRespSample {
  1: required list<Line> lines;
  2: required i64 totalCount;
}

struct LineResp {
  1: required i64 code;
  2: required LineRespSample data;
  3: required string msg;
}


struct ChangeDependenceWithRelationReq {
  1: string source(api.body="source");
  2: string target(api.body="target");
  3: string relation_type(api.body="relation_type");
  4: string scene (api.body="scene")
}

struct Indicator {
  1: bool isCycle
}

struct TopologyIndicatorResp {
  1: required i64 code;
  2: required Indicator data;
  3: required string msg;
}


struct TopologyIndicatorReq {
  1: string scene(api.body="source");
}


service LineService {
  LineResp Query(1: LineReq request) (api.get="/api/line/query")
  base.SampleResp ChangeDependence(1: ChangeDependenceWithRelationReq request) (api.post="/api/line/change-dependence")
  base.SampleResp Bulk(1: LineBulkReq request) (api.post="/api/line/bulk")
  base.SampleResp Delete(1: LineDeleteReq request) (api.post="/api/line/delete")
  TopologyIndicatorResp TopologyAnalyse(1: TopologyIndicatorReq request) (api.post="/api/line/topology_analyse")
}