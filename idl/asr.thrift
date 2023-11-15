namespace go asr

struct BaseRequest{
}

struct BaseResponse{
    1: required i64 code,
    2: required string message,
}

struct ASRResponse{
    1: BaseResponse base
    2: string data
}

service ASRService{
    ASRResponse ASRtoText(1:BaseRequest req)(api.post="/api/speech_recog")
}