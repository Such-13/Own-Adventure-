package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// Link structure to hold extracted link data
type Link struct {
	Href string
	Text string
}

func main() {
	// Check if the user provided a URL
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run link_parser.go <URL>")
		return
	}

	url := os.Args[1] // Get URL from command-line arguments

	// Fetch the HTML from the URL
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()

	// Parse the HTML
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}

	// Extract links from the HTML
	links := extractLinks(doc)

	// Print the extracted links
	for _, link := range links {
		fmt.Printf("Href: %s\nText: %s\n\n", link.Href, link.Text)
	}
}

// extractLinks traverses the HTML and extracts all <a> tags with href attributes.
func extractLinks(n *html.Node) []Link {
	var links []Link
	var crawler func(*html.Node)

	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			// Extract href attribute
			var href string
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					href = attr.Val
					break
				}
			}

			// Extract text content from the <a> tag
			text := extractText(node)

			// Append the link to the list
			links = append(links, Link{Href: href, Text: text})
		}

		// Recursively traverse child nodes
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			crawler(c)
		}
	}

	crawler(n)
	return links
}

// extractText extracts text content from a node, removing extra whitespace and nested tags.mport (

// extractText extracts visible text inside an <a> tag, removing extra spaces.
func extractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return strings.TrimSpace(n.Data)
	}

	var texts []string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text := extractText(c)
		if text != "" {
			texts = append(texts, text)
		}
	}

	return strings.Join(texts, " ") // Join text parts with a single space
}
