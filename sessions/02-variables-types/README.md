# Aşama 2 — Variables & Data Types

Bu aşamada Go'da değişkenlerin nasıl oluşturulduğu, temel veri tipleri, type inference, zero value, constant ve type conversion kavramları öğrenildi.

## Variable Tanımlama

Go'da değişkenler `var` ile tanımlanabilir:

```go
var name string = "Fatih"
var age int = 22
```

Burada type açıkça belirtilmiştir.

Go type'ı değerden de çıkarabilir:

```go
var age = 22
```

Function içerisinde kısa tanımlama için `:=` kullanılabilir:

```go
age := 22
```

`:=` yeni variable oluştururken, `=` mevcut variable'a yeni değer atar:

```go
age := 22
age = 23
```

## Temel Data Type'lar

Bu aşamada kullanılan temel type'lar:

```text
string  → Metin
int     → Tam sayı
float64 → Ondalıklı sayı
bool    → true / false
```

Örnek:

```go
name := "Fatih"      // string
age := 22            // int
height := 1.73       // float64
isStudent := true    // bool
```

Go **statically typed** bir dildir. Bir variable'ın type'ı belirlendikten sonra başka bir type doğrudan atanamaz.

## Type Inference

Go bazı durumlarda type'ı bizim yerimize belirleyebilir:

```go
age := 22
```

Compiler `22` değerinden `age` değişkeninin `int` olduğunu çıkarır.

Type inference kullanılması Go'yu dynamically typed yapmaz. Type yine compile aşamasında bellidir.

## Zero Value

Başlangıç değeri verilmeyen variable'lar type'larının zero value değeriyle başlar:

```text
string  → ""
int     → 0
float64 → 0
bool    → false
```

Örneğin:

```go
var age int
```

başlangıçta `0` değerine sahiptir.

## Constant

Değişmemesi gereken değerler `const` ile tanımlanabilir:

```go
const language = "Go"
```

Constant oluşturulduktan sonra değeri değiştirilemez.

## Type Conversion

Numeric type'lar açık şekilde dönüştürülebilir:

```go
age := 22
ageFloat := float64(age)
```

Burada:

```text
22 int
  ↓
float64(age)
  ↓
22 float64
```

String ve integer dönüşümlerinde `strconv` package'i kullanıldı:

```go
number, err := strconv.Atoi("25")
```

```text
"25" → string
        ↓
   strconv.Atoi
        ↓
      25 → int
```

Tersi için:

```go
ageText := strconv.Itoa(age)
```

```text
22 → int
     ↓
strconv.Itoa
     ↓
"22" → string
```

## Numeric Type'lar

Go farklı boyutlarda numeric type'lar sağlar:

```text
int8    int16    int32    int64
uint8   uint16   uint32   uint64

float32
float64
```

Ayrıca:

```text
byte → uint8
rune → int32
```

alias'ları bulunur.

## Yapılan Pratik

`exercises/basics/variables-types/` altında küçük bir uygulama oluşturuldu.

Pratikte:

- Variable declaration
- `:=`
- `const`
- Type inference
- `string`, `int`, `float64`, `bool`
- `int → float64`
- `string → int`
- `int → string`
- `%T` ile type kontrolü

uygulandı.

## Aşama Sonu Zihinsel Model

```text
Değer
  ↓
Variable
  ↓
Data Type
  ↓
Compiler Type Kontrolü
  ↓
Gerekirse Explicit Conversion
```

Go'da type sistemi mümkün olduğunca açık ve compile-time kontrollü çalışır.