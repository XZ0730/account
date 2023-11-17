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
	resp.Data = make(map[string]*consumption.TimeKeyConArray)
	timeKeyConArray := new(consumption.TimeKeyConArray)
	timeKeyConArray.Tmap = make(map[string][]*consumption.ConsumptionModel)
	for _, val := range cons {
		timeKeyConArray.Tmap[val.ConsumeTime[0:10]] = append(timeKeyConArray.Tmap[val.ConsumeTime], val)
	}
	resp.Data["data"] = timeKeyConArray
}
