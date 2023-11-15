package service

import (
	"fmt"
	"mime/multipart"
	"sync"

	"github.com/XZ0730/runFzu/pkg/errno"
	"github.com/XZ0730/runFzu/pkg/utils"
)

func (a *ASRService) ASRtoText(file *multipart.FileHeader) (data string, code int64, msg string) {

	f, err := file.Open()
	if err != nil {
		return "", errno.FileError.ErrorCode, errno.FileError.ErrorMsg
	}
	var wg sync.WaitGroup
	utils.ProxyURL = ""
	for i := 0; i < 1; i++ {
		fmt.Println("Main: Starting worker", i)
		wg.Add(1)
		go utils.ProcessOnce(i, &wg, f)
	}

	fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	data = utils.Data
	utils.Mx.Unlock()
	return data, errno.StatusSuccessCode, errno.StatusSuccessMsg
}
