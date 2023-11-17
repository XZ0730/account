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

service ASRService{
    ASRResponse ASRtoText(1:BaseRequest req)(api.post="/api/speech_recog")
    FileUploadResponse FileUpload(1:BaseRequest req)(api.post="/api/file")
}