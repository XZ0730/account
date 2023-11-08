package service

import (
	"time"

	"github.com/XZ0730/runFzu/biz/dal/db"
	"github.com/XZ0730/runFzu/biz/model/goal"
	"github.com/XZ0730/runFzu/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
)

func (g *GoalService) CreateGoal(user_id int64, req *goal.GoalCreateRequest) (code int64, msg string) {
	klog.Info("c_time:", req.GetCreateDate())
	c_time, err := time.Parse(time.DateTime, req.GetCreateDate())
	if err != nil {
		klog.Error("[goal] error:", err.Error())
		return errno.GoalTimeError.ErrorCode, errno.GoalTimeError.ErrorMsg
	}
	ddl, err := time.Parse(time.DateTime, req.GetDeadline())
	if err != nil {
		klog.Error("[goal] error:", err.Error())
		return errno.GoalTimeError.ErrorCode, errno.GoalTimeError.ErrorMsg
	}
	goal := db.NewGoal(user_id, req.GetGoalName(), req.GetMoney(), c_time, ddl)
	if err := db.CreateGoal(goal); err != nil {
		klog.Error("[goal]create error:", err.Error())
		return errno.GoalCreateError.ErrorCode, errno.GoalCreateError.ErrorMsg
	}
	return errno.StatusSuccessCode, errno.StatusSuccessMsg
}
