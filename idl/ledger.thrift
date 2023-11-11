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

struct ConsumptionModel{
    1: i64 consumptionId
    2: string consumptionName
    3: string description
    4: double amount
    5: i8 typeId
    6: string store
    7: string consumeTime
    8: string credential
}

struct LedgerConsumptionRel{
    1: i32 ledgerId
    2: i64 consumptionId
}

struct LedgerConsumptionListResponse{
    1: i64 code,
    2: string msg,
    3: map<string, list<ConsumptionModel>> data
}

struct LedgerBalanceResponse{
    1: i64 code,
    2: string msg,
    3: map<string, double> data
}

service LedgerService{
    LedgerBalanceResponse LedgerBalance(1:BaseRequest req)(api.get = "/api/ledger/balance")
    LedgerConsumptionListResponse LedgerConsumptionList(1:BaseRequest req)(api.get = "/api/ledger/consumption")
    LedgerCreateResponse LedgerUpdate(1:LedgerModel req)(api.put = "/api/ledger")
    LedgerListResponse LedgerList(1:BaseRequest req)(api.get = "/api/ledger")
    BaseResponse LedgerDelete(1:LedgerModel req)(api.delete = "/api/ledger")
    LedgerCreateResponse LedgerCreate(1:LedgerModel req)(api.post="/api/ledger")
    BaseResponse LedgerAddConsumption(1:BaseRequest req)(api.post = "/api/ledger/consumption")
    BaseResponse LedgerDeleteConsumption(1:BaseRequest req)(api.delete = "/api/ledger/consumption")
}
