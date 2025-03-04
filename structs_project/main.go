package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"structs_project/note"
)

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	// bufio ile içinde boşluk içeren metinlerde rahatça alınabilir
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	title := getUserInput("Note Title: ")
	content := getUserInput("Note Content: ")

	note, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	note.Display()

	err = note.Save()
	if err != nil {
		fmt.Println("\n Note couldnt saved!")
		return
	}
	// fmt.Println("\nNote Saved Successfully!")
}


/*
Note Title: AI is crazy 
Note Content: AI is going to take our places
Note saved to AI_is_crazy.json
*/