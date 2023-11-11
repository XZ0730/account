package db

import "time"

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
