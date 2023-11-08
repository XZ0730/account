package errno

var (
	// Success
	Success = NewErrNo(SuccessCode, "Success")

	ServiceError             = NewErrNo(ServiceErrorCode, "service is unable to start successfully")
	ServiceInternalError     = NewErrNo(ServiceErrorCode, "service internal error")
	ParamError               = NewErrNo(ParamErrorCode, "parameter error")
	AuthorizationFailedError = NewErrNo(AuthorizationFailedErrCode, "authorization failed")

	// User
	UserExistedError = NewErrNo(ParamErrorCode, "user existed")
	UserNameError    = NewErrNo(UserNameAuthErrorCode, "user name is not exist")
	PWDError         = NewErrNo(PwdErrorCode, "pwd not match")

	// Goal
	GoalTimeError   = NewErrNo(GoalTimeErrorCode, "time set error")
	GoalCreateError = NewErrNo(GoalCreateErrorCode, "goal create error")
	GoalGetError    = NewErrNo(GoalGetErrorCode, "goal get error")
	GoalDelError    = NewErrNo(GoalDelErrorCode, "goal del error")
)
