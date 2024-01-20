package main

import (
	"fmt"

	"github.com/ggerganov/whisper.cpp/bindings/go/pkg/whisper"
)

func Run_speech_recognition() string {
	var modelpath string
	var samples []float32

	// Load the model
	model, err := whisper.New(modelpath)
	if err != nil {
		panic(err)
	}
	defer model.Close()

	context, err := model.NewContext()
	if err != nil {
		panic(err)
	}
	if err := context.Process(samples, nil, nil); err != nil {
		panic(err)
	}

	var content string = ""
	for {
		segment, err := context.NextSegment()
		if err != nil {
			break
		}
		content = content + " " + segment.Text
		fmt.Printf("[%6s->%6s] %s\n", segment.Start, segment.End, segment.Text)
	}

	return content
}
