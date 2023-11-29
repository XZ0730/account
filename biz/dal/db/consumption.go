package db

import (
	"time"
)

type Consumption struct {
	ConsumptionId   int64 `gorm:"primary key"`
	Amount          float64
	ConsumptionName string
	Description     string
	TypeId          int8
	Store           string
	ConsumeTime     time.Time
	Credential      string
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

func CreateConsumption(consumption *Consumption) error {
	return DB.Table("t_consumption").Create(&consumption).Error
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


func GetConsumptionByLedgerIds(ledgerIds []*int64) []*Consumption {
	consumptions := make([]*Consumption, 0)
	DB.Table("t_consumption").
		Joins("JOIN t_ledger_consumption ON "+
			"t_ledger_consumption.consumption_id=t_consumption.consumption_id "+
			"AND t_ledger_consumption.ledger_id IN ?", ledgerIds).
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
