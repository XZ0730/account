namespace go goal

struct BaseRequest{
}

struct BaseResponse{
    1: required i64 code,
    2: required string message,
}

struct GoalCreateRequest{
    1: string goal_name
    2: double money
    3: string create_date
    4: string deadline
}

service GoalService{
    BaseResponse GoalCreate(1:GoalCreateRequest req)(api.post="/api/goal")

}
