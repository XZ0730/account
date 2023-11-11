package db

type LedgerConsumptionRel struct {
	LedgerId      int32
	ConsumptionId int64
}

func NewLedgerConsumptionRel(ledgerId int32, consumptionId int64) *LedgerConsumptionRel {
	return &LedgerConsumptionRel{LedgerId: ledgerId, ConsumptionId: consumptionId}
}

func CreateLedgerConsumptionRel(rel *LedgerConsumptionRel) error {
	return DB.Table("t_ledger_consumption").Create(&rel).Error
}
