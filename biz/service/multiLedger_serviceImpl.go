package service

import (
	"github.com/XZ0730/runFzu/biz/dal/db"
	"github.com/XZ0730/runFzu/biz/model/multiledger"
	"github.com/XZ0730/runFzu/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
)

func (m *MultiLedgerService) CreateMultiLedger(uid int64, req *multiledger.CreateMLRequest) (code int64, msg string) {
	ml := db.NewMultiLedger(req.GetMultiLedgerName(), req.GetDescription(), req.GetPassword())
	if err := db.CreateMultiLedger(ml); err != nil {
		klog.Error("[mul_l]error:", err.Error())
		return errno.CreateError.ErrorCode, errno.CreateError.ErrorMsg
	}
	if err := db.CreateM_user(ml.MultiLedgerId, uid); err != nil {
		klog.Error("[mul_l]error:", err.Error())
		return errno.CreateError.ErrorCode, errno.CreateError.ErrorMsg
	}
	return errno.StatusSuccessCode, errno.StatusSuccessMsg
}

func (m *MultiLedgerService) JoinMultiledger(uid int64, pwd string) (code int64, msg string) {

	id, err := db.GetMultiLedgerByPassword(pwd)
	if err != nil {
		klog.Error("[multi]error:", err.Error())
		return errno.GetError.ErrorCode, errno.GetError.ErrorMsg
	}
	if err = db.JudgeM_user(id, uid); err == nil {
		klog.Error("[multi]error: user exist")
		return errno.UserExistedError.ErrorCode, errno.UserExistedError.ErrorMsg
	}
	if err = db.CreateM_user(id, uid); err != nil {
		klog.Error("[multi]error:", err.Error())
		return errno.CreateError.ErrorCode, errno.CreateError.ErrorMsg
	}

	return errno.StatusSuccessCode, errno.StatusSuccessMsg
}
