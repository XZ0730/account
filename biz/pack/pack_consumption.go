package pack

import (
	"github.com/XZ0730/runFzu/biz/model/consumption"
)

func PackUpdateConsumptionResp(resp *consumption.ConsumptionUpdateResponse, req *consumption.ConsumptionModel, code int64, msg string) {
	resp.Data = make(map[string]*consumption.ConsumptionModel)
	resp.Data["consumption"] = req
	resp.Code = code
	resp.Msg = msg
}

func PackConsumptionByRangeResp(resp *consumption.GetConsumptionByRangeResponse, code int64, msg string, cons []*consumption.ConsumptionModel) {
	resp.Data = make(map[string][]*consumption.ConsumptionModel, 0)
	for _, val := range cons {
		resp.Data[val.ConsumeTime[0:10]] = append(resp.Data[val.ConsumeTime[0:10]], val)
	}
}

func PackSumRangeResp(resp *consumption.GetSumByRangeResponse, code int64, msg string, sum float64) {
	resp.Data = make(map[string]float64)
	resp.Data["sum"] = sum
	resp.Code = code
	resp.Msg = msg
}
