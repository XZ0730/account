// Code generated by hertz generator.

package ledger

import (
	"context"
	"github.com/XZ0730/runFzu/biz/model/base"
	ledger "github.com/XZ0730/runFzu/biz/model/ledger"
	"github.com/XZ0730/runFzu/biz/pack"
	"github.com/XZ0730/runFzu/biz/service"
	"github.com/XZ0730/runFzu/pkg/errno"
	"github.com/XZ0730/runFzu/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// LedgerCreate .
// @router /api/ledger [POST]
func LedgerCreate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ledger.LedgerModel
	resp := new(ledger.LedgerCreateResponse)
	baseResp := new(base.BaseResponse)

	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))
	err = c.BindAndValidate(&req)

	req.UserId = claim.UserId

	if err != nil {
		pack.PackBase(baseResp, errno.ParamErrorCode, errno.ParamError.ErrorMsg)
		c.JSON(consts.StatusOK, resp)
		return
	}
	code, msg := service.NewLedgerService().CreateLedger(claim.UserId, &req)
	pack.PackLedgerCreate(resp, code, msg, req)
	c.JSON(consts.StatusOK, resp)
}

// LedgerDelete .
// @router /api/ledger [DELETE]
func LedgerDelete(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ledger.LedgerModel
	baseResp := new(base.BaseResponse)

	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))
	err = c.BindAndValidate(&req)

	req.UserId = claim.UserId

	if err != nil {
		pack.PackBase(baseResp, errno.ParamErrorCode, errno.ParamError.ErrorMsg)
		c.JSON(consts.StatusOK, baseResp)
		return
	}

	code, msg := service.NewLedgerService().DeleteLedger(claim.UserId, &req)
	pack.PackBase(baseResp, code, msg)
	c.JSON(consts.StatusOK, baseResp)
}

// LedgerUpdate .
// @router /api/ledger [PUT]
func LedgerUpdate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ledger.LedgerModel
	resp := new(ledger.LedgerCreateResponse)
	baseResp := new(base.BaseResponse)

	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))
	err = c.BindAndValidate(&req)

	if err != nil {
		pack.PackBase(baseResp, errno.ParamErrorCode, errno.ParamError.ErrorMsg)
		c.JSON(consts.StatusOK, baseResp)
		return
	}

	req.UserId = claim.UserId
	code, msg := service.NewLedgerService().UpdateLedger(&req)
	pack.PackLedgerCreate(resp, code, msg, req)
	c.JSON(consts.StatusOK, resp)
}

// LedgerList .
// @router /api/ledger [GET]
func LedgerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ledger.LedgerModel
	resp := new(ledger.LedgerListResponse)
	baseResp := new(base.BaseResponse)
	token_byte := c.GetHeader("token")
	claim, _ := utils.CheckToken(string(token_byte))
	err = c.BindAndValidate(&req)

	req.UserId = claim.UserId

	if err != nil {
		pack.PackBase(baseResp, errno.ParamErrorCode, errno.ParamError.ErrorMsg)
		c.JSON(consts.StatusOK, baseResp)
		return
	}

	ledgers, code, msg := service.NewLedgerService().ListLedgers(claim.UserId)
	pack.PackLedgerList(resp, code, msg, ledgers)
	c.JSON(consts.StatusOK, resp)
}
