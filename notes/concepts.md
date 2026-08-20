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