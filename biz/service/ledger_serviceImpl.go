package service

import (
	"time"

	"github.com/XZ0730/runFzu/biz/dal/db"
	"github.com/XZ0730/runFzu/biz/model/ledger"
	"github.com/XZ0730/runFzu/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/sync/errgroup"
)

func (l *LedgerService) CreateLedger(user_id int64, req *ledger.LedgerModel) (code int64, msg string) {
	c_time, err := time.Parse(time.DateTime, req.GetCreateTime())
	if err != nil {
		klog.Error("[newLedger] error:", err.Error())
		return errno.TimeError.ErrorCode, errno.TimeError.ErrorMsg
	}

	u_time, err := time.Parse(time.DateTime, req.GetUpdateTime())
	if err != nil {
		klog.Error("[newLedger] error:", err.Error())
		return errno.TimeError.ErrorCode, errno.TimeError.ErrorMsg
	}

	newLedger := db.NewLedger(req.GetLedgerId(), user_id, req.GetLedgerName(), req.GetCoverMsg(), c_time, u_time)
	if err := db.CreateLedger(newLedger); err != nil {
		klog.Error("[newLedger]create error:", err.Error())
		return errno.CreateError.ErrorCode, errno.CreateError.ErrorMsg
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
		return errno.DelError.ErrorCode, errno.DelError.ErrorMsg
	}
	return errno.StatusSuccessCode, errno.StatusSuccessMsg
}

func (l *LedgerService) ListLedgers(user_id int64) (ledgerList []*ledger.LedgerModel, code int64, msg string) {
	ledgers, err := db.ListLedgers(user_id)
	if err != nil {
		klog.Info("[ledger]get error:", err.Error())
		return nil, errno.GetError.ErrorCode, errno.GetError.ErrorMsg
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
		klog.Info("[ledger]get error:", err.Error())
		return nil, errno.GetError.ErrorCode, errno.GetError.ErrorMsg
	}
	return list, errno.StatusSuccessCode, errno.StatusSuccessMsg
}

func (l *LedgerService) UpdateLedger(model *ledger.LedgerModel) (code int64, msg string) {
	c_time, err := time.Parse(time.DateTime, model.GetCreateTime())
	if err != nil {
		klog.Error("[ledger] error:", err.Error())
		return errno.TimeError.ErrorCode, errno.TimeError.ErrorMsg
	}
	u_time, err := time.Parse(time.DateTime, model.GetUpdateTime())
	if err != nil {
		klog.Error("[ledger] error:", err.Error())
		return errno.TimeError.ErrorCode, errno.TimeError.ErrorMsg
	}
	ledger := db.NewLedger(model.GetLedgerId(), model.GetUserId(), model.GetLedgerName(), model.GetCoverMsg(), c_time,
		u_time)

	if err = db.UpdateLedger(ledger); err != nil {
		klog.Error("[ledger] update error:", err.Error())
		return errno.UpdateError.ErrorCode, errno.UpdateError.ErrorMsg
	}

	return errno.StatusSuccessCode, errno.StatusSuccessMsg
}

func (l *LedgerService) CreateLedgerConsumption(ledgerId int32, consumptionId int64) (code int64, msg string) {
	rel := db.NewLedgerConsumptionRel(ledgerId, consumptionId)
	if err := db.CreateLedgerConsumptionRel(rel); err != nil {
		klog.Error("[LedgerConsumptionRel] create error:", err.Error())
		return errno.CreateError.ErrorCode, errno.CreateError.ErrorMsg
	}

	return errno.StatusSuccessCode, errno.StatusSuccessMsg
}

func (l *LedgerService) LedgerConsumptionList(ledgerId int32) (consumptionList []*ledger.ConsumptionModel, code int64, msg string) {
	cm := make([]*ledger.ConsumptionModel, 0)
	cl, err := db.ConsumptionList(ledgerId)

	if err != nil {
		klog.Error("[ledger_consumption]error:", err.Error())
		return nil, errno.GetError.ErrorCode, errno.GetError.ErrorMsg
	}
	var eg errgroup.Group
	for _, val := range cl {
		tmp := val
		eg.Go(func() error {
			vo_g := new(ledger.ConsumptionModel)
			vo_g.ConsumptionId = tmp.ConsumptionId
			vo_g.ConsumptionName = tmp.ConsumptionName
			vo_g.Description = tmp.Description
			vo_g.Amount = tmp.Amount
			vo_g.TypeId = tmp.TypeId
			vo_g.Store = tmp.Store
			vo_g.ConsumeTime = tmp.ConsumeTime.Format(time.DateTime)
			vo_g.Credential = tmp.Credential
			cm = append(cm, vo_g)
			return nil
		})
	}
	if err = eg.Wait(); err != nil {
		klog.Info("[ledger_consumption] get error:", err.Error())
		return nil, errno.GetError.ErrorCode, errno.GetError.ErrorMsg
	}
	return cm, errno.StatusSuccessCode, errno.StatusSuccessMsg
}

func (l *LedgerService) DeleteLedgerConsumption(ledgerId int32, consumptionId int64) (code int64, msg string) {
	if err := db.DeleteLedgerConsumption(ledgerId, consumptionId); err != nil {
		klog.Info("[ledger_consumption] delete error:", err.Error())
		return errno.DelErrorCode, errno.DelError.ErrorMsg
	}
	return errno.StatusSuccessCode, errno.StatusSuccessMsg
}

func (l *LedgerService) GetLedgerBalance(ledgerId int32) (balance float64, code int64, msg string) {
	balance, err := db.GetLedgerBalance(ledgerId)
	if err != nil {
		klog.Info("[ledger_consumption] get error:", err.Error())
		return 0, errno.GetErrorCode, errno.GetError.ErrorMsg
	}
	return balance, errno.StatusSuccessCode, errno.StatusSuccessMsg
}
