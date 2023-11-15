package pack

import "github.com/XZ0730/runFzu/biz/model/asr"

func PackASRResponse(resp *asr.ASRResponse, code int64, msg string, data string) {
	resp.Base = new(asr.BaseResponse)
	resp.Base.Code = code
	resp.Base.Message = msg
	resp.Data = data
}
