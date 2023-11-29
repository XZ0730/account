package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"sync"

	"github.com/XZ0730/runFzu/pkg/errno"
	"github.com/XZ0730/runFzu/pkg/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/sync/errgroup"
)

func (a *ASRService) ASRtoText(file *multipart.FileHeader) (data string, code int64, msg string) {

	f, err := file.Open()
	if err != nil {
		return "", errno.FileError.ErrorCode, errno.FileError.ErrorMsg
	}
	utils.Mx.Lock()
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
	utils.Data = ""
	utils.Mx.Unlock()
	return data, errno.StatusSuccessCode, errno.StatusSuccessMsg
}

func (a *ASRService) FileUpload(uid int64, files []*multipart.FileHeader) (data []string, code int64, msg string) {

	var eg errgroup.Group
	file_urls := make([]string, 0)
	for _, fileheader := range files {
		tmp := fileheader
		eg.Go(func() error {
			file, err := tmp.Open()
			if err != nil {
				klog.Error("[file_open]error:", err.Error())
				return err
			}
			code, url := utils.UploadToQiNiu(file, tmp, fmt.Sprint(uid))
			if code != 200 {
				klog.Error("[file_upload]error:upload failed | code:", code, "| url:", url)
				return errors.New("[file_upload]error:upload failed")
			}
			file_urls = append(file_urls, url)
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		klog.Info("[file]error:", err.Error())
		return nil, errno.FileError.ErrorCode, errno.FileError.ErrorMsg
	}
	return file_urls, errno.StatusSuccessCode, errno.StatusSuccessMsg
}
