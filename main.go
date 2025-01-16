package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	NOTE_PATH = "/your/preferred/path/to/notes"
)

func main() {
	noteContent := strings.Join(os.Args[1:], " ")

	noteTaker := NewNoteTaker(NOTE_PATH)

	if err := noteTaker.WriteNote(noteContent); err != nil {
		log.Fatalf("Error writing note: %v", err)
	}

	fmt.Println("Note successfully added!")
}
