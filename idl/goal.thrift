namespace go goal

struct BaseRequest{
}

struct BaseResponse{
    1: required i64 code,
    2: required string message,
}

struct GoalModel {
    1: i64 goalId
    2: string goalName
    3: i64 userId
    4: double money
    5: string createDate
    6: string deadline
    7: double saved_money
}

struct GoalCreateRequest{
    1: string goal_name
    2: double money
    3: string create_date
    4: string deadline
}
struct GoalListGetResponse{
    1: BaseResponse base
    2: map<string,list<GoalModel>> data
}

struct GoalDelRequest{
    1: i64 goalId
}

struct GoalPutRequest{
    1: i64 goalId
    2: string goalName
    3: string create_date
    4: string deadline
    5: double money
}

service GoalService{
    BaseResponse GoalCreate(1:GoalCreateRequest req)(api.post="/api/goal")
    GoalListGetResponse GoalListGet(1:BaseRequest req)(api.get="/api/goal")
    BaseResponse GoalDelete(1:GoalDelRequest req)(api.delete="/api/goal")
    BaseResponse GoalPut(1:GoalPutRequest req)(api.put="/api/goal")
}
