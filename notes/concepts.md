# Go Kavramları

Bu dosya Go öğrenme sürecinde karşılaşılan önemli kavramların kısa açıklamalarını içerir.

## Go Module

Bir Go projesinin üst seviye kimliği ve dependency yönetim sınırıdır.

Module bilgileri `go.mod` dosyasında tutulur.

```text
Module
  ↓
Packages
  ↓
.go Files
```

## Package

Birbiriyle ilişkili Go kodlarını gruplamak için kullanılan yapıdır.

Bir package birden fazla `.go` dosyasından oluşabilir.

## `package main`

Çalıştırılabilir Go uygulamalarında kullanılan özel package'dir.

`func main()` ile birlikte programın başlangıç yapısını oluşturur.

## Entry Point

Programın çalışmaya başladığı noktadır.

Go uygulamalarında:

```go
func main() {
}
```

entry point olarak kullanılır.

## Standard Library

Go kurulumu ile birlikte gelen hazır package koleksiyonudur.

Örnekler:

- `fmt`
- `net/http`
- `os`
- `time`
- `context`

## Exported Identifier

Go'da büyük harfle başlayan identifier'lar başka package'lerden erişilebilir.

```text
Println → Exported
println → Unexported
```

## Compiler

Go source code'unu işletim sisteminin ve CPU'nun çalıştırabileceği machine code'a dönüştürür.

```text
Source Code
    ↓
Compiler
    ↓
Binary
```

## Binary / Executable

`go build` sonucunda oluşturulan çalıştırılabilir programdır.

Windows üzerinde genellikle `.exe` uzantısına sahiptir.

## `go run`

Go kodunu derler ve hemen çalıştırır.

```powershell
go run .
```

## `go build`

Go kodunu derleyerek executable binary oluşturur.

```powershell
go build .
```

## `go fmt`

Go kodunu standart Go formatına göre düzenler.

```powershell
go fmt ./...
```
# Aşama 2 — Variables & Data Types

## Variable

Program içerisinde bir değeri tutmak için kullanılan isimlendirilmiş yapıdır.

```go
var age int = 22
```

## Static Typing

Go statically typed bir dildir. Variable'ın type'ı compile aşamasında bellidir ve sonradan başka bir type'a dönüşmez.

```go
age := 22
age = "22" // Hata
```

## Type Inference

Compiler'ın verilen değere bakarak variable'ın type'ını belirlemesidir.

```go
age := 22       // int
height := 1.73  // float64
```

## Short Variable Declaration `:=`

Function içerisinde kısa şekilde yeni variable oluşturmak için kullanılır.

```go
age := 22
```

`:=` yeni variable oluştururken `=` mevcut variable'a değer atar.

## Zero Value

Başlangıç değeri verilmeyen variable'lara Go otomatik olarak zero value verir.

```text
string  → ""
int     → 0
float64 → 0
bool    → false
```

## Constant

Program içerisinde değeri değişmemesi gereken değerler `const` ile tanımlanır.

```go
const language = "Go"
```

## Type Conversion

Bir değerin başka bir type olarak kullanılmasını sağlar.

```go
age := 22
ageFloat := float64(age)
```

## Parsing

Text olarak bulunan bir değerin yorumlanarak başka bir type'a çevrilmesidir.

```go
number, err := strconv.Atoi("25")
```

Burada `"25"` string değeri `25` integer değerine dönüştürülür.

## `strconv.Atoi`

String → int dönüşümü için kullanılır.

```go
number, err := strconv.Atoi("25")
```

## `strconv.Itoa`

Int → string dönüşümü için kullanılır.

```go
text := strconv.Itoa(25)
```

## `byte` ve `rune`

Go'daki iki önemli alias:

```text
byte → uint8
rune → int32
```

`byte` binary/veri işlemlerinde, `rune` ise Unicode karakterlerle çalışırken sık kullanılır.
# Aşama 3 — Control Flow

## Comparison Operators

İki değeri karşılaştırır ve `bool` sonuç üretir.

```text
==   Eşit
!=   Eşit değil
>    Büyük
<    Küçük
>=   Büyük veya eşit
<=   Küçük veya eşit
```

## Logical Operators

Birden fazla koşulu birleştirmek için kullanılır.

```text
&& → AND
|| → OR
!  → NOT
```

## if / else

Bir koşula göre hangi kod bloğunun çalışacağını belirler.

```go
if age >= 18 {
	fmt.Println("Reşit")
} else {
	fmt.Println("Reşit değil")
}
```

Birden fazla koşul için `else if` kullanılabilir.

## switch

Bir değerin farklı olasılıklarını kontrol etmek için kullanılır.

```go
switch role {
case "admin":
	fmt.Println("Admin")
case "user":
	fmt.Println("User")
default:
	fmt.Println("Bilinmeyen")
}
```

## for

Go'daki temel döngü yapısıdır.

```go
for i := 1; i <= 5; i++ {
	fmt.Println(i)
}
```

Go'da ayrı bir `while` keyword'ü bulunmaz. `for` farklı döngü biçimleri için kullanılır.

## break

İçinde bulunduğu döngüyü sonlandırır.

```go
if i == 8 {
	break
}
```

## continue

Döngünün mevcut turunun geri kalanını atlayarak sonraki tura geçer.

```go
if i == 4 {
	continue
}
```

## range

Slice gibi collection'ların elemanlarını dolaşmak için kullanılır.

```go
for index, value := range languages {
	fmt.Println(index, value)
}
```

## Blank Identifier `_`

Bir fonksiyondan, `range`'den vb. gelen ancak kullanılmayacak bir değeri yok saymak için kullanılır.

```go
for _, value := range languages {
	fmt.Println(value)
}
```

# Aşama 4 — Functions

## Function

Belirli bir işi yapan ve tekrar kullanılabilen kod bloğudur.

```go
func greet() {
	fmt.Println("Hello")
}
```

## Parameter ve Argument

Parameter, function tanımlanırken belirtilen değişkendir.

Argument ise function çağrılırken gönderilen gerçek değerdir.

```go
func greet(name string) {
	fmt.Println(name)
}

greet("Fatih")
```

```text
name    → Parameter
"Fatih" → Argument
```

## Return Value

Function yaptığı işlemin sonucunu geri döndürebilir.

```go
func add(a, b int) int {
	return a + b
}
```

Buradaki son `int`, function'ın return type'ıdır.

## Multiple Return Values

Go function'ları birden fazla değer döndürebilir.

```go
func getUser() (string, int) {
	return "Fatih", 22
}
```

Dönen değerler sırayla variable'lara atanır:

```go
name, age := getUser()
```

```text
"Fatih" → name
22      → age
```

## Named Return Values

Return değerlerine function tanımında isim verilebilir.

```go
func add(a, b int) (result int) {
	result = a + b
	return
}
```

## Function Scope

Function içerisinde oluşturulan variable'lar kendi scope'ları içerisinde kullanılabilir.

```go
func calculate() int {
	result := 10 + 20
	return result
}
```

Başka bir function `result` değişkenine doğrudan erişemez.

## Temel Function Akışı

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
# Aşama 5 — Arrays & Slices

## Array

Aynı type'taki birden fazla değeri sabit uzunlukta tutan veri yapısıdır.

```go
languages := [3]string{"Go", "Python", "Java"}
```

`[3]string`, 3 elemanlı bir `string` array'idir.

## Index

Array ve Slice elemanlarına konumları üzerinden erişilir.

```go
languages[0]
```

Go'da index `0`'dan başlar.

## Slice

Boyutu sabit olmayan, Go'da collection işlemlerinde sık kullanılan yapıdır.

```go
languages := []string{"Go", "Python", "Java"}
```

Temel syntax farkı:

```text
[3]string → Array
[]string  → Slice
```

## append

Slice'a eleman eklemek için kullanılır.

```go
languages = append(languages, "Rust")
```

`append`, güncellenmiş slice'ı geri döndürdüğü için sonuç tekrar variable'a atanır.

## len

Mevcut eleman sayısını verir.

```go
len(languages)
```

## cap

Slice'ın mevcut backing array üzerinde erişebildiği kapasiteyi verir.

```text
len → Mevcut eleman sayısı
cap → Mevcut kapasite
```

## Slicing

Bir slice'ın belirli bir bölümünü almak için kullanılır.

```go
selected := languages[1:4]
```

Kural:

```text
[start:end]

start → dahil
end   → dahil değil
```

## range

Slice elemanlarını sırayla dolaşmayı sağlar.

```go
for index, value := range languages {
	fmt.Println(index, value)
}
```

## Array ve Slice Farkı

```text
Array
→ Sabit uzunluk
→ [3]string

Slice
→ Esnek uzunluk
→ []string
→ append ile büyüyebilir
```
# Aşama 6 — Maps

## Map

Key-value mantığıyla veri tutan collection yapısıdır.

```go
scores := map[string]int{
	"Fatih": 90,
	"Ahmet": 75,
}
```

```text
map[string]int
     ↓      ↓
    key    value
```

## Değer Okuma

Bir value'ya key üzerinden erişilir.

```go
score := scores["Fatih"]
```

## Eleman Ekleme ve Güncelleme

Aynı syntax kullanılır:

```go
scores["Ayşe"] = 85
```

Key yoksa yeni eleman eklenir, varsa mevcut value güncellenir.

## Key Kontrolü

```go
score, ok := scores["Fatih"]
```

```text
score → Value
ok    → Key mevcut mu?
```

Key varsa `ok = true`, yoksa `ok = false` olur.

Bu kullanım **comma ok idiom** olarak bilinir.

## Zero Value

Olmayan bir key doğrudan okunursa value type'ın zero value'su döner.

Örneğin `map[string]int` için:

```go
scores["Unknown"]
```

key yoksa `0` döner.

Bu nedenle key'in gerçekten var olup olmadığını anlamak için `value, ok` kullanılabilir.

## delete

Map'ten key-value çiftini siler.

```go
delete(scores, "Fatih")
```

## len

Map'teki eleman sayısını verir.

```go
len(scores)
```

## range

Map'teki key-value çiftlerini dolaşır.

```go
for name, score := range scores {
	fmt.Println(name, score)
}
```

Map üzerinde `range` kullanırken belirli bir sıralamaya güvenilmemelidir.

## Temel Map Modeli

```text
Map
 ↓
Key → Value

scores["Fatih"]
       ↓
       90
```