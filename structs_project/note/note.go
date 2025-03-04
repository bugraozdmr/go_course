package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// eğer public erişimi olmazsa save ederken boş dosya çıkar
// struct Tags -> bu şekild efarklı bir şekilde kaydedilebilir
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// pointer gerek yok zaten basit bir yapıya sahip
func validateNote(title, content string) error {
	if title == "" {
		return errors.New("title cannot be empty")
	}
	if content == "" {
		return errors.New("content cannot be empty")
	}
	return nil
}

func New(title, content string) (Note, error) {
	if err := validateNote(title, content); err != nil {
		return Note{}, err
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Your note titled '%v' and the content is '%v'\n", note.Title, note.Content)
}

func (note Note) Save() error {
	// Dosya adını formatla
	fileName := note.formatFileName()

	// JSON'a dönüştür
	noteData, err := json.MarshalIndent(note, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling note: %v", err)
	}

	// Dosyaya yaz
	err = os.WriteFile(fileName, noteData, 0644)
	if err != nil {
		return fmt.Errorf("error saving note to file: %v", err)
	}

	fmt.Printf("Note saved to %s\n", fileName)
	return nil
}

// formatFileName format the title to create a valid and readable filename
func (note Note) formatFileName() string {
	sanitizedTitle := strings.ReplaceAll(note.Title, " ", "_")

	sanitizedTitle = strings.Map(func(r rune) rune {
		if strings.ContainsRune(`\/:*?"<>|`, r) {
			return -1
		}
		return r
	}, sanitizedTitle)

	if len(sanitizedTitle) > 30 {
		sanitizedTitle = sanitizedTitle[:30]
	}

	return fmt.Sprintf("%s.json", sanitizedTitle)
}
