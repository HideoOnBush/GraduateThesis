namespace go base

struct Result {
    1: required bool success;
    2: optional i64 code;
    3: optional string message;
}

struct SampleResp {
  1: required i64 code;
  2: required bool data;
  3: required string message;
}

typedef i64 Timestamp

service BaseService {
    bool Healthz() (api.get="/healthz")
}