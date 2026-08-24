# progress/roadmap.md
# ============================================

# Go Backend Learning Roadmap

**Güncel Aşama:** Aşama 1 — Go Ortamı ve Temeller

---

## Genel Yol Haritası

```text
Go Ortamı ve Temeller
        ↓
Variables & Data Types
        ↓
Control Flow
        ↓
Functions
        ↓
Arrays & Slices
        ↓
Maps
        ↓
Structs
        ↓
Methods
        ↓
Pointers
        ↓
Interfaces
        ↓
Error Handling
        ↓
Packages & Modules
        ↓
Generics
        ↓
Goroutines
        ↓
Channels
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

## Aşama 1 — Go Ortamı ve Temeller

Öğrenilecek konular:

- Go kurulumu
- `go version`
- Go source file yapısı
- `package main`
- `import`
- `func main()`
- `fmt`
- `go run`
- `go build`
- Temel Go program yapısı

**Durum:** Güncel Aşama

---

## Aşama 2 — Variables & Data Types

Öğrenilecek konular:

- Değişkenler
- `var`
- `:=`
- Constants
- `string`
- `int`
- `float`
- `bool`
- Type inference
- Zero values
- Type conversion

**Durum:** Bekliyor

---

## Aşama 3 — Control Flow

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- Comparison Operators
- Logical Operators
- `if / else / else if`
- `switch`
- `for`
- `break`
- `continue`
- `range`

---

## Aşama 4 — Functions

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- Function tanımlama ve çağırma
- Parameters ve Arguments
- Return Values
- Multiple Return Values
- Named Returns
- Function Scope

---

## Aşama 5 — Arrays & Slices

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- Arrays
- Index
- Slices
- `len`
- `cap`
- `append`
- Slicing
- `range`

---

## Aşama 6 — Maps

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- Map
- Key / Value
- Eleman okuma
- Eleman ekleme
- Eleman güncelleme
- `value, ok`
- Zero Value
- `delete`
- `len`
- `range`

---

## Aşama 7 — Structs

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- Struct
- Field
- Struct oluşturma
- Field erişimi ve güncelleme
- Struct + Function
- Slice of Structs
- `range` ile Struct dolaşma
- Exported / Unexported Fields

---

## Aşama 8 — Methods

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- Method
- Function ve Method farkı
- Receiver
- Value Receiver
- Method Parameters
- Method Return Values
- Struct + Method

---

## Aşama 9 — Pointers

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- Pointer
- Memory Address
- `&` Operator
- `*` Operator
- Dereferencing
- Pointer Types
- Pointer + Function
- Pointer Receiver
- Value Receiver vs Pointer Receiver

---

## Aşama 10 — Interfaces

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- Interface
- Method Signature
- Implicit Implementation
- Interface Parameter
- Interface Variable
- Concrete Type
- Polymorphism
- Birden fazla method içeren interface

---

## Aşama 11 — Error Handling

**Durum:** 🔄 Güncel Aşama

Öğrenilecek konular:

- `error` nedir?
- `(value, error)` yapısı
- `nil`
- `if err != nil`
- `errors.New()`
- `fmt.Errorf()`
- Error Propagation
- Hataları üst katmana iletme
---

## Aşama 12 — Packages & Modules

Öğrenilecek konular:

- Packages
- Exported identifiers
- Imports
- `go.mod`
- `go mod init`
- Go Modules
- Package organizasyonu

**Durum:** Bekliyor

---

## Aşama 13 — Generics

Öğrenilecek konular:

- Type parameters
- Constraints
- Generic functions
- Generic data structures

**Durum:** Bekliyor

---

## Aşama 14 — Goroutines

Öğrenilecek konular:

- Concurrency kavramı
- Goroutines
- `go` keyword
- Go scheduler temel mantığı
- Concurrent execution

**Durum:** Bekliyor

---

## Aşama 15 — Channels

Öğrenilecek konular:

- Channels
- Send / receive
- Unbuffered channels
- Buffered channels
- Channel closing
- `select`

**Durum:** Bekliyor

---

## Aşama 16 — Context

Öğrenilecek konular:

- `context.Context`
- Cancellation
- Timeout
- Deadline
- Request lifecycle

**Durum:** Bekliyor

---

## Aşama 17 — HTTP

Öğrenilecek konular:

- `net/http`
- HTTP server
- Handler
- Request
- Response
- Routing
- HTTP lifecycle

**Durum:** Bekliyor

---

## Aşama 18 — REST API

Öğrenilecek konular:

- REST mantığı
- HTTP methods
- Status codes
- JSON
- Request body
- Response body
- Validation
- API endpoint yapısı

**Durum:** Bekliyor

---

## Aşama 19 — Backend Mimarisi

Öğrenilecek konular:

- `cmd`
- `internal`
- Handler layer
- Service layer
- Repository layer
- Dependency management
- Separation of concerns
- Dependency injection temel mantığı

**Durum:** Bekliyor

---

## Aşama 20 — Database

Öğrenilecek konular:

- SQL temelleri
- PostgreSQL
- Database connection
- CRUD
- Repository pattern
- Transactions
- Connection management

**Durum:** Bekliyor

---

## Aşama 21 — Testing

Öğrenilecek konular:

- `testing` package
- Unit testing
- Table-driven tests
- HTTP tests
- Mocking
- Integration tests

**Durum:** Bekliyor

---

## Aşama 22 — Configuration

Öğrenilecek konular:

- Environment variables
- Configuration structs
- `.env`
- Uygulama konfigürasyonu
- Development / production ayarları

**Durum:** Bekliyor

---

## Aşama 23 — Logging

Öğrenilecek konular:

- Logging mantığı
- Structured logging
- Log levels
- Request logging
- Application logging

**Durum:** Bekliyor

---

## Aşama 24 — Docker

Öğrenilecek konular:

- Dockerfile
- Go uygulamasını build etme
- Containerization
- Multi-stage builds
- Docker image
- Docker container

**Durum:** Bekliyor

---

## Aşama 25 — Docker Compose

Öğrenilecek konular:

- Birden fazla servis
- Docker networks
- Environment variables
- Health checks
- Service dependencies
- Backend + Database ortamı

**Durum:** Bekliyor

---

## Aşama 26 — Kubernetes

Öğrenilecek konular:

- Pod
- Deployment
- Service
- ConfigMap
- Environment variables
- Readiness Probe
- Liveness Probe
- Resource requests
- Resource limits
- Scaling

**Durum:** Bekliyor

---

## Aşama 27 — Observability

Öğrenilecek konular:

- Metrics
- Application metrics
- Prometheus temel mantığı
- Kubernetes Metrics Server
- CPU / Memory metrics
- Horizontal Pod Autoscaler

**Durum:** Bekliyor

---

## Aşama 28 — Load Testing

Öğrenilecek konular:

- Load testing
- Virtual Users
- Throughput
- Latency
- Average
- Median
- P95
- k6
- Performans analizi

**Durum:** Bekliyor


