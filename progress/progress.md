# Go Backend Learning — Progress

**Son Güncelleme:** 2026-08-25  
**Güncel Aşama:** Aşama 14 — Generics

---

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
Concurrency                 ✅
        ↓
Generics                    ← GÜNCEL
        ↓
Context
        ↓
HTTP
        ↓
REST API
        ↓
Backend Mimarisi
        ↓
Database
        ↓
Testing
        ↓
Configuration
        ↓
Logging
        ↓
Docker
        ↓
Docker Compose
        ↓
Kubernetes
        ↓
Observability
        ↓
Load Testing
```

---

## Tamamlanan Son Aşamalar

### Aşama 11 — Error Handling ✅

- `error`
- `nil`
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
- `func main()` ve Entry Point
- Custom Package
- Exported / Unexported
- Module
- Module Path
- Package Path
- Import Path
- Import Alias
- `go.mod`
- `go mod init`
- `go mod tidy`
- `go.sum`

### Aşama 13 — Concurrency ✅

- Sequential vs Concurrent
- Concurrency vs Parallelism
- Goroutine
- `go` keyword
- Main Goroutine
- `sync.WaitGroup`
- `Add()`
- `Done()`
- `Wait()`
- Channel
- Channel Send / Receive
- Unbuffered Channel
- Buffered Channel
- `len()` / `cap()`
- Channel ile sonuç döndürme
- `close()`
- Channel üzerinde `range`
- WaitGroup + Channel
- Deadlock
- Race Condition
- Race Detector
- `sync.Mutex`
- `Lock()` / `Unlock()`

### Concurrency Pratiği

`exercises/basics/concurrency/` altında concurrent square calculator geliştirildi.

```text
3 Goroutine
     ↓
CalculateSquare()
     ↓
results Channel
     ↓
WaitGroup
     ↓
close(results)
     ↓
range
```

Uygulanan kavramlar:

- Birden fazla goroutine
- WaitGroup
- Buffered Channel
- Channel üzerinden sonuç gönderme
- `close()`
- `range`

---

## Güncel Aşama

### Aşama 14 — Generics 🔄

Sıradaki konular:

```text
Generics nedir?
        ↓
Type Parameters
        ↓
Type Arguments
        ↓
Type Constraints
        ↓
any
        ↓
comparable
        ↓
Generic Functions
        ↓
Generic Types
        ↓
Generic Structs
```

Bu aşamada Go'da farklı veri tipleriyle çalışabilen, tekrar kullanılabilir ve type-safe kod yazmayı öğreneceğiz.