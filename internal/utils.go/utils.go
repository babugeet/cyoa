package utils

import (
	"encoding/json"
	"os"
)

type FullStory map[string]StorySection

type StorySection struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func ParseJson(filename string) FullStory {
	fileByte, _ := os.ReadFile(filename)
	var StoryPoint FullStory
	json.Unmarshal(fileByte, &StoryPoint)
	// fmt.Println(StoryPoint)
	return StoryPoint

}
