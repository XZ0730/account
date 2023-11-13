// Code generated by hertz generator.

package multiledger

import (
	"context"

	"github.com/XZ0730/runFzu/biz/model/base"
	multiledger "github.com/XZ0730/runFzu/biz/model/multiledger"
	"github.com/XZ0730/runFzu/biz/pack"
	"github.com/XZ0730/runFzu/biz/service"
	"github.com/XZ0730/runFzu/pkg/errno"
	"github.com/XZ0730/runFzu/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CreateMultiledger .
// @router /api/multiLedger [POST]
func CreateMultiledger(ctx context.Context, c *app.RequestContext) {
	var err error
	var req multiledger.CreateMLRequest
	resp := new(base.BaseResponse)
	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.PackBase(resp, errno.ParamError.ErrorCode, errno.ParamError.ErrorMsg)
		c.JSON(consts.StatusOK, resp)
		return
	}
	code, msg := service.NewMultiLedgerService().CreateMultiLedger(claim.UserId, &req)
	pack.PackBase(resp, code, msg)
	c.JSON(consts.StatusOK, resp)
}

// JoinMultiledger .
// @router /api/multiLedger/join [POST]
func JoinMultiledger(ctx context.Context, c *app.RequestContext) {
	var err error
	var req multiledger.JoinMLRequest
	resp := new(base.BaseResponse)
	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.PackBase(resp, errno.ParamError.ErrorCode, errno.ParamError.ErrorMsg)
		c.JSON(consts.StatusOK, resp)
		return
	}
	code, msg := service.NewMultiLedgerService().JoinMultiledger(claim.UserId, req.Password)
	pack.PackBase(resp, code, msg)
	c.JSON(consts.StatusOK, resp)
}

// InsertMlConsumption .
// @router /api/multiLedger/consumption [POST]
func InsertMlConsumption(ctx context.Context, c *app.RequestContext) {
	var err error
	var req multiledger.InsertMlConsumReq
	resp := new(base.BaseResponse)
	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.PackBase(resp, errno.ParamError.ErrorCode, errno.ParamError.ErrorMsg)
		c.JSON(consts.StatusOK, resp)
		return
	}
	code, msg := service.NewMultiLedgerService().InsertMlConsumption(claim.UserId, &req)
	pack.PackBase(resp, code, msg)
	c.JSON(consts.StatusOK, resp)
}

// GetMulConsumption .
// @router /api/multiLedger/consumption [GET]
func GetMulConsumption(ctx context.Context, c *app.RequestContext) {
	var err error
	var req multiledger.GetMulConsumptionReq
	resp := new(multiledger.GetMulConsumptionResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.PackML_GetConsumption(resp, errno.ParamError.ErrorCode, errno.ParamError.ErrorMsg, nil)
		c.JSON(consts.StatusOK, resp)
		return
	}
	cm, code, msg := service.NewMultiLedgerService().GetMulConsumption(req.GetMultiLedgerId())
	pack.PackML_GetConsumption(resp, code, msg, cm)
	c.JSON(consts.StatusOK, resp)
}

// GetMultiLedgerList .
// @router /api/multiLedger [POST]
func GetMultiLedgerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req multiledger.BaseRequest
	resp := new(multiledger.GetMultiLedgerListResp)
	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.PackML_GetMLlist(resp, errno.ParamError.ErrorCode, errno.ParamError.ErrorMsg, nil)
		c.JSON(consts.StatusOK, resp)
		return
	}
	mm, code, msg := service.NewMultiLedgerService().GetMultiLedgerList(claim.UserId)
	pack.PackML_GetMLlist(resp, code, msg, mm)
	c.JSON(consts.StatusOK, resp)
}

// DelMultiLedger .
// @router /api/multiLedger [DELETE]
func DelMultiLedger(ctx context.Context, c *app.RequestContext) {
	var err error
	var req multiledger.DelMultiLedgerReq
	resp := new(base.BaseResponse)
	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.PackBase(resp, errno.ParamError.ErrorCode, errno.ParamError.ErrorMsg)
		c.JSON(consts.StatusOK, resp)
		return
	}
	code, msg := service.NewMultiLedgerService().DelMultiLedger(claim.UserId, req.GetMultiLedgerId())
	pack.PackBase(resp, code, msg)
	c.JSON(consts.StatusOK, resp)
}

// PutMultiLedger .
// @router /api/multiLedger [PUT]
func PutMultiLedger(ctx context.Context, c *app.RequestContext) {
	var err error
	var req multiledger.PutMultiLedgerReq
	resp := new(base.BaseResponse)
	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.PackBase(resp, errno.ParamError.ErrorCode, errno.ParamError.ErrorMsg)
		c.JSON(consts.StatusOK, resp)
		return
	}

	code, msg := service.NewMultiLedgerService().PutMultiLedger(claim.UserId, &req)
	pack.PackBase(resp, code, msg)
	c.JSON(consts.StatusOK, resp)
}

// DelMulConsumption .
// @router /api/multiLedger/consumption [DELETE]
func DelMulConsumption(ctx context.Context, c *app.RequestContext) {
	var err error
	var req multiledger.DelMulConsumptionReq
	resp := new(base.BaseResponse)
	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.PackBase(resp, errno.ParamError.ErrorCode, errno.ParamError.ErrorMsg)
		c.JSON(consts.StatusOK, resp)
		return
	}
	code, msg := service.NewMultiLedgerService().DelMulConsumption(claim.UserId, &req)
	pack.PackBase(resp, code, msg)
	c.JSON(consts.StatusOK, resp)
}
