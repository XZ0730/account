// Code generated by hertz generator. DO NOT EDIT.

package consumption

import (
	consumption "github.com/XZ0730/runFzu/biz/handler/consumption"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		_api.POST("/consumption", append(_createconsumptionMw(), consumption.CreateConsumption)...)
		_consumption := _api.Group("/consumption", _consumptionMw()...)
		_consumption.GET("/inout", append(_getoutmonthMw(), consumption.GetOutMonth)...)
		{
			_day := _consumption.Group("/day", _dayMw()...)
			_day.GET("/balance", append(_getdaybalanceMw(), consumption.GetDayBalance)...)
		}
		_consumption.GET("/sum", append(_getsumMw(), consumption.GetSum)...)
		_sum := _consumption.Group("/sum", _sumMw()...)
		_sum.GET("/balance", append(_getsumbalanceMw(), consumption.GetSumBalance)...)
		_api.PUT("/consumption", append(_updateconsumptionMw(), consumption.UpdateConsumption)...)
		_consumption0 := _api.Group("/consumption", _consumption0Mw()...)
		_consumption0.GET("/date", append(_getconsumptionbydateMw(), consumption.GetConsumptionByDate)...)
		_api.GET("/consumption", append(_getuseconsumptionMw(), consumption.GetUseConsumption)...)
		_consumption1 := _api.Group("/consumption", _consumption1Mw()...)
		{
			_balance := _consumption1.Group("/balance", _balanceMw()...)
			_balance.GET("/month", append(_getbalancebymonthMw(), consumption.GetBalanceByMonth)...)
			_balance.GET("/year", append(_getbalancebyyearMw(), consumption.GetBalanceByYear)...)
		}
		{
			_last := _consumption1.Group("/last", _lastMw()...)
			{
				_month := _last.Group("/month", _monthMw()...)
				_month.GET("/analysis", append(_getlastmonthmoneyMw(), consumption.GetLastMonthMoney)...)
			}
		}
		{
			_month0 := _consumption1.Group("/month", _month0Mw()...)
			_month0.GET("/map", append(_getlocalmonthconsumptionMw(), consumption.GetLocalMonthConsumption)...)
		}
		{
			_range := _consumption1.Group("/range", _rangeMw()...)
			_range.GET("/in", append(_getinbyrangeMw(), consumption.GetInByRange)...)
			_range.GET("/map", append(_getconsumptionbyrangeMw(), consumption.GetConsumptionByRange)...)
			_range.GET("/out", append(_getoutbyrangeMw(), consumption.GetOutByRange)...)
		}
		{
			_consumption2 := _api.Group("/consumption", _consumption2Mw()...)
			{
				_day0 := _consumption2.Group("/day", _day0Mw()...)
				_day0.GET("/out", append(_getdayoutMw(), consumption.GetDayOut)...)
			}
			{
				_in := _consumption2.Group("/in", _inMw()...)
				{
					_month1 := _in.Group("/month", _month1Mw()...)
					_month1.GET("/sum", append(_getsuminMw(), consumption.GetSumIn)...)
				}
			}
			{
				_out := _consumption2.Group("/out", _outMw()...)
				_out.GET("/sum", append(_gettotaloutMw(), consumption.GetTotalOut)...)
			}
		}
	}
}
