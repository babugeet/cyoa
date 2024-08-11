package main

import (
	"github.com/cyoa/internal/storyserver"
	"github.com/cyoa/internal/utils.go"
)

func main() {

	storycontent := utils.ParseJson("/workspaces/cyoa/story.json")
	storyserver.StartStory(storycontent)

}
