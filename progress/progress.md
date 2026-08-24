# Go Backend Learning — Progress

**Son Güncelleme:** 2026-08-24  
**Güncel Aşama:** Aşama 13 — Concurrency

## Güncel Konum

```text
Go Ortamı ve Temeller       ✅
        ↓
Variables & Data Types      ✅
        ↓
Control Flow                ✅
        ↓
Functions                   ✅
        ↓
Arrays & Slices             ✅
        ↓
Maps                        ✅
        ↓
Structs                     ✅
        ↓
Methods                     ✅
        ↓
Pointers                    ✅
        ↓
Interfaces                  ✅
        ↓
Error Handling              ✅
        ↓
Packages & Modules          ✅
        ↓
Concurrency                 ← GÜNCEL
        ↓
HTTP / REST API
        ↓
Database
        ↓
Testing
        ↓
Docker
        ↓
Kubernetes
```

## Tamamlanan Son Aşamalar

### Aşama 10 — Interfaces ✅

- Interface
- Method Signature
- Implicit Implementation
- Interface Parameter
- Interface Variable
- Concrete Type
- Polymorphism

### Aşama 11 — Error Handling ✅

- `error`
- `nil`
- `err != nil`
- `errors.New()`
- `fmt.Errorf()`
- `(value, error)`
- Error Propagation
- Error Wrapping
- `%w`
- Struct Zero Value

### Aşama 12 — Packages & Modules ✅

- Package
- `package main`
- `func main()` ve entry point
- Custom Package
- Aynı package içerisinde birden fazla `.go` dosyası
- Exported / Unexported Identifiers
- Module
- Module Path
- Package Path
- Import Path
- Import Alias
- `go.mod`
- `go mod init`
- `go mod tidy`
- `go.sum`

### Yapılan Pratik

`exercises/basics/packages/` altında:

- `user` package oluşturuldu.
- `User` struct oluşturuldu.
- `CreateUser()` ve `GetName()` yazıldı.
- `validation.go` ile aynı package farklı dosyalara bölündü.
- `IsAdult()` eklendi.
- Package başka bir `main` package'dan import edildi.

`exercises/basics/packages-final/` altında:

- `mathutil` package oluşturuldu.
- `math.go` ve `validation.go` aynı package altında kullanıldı.
- `Add()` oluşturuldu.
- `Multiply()` oluşturuldu.
- `IsPositive()` oluşturuldu.
- Import alias (`mathpkg`) kullanıldı.

## Güncel Aşama

### Aşama 13 — Concurrency 🔄

Sıradaki konular:

- Concurrency nedir?
- Sequential vs Concurrent çalışma
- Goroutine
- `go` keyword
- `time.Sleep()` ile ilk goroutine gözlemi
- `sync.WaitGroup`
- Channel
- Channel send / receive
- Buffered / Unbuffered Channel
- Goroutine'ler arası iletişim
- Temel concurrency problemleri