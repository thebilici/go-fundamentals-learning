# Go Backend Learning — Roadmap

**Son Güncelleme:** 2026-08-27  
**Durum:** ✅ Tamamlandı  
**Tamamlanan Aşama:** 17 / 17

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
Channels                    ✅
        ↓
Generics                    ✅
        ↓
Context                     ✅
        ↓
HTTP                        ✅
        ↓
Go Fundamentals Complete    ✅
```

Bu roadmap ile Go'nun temel dil özellikleri, veri yapıları, error handling yaklaşımı, concurrency modeli ve temel HTTP kullanımı uygulamalı olarak tamamlandı.

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

Temel çalışma modeli:

```text
Go Source Code
      ↓
Compiler
      ↓
Executable
```

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

Bu aşamada Go'nun statically typed yapısı ve değişkenlerin type'larla ilişkisi öğrenildi.

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

Programın çalışma akışını koşullara ve döngülere göre kontrol etme mantığı uygulandı.

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
- Function'ları birlikte kullanma

Örnek:

```go
func getUser() (string, int) {
    return "Fatih", 22
}
```

Temel çalışma modeli:

```text
Input
  ↓
Function
  ↓
Processing
  ↓
Return Value
```

Function'ların yalnızca kod tekrarını azaltmak için değil, farklı sorumlulukları birbirinden ayırmak için kullanılabileceği uygulandı.

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
- Slice filtering
- Yeni slice oluşturma

Temel filtering pattern'i:

```go
for _, value := range values {
    if condition {
        result = append(result, value)
    }
}
```

---

## Aşama 6 — Maps

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- Map
- Key / Value
- Map oluşturma
- `make`
- Eleman okuma
- Eleman ekleme
- Eleman güncelleme
- `value, ok` idiomu
- Map zero value davranışı
- `delete()`
- `len()`
- `range`
- Map filtering

Temel lookup pattern'i:

```go
value, ok := data[key]
```

Bu yapı ile bir key'in map içerisinde bulunup bulunmadığının kontrol edilmesi uygulandı.

---

## Aşama 7 — Structs

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- Struct
- Field
- Custom Type
- Struct tanımlama
- Struct oluşturma
- Field erişimi
- Field güncelleme
- Struct + Function
- Slice of Structs
- `range` ile Struct dolaşma
- Struct filtering
- Struct searching
- Zero value
- Exported Fields
- Unexported Fields

Örnek:

```go
type User struct {
    ID       int
    Name     string
    Age      int
    IsActive bool
}
```

Birden fazla ilişkili veriyi tek bir type altında modelleme ve struct collection'ları üzerinde işlem yapma mantığı uygulandı.

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
- Pointer Receiver'a giriş

Örnek:

```go
func (u User) GetName() string {
    return u.Name
}
```

Method sayesinde davranışların ilgili type'a bağlanması öğrenildi.

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
- Struct state'ini pointer ile değiştirme

Örnek:

```go
func updateAge(u *User) {
    u.Age = 25
}
```

Temel fark:

```text
Value Receiver
→ değer üzerinde çalışır

Pointer Receiver
→ gerçek state üzerinde değişiklik yapabilir
```

Pointer kullanımının özellikle state değiştiren method'larda neden gerekli olduğu uygulandı.

---

## Aşama 10 — Interfaces

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- Interface
- Method Signature
- Method Set
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
    Send() string
    GetType() string
}
```

Temel ilişki:

```text
Notifier
├── Send()
└── GetType()
       ↑
       │
 ┌─────┴─────┐
 │           │
Email       SMS
```

Bir type'ın interface'i kullanabilmesi için interface içerisinde tanımlanan method set'ini sağlaması gerektiği öğrenildi.

Uygulanan örnekler:

- `Greeter`
- `User`
- `Admin`
- `Notifier`
- `EmailNotifier`
- `SMSNotifier`
- `PaymentMethod`
- `CreditCard`
- `CashPayment`
- `Exporter`
- `JSONExporter`
- `CSVExporter`

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

Başarılı işlem:

```text
value, nil
```

Başarısız işlem:

```text
zero/default value, error
```

Error propagation:

```text
Alt Function
     ↓
   error
     ↓
Üst Function
     ↓
   error
     ↓
   main
```

Error wrapping:

```go
return fmt.Errorf("operation failed: %w", err)
```

Böylece bir error'ın kaybedilmeden üst katmana ek context ile taşınması uygulandı.

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
│
├── main package
│
└── custom package
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
- Import alias

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
- Race Condition
- Race Detector
- Shared State
- `sync.Mutex`
- `Lock()`
- `Unlock()`

Temel Goroutine:

```go
go task()
```

WaitGroup:

```text
Add()
  ↓
Goroutine
  ↓
Done()
  ↓
Wait()
```

Temel concurrency modeli:

```text
Main Goroutine
     │
     ├── Goroutine 1
     ├── Goroutine 2
     └── Goroutine 3
```

Birden fazla işin concurrent şekilde çalıştırılması ve shared state üzerinde oluşabilecek race condition problemlerinin kontrol edilmesi uygulandı.

---

## Aşama 14 — Channels

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- Channel
- Channel Send
- Channel Receive
- `make(chan T)`
- Unbuffered Channel
- Buffered Channel
- `len()`
- `cap()`
- Goroutine'ler arası iletişim
- Channel ile sonuç döndürme
- Birden fazla goroutine'in aynı channel'a veri göndermesi
- `close()`
- Channel üzerinde `range`
- WaitGroup + Channel
- Deadlock temel mantığı

Temel kullanım:

```go
ch <- value

value := <-ch
```

Temel iletişim modeli:

```text
Goroutine A
     │
     │ data
     ▼
   Channel
     │
     │ data
     ▼
Goroutine B
```

WaitGroup + Channel pattern:

```text
Worker 1 ──┐
Worker 2 ──┼──→ Channel ──→ Receiver
Worker 3 ──┘
     │
     ▼
 WaitGroup
     │
     ▼
 close(ch)
```

Bu aşamada goroutine oluşturmanın yanında goroutine'ler arasında güvenli veri iletişimi kurma mantığı öğrenildi.

---

## Aşama 15 — Generics

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- Generics
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

Yapılan pratikler:

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

Örnek:

```go
func GetValue[K comparable, V any](data map[K]V, key K) (V, bool)
```

Burada:

```text
K → Key tipi
V → Value tipi
```

mantığı uygulandı.

---

## Aşama 16 — Context

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

- `context.Context`
- `context.Background()`
- Cancellation
- `context.WithCancel()`
- `cancel()`
- `ctx.Done()`
- `ctx.Err()`
- Timeout
- `context.WithTimeout()`
- Deadline
- `context.WithDeadline()`
- `defer cancel()`
- Context + Goroutine
- Context Propagation
- Request Lifecycle temel mantığı

Temel lifecycle:

```text
Parent Context
      ↓
Child Operation
      ↓
Goroutine
      ↓
ctx.Done()
      ↓
Cancellation / Timeout
```

Context'in uzun süren veya birbirine bağlı işlemlerin lifecycle'ını kontrol etmek için nasıl kullanılabileceği öğrenildi.

---

# Go ve HTTP

## Aşama 17 — HTTP

**Durum:** ✅ Tamamlandı

Öğrenilen konular:

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
- Temel Routing
- Endpoint

Temel HTTP akışı:

```text
Client
   │
   │ HTTP Request
   ▼
Go HTTP Server
   │
   ▼
Handler
   │
   │ Request işle
   ▼
HTTP Response
   │
   ▼
Client
```

Bu aşamanın amacı kapsamlı bir REST API geliştirmek değil, HTTP'nin temel çalışma mantığını ve Go ile temel bir HTTP server oluşturmayı öğrenmekti.

Daha kapsamlı HTTP ve REST API kullanımı gerçek backend projesinde uygulanacaktır.

---

# Pekiştirme Süreci

Temel Go konuları tamamlandıktan sonra önemli konular bağımsız exercise'larla tekrar edildi.

Pekiştirilen ana konular:

```text
Functions
    ↓
Slices
    ↓
Maps
    ↓
Structs
    ↓
Methods
    ↓
Pointer Receivers
    ↓
Interfaces
    ↓
Error Handling
```

Yapılan exercise'larda farklı problem alanları kullanıldı:

- Maaş hesaplama
- Stok filtreleme
- Sıcaklık dönüşümü
- Sipariş yönetimi
- Kullanıcı yönetimi
- Banka hesabı
- Ürün ve stok yönetimi
- Payment sistemi
- Notification sistemi
- Exporter sistemi
- Ürün satışı

Bu çalışmalar sırasında özellikle:

- Function parameters
- Multiple return values
- Slice filtering
- `append`
- Map lookup
- `value, ok`
- Struct modelling
- Slice of structs
- Struct filtering
- Struct searching
- Value receiver
- Pointer receiver
- Interface implementation
- Polymorphism
- `(value, error)`
- `errors.New()`
- `fmt.Errorf()`
- Error propagation
- Error wrapping
- `%w`

konuları tekrar uygulandı.

---

# Roadmap Son Durumu

Planlanan 17 aşamanın tamamı tamamlandı.

```text
17 / 17 Aşama

████████████████████ 100%

Go Fundamentals Complete ✅
```

Bu repo kapsamında:

```text
Go Syntax
      +
Core Language Features
      +
Data Structures
      +
Functions
      +
Structs
      +
Methods
      +
Pointers
      +
Interfaces
      +
Error Handling
      +
Packages & Modules
      +
Concurrency
      +
Channels
      +
Generics
      +
Context
      +
HTTP Basics
```

konuları uygulamalı olarak çalışıldı.

---

# Repo Sınırı

Bu roadmap'in kapsamı burada tamamlanmaktadır.

Aşağıdaki konular bu reponun roadmap'inin bir parçası değildir:

```text
REST API Development
Backend Architecture
Database
Testing
Application Configuration
Structured Logging
Service-to-Service Communication
Docker
Docker Compose
Load Testing
Kubernetes
Metrics
Horizontal Pod Autoscaler
Observability
```

Bu konular bağımsız örnekler şeklinde bu repoya eklenmek yerine gerçek bir Go backend ve cloud-native projesi içerisinde öğrenilecektir.

---

# Sonraki Yol

Bu repo ile kazanılan Go temelleri bundan sonra gerçek bir proje içerisinde kullanılacaktır.

```text
go-backend-learning
        │
        ▼
Go Fundamentals Complete ✅
        │
        ▼
Gerçek Go Backend / Cloud-Native Projesi
        │
        ├── REST API
        │
        ├── Backend Architecture
        │
        ├── Database
        │
        ├── Testing
        │
        ├── Configuration
        │
        ├── Logging
        │
        ├── Service-to-Service Communication
        │
        ├── Concurrency'nin Gerçek Kullanımı
        │
        ├── Context'in Gerçek Kullanımı
        │
        ├── Docker
        │
        ├── Docker Compose
        │
        ├── Load Testing
        │
        ├── Kubernetes
        │
        ├── Metrics
        │
        ├── HPA
        │
        └── Observability
        │
        ▼
Production-Oriented Go Backend Temeli
```

Buradaki amaç artık yeni Go syntax'ı öğrenmek değil, bu roadmap boyunca öğrenilen kavramların gerçek bir sistem içerisinde neden ve nerede kullanıldığını görmek olacaktır.

---

# Öğrenme Prensibi

Bu roadmap boyunca her aşamada temel olarak aşağıdaki döngü takip edildi:

```text
Kavram
  ↓
Neden kullanılır?
  ↓
Küçük örnek
  ↓
Kodu inceleme
  ↓
Exercise
  ↓
Önceki konularla birleştirme
  ↓
Concepts
  ↓
Flashcards
  ↓
Questions
  ↓
Mistakes
  ↓
Progress güncelleme
  ↓
Commit
```

Amaç yalnızca çalışan kod yazmak değildi.

Amaç:

> **Kodun neden çalıştığını, hangi problemi çözdüğünü ve daha büyük bir Go uygulamasında nerede kullanılabileceğini anlayabilmek.**

---

# Final

```text
GO BACKEND LEARNING ROADMAP

Aşama 01  Go Basics             ✅
Aşama 02  Variables & Types     ✅
Aşama 03  Control Flow          ✅
Aşama 04  Functions             ✅
Aşama 05  Arrays & Slices       ✅
Aşama 06  Maps                  ✅
Aşama 07  Structs               ✅
Aşama 08  Methods               ✅
Aşama 09  Pointers              ✅
Aşama 10  Interfaces            ✅
Aşama 11  Error Handling        ✅
Aşama 12  Packages & Modules    ✅
Aşama 13  Concurrency           ✅
Aşama 14  Channels              ✅
Aşama 15  Generics              ✅
Aşama 16  Context               ✅
Aşama 17  HTTP                  ✅

17 / 17 COMPLETE

Go Fundamentals Complete ✅
```