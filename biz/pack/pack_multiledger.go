package pack

import "github.com/XZ0730/runFzu/biz/model/multiledger"

func PackML_GetConsumption(resp *multiledger.GetMulConsumptionResp, code int64, msg string, list []*multiledger.ConsumptionModel) {
	resp.Base = multiledger.NewBaseResponse()
	resp.Base.Code = code
	resp.Base.Message = msg
	resp.Data = make(map[string][]*multiledger.ConsumptionModel)
	resp.Data["list"] = list
}

func PackML_GetMLlist(resp *multiledger.GetMultiLedgerListResp, code int64, msg string, list []*multiledger.MultiledgerModel) {
	resp.Base = multiledger.NewBaseResponse()
	resp.Base.Code = code
	resp.Base.Message = msg
	resp.Data = make(map[string][]*multiledger.MultiledgerModel)
	resp.Data["list"] = list
}

func PackML_Balance(resp *multiledger.GetMultiBalanceResp, code int64, msg string, balance float64) {
	resp.Base = multiledger.NewBaseResponse()
	resp.Base.Code = code
	resp.Base.Message = msg
	resp.Balance = balance
}

func PackML_Users(resp *multiledger.GetMultiUsersResp, code int64, msg string, data []*multiledger.UserModel) {
	resp.Base = multiledger.NewBaseResponse()
	resp.Base.Code = code
	resp.Base.Message = msg
	resp.Data = make(map[string][]*multiledger.UserModel)
	resp.Data["list"] = data
}
