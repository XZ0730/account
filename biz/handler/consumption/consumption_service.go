// Code generated by hertz generator.

package consumption

import (
	"context"
	"strconv"
	"time"

	"github.com/XZ0730/runFzu/biz/model/base"
	"github.com/XZ0730/runFzu/biz/model/consumption"
	"github.com/XZ0730/runFzu/biz/pack"
	"github.com/XZ0730/runFzu/biz/service"
	"github.com/XZ0730/runFzu/pkg/errno"
	"github.com/XZ0730/runFzu/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/klog"
)

// UpdateConsumption .
// @router /api/consumption [PUT]
func UpdateConsumption(ctx context.Context, c *app.RequestContext) {
	var err error
	var req consumption.ConsumptionModel
	resp := new(consumption.ConsumptionUpdateResponse)
	baseResp := new(base.BaseResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.PackBase(baseResp, errno.ParamErrorCode, errno.ParamError.ErrorMsg)
		c.JSON(consts.StatusOK, baseResp)
		return
	}
	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))
	userId := claim.UserId
	code, msg := service.NewConsumptionService().UpdateConsumption(userId, &req)
	pack.PackUpdateConsumptionResp(resp, &req, code, msg)
	c.JSON(consts.StatusOK, resp)
}

// GetConsumptionByRange .
// @router /api/consumption/range/map [GET]
func GetConsumptionByRange(ctx context.Context, c *app.RequestContext) {
	start := c.Query("start")
	end := c.Query("end")
	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))
	consumptions, code, msg := service.NewConsumptionService().GetConsumptionsByRange(start, end, claim.UserId)
	resp := new(consumption.GetConsumptionByRangeResponse)
	pack.PackConsumptionByRangeResp(resp, code, msg, consumptions)
	c.JSON(consts.StatusOK, resp)
}

// GetConsumptionByDate .
// @router /api/consumption/date [GET]
func GetConsumptionByDate(ctx context.Context, c *app.RequestContext) {
	var err error
	resp := new(consumption.GetConsumptionByRangeResponse)
	date := c.Query("date")
	date_time, err := time.Parse(time.DateTime, date)
	if date == "" || err != nil {
		pack.PackConsumptionByRangeResp(resp, errno.ParamError.ErrorCode, errno.ParamError.ErrorMsg, nil)
		c.JSON(consts.StatusOK, resp)
		return
	}
	types := c.Query("type")
	the_type, err := strconv.ParseInt(types, 10, 32)
	if err != nil {
		pack.PackConsumptionByRangeResp(resp, errno.ParamError.ErrorCode, errno.ParamError.ErrorMsg, nil)
		c.JSON(consts.StatusOK, resp)
		return
	}
	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))
	code, msg, cm := service.NewConsumptionService().GetConsumptionByDate(claim.UserId, date_time, the_type)
	pack.PackConsumptionByRangeResp(resp, code, msg, cm)
	c.JSON(consts.StatusOK, resp)
}

// GetOutByRange .
// @router api/consumption/range/out [GET]
func GetOutByRange(ctx context.Context, c *app.RequestContext) {
	start := c.Query("start")
	end := c.Query("end")
	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))

	resp := new(consumption.GetSumByRangeResponse)
	code, msg, sum := service.NewConsumptionService().GetSumByRange(start, end, claim.UserId, -1)
	pack.PackSumRangeResp(resp, code, msg, sum)
	c.JSON(consts.StatusOK, resp)
}

// GetInByRange .
// @router /api/consumption/range/in [GET]
func GetInByRange(ctx context.Context, c *app.RequestContext) {
	start := c.Query("start")
	end := c.Query("end")
	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))

	resp := new(consumption.GetSumByRangeResponse)
	code, msg, sum := service.NewConsumptionService().GetSumByRange(start, end, claim.UserId, 1)
	pack.PackSumRangeResp(resp, code, msg, sum)
	c.JSON(consts.StatusOK, resp)
}

// GetLastMonthMoney .
// @router /api/consumption/last/month/analysis [GET]
func GetLastMonthMoney(ctx context.Context, c *app.RequestContext) {
	var err error
	resp := new(consumption.GetLastMonthMoneyResp)
	d := c.Query("date")
	currentTime, err := time.Parse(time.DateTime, d)
	sum := []float64{0}
	if err != nil {
		pack.PackLastMonthSumResp(resp, errno.ParamErrorCode, err.Error(), sum)
		c.JSON(consts.StatusOK, resp)
		return
	}

	beforeTime := currentTime.AddDate(0, -1, 0)
	end := currentTime.Format(time.DateTime)
	start := beforeTime.Format(time.DateTime)

	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))

	code, msg, sum := service.NewConsumptionService().GetConsumptionSumListByRange(start, end, claim.UserId)
	pack.PackLastMonthSumResp(resp, code, msg, sum)
	c.JSON(consts.StatusOK, resp)
}

// GetLocalMonthConsumption .
// @router /api/consumption/month/map [GET]
func GetLocalMonthConsumption(ctx context.Context, c *app.RequestContext) {
	var err error
	resp := new(consumption.GetConsumptionByRangeResponse)
	d := c.Query("date")
	days := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	year, err := strconv.Atoi(d[0:4])
	if err != nil {
		pack.PackConsumptionByRangeResp(resp, errno.ParamErrorCode, "时间格式错误", nil)
		c.JSON(consts.StatusOK, resp)
		return
	}
	if year%400 == 0 || year%4 == 0 && year%100 != 0 {
		days[2]++
	}

	x := "00:00:00"
	y := "23:59:59"

	a, err := strconv.Atoi(string(d[5]))
	if err != nil {
		pack.PackConsumptionByRangeResp(resp, errno.ParamErrorCode, "时间格式错误", nil)
		c.JSON(consts.StatusOK, resp)
	}
	b, err := strconv.Atoi(string(d[6]))
	if err != nil {
		pack.PackConsumptionByRangeResp(resp, errno.ParamErrorCode, "时间格式错误", nil)
		c.JSON(consts.StatusOK, resp)
	}

	m := 10*a + b
	cnt := strconv.Itoa(days[m])

	start := d[0:8] + "01" + " " + x
	end := d[0:8] + cnt + " " + y

	klog.Info(start)
	klog.Info(end)
	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))
	consumptions, code, msg := service.NewConsumptionService().GetConsumptionsByRange(start, end, claim.UserId)
	pack.PackConsumptionByRangeResp(resp, code, msg, consumptions)
	c.JSON(consts.StatusOK, resp)
}

// GetBalanceByMonth .
// @router /api/consumption/balance/month [GET]
func GetBalanceByMonth(ctx context.Context, c *app.RequestContext) {
	var err error
	resp := new(consumption.GetSumByRangeResponse)
	d := c.Query("date")
	currentTime, err := time.Parse(time.DateTime, d)

	if err != nil {
		pack.PackSumRangeResp(resp, errno.ParamErrorCode, "时间格式错误", 0)
		c.JSON(consts.StatusOK, resp)
		return
	}

	// 获取月初和月末
	monthStart := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, time.Local)
	monthEnd := monthStart.AddDate(0, 1, 0).Add(-time.Nanosecond)

	klog.Info(monthStart)
	klog.Info(monthEnd)

	start := monthStart.Format(time.DateTime)
	end := monthEnd.Format(time.DateTime)

	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))
	code, msg, sum := service.NewConsumptionService().GetSumByRange(start, end, claim.UserId, 0.0)

	pack.PackSumRangeResp(resp, code, msg, sum)
	c.JSON(consts.StatusOK, resp)
}

// GetBalanceByYear .
// @router /api/consumption/balance/year [GET]
func GetBalanceByYear(ctx context.Context, c *app.RequestContext) {
	var err error
	resp := new(consumption.GetSumByRangeResponse)
	d := c.Query("date")
	currentTime, err := time.Parse(time.DateTime, d)

	if err != nil {
		pack.PackSumRangeResp(resp, errno.ParamErrorCode, "时间格式错误", 0)
		c.JSON(consts.StatusOK, resp)
		return
	}

	// 获取年初和年末
	yearStart := time.Date(currentTime.Year(), 1, 1, 0, 0, 0, 0, time.Local)
	yearEnd := time.Date(currentTime.Year(), 12, 31, 23, 59, 59, 59, time.Local)
	klog.Info(yearStart)
	klog.Info(yearEnd)

	start := yearStart.Format(time.DateTime)
	end := yearEnd.Format(time.DateTime)

	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))
	code, msg, sum := service.NewConsumptionService().GetSumByRange(start, end, claim.UserId, 0.0)

	pack.PackSumRangeResp(resp, code, msg, sum)
	c.JSON(consts.StatusOK, resp)
}
