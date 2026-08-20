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

# Aşama 2 — Variables & Data Types

### 1. Variable nedir?

Program içerisinde bir değeri tutmak için kullanılan isimlendirilmiş yapıdır.

---

### 2. `var` ne işe yarar?

Variable tanımlamak için kullanılır.

```go
var age int = 22
```

---

### 3. `:=` ne işe yarar?

Function içerisinde kısa şekilde yeni variable oluşturur ve type inference yapar.

```go
age := 22
```

---

### 4. `:=` ile `=` farkı nedir?

```text
:= → Yeni variable oluşturur.
=  → Mevcut variable'a değer atar.
```

---

### 5. Go'daki temel data type'lar nelerdir?

```text
string
int
float64
bool
```

---

### 6. Static Typing nedir?

Variable type'larının compile aşamasında belirli olmasıdır.

```go
age := 22
age = "22" // Hata
```

---

### 7. Type Inference nedir?

Compiler'ın verilen değerden type'ı otomatik belirlemesidir.

```go
age := 22 // int
```

---

### 8. Zero Value nedir?

Başlangıç değeri verilmeyen variable'ın otomatik aldığı değerdir.

```text
string  → ""
int     → 0
float64 → 0
bool    → false
```

---

### 9. `const` ne işe yarar?

Değeri sonradan değiştirilemeyen constant tanımlar.

```go
const language = "Go"
```

---

### 10. Type Conversion nedir?

Bir değeri başka bir type olarak kullanmak için yapılan açık dönüşümdür.

```go
ageFloat := float64(age)
```

---

### 11. `strconv.Atoi` ne yapar?

String içerisindeki decimal integer değerini `int` olarak parse eder.

```text
"25" → 25
```

---

### 12. `strconv.Itoa` ne yapar?

`int` değerini decimal string'e dönüştürür.

```text
25 → "25"
```

---

### 13. `byte` hangi type'ın alias'ıdır?

```text
byte → uint8
```

---

### 14. `rune` hangi type'ın alias'ıdır?

```text
rune → int32
```

---

### 15. `int` ve `int64` aynı type mıdır?

Hayır. İkisi farklı type'lardır ve gerektiğinde explicit conversion yapılmalıdır.

```go
var number int = 22
converted := int64(number)
```