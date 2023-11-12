package db

import (
	"time"
)

type MultiLedger struct {
	MultiLedgerId   int64 `gorm:"primary_key"`
	MultiLedgerName string
	Description     string
	Password        string
	ModifyTime      time.Time
}

type M_User struct {
	MultiLedgerId int64
	UserId        int64
}

type M_Consumption struct {
	MultiLedgerId int64
	ConsumptionId int64
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

func NewM_Consumption(mid, uid, cid int64) *M_Consumption {
	return &M_Consumption{
		UserId:        uid,
		MultiLedgerId: mid,
		ConsumptionId: cid,
	}
}

func CreateMultiLedger(ml *MultiLedger) error {

	return DB.Table("t_multi_ledger").Create(&ml).Error
}

func GetMultiLedgerByPassword(password string) (id int64, err error) {
	ml := new(MultiLedger)
	err = DB.Table("t_multi_ledger").Where("password=?", password).First(&ml).Error
	id = ml.MultiLedgerId
	return
}

func CreateM_user(mid, uid int64) error {
	m_user := NewM_User(mid, uid)
	return DB.Table("t_multi_ledger_user").Create(&m_user).Error
}

func JudgeM_user(mid, uid int64) error {
	ml := new(MultiLedger)
	return DB.Table("t_multi_ledger_user").Where("multi_ledger_id=? AND user_id=?", mid, uid).First(ml).Error
}

func CreateM_Consumption(mid, uid, cid int64) error {
	return DB.Table("t_multi_ledger_consumption").Create(NewM_Consumption(mid, uid, cid)).Error
}

func JudgeM_consumption(mid, uid, cid int64) error {
	return DB.Table("t_multi_ledger_consumption").Where("multi_ledger_id=? AND consumption_id=? AND user_id=?", mid, cid, uid).First(NewM_Consumption(mid, uid, cid)).Error
}

func GetMl_Consumption(mid int64) ([]*Consumption, error) {
	clist := make([]*Consumption, 0)
	err := DB.Table("t_consumption").Joins("JOIN t_multi_ledger_consumption ON t_multi_ledger_consumption.consumption_id=t_consumption.consumption_id AND t_multi_ledger_consumption.multi_ledger_id=?", mid).Find(&clist).Error
	return clist, err
}

func GetMltiledgerList(uid int64) ([]*MultiLedger, error) {
	mlist := make([]*MultiLedger, 0)
	err := DB.Table("t_multi_ledger").Joins("JOIN t_multi_ledger_user ON t_multi_ledger_user.multi_ledger_id = t_multi_ledger.multi_ledger_id AND t_multi_ledger_user.user_id=?", uid).Find(&mlist).Error
	return mlist, err
}
