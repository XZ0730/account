package pack

import "github.com/XZ0730/runFzu/biz/model/goal"

func PackGoallist(resp *goal.GoalListGetResponse, code int64, msg string, goals []*goal.GoalModel) {
	resp.Base = &goal.BaseResponse{
		Code:    code,
		Message: msg,
	}
	resp.Data = make(map[string][]*goal.GoalModel)
	resp.Data["list"] = goals
}
