package utils

import (
	"fmt"
	"mime/multipart"
	"os"
	"sync"
	"time"

	"github.com/XZ0730/runFzu/pkg/constants"
	"github.com/tencentcloud/tencentcloud-speech-sdk-go/asr"
	"github.com/tencentcloud/tencentcloud-speech-sdk-go/common"
)

var (
	Data string
	Mx   sync.Mutex
)

type MySpeechRecognitionListener struct {
	ID int
}

// OnRecognitionStart implementation of SpeechRecognitionListener
func (listener *MySpeechRecognitionListener) OnRecognitionStart(response *asr.SpeechRecognitionResponse) {
	fmt.Printf("%s|%s|OnRecognitionStart\n", time.Now().Format("2006-01-02 15:04:05"), response.Result.VoiceTextStr)
}

// OnSentenceBegin implementation of SpeechRecognitionListener
func (listener *MySpeechRecognitionListener) OnSentenceBegin(response *asr.SpeechRecognitionResponse) {
	fmt.Printf("%s|%s|OnSentenceBegin: %v\n", time.Now().Format("2006-01-02 15:04:05"), response.Result.VoiceTextStr, response.Message)
}

// OnRecognitionResultChange implementation of SpeechRecognitionListener
func (listener *MySpeechRecognitionListener) OnRecognitionResultChange(response *asr.SpeechRecognitionResponse) {
	fmt.Printf("%s|%s|OnRecognitionResultChange: %v\n", time.Now().Format("2006-01-02 15:04:05"), response.Result.VoiceTextStr, response.Message)
}

// OnSentenceEnd implementation of SpeechRecognitionListener
func (listener *MySpeechRecognitionListener) OnSentenceEnd(response *asr.SpeechRecognitionResponse) {
	fmt.Printf("%s|%s|OnSentenceEnd: %v\n", time.Now().Format("2006-01-02 15:04:05"), response.Result.VoiceTextStr, response)

	Data += response.Result.VoiceTextStr
}

// OnRecognitionComplete implementation of SpeechRecognitionListener
func (listener *MySpeechRecognitionListener) OnRecognitionComplete(response *asr.SpeechRecognitionResponse) {
	fmt.Printf("%s|%s|OnRecognitionComplete\n", time.Now().Format("2006-01-02 15:04:05"), response.Result.VoiceTextStr)
}

// OnFail implementation of SpeechRecognitionListener
func (listener *MySpeechRecognitionListener) OnFail(response *asr.SpeechRecognitionResponse, err error) {
	fmt.Printf("%s|%s|OnFail: %v\n", time.Now().Format("2006-01-02 15:04:05"), response.Result.VoiceTextStr, err)
}

func ProcessLoop(id int, wg *sync.WaitGroup, file *os.File) {
	defer wg.Done()
	for {
		err := Process(id, file)
		if err != nil {
			return
		}
	}
}

var ProxyURL string

func ProcessOnce(id int, wg *sync.WaitGroup, file multipart.File) {
	defer wg.Done()
	Process(id, file)
}

func Process(id int, file multipart.File) error {
	audio := file
	defer audio.Close()
	listener := &MySpeechRecognitionListener{
		ID: id,
	}
	credential := common.NewCredential(constants.SecretID, constants.SecretKey)
	recognizer := asr.NewSpeechRecognizer(constants.AppID, credential, constants.EngineModelType, listener)
	recognizer.ProxyURL = ProxyURL
	recognizer.VoiceFormat = asr.AudioFormatAAC
	err := recognizer.Start()
	if err != nil {
		fmt.Printf("%s|recognizer start failed, error: %v\n", time.Now().Format("2006-01-02 15:04:05"), err)
		return err
	}
	for {
		data := make([]byte, constants.SliceSize)
		n, err := audio.Read(data)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Printf("read file error: %v\n", err)
			break
		}
		if n <= 0 {
			break
		}
		err = recognizer.Write(data)
		if err != nil {
			break
		}
	}
	recognizer.Stop()
	return nil
}
