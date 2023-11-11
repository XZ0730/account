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

func NewConsumption() *Consumption {
	return &Consumption{}
}

func JudgeConsumption(cid int64) error {
	return DB.Table("t_consumption").Where("consumption_id=?", cid).First(NewConsumption()).Error
}
