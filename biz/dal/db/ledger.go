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
