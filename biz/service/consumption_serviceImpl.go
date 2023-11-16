package service

import (
	"github.com/XZ0730/runFzu/biz/dal/db"
	"github.com/XZ0730/runFzu/biz/model/consumption"
	"github.com/XZ0730/runFzu/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/sync/errgroup"
	"time"
)

func (c *ConsumptionService) UpdateConsumption(userId int64, model *consumption.ConsumptionModel) (int64, string) {
	if err := db.JudgeConsumption(model.ConsumptionId); err != nil {
		klog.Error("[consumption] update error", err.Error())
		return errno.NotExistErrorCode, errno.NotExistError.ErrorMsg
	}

	ledgerId, err := db.GetLedgerIdByConsumption(model.ConsumptionId)
	if err != nil {
		klog.Error("[consumption] error", err.Error())
		return errno.NotExistErrorCode, errno.NotExistError.ErrorMsg
	}

	if err = db.JudgeUserHaveLedger(ledgerId, userId); err != nil {
		klog.Error("[consumption] error", err.Error())
		return errno.NotExistErrorCode, errno.NotExistError.ErrorMsg
	}

	c_time, err := time.Parse(time.DateTime, model.GetConsumeTime())
	if err != nil {
		klog.Error("[consumption] error:", err.Error())
		return errno.TimeError.ErrorCode, errno.TimeError.ErrorMsg
	}
	con := db.NewConsumptionAllParams(model.ConsumptionId, model.GetAmount(), model.GetConsumptionName(), model.GetDescription(), model.GetTypeId(), model.GetStore(), c_time, model.GetCredential())
	if err = db.UpdateConsumption(con); err != nil {
		klog.Error("[consumption] update error:", err.Error())
		return errno.UpdateErrorCode, errno.UpdateError.ErrorMsg
	}
	return errno.StatusSuccessCode, errno.SuccessMsg
}

func (c *ConsumptionService) GetConsumptionsByRange(start string, end string, userId int64) ([]*consumption.ConsumptionModel, int64, string) {
	ledgerIds := db.GetLedgersByUserId(userId)
	consumptions := db.GetConByRange(start, end, ledgerIds)

	var eg errgroup.Group
	list := make([]*consumption.ConsumptionModel, 0)

	for _, val := range consumptions {
		tmp := val
		eg.Go(func() error {
			vo_g := new(consumption.ConsumptionModel)
			vo_g.ConsumptionId = tmp.ConsumptionId
			vo_g.ConsumptionName = tmp.ConsumptionName
			vo_g.Amount = tmp.Amount
			vo_g.Description = tmp.Description
			vo_g.TypeId = tmp.TypeId
			vo_g.ConsumeTime = tmp.ConsumeTime.Format(time.DateTime)
			vo_g.Store = tmp.Store
			vo_g.Credential = tmp.Credential
			list = append(list, vo_g)
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		klog.Info("[consumption]get error:", err.Error())
		return nil, errno.GetError.ErrorCode, errno.GetError.ErrorMsg
	}
	return list, errno.StatusSuccessCode, errno.StatusSuccessMsg
}
