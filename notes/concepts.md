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