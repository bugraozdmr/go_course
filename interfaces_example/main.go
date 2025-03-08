package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todo-app/todo"
)

/*
interface ile bir metodta herhangi tür alınabilir diyebiliriz ve aynı zamanda metod içinde
switch kullanarak şu gelirse şu diye ayırabiliriz (type)
value.(int) => gibi
*/


func main() {
	service := todo.NewService() // Todo servisini oluştur (interface üzerinden)
	reader := bufio.NewReader(os.Stdin)

	for {
		// Kullanıcıdan başlık al
		fmt.Print("TODO Title: ")
		title, _ := reader.ReadString('\n')
		title = strings.TrimSpace(title) // Boşlukları temizle

		// Kullanıcıdan içerik al
		fmt.Print("TODO Content: ")
		content, _ := reader.ReadString('\n')
		content = strings.TrimSpace(content) // Boşlukları temizle

		// Yeni TODO ekle
		service.Add(title, content)

		// JSON dosyasına kaydet
		if err := service.SaveToFile("out.json"); err != nil {
			fmt.Println("Hata:", err)
			return
		}

		fmt.Println("✅ TODO eklendi ve out.json dosyasına kaydedildi!")

		// Kullanıcıya devam etmek isteyip istemediğini sor
		fmt.Print("Yeni bir TODO eklemek istiyor musun? (e/h): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if choice != "e" {
			fmt.Println("👋 Programdan çıkılıyor...")
			break
		}
	}
}
