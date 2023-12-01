package utils

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/XZ0730/runFzu/config"
	"github.com/cloudwego/kitex/pkg/klog"
)

func GetMysqlDSN() string {
	if config.Mysql == nil {
		klog.Fatal("config not found")
	}

	dsn := strings.Join([]string{config.Mysql.Username, ":", config.Mysql.Password, "@tcp(", config.Mysql.Addr, ")/", config.Mysql.Database, "?charset=" + config.Mysql.Charset + "&parseTime=true"}, "")

	return dsn
}

func GetFileContentAsBase64(path string) string {
	srcByte, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(srcByte)
}

// func GetMQUrl() string {
// 	if config.RabbitMQ == nil {
// 		klog.Fatal("config not found")
// 	}

// 	url := strings.Join([]string{"amqp://", config.RabbitMQ.Username, ":", config.RabbitMQ.Password, "@", config.RabbitMQ.Addr, "/"}, "")

// 	return url
// }
