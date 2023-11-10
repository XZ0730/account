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
	if err := db.CreateM_user(ml.Id, uid); err != nil {
		klog.Error("[mul_l]error:", err.Error())
		return errno.CreateError.ErrorCode, errno.CreateError.ErrorMsg
	}
	return errno.StatusSuccessCode, errno.StatusSuccessMsg
}
