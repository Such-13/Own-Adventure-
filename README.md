# Own-Adventure-
A Go project for learning -Go functions vs. methods, HTTP request handling, ServeHTTP, handler registration, and pointers in methods. 

# Choose Your Own Adventure - Go Web App

This is a web-based "Choose Your Own Adventure" game built using Go. The user navigates through different story arcs based on options presented at the end of each story page. The story is read from a JSON file and displayed in a web page using HTML templates.

## Features
- Dynamic story navigation via URL paths (e.g., `/new-york`, `/denver`).
- JSON file-based story content.
- Simple, reusable HTML template to render story pages.

## Requirements
- Go 1.18 or higher
- A browser to view the app

# Choose Your Own Adventure - Go Web App

## Installation

1. Clone or download the project to your local machine:
   git clone [https://github.com/your-username/choose-your-own-adventure.git](https://github.com/Such-13/Own-Adventure-.git)
   cd choose-your-own-adventure

2. Run `go mod tidy` to install necessary dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application directly:
   ```bash
   go run .
   ```
   This will start a web server on `http://localhost:8080`.


## File Structure

```
choose-your-own-adventure/
├── go.mod               # Go modules file
├── gopher.json          # JSON file containing story content
├── main.go              # Main Go file to run the server
├── story.go             # Go file to handle story loading and parsing JSON
├── templates/           # Folder containing HTML templates
│   └── story.html       # HTML template for displaying the story
└── README.md            # This README file
```
## Running the App
Once the app is running, open your browser and visit:
http://localhost:8080


## Next Steps
- Add More Arcs: You can extend the gopher.json file with more arcs and choices to create a longer adventure.
- Styling: Add your own CSS in story.html to improve the appearance of the game.
- Expand Functionality: Add features like saving the user's progress or adding more interactive elements.

## License
 This project is open-source.

## Contributing
 Feel free to fork this project, create pull requests, and share improvements! Contributions are welcome.
