// Copyright (c) 2024, https://github.com/skys-mission and SoyMilkWhisky

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/skys-mission/gout/cgo/ai/audio/vosk"
)

func main() {
	args, err := setFlag()
	if err != nil {
		fmt.Printf("Error setting flags: %v\n", err)
		os.Exit(-1)
	}

	results, err := vosk.RecognizeWithProgress(&vosk.RecognizeParams{
		AudioFilePath: args.AudioFilePath,
		ModelPath:     args.ModelPath,
		SampleRate:    args.SampleRate,
	}, func(progress float64) {
		fmt.Printf("Processing audio-to-text: %.2f%% completed\n", progress)
	})
	if err != nil {
		fmt.Printf("Error during audio recognition: %v\n", err)
		return
	}

	jsonData, err := json.Marshal(results)
	if err != nil {
		fmt.Printf("Error serializing results to JSON: %v\n", err)
		return
	}

	// 如果没有指定输出文件路径，则直接以字符串格式打印 JSON
	if args.OutJsonFilePath == "" {
		fmt.Printf("Recognition Results (JSON):\n%s\n", string(jsonData))
		return
	}

	// 美化 JSON 数据
	prettyJSON := new(bytes.Buffer)
	err = json.Indent(prettyJSON, jsonData, "", "    ")
	if err != nil {
		fmt.Printf("Error formatting JSON: %v\n", err)
		return
	}

	// 写入到指定文件路径
	outFile, err := os.Create(args.OutJsonFilePath)
	if err != nil {
		fmt.Printf("Error creating output file (%s): %v\n", args.OutJsonFilePath, err)
		return
	}
	defer outFile.Close()

	_, err = outFile.Write(prettyJSON.Bytes())
	if err != nil {
		fmt.Printf("Error writing JSON to file (%s): %v\n", args.OutJsonFilePath, err)
		return
	}

	fmt.Printf("Recognition results successfully written to file: %s\n", args.OutJsonFilePath)
}
