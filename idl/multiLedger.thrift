namespace go multiledger

struct BaseRequest{
}

struct BaseResponse{
    1: required i64 code,
    2: required string message,
}

struct CreateMLRequest{
    1:string description
    2:string password
    3:string multiLedgerName
}

struct JoinMLRequest{
    1:string password
}

struct InsertMlConsumReq{
    1: i64 consId
    2: i64 multiLedgerId
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
# MultiLedgerId
# MultiLedgerName
# Description
# Password
# ModifyTime
struct MultiledgerModel{
    1: i64 multiLedgerId
    2: string multiLedgerName
    3: string description
    4: string password
    5: string modifyTime
}

struct GetMulConsumptionReq{
    1: i64  multiLedgerId
}

struct GetMulConsumptionResp{
    1: BaseResponse  base
    2: map<string,list<ConsumptionModel>> data
}

struct GetMultiLedgerListResp{
    1: BaseResponse  base
    2: map<string,list<MultiledgerModel>> data
}

struct DelMultiLedgerReq{
    1: i64  multiLedgerId
}

struct PutMultiLedgerReq{
    1: i64 multiLedgerId
    2: string multiLedgerName
    3: string description
    4: string password
}

struct DelMulConsumptionReq{
    1: i64 consId
    2: i64 multiLedgerId
}

service MultiLedgerService{
    BaseResponse CreateMultiledger(1:CreateMLRequest req)(api.post="/api/multiLedger")
    BaseResponse JoinMultiledger(1:JoinMLRequest req)(api.post="/api/multiLedger/join") 
    BaseResponse InsertMlConsumption(1:InsertMlConsumReq req)(api.post="/api/multiLedger/consumption")

    GetMulConsumptionResp GetMulConsumption(1:GetMulConsumptionReq req)(api.get="/api/multiLedger/consumption")
    GetMultiLedgerListResp GetMultiLedgerList(1:BaseRequest req)(api.get="/api/multiLedger")

    BaseResponse DelMultiLedger(1:DelMultiLedgerReq req)(api.delete="/api/multiLedger")
    BaseResponse PutMultiLedger(1:PutMultiLedgerReq req)(api.put="/api/multiLedger")
    BaseResponse DelMulConsumption(1:DelMulConsumptionReq req)(api.delete="/api/multiLedger/consumption")
}