package db

import (
	"time"
)

type Ledger struct {
	LedgerId   int32
	UserId     int64
	LedgerName string
	CoverMsg   string
	CreateTime time.Time
	UpdateTime time.Time
}

func NewLedger(ledgerId int32, userId int64, ledgerName string, cover string, createTime time.Time, updateTime time.Time) *Ledger {
	return &Ledger{LedgerId: ledgerId, UserId: userId, LedgerName: ledgerName, CoverMsg: cover, CreateTime: createTime, UpdateTime: updateTime}
}

func CreateLedger(ledger *Ledger) error {
	return DB.Table("t_ledger").Create(&ledger).Error
}

func DeleteLedger(ledger *Ledger) error {
	return DB.Table("t_ledger").Where("user_id = ? and ledger_id = ?", ledger.UserId, ledger.LedgerId).Delete(&ledger).Error
}
func ListLedgers(userId int64) ([]Ledger, error) {
	ledgers := make([]Ledger, 0)
	err := DB.Table("t_ledger").Where("user_id=?", userId).Find(&ledgers).Error
	if err != nil {
		return nil, err
	}
	return ledgers, nil
}

func UpdateLedger(ledger *Ledger) error {
	return DB.Table("t_ledger").Where("ledger_id=? AND user_id=?", ledger.LedgerId, ledger.UserId).
		Updates(&Ledger{LedgerName: ledger.LedgerName, CreateTime: ledger.CreateTime, UpdateTime: ledger.UpdateTime,
			CoverMsg: ledger.CoverMsg}).Error
}

func CheckUserLedger(userId int64, ledgerId int64) bool {
	ledger := Ledger{}
	err := DB.Table("t_ledger").Where("user_id = ? and ledger_id = ?", userId, ledgerId).
		First(&ledger).Error

	return err == nil
}

func ConsumptionList(ledgerId int32) ([]*Consumption, error) {
	consumptions := make([]*Consumption, 0)
	err := DB.Table("t_consumption").
		Joins("JOIN t_ledger_consumption ON t_ledger_consumption.consumption_id=t_consumption.consumption_id AND t_ledger_consumption.ledger_id=?", ledgerId).
		Find(&consumptions).
		Error

	return consumptions, err
}

func DeleteLedgerConsumption(ledgerId int32, consumptionId int64) error {
	return DB.Table("t_ledger_consumption").
		Where("ledger_id = ? and consumption_id = ?", ledgerId, consumptionId).
		Delete(&Ledger{}).
		Error
}
