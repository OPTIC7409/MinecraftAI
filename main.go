package main

import (
	"fmt"
	"log"
	"minecraft-ai/capture"
	"minecraft-ai/process"
	"minecraft-ai/analyze"
	"minecraft-ai/suggest"
)

func main() {
	img, err := capture.CaptureScreen(0)
	if err != nil {
		log.Fatalf("Error capturing screen: %v", err)
	}

	process.ProcessImage(img)

	health := analyze.ExtractHealth(img)

	action := suggest.SuggestAction(health)
	fmt.Println("Suggestion:", action)

	// Optionally, integrate with UI (if implemented)
	// ui.DisplaySuggestion(action)
}
