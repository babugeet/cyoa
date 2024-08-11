package storyserver

import (
	"fmt"

	"github.com/cyoa/internal/utils.go"
)

func StartStory(storycontent utils.FullStory) {

	fmt.Println(storycontent["intro"].Title)
	fmt.Println(storycontent["intro"].Story)
	fmt.Println(storycontent["intro"].Options)
}

func PrintStory(storycontent utils.FullStory)