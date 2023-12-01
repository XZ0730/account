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

struct FileUploadResponse{
    1: BaseResponse base
    2: map<string,list<string>> data
}

struct ConsumptionModel{
    1: i64 consumptionId
    2: string consumptionName
    3: string description
    4: double amount
    5: i8 typeId
    6: string store
    7: string consumeTime
    8: string credential
}

struct OCRReceipt{
    1:i64 code
    2:string msg
    3:ConsumptionModel data
}

service ASRService{
    ASRResponse ASRtoText(1:BaseRequest req)(api.post="/api/speech_recog")
    FileUploadResponse FileUpload(1:BaseRequest req)(api.post="/api/file")
}