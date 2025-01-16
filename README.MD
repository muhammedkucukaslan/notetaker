# Note Taker

A simple command-line tool written in Go that helps you maintain daily notes in Markdown format. It automatically creates dated files and adds timestamped notes with checkbox formatting.

## Features

- Creates daily Markdown files (YYYY-MM-DD.md format)
- Adds notes with checkbox formatting (`- [ ] Your note`)
- Automatically creates directories if they don't exist
- Simple command-line interface
- Works with Obsidian or any other Markdown editor

## Installation

1. Clone the repository or download the code
2. Modify the `NOTE_PATH` constant in `main.go` to your preferred notes location:
```go
const (
    NOTE_PATH = "/your/preferred/path/to/notes"
)
```

3. Build the executable:
```bash
go build -o notetaker
```

4. (Optional) Move the executable to your system's PATH for global access:
```bash
sudo mv notetaker /usr/local/bin/
```

## Usage

Add a new note:
```bash
notetaker "Your note content here"
```

This will:
- Create a file with today's date (if it doesn't exist)
- Add your note with a checkbox format
- Place it in your specified notes directory

Example output file (`2024-01-16.md`):
```markdown
- [ ] Your note content here
```
 I built it for Obsidian. So, the file format is `.md`. You can change it to .txt or any other format by adjusting  `getTodayFilePath()`  in `taker.go`.

## Requirements

- Go 1.16 or higher

## Contributing

Feel free to open issues or submit pull requests if you have suggestions for improvements.

