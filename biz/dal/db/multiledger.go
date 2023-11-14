package db

import (
	"time"

	"github.com/XZ0730/runFzu/biz/model/multiledger"
	"github.com/cloudwego/kitex/pkg/klog"
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

func DelMultiLedger(mid int64) error {
	return DB.Table("t_multi_ledger").Where("multi_ledger_id=?", mid).Unscoped().Delete(&MultiLedger{MultiLedgerId: mid}).Error
}

func DelMultiLedgerConsumption(mid int64) error {
	clist := make([]*Consumption, 0)
	DB.Table("t_consumption").Joins("JOIN t_multi_ledger_consumption ON t_multi_ledger_consumption.consumption_id = t_consumption.consumption_id AND t_multi_ledger_consumption.multi_ledger_id=?", mid).Find(&clist)
	for _, v := range clist {
		err := DB.Table("t_consumption").Where("consumption_id=?", v.ConsumptionId).Unscoped().Delete(v).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateMultiLedger(ml *MultiLedger) error {
	return DB.Table("t_multi_ledger").Where("multi_ledger_id=?", ml.MultiLedgerId).Updates(&ml).Error
}

func DelSpecialConsumption(uid, mid, cid int64) error {
	m := NewM_Consumption(mid, uid, cid)
	err := DB.Table("t_multi_ledger_consumption").Where("multi_ledger_id=? AND consumption_id=?", mid, cid).Unscoped().Delete(&m).Error
	if err != nil {
		klog.Error("[multi_db]error:", err.Error)
		return err
	}

	err = DB.Table("t_consumption").Where("consumption_id=?", cid).Delete(&Consumption{}).Error
	if err != nil {
		klog.Error("[multi_db]error:", err.Error)
		return err
	}
	return nil
}

func GetMultiLedgerBalance(ledgerId int64) (float64, error) {
	balance := 0.0
	consumptions := make([]*Consumption, 0)
	err := DB.Table("t_consumption").
		Joins("JOIN t_multi_ledger_consumption ON "+
			"t_multi_ledger_consumption.consumption_id=t_consumption.consumption_id "+
			"AND t_multi_ledger_consumption.multi_ledger_id=?", ledgerId).
		Find(&consumptions).Error

	if err != nil {
		return 0, err
	}

	for _, c := range consumptions {
		balance += c.Amount
	}
	return balance, err
}

func CreateMl_Consumption(uid, mid, cid int64) error {
	mlc := NewM_Consumption(mid, uid, cid)
	return DB.Table("t_multi_ledger_consumption").Create(&mlc).Error
}

func GetML_Users(mid int64) ([]*multiledger.UserModel, error) {
	um := make([]*multiledger.UserModel, 0)
	err := DB.Table("t_user").Joins("JOIN t_multi_ledger_user ON t_multi_ledger_user.user_id=t_user.user_id AND t_multi_ledger_user.multi_ledger_id=?", mid).Find(&um).Error
	return um, err
}
