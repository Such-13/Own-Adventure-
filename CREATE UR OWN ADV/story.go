package main

import (
	"encoding/json"
	"os"
)

// Option represents a choice in the story
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

// StoryArc represents a single part of the story
type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

// Story is the full adventure, where each key is a story arc
type Story map[string]StoryArc

// LoadStory reads the JSON file and decodes it into a Story struct
func LoadStory(filename string) (Story, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var story Story
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&story)
	if err != nil {
		return nil, err
	}

	return story, nil
}
