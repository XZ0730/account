namespace go consumption

struct BaseRequest{
}

struct BaseResponse{
    1: required i64 code,
    2: required string message,
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

struct ConsumptionUpdateResponse{
    1: i64 code
    2: string msg
    3: map<string, ConsumptionModel> data
}

struct TimeKeyConArray{
    1:map<string, list<ConsumptionModel>> tmap
}

struct GetConsumptionByRangeResponse{
    1: i64 code
    2: string msg
    3: map<string, TimeKeyConArray> data
}

struct GetSumByRangeResponse{
    1: i64 code
    2: string msg
    3: map<string, double> data
}

service ConsumptionService{
    GetSumByRangeResponse GetSumByRange(1:BaseRequest req)(api.get = "/api/consumption/range/in")
    GetConsumptionByRangeResponse GetConsumptionByRange(1: BaseRequest req)(api.get = "/api/consumption/range/map")
    ConsumptionUpdateResponse UpdateConsumption(1:ConsumptionModel req)(api.put = "/api/consumption")

    GetConsumptionByRangeResponse GetConsumptionByDate(1:BaseRequest req)(api.get="/api/consumption/date")
}