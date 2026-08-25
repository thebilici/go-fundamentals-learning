# Go Backend Learning — Progress

**Son Güncelleme:** 2026-08-25  
**Güncel Aşama:** Aşama 16 — Context

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
Channels                    ✅
        ↓
Generics                    ✅
        ↓
Context                     ← GÜNCEL
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

### Aşama 13 — Concurrency ✅

- Sequential vs Concurrent
- Concurrency vs Parallelism
- Goroutine
- `go` keyword
- `sync.WaitGroup`
- `Add()`
- `Done()`
- `Wait()`
- Race Condition
- Race Detector
- `sync.Mutex`
- `Lock()` / `Unlock()`

### Aşama 14 — Channels ✅

- Channel
- Channel Send / Receive
- Unbuffered Channel
- Buffered Channel
- `len()` / `cap()`
- Goroutine'ler arası iletişim
- Channel ile sonuç döndürme
- `close()`
- Channel üzerinde `range`
- WaitGroup + Channel

### Aşama 15 — Generics ✅

- Generics
- Type Parameters
- Type Arguments
- Type Constraints
- `any`
- `comparable`
- Generic Functions
- Birden fazla Type Parameter
- Generic Structs
- Custom Constraints
- Generic Slice
- Generic Map
- Type Inference

### Generics Pratiği

Generic function ve veri yapıları oluşturuldu:

```text
GetFirst[T any]
        ↓
Generic Slice

Contains[T comparable]
        ↓
Generic karşılaştırma

Response[T any]
        ↓
Generic Struct

Number Constraint
        ↓
Generic Sum

GetValue[K comparable, V any]
        ↓
Generic Map
```

Özellikle:

```go
func GetValue[K comparable, V any](data map[K]V, key K) (V, bool)
```

üzerinden:

```text
K → Key tipi
V → Value tipi
```

mantığı uygulandı.

---

## Güncel Aşama

### Aşama 16 — Context 🔄

Sıradaki konular:

```text
Context nedir?
        ↓
context.Context
        ↓
context.Background()
        ↓
Cancellation
        ↓
context.WithCancel()
        ↓
Timeout
        ↓
context.WithTimeout()
        ↓
Deadline
        ↓
Context Propagation
        ↓
Request Lifecycle
```

Bu aşamada uzun süren işlemleri iptal etme, timeout belirleme ve ileride HTTP request lifecycle'ını yönetmek için kullanılan Go `context` yapısını öğreneceğiz.