package storyserver

import (
	"fmt"
	"os"

	"github.com/cyoa/internal/utils.go"
)

func StartStory(storycontent utils.FullStory, arc string) {
	if _, ok := storycontent[arc]; ok {
		PrintTitle(storycontent, arc)
		PrintStory(storycontent, arc)
		PrintOptions(storycontent, arc)
		fmt.Println()
	} else {
		fmt.Println("Entered input is not part of current story line ")
	}

}

func PrintStory(storycontent utils.FullStory, arc string) {
	fmt.Println(storycontent[arc].Story)

}
func PrintOptions(storycontent utils.FullStory, arc string) {
	if len(storycontent[arc].Options) <= 0 {
		os.Exit(0)
	}
	// fmt.Println(storycontent[arc].Options)
	for _, v := range storycontent[arc].Options {
		fmt.Printf("Choose option %s for %s \n", v.Arc, v.Text)
	}
	// fmt.Println(storycontent[arc].Options)
}

func PrintTitle(storycontent utils.FullStory, arc string) {
	fmt.Println(storycontent[arc].Title)
}

func TakeUserInput() string {
	var input string
	fmt.Scanln(&input)
	return input
}

func ShowStoryAccordingtoUserInput(storycontent utils.FullStory) {
	input := TakeUserInput()
	StartStory(storycontent, input)
	// StartStory(storycontent, TakeUserInput())
}
