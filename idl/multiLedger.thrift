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

service MultiLedgerService{
    BaseResponse CreateMultiledger(1:CreateMLRequest req)(api.post="/api/multiLedger")
    BaseResponse JoinMultiledger(1:JoinMLRequest req)(api.post="/api/multiLedger/join")   
}