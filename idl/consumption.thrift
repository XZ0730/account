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


struct GetConsumptionByRangeResponse{
    1: i64 code
    2: string msg
    3: map<string, list<ConsumptionModel>> data
}

struct GetSumByRangeResponse{
    1: i64 code
    2: string msg
    3: map<string, double> data
}

struct GetLastMonthMoneyResp{
    1: i64 code
    2: string msg
    3: map<string, list<double>> data;
}

struct GetUserConsumptionResp{
    1: i64 code
    2: string msg
    3: list<ConsumptionModel> list;
}

service ConsumptionService{
    GetSumByRangeResponse GetDayOut(1:BaseRequest req)(api.get = "/api/consumption/day/out")
    GetUserConsumptionResp GetUseConsumption(1:BaseRequest req)(api.get = "/api/consumption")
    GetSumByRangeResponse GetBalanceByMonth(1:BaseRequest req)(api.get = "/api/consumption/balance/month")
    GetSumByRangeResponse GetBalanceByYear(1:BaseRequest req)(api.get = "/api/consumption/balance/year")
    GetLastMonthMoneyResp GetLastMonthMoney(1:BaseRequest req)(api.get = "/api/consumption/last/month/analysis")
    GetConsumptionByRangeResponse GetLocalMonthConsumption(1:BaseRequest req)(api.get = "/api/consumption/month/map")
    GetSumByRangeResponse GetOutByRange(1:BaseRequest req)(api.get = "/api/consumption/range/out")
    GetSumByRangeResponse GetInByRange(1:BaseRequest req)(api.get = "/api/consumption/range/in")
    GetConsumptionByRangeResponse GetConsumptionByRange(1: BaseRequest req)(api.get = "/api/consumption/range/map")
    ConsumptionUpdateResponse UpdateConsumption(1:ConsumptionModel req)(api.put = "/api/consumption")

    GetConsumptionByRangeResponse GetConsumptionByDate(1:BaseRequest req)(api.get="/api/consumption/date")
}