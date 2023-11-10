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



service MultiLedger{
    BaseResponse CreateMultiledger(1:CreateMLRequest req)(api.post="/api/multiLedger")
}