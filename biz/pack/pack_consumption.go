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
