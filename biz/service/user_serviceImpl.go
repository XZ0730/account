package service

import (
	"github.com/XZ0730/runFzu/biz/dal/db"
	"github.com/XZ0730/runFzu/biz/model/base"
	"github.com/XZ0730/runFzu/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
)

func (u *UserService) Login(user_name string, pwd string) (code int64, message string, user_id int64) {

	user, err := db.GetUserInfoByUserName(user_name)
	if err != nil {
		klog.Error(err.Error())
		return errno.UserNameError.ErrorCode, errno.UserNameError.ErrorMsg, -1
	}
	if user.Password != pwd {
		klog.Error("pwd not match")
		return errno.PWDError.ErrorCode, errno.PWDError.ErrorMsg, -1
	}

	return 200, "", user.UserId
}

func (u *UserService) Register(req *base.RegisterReq) (code int64, message string) {
	if !db.JudgeUser(req.GetUsername()) {
		klog.Error("error:[user] error: exist")
		return 20002, "[user] error: exist"
	}
	err := db.CreateUser(req.GetUsername(), req.GetPassword())
	if err != nil {
		return 20001, "[user] error:" + err.Error()
	}
	return 200, ""
}
