package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"tilte"` // Struct tags

}

func (todo Todo) Display() {
	fmt.Printf("Your todo text is %v\n", todo.Text)
}

func (todo Todo) Save() error {

	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile("todo.json", json, 0644)
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("invalid input: Text cannot be empty")
	}
	return Todo{
		Text: text,
	}, nil
}
