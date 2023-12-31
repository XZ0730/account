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
		return errno.TimeError.ErrorCode, errno.TimeError.ErrorMsg
	}
	ddl, err := time.Parse(time.DateTime, req.GetDeadline())
	if err != nil {
		klog.Error("[goal] error:", err.Error())
		return errno.TimeError.ErrorCode, errno.TimeError.ErrorMsg
	}
	goal := db.NewGoal(user_id, req.GetGoalName(), req.GetMoney(), c_time, ddl)
	if err := db.CreateGoal(goal); err != nil {
		klog.Error("[goal]create error:", err.Error())
		return errno.CreateError.ErrorCode, errno.CreateError.ErrorMsg
	}
	return errno.StatusSuccessCode, errno.StatusSuccessMsg
}

func (g *GoalService) GetGoals(user_id int64) (goal_list []*goal.GoalModel, code int64, msg string) {
	ledgerId := db.GetLedgersByUserId(user_id)
	var sum float64
	for _, id := range ledgerId {
		balance, _ := db.GetLedgerBalance(int32(*id))
		sum += balance
	}

	goals, err := db.GetGoalList(user_id)
	if err != nil {
		klog.Info("[goal]get error:", err.Error())
		return nil, errno.GetError.ErrorCode, errno.GetError.ErrorMsg
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
			vo_g.SavedMoney = sum
			list = append(list, vo_g)
			return nil
		})
	}
	if err = eg.Wait(); err != nil {
		klog.Info("[goal]get error:", err.Error())
		return nil, errno.GetError.ErrorCode, errno.GetError.ErrorMsg
	}
	return list, errno.StatusSuccessCode, errno.StatusSuccessMsg
}

func (g *GoalService) DelGoal(user_id int64, goal_id int64) (code int64, msg string) {
	if err := db.DelGoal(user_id, goal_id); err != nil {
		klog.Error("[goal] delete error:", err.Error())
		return errno.DelError.ErrorCode, errno.DelError.ErrorMsg
	}
	return errno.StatusSuccessCode, errno.StatusSuccessMsg
}

func (g *GoalService) UpdateGoal(user_id int64, req *goal.GoalPutRequest) (code int64, msg string) {
	klog.Info("c_time:", req.GetCreateDate())
	c_time, err := time.Parse(time.DateTime, req.GetCreateDate())
	if err != nil {
		klog.Error("[goal] error:", err.Error())
		return errno.TimeError.ErrorCode, errno.TimeError.ErrorMsg
	}
	ddl, err := time.Parse(time.DateTime, req.GetDeadline())
	if err != nil {
		klog.Error("[goal] error:", err.Error())
		return errno.TimeError.ErrorCode, errno.TimeError.ErrorMsg
	}
	goal := db.NewGoal(user_id, req.GetGoalName(), req.GetMoney(), c_time, ddl)
	goal.GoalId = req.GetGoalId()
	if err = db.UpdateGoal(goal); err != nil {
		klog.Error("[goal] update error:", err.Error())
		return errno.UpdateError.ErrorCode, errno.UpdateError.ErrorMsg
	}

	return errno.StatusSuccessCode, errno.StatusSuccessMsg
}
