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

# ConsumptionId
# Amount
# ConsumptionName
# Description
# TypeId
# Store
# ConsumeTime
# Credential
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

struct GetMulConsumptionReq{
    1: i64  multiLedgerId
}

struct GetMulConsumptionResp{
    1: BaseResponse  base
    2: map<string,list<ConsumptionModel>> data
}

service MultiLedgerService{
    BaseResponse CreateMultiledger(1:CreateMLRequest req)(api.post="/api/multiLedger")
    BaseResponse JoinMultiledger(1:JoinMLRequest req)(api.post="/api/multiLedger/join") 
    BaseResponse InsertMlConsumption(1:InsertMlConsumReq req)(api.post="/api/multiLedger/consumption")

    GetMulConsumptionResp GetMulConsumption(1:GetMulConsumptionReq req)(api.get="/api/multiLedger/consumption")
}