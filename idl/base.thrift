namespace go base

struct Result {
    1: required bool success;
    2: optional i64 code;
    3: optional string message;
}

typedef i64 Timestamp

service BaseService {
    bool Healthz() (api.get="/healthz")
}