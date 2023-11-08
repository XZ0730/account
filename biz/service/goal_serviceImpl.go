package service

import (
	"time"

	"github.com/XZ0730/runFzu/biz/dal/db"
	"github.com/XZ0730/runFzu/biz/model/goal"
	"github.com/XZ0730/runFzu/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/sync/errgroup"
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

func (g *GoalService) GetGoals(user_id int64) (goal_list []*goal.GoalModel, code int64, msg string) {
	goals, err := db.GetGoalList(user_id)
	if err != nil {
		klog.Info("[goal]get error:", err.Error())
		return nil, errno.GoalGetError.ErrorCode, errno.GoalGetError.ErrorMsg
	}
	var eg errgroup.Group
	list := make([]*goal.GoalModel, 0)
	for _, val := range goals {
		tmp := val
		eg.Go(func() error {
			vo_g := new(goal.GoalModel)
			vo_g.GoalId = tmp.GoalId
			vo_g.GoalName = tmp.GoalName
			vo_g.Money = tmp.Money
			vo_g.UserId = tmp.UserId
			vo_g.CreateDate = tmp.CreateDate.Format(time.DateTime)
			vo_g.Deadline = tmp.Deadline.Format(time.DateTime)
			list = append(list, vo_g)
			return nil
		})
	}
	if err = eg.Wait(); err != nil {
		klog.Info("[goal]get error:", err.Error())
		return nil, errno.GoalGetError.ErrorCode, errno.GoalGetError.ErrorMsg
	}
	return list, errno.StatusSuccessCode, errno.StatusSuccessMsg
}
