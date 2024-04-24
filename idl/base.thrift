namespace go base

struct SampleResp {
  1: required i64 code;
  2: required bool success;
  3: required string message;
}

typedef i64 Timestamp

service BaseService {
    bool Healthz() (api.get="/healthz")
}