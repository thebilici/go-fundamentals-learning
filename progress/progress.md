# Go Backend Learning — Progress

**Son Güncelleme:** 2026-08-27  
**Durum:** ✅ Tamamlandı

---

## Genel Durum

`go-backend-learning` kapsamında planlanan Go öğrenme süreci tamamlandı.

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

Bu repo ile Go'nun temel dil özellikleri, veri yapıları, concurrency modeli ve temel HTTP kullanımı uygulamalı olarak çalışıldı.

---

# Tamamlanan Aşamalar

## Aşama 1 — Go Ortamı ve Temeller ✅

Go geliştirme ortamı kuruldu ve Go programlarının temel çalışma yapısı öğrenildi.

Çalışılan konular:

- Go kurulumu
- `go version`
- `go mod init`
- `go.mod`
- Module mantığı
- `package main`
- `import`
- `func main()`
- `fmt.Println`
- `go run`
- `go build`
- `go fmt`
- Executable oluşturma

Temel çalışma akışı:

```text
.go source code
      ↓
go run / go build
      ↓
Go Compiler
      ↓
Executable
```

---

## Aşama 2 — Variables & Data Types ✅

Go'daki değişken tanımlama ve temel veri tipleri öğrenildi.

Çalışılan konular:

- `var`
- `:=`
- `const`
- `string`
- `int`
- `float64`
- `bool`
- Type inference
- Type conversion
- `float64()`
- `strconv.Atoi()`
- `strconv.Itoa()`
- `fmt.Printf`
- `%T`

Temel olarak Go'nun statically typed bir dil olması ve değişkenlerin belirli bir type'a sahip olması üzerinde duruldu.

---

## Aşama 3 — Control Flow ✅

Programın çalışma akışını kontrol eden yapılar öğrenildi.

Çalışılan konular:

- `if`
- `else`
- Comparison operators
- Logical operators
- `switch`
- `for`
- `range`
- `break`
- `continue`

Go'da klasik `while` yerine `for` kullanımı ve collection'lar üzerinde `range` ile dolaşma uygulandı.

---

## Aşama 4 — Functions ✅

Kod içerisindeki sorumlulukları function'lara ayırma ve veri aktarımı çalışıldı.

Çalışılan konular:

- Function declaration
- Parameters
- Arguments
- Return values
- Multiple return values
- Function scope
- Type conversion
- Function composition

Özellikle:

```go
func example() (string, int, bool)
```

gibi birden fazla değer döndüren function'lar üzerinde çalışıldı.

Exercise'larda:

```text
Input
  ↓
Function
  ↓
Processing
  ↓
Return Value
```

mantığı pekiştirildi.

---

## Aşama 5 — Arrays & Slices ✅

Go'daki collection yapılarından array ve slice öğrenildi.

Çalışılan konular:

- Arrays
- Fixed size
- Slices
- `append`
- `len`
- `cap`
- Slicing
- `range`
- Index / value
- Slice filtering
- Yeni slice oluşturma

Özellikle:

```go
for _, value := range values {
    if condition {
        result = append(result, value)
    }
}
```

pattern'i farklı exercise'larda uygulandı.

---

## Aşama 6 — Maps ✅

Key-value tabanlı veri saklama yapısı öğrenildi.

Çalışılan konular:

- `map`
- Key / Value
- Map initialization
- `make`
- Value ekleme
- Value güncelleme
- `delete`
- `len`
- `range`
- Map lookup
- `value, ok`

Özellikle:

```go
value, ok := data[key]
```

pattern'i ile bir key'in map içerisinde bulunup bulunmadığının kontrol edilmesi çalışıldı.

---

## Aşama 7 — Structs ✅

Birbiriyle ilişkili verileri kendi veri tipleri içerisinde modellemek için struct yapısı öğrenildi.

Çalışılan konular:

- `struct`
- Custom types
- Struct fields
- Struct initialization
- Field access
- Exported / unexported fields
- Slice of structs
- Struct filtering
- Struct searching
- Zero value

Örnek yapı:

```go
type User struct {
    ID       int
    Name     string
    Age      int
    IsActive bool
}
```

Struct'ların slice içerisinde tutulması ve belirli özelliklere göre filtrelenmesi uygulandı.

---

## Aşama 8 — Methods ✅

Function'ların belirli bir type'a bağlanması ve receiver yapısı öğrenildi.

Çalışılan konular:

- Methods
- Receiver
- Value receiver
- Pointer receiver
- Struct methods
- State modification

Temel fark:

```text
Value Receiver
→ değeri okumak / kopya üzerinde çalışmak

Pointer Receiver
→ gerçek struct state'ini değiştirmek
```

Örneğin:

```go
func (p *Product) AddStock(amount int) {
    p.Stock += amount
}
```

ile struct içerisindeki gerçek state'in değiştirilmesi uygulandı.

---

## Aşama 9 — Pointers ✅

Memory address ve pointer mantığı öğrenildi.

Çalışılan konular:

- Pointer
- Memory address
- `&`
- `*`
- Dereferencing
- Pointer parameters
- Pointer receivers
- State modification

Temel ilişki:

```text
Variable
   ↓
Memory Address
   ↓
Pointer
   ↓
Original Value
```

Pointer kullanımının özellikle bir değerin gerçek state'ini değiştirmek gerektiğinde neden önemli olduğu incelendi.

---

## Aşama 10 — Interfaces ✅

Farklı type'ların ortak davranışlar üzerinden kullanılmasını sağlayan interface yapısı öğrenildi.

Çalışılan konular:

- Interface declaration
- Method sets
- Implicit implementation
- Interface parameters
- Multiple implementations
- Polymorphism

Örnek:

```go
type Notifier interface {
    Send() string
    GetType() string
}
```

Bir type'ın `Notifier` interface'ini sağlayabilmesi için interface içerisinde belirtilen methodların tamamına sahip olması gerektiği uygulandı.

```text
Notifier
├── Send()
└── GetType()

EmailNotifier
├── Send()    ✅
└── GetType() ✅

SMSNotifier
├── Send()    ✅
└── GetType() ✅
```

Aynı interface üzerinden farklı concrete type'ların tek bir function tarafından kullanılabilmesi çalışıldı.

---

## Aşama 11 — Error Handling ✅

Go'nun error handling yaklaşımı öğrenildi.

Çalışılan konular:

- `error`
- `errors.New()`
- `fmt.Errorf()`
- `nil`
- `(value, error)`
- `err != nil`
- Error propagation
- Error wrapping
- `%w`

Temel pattern:

```go
value, err := operation()

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

Alt seviyedeki error'ın üst seviyeye taşınması da uygulandı.

```text
SellProduct()
      ↓
error
      ↓
ProcessOrder()
      ↓
fmt.Errorf(... %w ...)
      ↓
main()
```

Böylece error'ın kaybedilmeden ek context ile üst katmana taşınması öğrenildi.

---

## Aşama 12 — Packages & Modules ✅

Go kodunun birden fazla package içerisinde organize edilmesi öğrenildi.

Çalışılan konular:

- Packages
- Custom packages
- Imports
- Exported identifiers
- Unexported identifiers
- Modules
- `go.mod`
- Package organization

Temel yapı:

```text
Module
│
├── main package
│
└── custom package
```

Kodun tek bir dosyada tutulması yerine farklı sorumluluklara ayrılmasının temelleri çalışıldı.

---

## Aşama 13 — Concurrency ✅

Go'nun concurrency modeli öğrenildi.

Çalışılan konular:

- Sequential execution
- Concurrent execution
- Concurrency vs Parallelism
- Goroutines
- `go` keyword
- Main Goroutine
- `sync.WaitGroup`
- `Add()`
- `Done()`
- `Wait()`
- Shared state
- Race Condition
- Race Detector
- `sync.Mutex`
- `Lock()`
- `Unlock()`

Temel çalışma modeli:

```text
Main Goroutine
     │
     ├── Goroutine 1
     ├── Goroutine 2
     └── Goroutine 3
```

Birden fazla işin concurrent olarak çalıştırılması ve goroutine'lerin tamamlanmasının beklenmesi uygulandı.

Shared state üzerinde oluşabilecek race condition problemi ve Mutex ile korunması incelendi.

---

## Aşama 14 — Channels ✅

Goroutine'ler arasında veri iletişimi için kullanılan channel yapısı öğrenildi.

Çalışılan konular:

- Channel
- `make(chan T)`
- Channel send
- Channel receive
- Unbuffered Channel
- Buffered Channel
- `len()`
- `cap()`
- Goroutine communication
- Channel ile sonuç döndürme
- `close()`
- Channel üzerinde `range`
- WaitGroup + Channel
- Deadlock temel mantığı

Temel model:

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

Concurrency'nin yalnızca aynı anda iş çalıştırmak olmadığı, goroutine'ler arasında güvenli iletişim kurulmasının da önemli olduğu görüldü.

---

## Aşama 15 — Generics ✅

Farklı type'larla çalışabilen type-safe yapılar oluşturmak için generics öğrenildi.

Çalışılan konular:

- Generics
- Type Parameters
- Type Arguments
- Type Constraints
- `any`
- `comparable`
- Generic Functions
- Multiple Type Parameters
- Generic Structs
- Generic Slice
- Generic Map
- Custom Constraints
- Type Inference

Yapılan temel pratikler:

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

## Aşama 16 — Context ✅

Go'da uzun süren ve birbirine bağlı işlemlerin yaşam döngüsünü yönetmek için kullanılan `context` yapısı öğrenildi.

Çalışılan konular:

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
- Request lifecycle temel mantığı

Temel akış:

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

Context'in özellikle ileride HTTP request'leri, database işlemleri ve servisler arası çağrılarda kullanılabilecek temel bir lifecycle yönetim mekanizması olduğu görüldü.

---

## Aşama 17 — HTTP ✅

HTTP'nin temel çalışma mantığı ve Go'nun `net/http` package'ı öğrenildi.

Çalışılan konular:

- HTTP temel mantığı
- Client
- Server
- Request
- Response
- HTTP Methods
- Status Codes
- Headers
- Body
- `net/http`
- HTTP Server
- Handler
- `http.ResponseWriter`
- `*http.Request`
- Temel routing
- Endpoint oluşturma

Go ile temel bir HTTP server oluşturuldu.

Temel request lifecycle:

```text
Client
   │
   │ HTTP Request
   ▼
Go HTTP Server
   │
   ▼
Router / Handler
   │
   │ işle
   ▼
HTTP Response
   │
   ▼
Client
```

Bu aşamada amaç kapsamlı bir REST API geliştirmek değil, HTTP'nin temel çalışma modelini ve Go ile bir HTTP server'ın nasıl oluşturulduğunu anlamaktı.

Backend geliştirme tarafındaki daha kapsamlı HTTP kullanımları gerçek proje içerisinde ele alınacaktır.

---

# Tekrar ve Pekiştirme Çalışmaları

Temel Go konuları tamamlandıktan sonra önemli konular bağımsız exercise'larla tekrar edildi.

Tekrar akışı:

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

Bu tekrar sürecinde farklı senaryolar kullanıldı.

Örnek problem alanları:

```text
Maaş hesaplama
Stok yönetimi
Sıcaklık filtreleme
Sipariş yönetimi
Kullanıcı yönetimi
Banka hesabı
Ürün yönetimi
Ödeme sistemleri
Notification sistemi
Exporter sistemi
Ürün satışı
```

Tekrar edilen başlıca kavramlar:

- Function parameters
- Return values
- Multiple return values
- Function scope
- Slice iteration
- Slice filtering
- `append`
- Map lookup
- `value, ok`
- Map update
- Struct modelling
- Slice of structs
- Struct filtering
- Struct searching
- Methods
- Value receiver
- Pointer receiver
- State modification
- Interface implementation
- Multiple interface methods
- Polymorphism
- `(value, error)`
- `errors.New()`
- `fmt.Errorf()`
- Error propagation
- Error wrapping
- `%w`

---

# Repo Son Durumu

`go-backend-learning` reposu tamamlandı.

Repo kapsamında:

```text
Go Fundamentals
        +
Variables & Types
        +
Control Flow
        +
Functions
        +
Arrays & Slices
        +
Maps
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

konuları uygulamalı olarak öğrenildi.

---

# Kazanımlar

Repo sonunda temel olarak:

```text
Go kodu yazabilme
        ↓
Verileri uygun type'larla modelleyebilme
        ↓
Function'larla sorumlulukları ayırabilme
        ↓
Slice ve Map'lerle veri işleyebilme
        ↓
Struct'larla domain modelleri oluşturabilme
        ↓
Method'larla davranış ekleyebilme
        ↓
Pointer'larla state değiştirebilme
        ↓
Interface'lerle ortak davranış tanımlayabilme
        ↓
Error'ları açık şekilde yönetebilme
        ↓
Package'larla kodu organize edebilme
        ↓
Goroutine'lerle concurrent işlemler oluşturabilme
        ↓
Channel'larla goroutine iletişimi kurabilme
        ↓
Generics ile reusable type-safe kod yazabilme
        ↓
Context ile işlem lifecycle'ını kontrol edebilme
        ↓
Temel bir Go HTTP server oluşturabilme
```

seviyesine gelmek hedeflendi.

---

# Sonraki Adım

Bu repo tamamlanmıştır ve yeni konu eklenmesi planlanmamaktadır.

Bundan sonraki aşamada Go kavramları ayrı ayrı çalışılmak yerine gerçek bir backend ve cloud-native proje içerisinde kullanılacaktır.

```text
go-backend-learning
        ↓
Go Fundamentals Complete ✅
        ↓
Gerçek Go Backend Projesi
        ↓
REST API
        ↓
Backend Architecture
        ↓
Database
        ↓
Testing
        ↓
Configuration
        ↓
Logging
        ↓
Service-to-Service Communication
        ↓
Concurrency'nin Gerçek Kullanımı
        ↓
Context'in Gerçek Kullanımı
        ↓
Docker
        ↓
Docker Compose
        ↓
Load Testing
        ↓
Kubernetes
        ↓
Metrics
        ↓
Horizontal Pod Autoscaler
        ↓
Observability
```

Yeni projedeki temel amaç yeni Go syntax'ı öğrenmek değil, bu repoda öğrenilen Go kavramlarını gerçek bir backend sisteminin parçaları olarak kullanmak ve neden gerekli olduklarını uygulama içerisinde görmektir.

---

# Final Durum

```text
GO BACKEND LEARNING

17 / 17 Aşama Tamamlandı

████████████████████ 100%

Go Fundamentals Complete ✅
```
```