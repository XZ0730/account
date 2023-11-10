namespace go ledger

struct BaseRequest{
}

struct BaseResponse{
    1: required i64 code,
    2: required string message,
}

struct LedgerModel{
    1: i32 ledgerId,
    2: i64 userId,
    3: string ledgerName,
    4: string coverMsg,
    5: string createTime,
    6: string updateTime,
}

struct LedgerCreateResponse{
    1: i64 code,
    2: string msg,
    3: LedgerModel data
}

struct LedgerListResponse{
    1: i64 code,
    2: string msg,
    3: map<string, list<LedgerModel>> data
}

service LedgerService{
    LedgerCreateResponse LedgerUpdate(1:LedgerModel req)(api.put = "/api/ledger")
    LedgerListResponse LedgerList(1:BaseRequest req)(api.get = "/api/ledger")
    BaseResponse LedgerDelete(1:LedgerModel req)(api.delete = "/api/ledger")
    LedgerCreateResponse LedgerCreate(1:LedgerModel req)(api.post="/api/ledger")
}
