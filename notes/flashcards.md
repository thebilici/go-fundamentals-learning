# Go Flashcards

## Aşama 1 — Go Ortamı ve Temeller

### 1. Go Module nedir?

Bir Go projesinin üst seviye kimliği ve dependency yönetim sınırıdır.

---

### 2. `go.mod` ne işe yarar?

Module path bilgisini, kullanılan Go sürümünü ve gerekli dependency bilgilerini tutar.

---

### 3. Package nedir?

Birbiriyle ilişkili Go kodlarını gruplamak için kullanılan yapıdır.

Bir package birden fazla `.go` dosyasından oluşabilir.

---

### 4. `package main` ne anlama gelir?

Çalıştırılabilir bir Go programının oluşturulmasında kullanılan özel package'dir.

---

### 5. `func main()` ne işe yarar?

Programın entry point'idir. Programın çalışması buradan başlar.

---

### 6. Standard Library nedir?

Go kurulumu ile birlikte gelen hazır package koleksiyonudur.

Örnek:

```text
fmt
net/http
os
time
context
```

---

### 7. `fmt.Println()` ne yapar?

`fmt` package'indeki `Println` fonksiyonunu kullanarak terminale çıktı yazar.

---

### 8. `go run .` ne yapar?

Mevcut package'i derler ve hemen çalıştırır.

```powershell
go run .
```

---

### 9. `go build .` ne yapar?

Mevcut package'i derleyerek executable binary oluşturur.

```powershell
go build .
```

---

### 10. `go run` ve `go build` arasındaki fark nedir?

```text
go run
→ Derle + Çalıştır

go build
→ Derle + Binary oluştur
```

---

### 11. Source Code ve Binary arasındaki fark nedir?

```text
main.go
→ İnsan tarafından okunabilen Source Code

main.exe
→ Compiler tarafından oluşturulan çalıştırılabilir Binary
```

---

### 12. Exported Identifier nedir?

Büyük harfle başlayan ve başka package'lerden erişilebilen identifier'dır.

```text
Hello   → Exported
hello   → Unexported
```

---

### 13. `go fmt` ne işe yarar?

Go kodunu standart Go formatına göre düzenler.

```powershell
go fmt ./...
```

---

### 14. Module ve Package aynı şey midir?

Hayır.

```text
Module
  ↓
Bir veya daha fazla Package
  ↓
.go Files
```

Module daha üst seviyedeki yapıdır.

---

### 15. Go neden kullanılmayan import'lara izin vermez?

Gereksiz kodu engellemek ve hataların compile aşamasında fark edilmesini sağlamak için kullanılmayan import'ları compiler error olarak değerlendirir.