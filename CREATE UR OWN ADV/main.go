package main

import (
	"html/template"
	"log"
	"net/http"
)

// StoryHandler is an HTTP handler for rendering the story
type StoryHandler struct {
	Story Story
}

// ServeHTTP handles HTTP requests and serves the story page
func (h StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	arc := r.URL.Path[1:] // Get the arc name from the URL (e.g., /new-york)
	if arc == "" {
		arc = "intro" // Default to "intro" if no arc is specified
	}

	// Find the requested story arc
	storyArc, ok := h.Story[arc]
	if !ok {
		http.Error(w, "Story not found", http.StatusNotFound)
		return
	}

	// Load the HTML template and render it
	tmpl := template.Must(template.ParseFiles("templates/story.html"))
	err := tmpl.Execute(w, storyArc)
	if err != nil {
		log.Printf("Error rendering template: %v", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
}

func main() {
	// Load the story from JSON
	story, err := LoadStory("gopher.json")
	if err != nil {
		log.Fatalf("Failed to load story: %v", err)
	}

	// Start the web server
	handler := StoryHandler{Story: story}
	log.Println("Starting the server on :8080...")
	http.ListenAndServe(":8080", handler)
}
