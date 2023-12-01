package service

import (
	"time"

	"github.com/XZ0730/runFzu/biz/dal/db"
	"github.com/XZ0730/runFzu/biz/model/consumption"
	"github.com/XZ0730/runFzu/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/sync/errgroup"
)

func (c *ConsumptionService) GetConsumptionsByUserId(userId int64) (int64, string, []*consumption.ConsumptionModel) {
	ledgerIds := db.GetLedgersByUserId(userId)
	consumptions := db.GetConsumptionByLedgerIds(ledgerIds)

	cons := make([]*consumption.ConsumptionModel, 0)
	for _, val := range consumptions {
		tmp := new(consumption.ConsumptionModel)
		tmp.ConsumptionName = val.ConsumptionName
		tmp.ConsumptionId = val.ConsumptionId
		tmp.Amount = val.Amount
		tmp.ConsumeTime = val.ConsumeTime.Format(time.DateTime)
		tmp.Description = val.Description
		tmp.TypeId = val.TypeId
		tmp.Store = val.Store
		tmp.Credential = val.Credential
		cons = append(cons, tmp)
	}

	return errno.SuccessCode, errno.SuccessMsg, cons
}

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

	list := make([]*consumption.ConsumptionModel, 0)

	for _, val := range consumptions {
		tmp := new(consumption.ConsumptionModel)
		tmp.ConsumptionName = val.ConsumptionName
		tmp.ConsumptionId = val.ConsumptionId
		tmp.Amount = val.Amount
		tmp.ConsumeTime = val.ConsumeTime.Format(time.DateTime)
		tmp.Description = val.Description
		tmp.TypeId = val.TypeId
		tmp.Store = val.Store
		tmp.Credential = val.Credential
		list = append(list, tmp)
	}

	return list, errno.StatusSuccessCode, errno.StatusSuccessMsg
}

func (c *ConsumptionService) GetSumByRange(start string, end string, userId int64, op float64) (int64, string, float64) {
	ledgerIds := db.GetLedgersByUserId(userId)
	consumptions := db.GetConByRange(start, end, ledgerIds)

	sum := 0.0
	for _, val := range consumptions {
		x := val.Amount
		if op > 0 && x > 0 {
			sum += x
		} else if op < 0 && x < 0 {
			sum += x
		} else {
			sum += x
		}
	}

	return errno.StatusSuccessCode, errno.StatusSuccessMsg, sum
}

func (c *ConsumptionService) GetConsumptionSumListByRange(start string, end string, userId int64) (int64, string, []float64) {
	ledgerIds := db.GetLedgersByUserId(userId)
	consumptions := db.GetConByRange(start, end, ledgerIds)

	var sum []float64
	for _, val := range consumptions {
		if val.Amount < 0 {
			sum = append(sum, val.Amount)
		}
	}

	return errno.SuccessCode, errno.SuccessMsg, sum
}

func (c *ConsumptionService) GetConsumptionByDate(uid int64, date time.Time, the_type int64) (int64, string, []*consumption.ConsumptionModel) {
	list := make([]*consumption.ConsumptionModel, 0)
	var start time.Time
	var end time.Time
	ledger_id := db.GetLedgersByUserId(uid)
	if the_type == 1 {
		start = time.Date(date.Year(), 1, 1, 0, 0, 0, 0, time.Local)
		end = time.Date(date.Year()+1, 1, 1, 0, 0, 0, 0, time.Local)
	} else if the_type == 2 {
		start = time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.Local)
		end = time.Date(date.Year(), date.Month()+1, 1, 0, 0, 0, 0, time.Local)
	} else {
		start = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local)
		end = time.Date(date.Year(), date.Month(), date.Day()+1, 0, 0, 0, 0, time.Local)
	}
	consumptions := db.GetConByRange(start.Format(time.DateOnly), end.Format(time.DateOnly), ledger_id)
	var eg errgroup.Group
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
		return errno.GetError.ErrorCode, errno.GetError.ErrorMsg, nil
	}
	return errno.StatusSuccessCode, errno.StatusSuccessMsg, list
}

func (c *ConsumptionService) CreateConsumption(req *consumption.CreateConsumptionReq) (int64, int64, string) {

	con_time, err := time.Parse(time.DateTime, req.GetConsumeTime())
	if err != nil {
		klog.Error("[consumption]create:", err.Error())
		return -1, errno.TimeError.ErrorCode, errno.TimeError.ErrorMsg
	}
	consumption := db.NewConsumption()
	consumption.Amount = req.GetAmount()
	consumption.ConsumptionName = req.GetConsumptionName()
	consumption.Credential = req.GetCredential()
	consumption.Store = req.GetStore()
	consumption.Description = req.GetDescription()
	consumption.TypeId = req.GetTypeId()
	consumption.ConsumeTime = con_time
	id, err := db.CreateConsumption(consumption)
	if err != nil {
		klog.Error("[consumption]create:", err.Error())
		return -1, errno.CreateError.ErrorCode, errno.CreateError.ErrorMsg
	}

	return id, errno.StatusSuccessCode, errno.StatusSuccessMsg
}

func (c *ConsumptionService) GetSumBalance(userId int64) (int64, string, float64) {
	ledgerIds := db.GetLedgersByUserId(userId)
	consumptions := db.GetConSumByRange(ledgerIds)

	sum := 0.0
	for _, val := range consumptions {
		sum += val.Amount
	}

	return errno.StatusSuccessCode, errno.StatusSuccessMsg, sum
}

// 获取所有的支出/收入
func (c *ConsumptionService) GetSum(userId int64, op float64) (int64, string, float64) {
	ledgerIds := db.GetLedgersByUserId(userId)
	consumptions := db.GetConSumByRange(ledgerIds)

	sum := 0.0
	for _, val := range consumptions {
		x := val.Amount
		if x*op > 0 {
			sum += x
		}
	}

	return errno.StatusSuccessCode, errno.StatusSuccessMsg, sum
}

func (c *ConsumptionService) GetDayBalance(start string, end string, userId int64) (int64, string, float64) {
	ledgerIds := db.GetLedgersByUserId(userId)
	consumptions := db.GetConByRange(start, end, ledgerIds)

	sum := 0.0
	for _, val := range consumptions {
		sum += val.Amount
	}
	return errno.StatusSuccessCode, errno.StatusSuccessMsg, sum
}
