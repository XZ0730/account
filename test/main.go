package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/XZ0730/runFzu/pkg/constants"
)

func main() {

	url1 := constants.Url
	fmt.Println(url1)
	// image 可以通过 GetFileContentAsBase64("C:\fakepath\B.png") 方法获取
	base64 := GetFileContentAsBase64("/root/goproject/account/test/piao5.jpg")
	image := url.QueryEscape(base64)
	payload := strings.NewReader("image=" + image + "&verify_parameter=false&probability=false&location=false")
	client := &http.Client{}
	req, err := http.NewRequest("POST", url1, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

/**
 * 获取文件base64编码
 * @param string  path 文件路径
 * @return string base64编码信息，不带文件头
 */
func GetFileContentAsBase64(path string) string {
	srcByte, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(srcByte)
}

/**
 * 使用 AK，SK 生成鉴权签名（Access Token）
 * @return string 鉴权签名信息（Access Token）
 */
