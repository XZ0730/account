package db

import (
	"time"
)

type Consumption struct {
	ConsumptionId   int64 `gorm:"primaryKey"`
	Amount          float64
	ConsumptionName string
	Description     string
	TypeId          int8
	Store           string
	ConsumeTime     time.Time
	Credential      string
}

type MultiLedgerConsumptionRel struct {
	MultiLedgerId int64
	ConsumptionId int64
	UserId        int64
}

func NewConsumptionAllParams(consumptionId int64, amount float64, consumptionName, description string, typeId int8, store string, consumeTime time.Time, credential string) *Consumption {
	return &Consumption{ConsumptionId: consumptionId, Amount: amount, ConsumptionName: consumptionName, Description: description, TypeId: typeId, Store: store, ConsumeTime: consumeTime, Credential: credential}
}

func NewConsumption() *Consumption {
	return &Consumption{}
}

func JudgeConsumption(cid int64) error {
	return DB.Table("t_consumption").Where("consumption_id=?", cid).First(NewConsumption()).Error
}

func CreateConsumption(consumption *Consumption) (int64, error) {
	cons := &Consumption{
		Amount:          consumption.Amount,
		ConsumptionId:   consumption.ConsumptionId,
		ConsumptionName: consumption.ConsumptionName,
		Description:     consumption.Description,
		TypeId:          consumption.TypeId,
		Store:           consumption.Store,
		ConsumeTime:     consumption.ConsumeTime,
		Credential:      consumption.Credential,
	}
	err := DB.Table("t_consumption").Create(cons).Error
	return cons.ConsumptionId, err
}

func UpdateConsumption(consumption *Consumption) error {
	return DB.Table("t_consumption").Where("consumption_id = ?", consumption.ConsumptionId).Save(&consumption).Error
}

func GetConByRange(start string, end string, ledgerId []*int64) []*Consumption {
	consumptions := make([]*Consumption, 0)
	DB.Table("t_consumption").
		Joins("JOIN t_ledger_consumption ON "+
			"t_ledger_consumption.consumption_id=t_consumption.consumption_id "+
			"AND t_ledger_consumption.ledger_id IN ? AND t_consumption.consume_time between ? AND ?", ledgerId, start, end).
		Find(&consumptions)
	return consumptions
}

// 获取所有的账单
func GetConSumByRange(ledgerId []*int64) []*Consumption {
	consumptions := make([]*Consumption, 0)
	DB.Table("t_consumption").
		Joins("JOIN t_ledger_consumption ON "+
			"t_ledger_consumption.consumption_id=t_consumption.consumption_id "+
			"AND t_ledger_consumption.ledger_id IN ?", ledgerId).
		Find(&consumptions)
	return consumptions
}

func GetConsumptionByLedgerIds(ledgerIds []*int64) []*Consumption {
	consumptions := make([]*Consumption, 0)
	DB.Table("t_consumption").
		Joins("JOIN t_ledger_consumption ON "+
			"t_ledger_consumption.consumption_id=t_consumption.consumption_id "+
			"AND t_ledger_consumption.ledger_id IN ?", ledgerIds).
		Find(&consumptions)

	return consumptions
}

func GetConsumptionsIdsOfMultiledgerByUserId(userId int64) []*int64 {
	ids := make([]*int64, 0)
	DB.Table("t_multi_ledger_consumption").Where("user_id = ?", userId).Pluck("consumption_id", &ids)
	return ids
}

func GetConsumptionsOfMultiledgerByConIds(ids []*int64) []*Consumption {
	cons := make([]*Consumption, 0)
	DB.Table("t_consumption").Where("consumption_id IN ?", ids).First(&cons)
	return cons
}

func GetConsumptionsOfMultiledgerByRangeAndConIds(start string, end string, ids []*int64) []*Consumption {
	cons := make([]*Consumption, 0)
	DB.Table("t_consumption").Where("consumption_id IN ? AND consume_time between ? and ?", ids, start, end).First(cons)
	return cons
}
