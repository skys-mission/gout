package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	vosk "github.com/skys-mission/gout/cgo/ai/audio/vosk/lib"
)

func main() {
	//var filename string
	//flag.StringVar(&filename, "f", "", "file to transcribe")
	//flag.Parse()

	model, err := vosk.NewModel("..\\demo\\vosk-model")
	if err != nil {
		log.Fatal(err)
	}
	//
	//// we can check if word is in the vocabulary
	//// fmt.Println(model.FindWord("air"))
	//
	sampleRate := 16000.0
	rec, err := vosk.NewRecognizer(model, sampleRate)
	if err != nil {
		log.Fatal(err)
	}
	rec.SetWords(1)

	file, err := os.Open("..\\demo\\test.wav")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, 4096)

	for {
		_, err := file.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}

			break
		}

		if rec.AcceptWaveform(buf) != 0 {
			fmt.Println(rec.Result())
		}
	}

	// Unmarshal example for final result
	var jres map[string]interface{}
	json.Unmarshal([]byte(rec.FinalResult()), &jres)
	fmt.Println(jres["text"])
}
