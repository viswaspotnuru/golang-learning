package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()

	n, err := note.New(title, content)
	if err != nil {
		log.Println("Error creating note:", err)
		return
	}

	n.Display()
	err = n.Save()

	if err != nil {
		fmt.Printf("The file is not saved")
	}

	fmt.Printf("Successfully saved.")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
