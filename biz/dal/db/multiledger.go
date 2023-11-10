package db

import "time"

type MultiLedger struct {
	Id              int64
	MultiLedgerId   int64
	MultiLedgerName string
	Description     string
	Password        string
	ModifyTime      time.Time
}

type M_User struct {
	MultiLedgerId int64
	UserId        int64
}

func NewMultiLedger(name string, desc string, pwd string) *MultiLedger {
	return &MultiLedger{
		MultiLedgerName: name,
		Description:     desc,
		Password:        pwd,
		ModifyTime:      time.Now(),
	}
}
func NewM_User(mid, uid int64) *M_User {
	return &M_User{
		MultiLedgerId: mid,
		UserId:        uid,
	}
}

func CreateMultiLedger(ml *MultiLedger) error {

	return DB.Table("t_multi_ledger").Create(&ml).Error
}

func CreateM_user(mid, uid int64) error {
	m_user := NewM_User(mid, uid)
	return DB.Table("t_multi_ledger_user").Create(&m_user).Error
}
