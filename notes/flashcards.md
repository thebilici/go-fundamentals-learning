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

# Aşama 3 — Control Flow

### 1. Comparison Operator'lar ne üretir?

`true` veya `false` değerinde bir `bool` sonuç üretir.

---

### 2. `=` ile `==` farkı nedir?

```text
=  → Değer atar
== → İki değeri karşılaştırır
```

---

### 3. Temel Logical Operator'lar nelerdir?

```text
&& → AND
|| → OR
!  → NOT
```

---

### 4. `if` ne işe yarar?

Bir koşul `true` olduğunda belirli bir kod bloğunu çalıştırır.

---

### 5. `else` ne zaman çalışır?

`if` koşulu `false` olduğunda çalışır.

---

### 6. `else if` neden kullanılır?

Birden fazla koşulu sırayla kontrol etmek için kullanılır.

---

### 7. `switch` ne işe yarar?

Bir değeri birden fazla olasılıkla karşılaştırmayı kolaylaştırır.

---

### 8. Go'nun temel döngüsü nedir?

`for` döngüsüdür.

```go
for i := 1; i <= 5; i++ {
	fmt.Println(i)
}
```

---

### 9. Go'da `while` var mı?

Hayır. While benzeri kullanım `for` ile yapılır.

```go
for count <= 5 {
	count++
}
```

---

### 10. `break` ne yapar?

Döngüyü tamamen sonlandırır.

---

### 11. `continue` ne yapar?

Mevcut turu atlar ve sonraki tura geçer.

---

### 12. `break` ve `continue` farkı nedir?

```text
break    → Döngüyü bitirir
continue → Sadece mevcut turu atlar
```

---

### 13. `range` ne işe yarar?

Collection içerisindeki elemanları sırayla dolaşmayı sağlar.

---

### 14. `range` ile hangi değerleri alabiliriz?

Örneğin slice üzerinde:

```go
for index, value := range languages {
}
```

ile index ve value alınabilir.

---

### 15. `_` nedir?

Blank Identifier'dır. Kullanmak istemediğimiz değeri yok saymamızı sağlar.

```go
for _, value := range languages {
	fmt.Println(value)
}
```

# Aşama 4 — Functions

### 1. Function nedir?

Belirli bir işi yapan ve gerektiğinde tekrar çağrılabilen kod bloğudur.

---

### 2. Go'da function hangi keyword ile tanımlanır?

`func` keyword'ü ile.

```go
func greet() {
}
```

---

### 3. Parameter nedir?

Function tanımlanırken belirtilen değişkendir.

```go
func greet(name string) {
}
```

Burada `name` parameter'dır.

---

### 4. Argument nedir?

Function çağrılırken gönderilen gerçek değerdir.

```go
greet("Fatih")
```

Burada `"Fatih"` argument'tır.

---

### 5. Parameter ile argument arasındaki fark nedir?

```text
func greet(name string)
           ↑
       Parameter

greet("Fatih")
       ↑
    Argument
```

---

### 6. Return Value nedir?

Function'ın ürettiği sonucu çağrıldığı yere geri göndermesidir.

```go
func add(a, b int) int {
	return a + b
}
```

---

### 7. Return Type nerede yazılır?

Parameter listesinden sonra yazılır.

```go
func add(a, b int) int
                   ↑
              Return Type
```

---

### 8. Bir function birden fazla değer döndürebilir mi?

Evet.

```go
func getUser() (string, int) {
	return "Fatih", 22
}
```

---

### 9. Multiple Return değerleri nasıl alınır?

Sırayla variable'lara atanır.

```go
name, age := getUser()
```

```text
"Fatih" → name
22      → age
```

---

### 10. Birden fazla Return Type nasıl belirtilir?

Parantez içerisinde yazılır.

```go
func example() (string, int, bool) {
	return "Fatih", 22, true
}
```

---

### 11. `_` Multiple Return ile neden kullanılır?

İhtiyacımız olmayan değeri yok saymak için.

```go
name, _ := getUser()
```

---

### 12. Named Return nedir?

Return değerine function tanımında isim verilmesidir.

```go
func add(a, b int) (result int) {
	result = a + b
	return
}
```

---

### 13. Function Scope nedir?

Function içerisinde oluşturulan local variable'ların yalnızca kendi scope'larında erişilebilir olmasıdır.

---

### 14. `:=` ile oluşturulan şey variable mıdır?

Evet.

```go
name := "Fatih"
```

`name` normal bir variable'dır.

---

### 15. `var`, `:=` ve `const` arasındaki temel fark nedir?

```text
var   → Variable tanımlar
:=    → Function içinde kısa variable tanımlar
const → Constant tanımlar
```

---

### 16. Aynı type'taki parameter'lar kısa nasıl yazılır?

```go
func add(a, b int) int {
	return a + b
}
```

Hem `a` hem `b`, `int` type'ındadır.

---

### 17. Function sonucu variable'a alınmak zorunda mı?

Hayır.

```go
fmt.Println(add(5, 10))
```

şeklinde doğrudan kullanılabilir.

---

### 18. `return` çalıştığında ne olur?

Değer çağıran yere gönderilir ve function'ın çalışması sona erer.

---

### 19. Function'ın temel veri akışı nasıldır?

```text
Argument
   ↓
Parameter
   ↓
Function
   ↓
İşlem
   ↓
Return
   ↓
Sonuç
```

---

### 20. `getUser()` örneğinde aşağıdaki kod ne yapar?

```go
name, age := getUser()
```

`getUser()` tarafından döndürülen değerleri sırasıyla `name` ve `age` variable'larına atar.