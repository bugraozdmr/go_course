package todo

import (
	"encoding/json"
	"os"
	"time"
)

// Todo struct'ı
type Todo struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

// TodoService arayüzü
type TodoService interface {
	Add(title, content string)
	SaveToFile(filename string) error
}

// Service, TodoService'in uygulanmış hali
type Service struct {
	Todos []Todo
}

// Yeni TODO ekler
func (s *Service) Add(title, content string) {
	newTodo := Todo{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}
	s.Todos = append(s.Todos, newTodo)
}

// TODO listesini JSON dosyasına kaydeder
func (s *Service) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Okunaklı JSON
	return encoder.Encode(s.Todos)
}

// bu şekilde eğer düzgün tanımlanmazsa metodlar kızacaktır
func NewService() TodoService {
	return &Service{}
}