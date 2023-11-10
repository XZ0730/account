package service

import (
	"time"

	"github.com/XZ0730/runFzu/biz/dal/db"
	"github.com/XZ0730/runFzu/biz/model/ledger"
	"github.com/XZ0730/runFzu/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
<<<<<<< HEAD
	"golang.org/x/sync/errgroup"
	"time"
=======
>>>>>>> c0629d3500484cf68c01991bb08fac3a35bbf982
)

func (l *LedgerService) CreateLedger(user_id int64, req *ledger.LedgerModel) (code int64, msg string) {
	c_time, err := time.Parse(time.DateTime, req.GetCreateTime())
	if err != nil {
		klog.Error("[newLedger] error:", err.Error())
		return errno.TimeError.ErrorCode, errno.TimeError.ErrorMsg
	}
<<<<<<< HEAD

	u_time, err := time.Parse(time.DateTime, req.GetUpdateTime())
=======
	u_time, err := time.Parse(time.DateTime, req.GetCreateTime())
>>>>>>> c0629d3500484cf68c01991bb08fac3a35bbf982
	if err != nil {
		klog.Error("[newLedger] error:", err.Error())
		return errno.TimeError.ErrorCode, errno.TimeError.ErrorMsg
	}
<<<<<<< HEAD

	newLedger := db.NewLedger(req.GetLedgerId(), user_id, req.GetLedgerName(), req.GetCoverMsg(), c_time, u_time)
=======
	newLedger := db.NewLedger(req.GetLedgerId(), user_id, req.GetLedgerName(), req.GetCover(), c_time, u_time)
>>>>>>> c0629d3500484cf68c01991bb08fac3a35bbf982
	if err := db.CreateLedger(newLedger); err != nil {
		klog.Error("[newLedger]create error:", err.Error())
		return errno.LedgerCreateError.ErrorCode, errno.LedgerCreateError.ErrorMsg
	}
	return errno.StatusSuccessCode, errno.StatusSuccessMsg
}

func (l *LedgerService) DeleteLedger(user_id int64, req *ledger.LedgerModel) (code int64, msg string) {
	c_time, err := time.Parse(time.DateTime, req.GetCreateTime())
	if err != nil {
		klog.Error("[newLedger] error:", err.Error())
		return errno.TimeError.ErrorCode, errno.TimeError.ErrorMsg
	}
	u_time, err := time.Parse(time.DateTime, req.GetCreateTime())
	if err != nil {
		klog.Error("[newLedger] error:", err.Error())
		return errno.TimeError.ErrorCode, errno.TimeError.ErrorMsg
	}

	newLedger := db.NewLedger(req.GetLedgerId(), user_id, req.GetLedgerName(), req.GetCoverMsg(), c_time, u_time)
	if err := db.DeleteLedger(newLedger); err != nil {
		klog.Error("[newLedger]delete error:", err.Error())
		return errno.LedgerDeleteError.ErrorCode, errno.LedgerDeleteError.ErrorMsg
	}
	return errno.StatusSuccessCode, errno.StatusSuccessMsg
}

func (l *LedgerService) ListLedgers(user_id int64) (ledgerList []*ledger.LedgerModel, code int64, msg string) {
	ledgers, err := db.ListLedgers(user_id)
	if err != nil {
		klog.Info("[ledger]get error:", err.Error())
		return nil, errno.LedgerGetError.ErrorCode, errno.LedgerGetError.ErrorMsg
	}
	var eg errgroup.Group
	list := make([]*ledger.LedgerModel, 0)

	for _, val := range ledgers {
		tmp := val
		eg.Go(func() error {
			vo_g := new(ledger.LedgerModel)
			vo_g.LedgerId = tmp.LedgerId
			vo_g.LedgerName = tmp.LedgerName
			vo_g.CoverMsg = tmp.CoverMsg
			vo_g.UserId = tmp.UserId
			vo_g.CreateTime = tmp.CreateTime.Format(time.DateTime)
			vo_g.UpdateTime = tmp.UpdateTime.Format(time.DateTime)
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
