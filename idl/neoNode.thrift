namespace go neoNode

include "base.thrift"
include "line.thrift"

struct NeoNode {
    1: string name;
    2: bool isCore;
    3: string scene;
}

service NeoNodeService {
    base.SampleResp EsToNeo (1: line.LineReq request) (api.get="/api/neo_node/es_to_neo")
    base.SampleResp EmptyNeo () (api.get="/api/neo_node/empty_neo")
}