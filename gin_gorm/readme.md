# API Projesi  

Bu proje, **Gin** ve **GORM** kullanılarak geliştirilmiştir.  

## Kurulum  

1. Depoyu klonlayın:  
   git clone <repo-link>  
   cd <proje-dizini>  

2. Bağımlılıkları yükleyin:  
   go mod tidy  

3. Veritabanı migrasyonlarını çalıştırın:  
   go run main.go migrate  

4. Sunucuyu başlatın:  
   go run main.go  

## Kullanılan Teknolojiler  
- **Gin** - [https://gin-gonic.com/](https://gin-gonic.com/)  
- **GORM** - [https://gorm.io/](https://gorm.io/)  
- **godotenv** - [https://github.com/joho/godotenv](https://github.com/joho/godotenv)  
- **sqlite3** - [https://pkg.go.dev/github.com/mattn/go-sqlite3](https://pkg.go.dev/github.com/mattn/go-sqlite3)  
- **mysql** - [MySQL](go get gorm.io/driver/mysql
) 


Bu dosya, tüm kurulum ve çalışma adımlarını tek bir yerden sunuyor. `migrate/migrate.go` dosyasını çalıştırarak veritabanı migrasyonlarını başlatabilir ve ardından sunucuyu çalıştırabilirsiniz.

