package service

import (
	"github.com/XZ0730/runFzu/biz/dal/db"
	"github.com/XZ0730/runFzu/biz/model/consumption"
	"github.com/XZ0730/runFzu/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
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
