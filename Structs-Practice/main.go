package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	text := getTodoText()
	userNote, err := note.New(title, content)
	if err != nil {
		log.Println("Error creating note:", err)
		return
	}
	err = outputData(userNote)
	if err != nil {
		fmt.Printf("note is not saved")
	}
	todoText, err := todo.New(text)
	if err != nil {
		log.Println("Error creating todo text:", err)
		return
	}
	err = outputData(todoText)

	if err != nil {
		fmt.Printf("note is not saved")
	}

}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving process failed")
		return err
	}
	fmt.Println("Successfully saved.")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func getTodoText() string {
	text := getUserInput("Todo text: ")

	return text
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
