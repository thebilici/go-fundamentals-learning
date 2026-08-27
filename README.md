# Go Backend Learning

Go programlama dilini temelden başlayarak öğrenmek ve temel Go kavramlarını uygulamalı olarak pekiştirmek amacıyla oluşturulmuş bir öğrenme reposudur.

Bu reponun amacı yalnızca Go syntax'ını öğrenmek değil; Go'nun type sistemi, veri yapıları, function yapısı, pointer mantığı, interface yaklaşımı, error handling modeli, concurrency yapısı, generics ve context gibi temel kavramlarını küçük örnekler ve bağımsız alıştırmalar üzerinden öğrenmektir.

Repo boyunca her konu önce küçük örneklerle incelenir, ardından bağımsız exercise'lar ile tekrar uygulanır.

Bu repo tamamlandıktan sonra öğrenilen Go temelleri ayrı bir gerçek backend ve cloud-native projesinde kullanılacaktır.

---

## Amaç

Öğrenme süreci aşağıdaki yolu takip eder:

```text
Go Temelleri
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
Concurrency
      ↓
Channels
      ↓
Generics
      ↓
Context
      ↓
HTTP Temelleri
      ↓
Go Fundamentals Complete
```

Bu aşamaların amacı daha büyük bir backend projesine geçmeden önce Go'nun temel yapılarını anlamak ve bağımsız olarak kullanabilecek seviyeye gelmektir.

---

## Repo Yapısı

```text
go-backend-learning/
│
├── README.md
│
├── sessions/
│
├── exercises/
│
├── notes/
│   ├── concepts.md
│   ├── flashcards.md
│   ├── mistakes.md
│   └── questions.md
│
└── progress/
    ├── roadmap.md
    └── progress.md
```

### `sessions/`

Go kavramlarını öğrenirken oluşturulan küçük ve konu odaklı örnekleri içerir.

Konular birbirinden ayrılmış session'lar halinde ilerler.

Örnek yapı:

```text
sessions/
│
├── 01-go-basics/
├── 02-variables-types/
├── 03-control-flow/
├── 04-functions/
├── 05-arrays-slices/
├── 06-maps/
├── 07-structs/
├── 08-methods/
├── 09-pointers/
├── 10-interfaces/
├── 11-errors/
├── 12-packages/
├── 13-concurrency/
├── 14-channel/
├── 15-generics/
├── 16-context/
└── HTTP/
```

Her session'ın amacı ilgili kavramı izole bir ortamda anlayarak daha büyük sistemlerde kullanılmadan önce temel çalışma mantığını öğrenmektir.

### `exercises/`

Öğrenilen konuların bağımsız olarak tekrar uygulanması için hazırlanan alıştırmaları içerir.

Amaç örnek kodları doğrudan kopyalamak yerine problemi anlayıp çözümü yeniden oluşturabilmektir.

Exercise'larda farklı konular birlikte kullanılabilir:

```text
Functions
    +
Slices
    +
Maps
    +
Structs
    +
Methods
    +
Interfaces
    +
Errors
```

Böylece daha önce öğrenilen kavramların unutulmaması ve birlikte kullanılabilmesi hedeflenir.

### `notes/`

Öğrenme sürecinde oluşan tekrar materyallerini içerir.

```text
concepts.md   → Öğrenilen kavramların kısa açıklamaları
flashcards.md → Hızlı tekrar kartları
mistakes.md   → Öğrenme sırasında yapılan gerçek hatalar
questions.md  → Tekrar ve kontrol soruları
```

Bu dosyalar yalnızca teorik bilgi içermek yerine öğrenme sürecinde gerçekten karşılaşılan noktaları kaydetmek için kullanılır.

### `progress/`

Repo içerisindeki öğrenme sürecinin takibini sağlar.

```text
roadmap.md  → Baştan sona öğrenme planı
progress.md → Tamamlanan ve güncel aşamalar
```

---

## Öğrenme Yöntemi

Her konu aşağıdaki öğrenme döngüsüyle ele alınır:

```text
Kavramı Öğren
      ↓
Neden Kullanıldığını Anla
      ↓
Küçük Bir Örnek Yaz
      ↓
Kodu Çalıştır
      ↓
Çıktıyı İncele
      ↓
Bağımsız Exercise Yap
      ↓
Önceki Konularla Birleştir
      ↓
Hataları Analiz Et
      ↓
Notları Güncelle
      ↓
Progress'i Güncelle
      ↓
Commit & Push
```

Temel amaç kod ezberlemek değil, yazılan kodun neden ve nasıl çalıştığını anlamaktır.

Özellikle ilerleyen konularda önceki konuların tekrar kullanılması hedeflenmektedir.

Örneğin:

```text
Struct
  ↓
Method
  ↓
Pointer Receiver
  ↓
Interface
  ↓
Error Handling
```

şeklinde kavramlar birbirinin üzerine inşa edilir.

---

## Şu Ana Kadar Öğrenilenler

```text
Go Ortamı ve Temeller       ✅
Variables & Data Types      ✅
Control Flow                ✅
Functions                   ✅
Arrays & Slices             ✅
Maps                        ✅
Structs                     ✅
Methods                     ✅
Pointers                    ✅
Interfaces                  ✅
Error Handling              ✅
Packages & Modules          ✅
Concurrency                 ✅
Channels                    ✅
Generics                    ✅
Context                     ✅
HTTP                        🔄
```

---

## Öğrenilen Temel Kavramlar

### Go Temelleri

```text
Go kurulumu
go version
go run
go build
go fmt
go.mod
package main
import
func main
```

### Variables & Types

```text
var
:=
const

string
int
float64
bool

type conversion
strconv
fmt.Printf
```

### Control Flow

```text
if / else
switch
for
range
break
continue
```

### Functions

```text
Function declaration
Parameters
Return values
Multiple return values
Function scope
```

Function'ların yalnızca kod tekrarını azaltmak için değil, program içerisindeki sorumlulukları birbirinden ayırmak için kullanılabileceği uygulamalı olarak çalışıldı.

### Arrays & Slices

```text
Arrays
Slices
append
len
cap
Slicing
range
Filtering
```

Slice içerisindeki verileri işleme, filtreleme ve yeni slice üretme üzerine farklı exercise'lar yapıldı.

### Maps

```text
map[key]value
make
Key / Value
Lookup
value, ok
Update
Delete
range
```

Map içerisindeki verilere doğrudan key üzerinden erişim ve `value, ok` pattern'i çalışıldı.

### Structs

```text
Custom data structures
Struct fields
Struct initialization
Slice of structs
Struct filtering
Struct searching
Zero value
```

Birden fazla ilişkili veriyi tek bir type altında modelleme mantığı uygulandı.

### Methods

```text
Receiver
Value receiver
Pointer receiver
Struct methods
State modification
```

Özellikle value receiver ile pointer receiver arasındaki fark üzerinde duruldu.

```text
Value Receiver
→ değeri okumak

Pointer Receiver
→ gerçek struct state'ini değiştirmek
```

### Pointers

```text
Memory address
&
*
Pointer dereferencing
Pointer parameters
Pointer receivers
```

Pointer kullanımının yalnızca syntax olarak değil, bir değerin gerçek state'ini değiştirmek açısından neden gerekli olduğu incelendi.

### Interfaces

```text
Interface declaration
Method sets
Implicit implementation
Interface parameters
Polymorphism
```

Bir interface'in concrete type'a method vermediği, bunun yerine ilgili type'ın sağlaması gereken davranışları tanımladığı öğrenildi.

Örneğin:

```text
Notifier
├── Send()
└── GetType()
```

Bir type'ın `Notifier` olarak kullanılabilmesi için interface içerisinde tanımlanan methodların tamamını sağlaması gerektiği uygulandı.

### Error Handling

```text
error
errors.New
fmt.Errorf
nil
value, error
err != nil
Error propagation
Error wrapping
%w
```

Go'da exception tabanlı bir yaklaşım yerine hataların açık şekilde return value olarak taşınması uygulandı.

Temel akış:

```text
Başarılı işlem
→ value, nil

Başarısız işlem
→ zero/default value, error
```

Ayrıca alt katmandaki bir error'ın üst katmana ek bağlam verilerek taşınması çalışıldı:

```text
SellProduct()
      ↓
error
      ↓
ProcessOrder()
      ↓
fmt.Errorf("order processing failed: %w", err)
      ↓
main()
```

### Packages & Modules

```text
Packages
Imports
Custom packages
Exported identifiers
go.mod
Modules
Package organization
```

Kodun tek bir dosyada tutulması yerine farklı sorumlulukların package'lara ayrılması incelendi.

### Concurrency

```text
Goroutines
go keyword
Concurrent execution
sync.WaitGroup
sync.Mutex
Shared state
```

Go'nun concurrency modeli ve birden fazla işin eş zamanlı ilerletilmesi üzerine temel örnekler uygulandı.

### Channels

```text
Channels
make(chan T)
Send
Receive
Buffered channels
Unbuffered channels
Channel closing
range over channel
```

Goroutine'ler arasında güvenli veri iletişimi için channel kullanımı çalışıldı.

### Generics

```text
Type parameters
Generic functions
Generic types
Constraints
any
comparable
Generic structs
Generic maps
```

Aynı algoritmanın farklı veri tipleriyle type-safe şekilde kullanılabilmesi incelendi.

### Context

```text
context.Context
context.Background
context.WithCancel
context.WithTimeout
context.WithDeadline
cancel()
Done()
Err()
Context propagation
```

Context'in uzun süren veya birbirine bağlı işlemlerin yaşam döngüsünü kontrol etmek için nasıl kullanılabileceği öğrenildi.

---

## Güncel Aşama

### Aşama 17 — HTTP

Go fundamentals aşamasının son bölümünde HTTP'nin temel çalışma mantığı ve Go'nun `net/http` package'ı incelenmektedir.

Şu ana kadar:

```text
HTTP
  ↓
Client / Server
  ↓
Request / Response
  ↓
HTTP Methods
  ↓
Status Codes
  ↓
Headers
  ↓
Body
  ↓
net/http
  ↓
HTTP Server
  ↓
Handler
```

konularına giriş yapıldı.

Go ile temel bir HTTP server oluşturuldu.

```text
Client
   │
   │ GET /hello
   ▼
Go HTTP Server :8080
   │
   ▼
Handler
   │
   ▼
HTTP Response
   │
   ▼
Client
```

Bu aşamada temel seviyede:

```text
Routing
Multiple endpoints
Request
Response
JSON
```

konuları uygulanarak HTTP'nin Go tarafındaki çalışma mantığı pekiştirilecektir.

HTTP burada kapsamlı bir backend uygulaması geliştirmek için değil, Go ile gerçek bir network uygulamasına geçişi anlamak için ele alınmaktadır.

---

## Bu Reponun Kapsamı

Bu repo özellikle Go dilinin temel ve orta seviye kavramlarını öğrenmeye odaklanmaktadır.

Repo kapsamı:

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
Packages
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

Database, kapsamlı REST API geliştirme, backend architecture, Docker, Kubernetes, observability ve load testing gibi konular bu reponun ana kapsamına dahil değildir.

Bu konular Go fundamentals tamamlandıktan sonra ayrı bir gerçek backend ve cloud-native proje içerisinde uygulanacaktır.

---

## Sonraki Adım

Bu repo tamamlandıktan sonra öğrenilen kavramlar ayrı bir Go cloud-native projesinde birleştirilecektir.

Genel geçiş:

```text
go-backend-learning
        │
        │
        ▼
Go Fundamentals Complete
        │
        ▼
Gerçek Go Backend Projesi
        │
        ▼
HTTP / REST API
        │
        ▼
Backend Architecture
        │
        ▼
Database
        │
        ▼
Testing
        │
        ▼
Configuration & Logging
        │
        ▼
Service-to-Service Communication
        │
        ▼
Docker
        │
        ▼
Docker Compose
        │
        ▼
Load Testing
        │
        ▼
Kubernetes
        │
        ▼
Observability & Scaling
```

Bu ikinci aşamada amaç Go kavramlarını ayrı ayrı öğrenmek değil, daha önce öğrenilmiş kavramların gerçek bir backend sisteminde neden ve nerede kullanıldığını görmek olacaktır.

---

## İlerleme Takibi

Detaylı öğrenme planı:

```text
progress/roadmap.md
```

Güncel ilerleme durumu:

```text
progress/progress.md
```

---

## Temel Hedef

Bu repo tamamlandığında hedef yalnızca Go syntax'ını bilmek değildir.

Amaç:

```text
Bir problemi analiz edebilmek
        ↓
Uygun veri yapısını seçebilmek
        ↓
Function'lara ayırabilmek
        ↓
Struct ile modelleyebilmek
        ↓
Method ve pointer receiver kullanabilmek
        ↓
Interface ile davranış tanımlayabilmek
        ↓
Error'ları doğru şekilde yönetebilmek
        ↓
Package'larla kodu organize edebilmek
        ↓
Concurrency ve channels mantığını anlayabilmek
        ↓
Generics kullanabilmek
        ↓
Context ile işlem yaşam döngüsünü yönetebilmek
        ↓
Temel HTTP uygulaması oluşturabilmek
```

seviyesine ulaşmaktır.

Bu temel tamamlandıktan sonra bir sonraki hedef, öğrenilen Go bilgisini gerçek bir backend ve cloud-native sistem geliştirirken kullanmaktır.