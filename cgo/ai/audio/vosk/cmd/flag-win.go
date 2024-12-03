//go:build windows

package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
)

// 自定义帮助信息
func printHelp() {
	var buf bytes.Buffer
	// 将 flag.PrintDefaults 的输出捕获到 buffer 中
	flag.CommandLine.SetOutput(&buf)
	flag.PrintDefaults()
	flag.CommandLine.SetOutput(os.Stderr) // 恢复默认输出

	// 按照期望顺序打印所有内容
	fmt.Printf("Usage: %s \n", os.Args[0])
	fmt.Println("Options:")
	fmt.Print(buf.String())
	fmt.Println("Example:")
	fmt.Printf("example 1:\n%s -audio=\".\\audio.wav\" -model=\".\\vosk-model\" -rate=16000\n", os.Args[0])
	fmt.Printf("example 2:\n%s -audio=\".\\audio.wav\" -model=\".\\vosk-model\" -rate=16000 -out=\".\\putout.json\"\n", os.Args[0])
}

type params struct {
	AudioFilePath   string
	ModelPath       string
	SampleRate      float64
	OutJsonFilePath string
}

func setFlag() (*params, error) {
	// 定义命令行标志
	audioFilePath := flag.String("audio", "", "Path to the audio file (required)")
	modelPath := flag.String("model", "", "Path to the Vosk model (required)")
	sampleRate := flag.Float64("rate", 16000, "Audio sample rate")
	outJsonFilePath := flag.String("out", "", "Output json file path (Disabled by default, outputs to stdout)")
	flag.Usage = printHelp
	// 解析命令行标志
	flag.Parse()
	if *audioFilePath == "" || *modelPath == "" {
		return nil, fmt.Errorf("audio or model is nil")
	}
	return &params{
		AudioFilePath:   *audioFilePath,
		ModelPath:       *modelPath,
		SampleRate:      *sampleRate,
		OutJsonFilePath: *outJsonFilePath}, nil
}
