package service

import (
	"github.com/XZ0730/runFzu/biz/dal/db"
	"github.com/XZ0730/runFzu/biz/model/ledger"
	"github.com/XZ0730/runFzu/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

func (l *LedgerService) CreateLedger(user_id int64, req *ledger.LedgerModel) (code int64, msg string) {
	klog.Info("c_time:", req.GetCreateTime())
	c_time, err := time.Parse(time.DateTime, req.GetCreateTime())
	if err != nil {
		klog.Error("[newLedger] error:", err.Error())
		return errno.TimeError.ErrorCode, errno.TimeError.ErrorMsg
	}
	_, err = time.Parse(time.DateTime, req.GetCreateTime())
	if err != nil {
		klog.Error("[newLedger] error:", err.Error())
		return errno.TimeError.ErrorCode, errno.TimeError.ErrorMsg
	}

	u_time, err := time.Parse(time.DateTime, req.GetCreateTime())
	if err != nil {
		klog.Error("[newLedger] error:", err.Error())
		return errno.TimeError.ErrorCode, errno.TimeError.ErrorMsg
	}
	_, err = time.Parse(time.DateTime, req.GetCreateTime())
	if err != nil {
		klog.Error("[newLedger] error:", err.Error())
		return errno.TimeError.ErrorCode, errno.TimeError.ErrorMsg
	}

	newLedger := db.NewLedger(req.GetLedgerId(), req.GetUserId(), req.GetLedgerName(), req.GetCover(), c_time, u_time)
	if err := db.CreateLedger(newLedger); err != nil {
		klog.Error("[newLedger]create error:", err.Error())
		return errno.LedgerCreateError.ErrorCode, errno.LedgerCreateError.ErrorMsg
	}
	return errno.StatusSuccessCode, errno.StatusSuccessMsg
}
