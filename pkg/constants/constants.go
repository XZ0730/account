package constants

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	JWTValue = "MTAxNTkwMTg1Mw=="

	// snowflake
	SnowflakeWorkerID     = 0
	SnowflakeDatacenterID = 0

	MaxConnections  = 1000
	MaxQPS          = 100
	MaxVideoSize    = 300000
	MaxListLength   = 100
	MaxIdleConns    = 10
	MaxGoroutines   = 10
	MaxOpenConns    = 100
	ConnMaxLifetime = 10 * time.Second

	// AppID AppID
	AppID = "1317777829"
	// SecretID SecretID
	SecretID = "AKIDsdCMXVFk14yiKrCNL3Zs5vStK8saN0Tr"
	// SecretKey SecretKey
	SecretKey = "JMFwRf3GG4eAYnrjMkKl2B71lKfonfMq"
	// EngineModelType EngineModelType
	EngineModelType = "16k_zh"
	// SliceSize SliceSize
	SliceSize = 6400

	BaiDuAppID     = "XluSFusWhyG4dqtSLOqca29r"
	BaiDuSecretKey = "Bd8xmaz0YuNDrdA4Mo7IEGSuLt4k1eQu"
)

var Url = "https://aip.baidubce.com/rest/2.0/ocr/v1/shopping_receipt?access_token=" + GetAccessToken()

// 24.acc4147dc8dd2b95d07b51f48dd3aa18.2592000.1704028691.282335-44118298
// 24.dadf134b6a3780fe51cf16cc66ecc5ff.2592000.1704028658.282335-44118298
// 24.f6d17437a792a18e1d59c63b9d1db1e6.2592000.1704028751.282335-44118298
func GetAccessToken() string {
	url := "https://aip.baidubce.com/oauth/2.0/token"
	postData := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", BaiDuAppID, BaiDuSecretKey)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(postData))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	accessTokenObj := map[string]string{}
	json.Unmarshal([]byte(body), &accessTokenObj)
	return accessTokenObj["access_token"]
}
