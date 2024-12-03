package vosk

import (
	"encoding/json"
	"io"
	"os"

	"github.com/skys-mission/gout/cgo/ai/audio/vosk/lib"
)

type LogLevel int

const (
	LogLevelError  LogLevel = -1
	LogLevelInfo   LogLevel = 0
	LogLevelDetail LogLevel = 1
)

type Word struct {
	Conf  float64 `json:"conf"`  // 置信度
	Start float64 `json:"start"` // 开始时间
	End   float64 `json:"end"`   // 结束时间
	Word  string  `json:"word"`  // 单词
}

type VResult struct {
	Result []*Word `json:"result"` // 存储单词的切片
	Text   string  `json:"text"`   // 文本内容
}

type RecognizeParams struct {
	AudioFilePath string
	ModelPath     string
	SampleRate    float64
}

// ProgressCallback 是进度回调函数类型
type ProgressCallback func(progress float64)

func RecognizeWithProgress(params *RecognizeParams, progressCb ProgressCallback) (results []*VResult, err error) {
	model, err := lib.NewModel(params.ModelPath)
	if err != nil {
		return nil, err
	}
	vr, err := lib.NewRecognizer(model, params.SampleRate)
	if err != nil {
		return nil, err
	}
	vr.SetWords(1)
	file, err := os.Open(params.AudioFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 获取文件大小
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	totalSize := fileInfo.Size()

	bufferSize := 4096
	if totalSize > 1024*1024 { // 1MB
		bufferSize = 8192 // 64KB
	}

	buf := make([]byte, bufferSize)

	results = make([]*VResult, 0)
	processedBytes := 0
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}
		processedBytes += n
		// 计算进度百分比
		if progressCb != nil {
			progress := float64(processedBytes) / float64(totalSize) * 100
			progressCb(progress)
		}
		if vr.AcceptWaveform(buf) != 0 {
			res := new(VResult)
			err = json.Unmarshal([]byte(vr.Result()), res)
			if err != nil {
				return nil, err
			}
			//if res.Result == nil {
			//	res.Result = make([]*Word, 0)
			//}
			results = append(results, res)
		}
	}
	if vr.AcceptWaveform(buf) != 0 {
		res := new(VResult)
		err = json.Unmarshal([]byte(vr.Result()), res)
		if err != nil {
			return nil, err
		}
		//if res.Result == nil {
		//	res.Result = make([]*Word, 0)
		//}
		results = append(results, res)
	}
	return results, nil
}

func Recognize(params *RecognizeParams) (results []*VResult, err error) {
	model, err := lib.NewModel(params.ModelPath)
	if err != nil {
		return nil, err
	}
	vr, err := lib.NewRecognizer(model, params.SampleRate)
	if err != nil {
		return nil, err
	}
	vr.SetWords(1)
	file, err := os.Open(params.AudioFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 获取文件大小
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	totalSize := fileInfo.Size()

	bufferSize := 4096
	if totalSize > 1024*1024 { // 1MB
		bufferSize = 8192 // 64KB
	}

	buf := make([]byte, bufferSize)

	results = make([]*VResult, 0)
	for {
		_, err := file.Read(buf)
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}
		if vr.AcceptWaveform(buf) != 0 {
			res := new(VResult)
			err = json.Unmarshal([]byte(vr.Result()), res)
			if err != nil {
				return nil, err
			}
			//if res.Result == nil {
			//	res.Result = make([]*Word, 0)
			//}
			results = append(results, res)
		}
	}
	if vr.AcceptWaveform(buf) != 0 {
		res := new(VResult)
		err = json.Unmarshal([]byte(vr.Result()), res)
		if err != nil {
			return nil, err
		}
		//if res.Result == nil {
		//	res.Result = make([]*Word, 0)
		//}
		results = append(results, res)
	}
	return results, nil
}

func SetLogLevel(level LogLevel) {
	lib.SetLogLevel(int(level))
}
