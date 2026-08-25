# Go Backend Learning Roadmap

**Son Güncelleme:** 2026-08-25  
**Güncel Aşama:** Aşama 14 — Generics

---

# Genel Yol Haritası

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

# Go Temelleri

## Aşama 1 — Go Ortamı ve Temeller

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- Go kurulumu
- `go version`
- Go source file yapısı
- `package main`
- `import`
- `func main()`
- `fmt`
- `fmt.Println()`
- `go run`
- `go build`
- `go fmt`
- Executable mantığı
- Temel Go program yapısı
- `go.mod` ile ilk tanışma

---

## Aşama 2 — Variables & Data Types

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- Değişkenler
- `var`
- `:=`
- Constants
- `const`
- `string`
- `int`
- `float64`
- `bool`
- Type inference
- Zero values
- Type conversion
- `fmt.Printf()`
- `%T`
- `strconv.Atoi()`
- `strconv.Itoa()`

---

## Aşama 3 — Control Flow

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- Comparison Operators
- Logical Operators
- `if`
- `else`
- `else if`
- `switch`
- `for`
- `break`
- `continue`
- `range`
- Slice üzerinde `range`

---

## Aşama 4 — Functions

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- Function tanımlama
- Function çağırma
- Parameters
- Arguments
- Return Values
- Multiple Return Values
- Named Returns
- Function Scope
- Return type
- Birden fazla değer döndürme

Örnek:

```go
func getUser() (string, int) {
	return "Fatih", 22
}
```

---

## Aşama 5 — Arrays & Slices

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- Array
- Array uzunluğu
- Index
- Slice
- Array vs Slice
- `len()`
- `cap()`
- `append()`
- Slice capacity
- Slicing
- `[start:end]`
- `range`
- Slice üzerinde iteration

---

## Aşama 6 — Maps

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- Map
- Key / Value
- Map oluşturma
- Eleman okuma
- Eleman ekleme
- Eleman güncelleme
- `value, ok` idiomu
- Map zero value davranışı
- `delete()`
- `len()`
- `range`

---

## Aşama 7 — Structs

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- Struct
- Field
- Struct tanımlama
- Struct oluşturma
- Field erişimi
- Field güncelleme
- Struct + Function
- Slice of Structs
- `range` ile Struct dolaşma
- Exported Fields
- Unexported Fields

Örnek:

```go
type User struct {
	Name     string
	Age      int
	IsActive bool
}
```

---

## Aşama 8 — Methods

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- Method
- Function vs Method
- Receiver
- Value Receiver
- Method Parameters
- Method Return Values
- Struct + Method
- Receiver üzerinden field erişimi

Örnek:

```go
func (u User) GetName() string {
	return u.Name
}
```

---

## Aşama 9 — Pointers

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- Pointer
- Memory Address
- `&` operator
- `*` operator
- Dereferencing
- Pointer Types
- Pointer + Function
- Pointer Receiver
- Value Receiver vs Pointer Receiver
- Struct değerlerini pointer ile değiştirme

Örnek:

```go
func updateAge(u *User) {
	u.Age = 25
}
```

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
- Bir struct'ın interface'i implicit implement etmesi
- Farklı concrete type'ların aynı interface üzerinden kullanılması

Örnek:

```go
type Notifier interface {
	send() string
}
```

Uygulanan örnekler:

- `EmailNotifier`
- `SMSNotifier`
- `Greeter`
- `User`
- `Admin`

---

## Aşama 11 — Error Handling

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- `error`
- `(value, error)`
- `nil`
- `err != nil`
- `errors.New()`
- `fmt.Errorf()`
- Error Propagation
- Error Wrapping
- `%w`
- Error Handling vs Propagation
- Struct Zero Value
- Hatanın function katmanları arasında taşınması

Temel pattern:

```go
result, err := operation()

if err != nil {
	return err
}
```

Error wrapping:

```go
return fmt.Errorf("operation failed: %w", err)
```

---

## Aşama 12 — Packages & Modules

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- Package
- `package main`
- `func main()` ve Entry Point
- Custom Package
- Bir package içerisinde birden fazla `.go` dosyası
- Exported Identifiers
- Unexported Identifiers
- Module
- Module Path
- Package Path
- Import Path
- Import Alias
- `go.mod`
- `go mod init`
- `go mod tidy`
- `go.sum`

Temel ilişki:

```text
Module
└── Package
    ├── file.go
    └── file.go
```

Import path:

```text
Module Path
+
Package Path
=
Import Path
```

Yapılan pratikler:

- `user` package
- `User`
- `CreateUser()`
- `GetName()`
- `validation.go`
- `mathutil` package
- `Add()`
- `Multiply()`
- `IsPositive()`
- Import alias (`mathpkg`)

---

# Go İleri Dil Özellikleri

## Aşama 13 — Concurrency

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- Concurrency
- Sequential vs Concurrent
- Concurrency vs Parallelism
- Goroutine
- `go` keyword
- Main Goroutine
- `time.Sleep()`
- `sync.WaitGroup`
- `Add()`
- `Done()`
- `Wait()`
- Channel
- Channel Send
- Channel Receive
- Unbuffered Channel
- Buffered Channel
- `len()`
- `cap()`
- Goroutine'ler arası iletişim
- Channel ile sonuç döndürme
- Birden fazla goroutine'in aynı channel'a veri göndermesi
- `close()`
- Channel üzerinde `range`
- WaitGroup + Channel pattern
- Deadlock
- Race Condition
- Race Detector
- `sync.Mutex`
- `Lock()`
- `Unlock()`

### Temel Goroutine

```go
go task()
```

### WaitGroup

```text
Add()
  ↓
Goroutine
  ↓
Done()
  ↓
Wait()
```

### Channel

```go
ch <- value

value := <-ch
```

### Concurrency Pattern

```text
Worker 1 ──┐
Worker 2 ──┼──→ Channel ──→ Receiver
Worker 3 ──┘
     │
     ↓
  WaitGroup
     │
     ↓
 close(ch)
```

Yapılan final exercise:

```text
CalculateSquare(2)
CalculateSquare(4)
CalculateSquare(6)
        ↓
    Goroutines
        ↓
 results channel
        ↓
      main
```

---

## Aşama 15 — Generics

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- Generics nedir?
- Type Parameters
- Type Arguments
- Type Constraints
- `any`
- `comparable`
- Generic Functions
- Birden fazla Type Parameter
- Generic Slices
- Generic Maps
- Generic Structs
- Custom Constraints
- Type Inference
- `K` / `V` mantığı
---

## Aşama 16 — Context

**Durum:** 🔄 Güncel Aşama

Öğrenilecek konular:

- `context` package
- `context.Context`
- `context.Background()`
- Cancellation
- `context.WithCancel()`
- Timeout
- `context.WithTimeout()`
- Deadline
- `context.WithDeadline()`
- `ctx.Done()`
- `ctx.Err()`
- Context Propagation
- Request Lifecycle
- Goroutine + Context
- Context ne zaman kullanılmalı?
---

# Go Backend

## Aşama 16 — HTTP

**Durum:** ⏳ Bekliyor

Öğrenilecek konular:

- HTTP nedir?
- Client
- Server
- Request
- Response
- HTTP Request Lifecycle
- HTTP Methods
- GET
- POST
- PUT
- PATCH
- DELETE
- Status Codes
- Headers
- Request Body
- Response Body
- `net/http`
- HTTP Server
- `http.ListenAndServe()`
- Handler
- `http.Handler`
- `http.HandlerFunc`
- `http.ResponseWriter`
- `*http.Request`
- Routing
- Endpoint
- Query Parameters
- Path mantığı

İlk hedef:

```text
Client
   ↓
HTTP Request
   ↓
Go HTTP Server
   ↓
Handler
   ↓
HTTP Response
   ↓
Client
```

---

## Aşama 17 — REST API

**Durum:** ⏳ Bekliyor

Öğrenilecek konular:

- REST nedir?
- Resource mantığı
- Endpoint tasarımı
- HTTP Methods
- GET
- POST
- PUT
- PATCH
- DELETE
- HTTP Status Codes
- JSON
- `encoding/json`
- JSON Encode
- JSON Decode
- Struct Tags
- Request Body
- Response Body
- URL Parameters
- Query Parameters
- Validation
- Error Response
- API Response yapısı
- CRUD endpoint'leri

Örnek hedef:

```text
GET    /users
GET    /users/{id}
POST   /users
PUT    /users/{id}
DELETE /users/{id}
```

---

## Aşama 18 — Backend Mimarisi

**Durum:** ⏳ Bekliyor

Öğrenilecek konular:

- Go backend proje organizasyonu
- `cmd`
- `internal`
- `config`
- `handler`
- `service`
- `repository`
- `model`
- Handler Layer
- Service Layer
- Repository Layer
- Separation of Concerns
- Dependency Management
- Dependency Injection temel mantığı
- Interface tabanlı dependency
- Application bootstrap
- Server başlangıç yapısı

Hedef mimari:

```text
Client
  ↓
Handler
  ↓
Service
  ↓
Repository
  ↓
Database
```

---

## Aşama 19 — Database

**Durum:** ⏳ Bekliyor

Öğrenilecek konular:

- Relational Database mantığı
- SQL temelleri
- PostgreSQL
- Table
- Row
- Primary Key
- Foreign Key
- Database Connection
- Go database bağlantısı
- `database/sql`
- Driver mantığı
- CRUD
- SELECT
- INSERT
- UPDATE
- DELETE
- Parameterized Queries
- Repository Pattern
- Context ile database işlemleri
- Transactions
- Connection Pool
- Connection Management
- Database error handling

Hedef:

```text
HTTP Request
     ↓
Handler
     ↓
Service
     ↓
Repository
     ↓
PostgreSQL
```

---

## Aşama 20 — Testing

**Durum:** ⏳ Bekliyor

Öğrenilecek konular:

- Test neden yazılır?
- `testing` package
- `_test.go`
- `go test`
- Unit Testing
- Test function
- Table-Driven Tests
- Test Cases
- HTTP Tests
- `httptest`
- Handler Testing
- Service Testing
- Interface ile dependency değiştirme
- Mocking temel mantığı
- Integration Tests
- Error case testleri
- Test coverage

Komutlar:

```bash
go test ./...
go test -v ./...
```

---

# Application Configuration & Operations

## Aşama 21 — Configuration

**Durum:** ⏳ Bekliyor

Öğrenilecek konular:

- Configuration nedir?
- Environment Variables
- `os.Getenv()`
- Configuration Struct
- `.env` mantığı
- Port configuration
- Database configuration
- Development environment
- Production environment
- Secret değerleri koddan ayırma
- Default values
- Configuration validation

Temel hedef:

```text
Environment Variables
        ↓
Config
        ↓
Application
```

---

## Aşama 22 — Logging

**Durum:** ⏳ Bekliyor

Öğrenilecek konular:

- Logging nedir?
- Logging neden gereklidir?
- Standard logging
- Structured Logging
- Log Levels
- Debug
- Info
- Warn
- Error
- Request Logging
- Application Logging
- Error Logging
- Contextual log bilgileri
- Production logging mantığı

---

# Containerization

## Aşama 23 — Docker

**Durum:** ⏳ Bekliyor

Öğrenilecek konular:

- Container nedir?
- Docker nedir?
- Image
- Container
- Dockerfile
- `FROM`
- `WORKDIR`
- `COPY`
- `RUN`
- `EXPOSE`
- `CMD`
- Go uygulamasını build etme
- Docker image oluşturma
- Container çalıştırma
- Port Mapping
- Environment Variables
- Multi-stage Builds
- Küçük production image oluşturma

Hedef:

```text
Go Source Code
      ↓
Docker Build
      ↓
Docker Image
      ↓
Container
      ↓
HTTP Server
```

---

## Aşama 24 — Docker Compose

**Durum:** ⏳ Bekliyor

Öğrenilecek konular:

- Docker Compose nedir?
- `compose.yaml`
- Services
- Build Context
- Ports
- Environment Variables
- `env_file`
- Docker Networks
- Service Name DNS
- Birden fazla servis
- Health Checks
- Service Dependencies
- `depends_on`
- Backend + Database ortamı
- Container'lar arası iletişim

Hedef:

```text
Client
  ↓
Go Backend Container
  ↓
PostgreSQL Container
```

---

# Kubernetes & Cloud-Native

## Aşama 25 — Kubernetes

**Durum:** ⏳ Bekliyor

Öğrenilecek konular:

- Kubernetes nedir?
- Cluster
- Node
- Pod
- Deployment
- Replica
- Service
- ClusterIP
- ConfigMap
- Environment Variables
- Container Image
- `kubectl`
- Kubernetes YAML
- Readiness Probe
- Liveness Probe
- Resource Requests
- Resource Limits
- Rolling Update
- Rollout
- Scaling
- Pod lifecycle
- Service discovery

Temel mimari:

```text
Client
   ↓
Service
   ↓
Deployment
   ↓
Pods
```

---

## Aşama 26 — Observability

**Durum:** ⏳ Bekliyor

Öğrenilecek konular:

- Observability nedir?
- Metrics
- Logs
- Application Metrics
- Infrastructure Metrics
- CPU Metrics
- Memory Metrics
- Prometheus temel mantığı
- Metrics endpoint
- Kubernetes Metrics Server
- `kubectl top`
- Resource kullanımı
- Horizontal Pod Autoscaler
- HPA
- CPU based scaling
- Monitoring temel mantığı

Hedef akış:

```text
Application
     ↓
Metrics
     ↓
Monitoring
     ↓
Scaling Decision
     ↓
HPA
     ↓
Replica Count
```

---

## Aşama 27 — Load Testing

**Durum:** ⏳ Bekliyor

Öğrenilecek konular:

- Load Testing nedir?
- Load Test neden yapılır?
- k6
- Virtual Users
- VU
- Iteration
- Request Rate
- Throughput
- Latency
- Average
- Median
- P95
- P99 temel mantığı
- Error Rate
- Checks
- Thresholds
- Stages
- Ramp-up
- Ramp-down
- Stress Testing
- Performans analizi
- Kubernetes altında load testing
- HPA davranışını yük altında gözlemleme

Hedef:

```text
k6
 ↓
HTTP Load
 ↓
Go Backend
 ↓
CPU / Memory Usage
 ↓
Metrics
 ↓
HPA
 ↓
Replica Scaling
```

---

# Nihai Hedef

Roadmap tamamlandığında aşağıdaki yapıyı sıfırdan anlayarak oluşturabilecek seviyeye gelmek:

```text
                        Client
                          │
                          │ HTTP
                          ▼
                   ┌─────────────┐
                   │ Go Backend  │
                   └─────────────┘
                          │
                    Handler Layer
                          │
                     Service Layer
                          │
                   Repository Layer
                          │
                          ▼
                    ┌──────────┐
                    │PostgreSQL│
                    └──────────┘

                          │
                          ▼

                    Docker Image
                          │
                          ▼
                     Containers
                          │
                          ▼
                     Kubernetes
                          │
                ┌─────────┴─────────┐
                ▼                   ▼
              Pod 1               Pod 2
                │                   │
                └─────────┬─────────┘
                          │
                       Service
                          │
                          ▼
                       Metrics
                          │
                          ▼
                         HPA
                          ▲
                          │
                      k6 Load Test
```

---

# Öğrenme Prensibi

Her aşamada aşağıdaki döngü takip edilir:

```text
Kavram
  ↓
Neden kullanılır?
  ↓
Küçük örnek
  ↓
Satır satır inceleme
  ↓
Exercise
  ↓
Gerçek backend kullanımına bağlama
  ↓
Concepts
  ↓
Flashcards
  ↓
Questions
  ↓
Mistakes
  ↓
Progress güncelle
  ↓
Commit
```

Amaç yalnızca çalışan kod yazmak değildir.

Amaç:

> **Kodun neden çalıştığını, hangi problemi çözdüğünü ve gerçek bir backend sisteminde nerede kullanılacağını anlayabilmek.**