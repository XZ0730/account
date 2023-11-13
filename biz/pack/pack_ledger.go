package pack

import (
	"github.com/XZ0730/runFzu/biz/model/ledger"
)

func PackLedgerCreate(resp *ledger.LedgerCreateResponse, code int64, msg string, model ledger.LedgerModel) {
	resp.Code = code
	resp.Msg = msg
	resp.Data = &model
}

func PackLedgerList(resp *ledger.LedgerListResponse, code int64, msg string, ledgers []*ledger.LedgerModel) {
	resp.Code = code
	resp.Msg = msg
	resp.Data = make(map[string][]*ledger.LedgerModel)
	resp.Data["list"] = ledgers
}

func PackConsumptionList(resp *ledger.LedgerConsumptionListResponse, code int64, msg string, list []*ledger.ConsumptionModel) {
	resp.Code = code
	resp.Msg = msg
	resp.Data = make(map[string][]*ledger.ConsumptionModel)
	resp.Data["list"] = list
}

func PackLedgerBalanceResponse(resp *ledger.LedgerBalanceResponse, balance float64, code int64, msg string) {
	resp.Code = code
	resp.Msg = msg
	resp.Data = make(map[string]float64)
	resp.Data["balance"] = balance
}
