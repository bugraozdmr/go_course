# Vergili Fiyat Hesaplama (Go)

## ✨ Proje Amacı
Bu Go projesi, belirli fiyatlar üzerinde farklı vergi oranları uygulayarak yeni fiyatları hesaplar ve sonucu JSON formatında kaydeder.

---

## 📂 Proje Modülleri
Kod farklı **paketlere** ayrılmıştır:

1. **`iomanager`** - Giriş/Çıkış arayüzünü tanımlar.
2. **`filemanager`** - Dosyadan veri okur/yazar.
3. **`cmdmanager`** - Konsoldan veri okur/yazar.
4. **`conversion`** - String'ten float'a dönüşüm yapar.
5. **`prices`** - Vergili fiyatları hesaplar.
6. **`main.go`** - Programı çalıştırır.

---

## 👀 `main.go` - Ana Program
**Ana program**, `prices.txt` dosyasından fiyatları okur ve farklı vergi oranları için yeni fiyatlar hesaplar.

### 📝 Örnek Kod:
```go
fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
err := priceJob.Process()
```

### 📃 Örnek `prices.txt` Dosyası:
```
100
200
300
```

### 📝 Örnek JSON Çıktısı (`result_10.json`):
```json
{
    "tax_rate": 0.1,
    "input_prices": [100, 200, 300],
    "tax_included_prices": {
        "100.00": "110.00",
        "200.00": "220.00",
        "300.00": "330.00"
    }
}
```

---

## 🔧 `iomanager` - Giriş/Çıkış Yönetimi
Dosya veya konsoldan veri almak için ortak bir **arayüz** tanımlar:
```go
type IOManager interface {
    ReadLines() ([]string, error)  // Satır satır okuma
    WriteResult(data interface{}) error // Sonucu yazma
}
```

---

## 📂 `prices` - Vergili Fiyat Hesaplama
### 📝 `LoadData()` - Fiyatları Yükler
```go
func (job *taxIncludedPriceJob) LoadData() error {
    lines, err := job.IOManager.ReadLines()
    prices, err := conversion.StringsToFloat(lines)
    job.InputPrices = prices
    return nil
}
```

### 📝 `Process()` - Vergili Fiyat Hesaplama ve Kaydetme
```go
func (job taxIncludedPriceJob) Process() error {
    err := job.LoadData()
    for _, price := range job.InputPrices {
        taxIncludedPrice := price * (1 + job.TaxRate)
        job.TaxIncludedPrices[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
    }
    return job.IOManager.WriteResult(job)
}
```

---

## 🔧 `cmdmanager` - Konsoldan Veri Okuma
```go
func (cmd CMDManager) ReadLines() ([]string, error) {
    fmt.Println("Please enter your prices. Confirm every price with ENTER")
    var prices []string
    for {
        var price string
        fmt.Print("Price: ")
        fmt.Scan(&price)
        if price == "0" {
            break
        }
        prices = append(prices, price)
    }
    return prices, nil
}
```

### 📝 Örnek Kullanıcı Girişi:
```
Price: 100
Price: 200
Price: 0
```

### 📝 Konsol Çıktısı:
```json
{"tax_rate":0.1,"input_prices":[100,200],"tax_included_prices":{"100.00":"110.00","200.00":"220.00"}}
```

---

## 📂 `filemanager` - Dosya Okuma/Yazma
### 📝 `ReadLines()` - Dosyadan Okuma
```go
func (fm FileManager) ReadLines() ([]string, error) {
    file, err := os.Open(fm.InputFilePath)
    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    file.Close()
    return lines, nil
}
```

### 📝 `WriteResult()` - JSON Dosyasına Kaydetme
```go
func (fm FileManager) WriteResult(data interface{}) error {
    file, err := os.Create(fm.OutputFilePath)
    encoder := json.NewEncoder(file)
    err = encoder.Encode(data)
    file.Close()
    return nil
}
```

---

## ✨ Özet
✅ **Dosya veya konsoldan fiyat okuma**
✅ **Vergi oranını uygulayarak yeni fiyatlar hesaplama**
✅ **Sonucu JSON formatında dosyaya veya ekrana yazma**
✅ **Modüler yapı sayesinde esnek ve geliştirilebilir bir sistem**

