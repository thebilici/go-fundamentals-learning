# Go Backend Learning

Go programlama dilini temelden başlayarak backend ve cloud-native geliştirme seviyesine kadar öğrenmek amacıyla oluşturulmuş uygulamalı bir öğrenme reposudur.

Bu reponun amacı yalnızca Go syntax'ını öğrenmek değil; Go'nun type sistemi, veri yapıları, error handling yaklaşımı, concurrency modeli, context yapısı ve backend geliştirme mantığını uygulayarak öğrenmektir.

Öğrenilen konular ilerleyen aşamalarda gerçek bir Go backend projesinde birleştirilerek Docker, Kubernetes, observability ve load testing seviyesine taşınacaktır.

---

## Amaç

Öğrenme süreci aşağıdaki genel yolu takip eder:

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
├── project/
│   ├── service-a/
│   └── service-b/
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

Yeni Go kavramlarını öğrenirken oluşturulan küçük örnekleri ve deneyleri içerir.

Her aşamada konu önce burada öğrenilir ve temel çalışma mantığı anlaşılır.

### `exercises/`

Öğrenilen konuların bağımsız olarak uygulanması için hazırlanan alıştırmaları içerir.

Amaç örnek kodu doğrudan kopyalamadan, öğrenilen kavramları yeniden uygulayabilmektir.

### `project/`

Öğrenilen Go ve backend kavramlarının gerçek bir cloud-native uygulamada birleştirileceği ana proje alanıdır.

Proje ilerleyen aşamalarda iki servisli bir yapıya dönüşecektir:

```text
Client
  ↓
Service A
  ↓
Service B
```

Bu proje üzerinde aşamalı olarak:

```text
Go HTTP Server
      ↓
REST API
      ↓
Backend Mimarisi
      ↓
Servisler Arası İletişim
      ↓
Database
      ↓
Testing
      ↓
Configuration & Logging
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
```

konuları uygulanacaktır.

### `notes/`

Öğrenme sürecinde oluşan tekrar materyallerini içerir.

```text
concepts.md   → Kavram özetleri
flashcards.md → Hızlı tekrar kartları
mistakes.md   → Öğrenme sırasında yapılan gerçek hatalar
questions.md  → Tekrar ve kontrol soruları
```

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
Kodu Çalıştır ve İncele
      ↓
Bağımsız Exercise Yap
      ↓
Ana Projede Uygula
      ↓
Notları Güncelle
      ↓
Progress'i Güncelle
      ↓
Commit & Push
```

Temel amaç kod ezberlemek değil, yazılan kodun neden ve nasıl çalıştığını anlamaktır.

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

Bu aşamalarda temel olarak:

- Go module ve package yapısı
- Variables ve type sistemi
- Control flow
- Functions ve multiple return values
- Arrays, slices ve maps
- Struct ve custom type oluşturma
- Methods ve receiver
- Pointers ve pointer receiver
- Interfaces ve polymorphism
- Error handling, propagation ve wrapping
- Custom packages ve modules
- Goroutines ve concurrency
- WaitGroup ve Mutex
- Channels ve goroutine iletişimi
- Generic functions ve generic types
- Type parameters ve constraints
- `any` ve `comparable`
- Generic structs ve maps
- `context.Context`
- Cancellation
- Timeout ve deadline
- Context propagation

konuları uygulamalı olarak öğrenildi.

---

## Güncel Aşama

**Aşama 17 — HTTP**

Bu aşamada HTTP'nin temel çalışma mantığı ve Go'nun `net/http` package'ı öğrenilmektedir.

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

Go ile ilk HTTP server oluşturuldu:

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

Bu aşamada devamında routing, birden fazla endpoint ve request/response işlemleri uygulanacaktır.

---

## Ana Proje

HTTP aşamasından itibaren öğrenilen backend konuları yalnızca küçük örneklerde bırakılmayacak, `project/` altında gerçek projeye uygulanacaktır.

Hedef mimari:

```text
                    Client
                      │
                      ▼
                  Service A
                  Go Backend
                      │
                      │ HTTP
                      ▼
                  Service B
                  Go Backend
```

İlerleyen aşamalarda:

```text
                     Client
                       │
                       ▼
                   Service A
                       │
                       ▼
                   Service B
                       │
                       ▼
                    Database

                       +
                       
                 Docker Compose

                       +

                      k6
                 Load Testing

                       +

                  Kubernetes
                       │
              ┌────────┴────────┐
              │                 │
           Metrics            HPA
```

yapısına ulaşılması hedeflenmektedir.

Load testing aşamasında CPU yükü oluşturabilen endpoint'ler üzerinden k6 ile stres testi yapılacak ve Kubernetes ortamında resource kullanımı ile horizontal scaling davranışı gözlemlenecektir.

---

## İlerleme Takibi

Detaylı roadmap:

```text
progress/roadmap.md
```

Güncel ilerleme durumu:

```text
progress/progress.md
```

---

## Temel Hedef

Repo tamamlandığında amaç yalnızca Go syntax'ını bilen biri olmak değil;

```text
Go
+
Backend Development
+
REST API
+
Database
+
Testing
+
Docker
+
Docker Compose
+
Kubernetes
+
Observability
+
Load Testing
```

konularını birlikte kullanabilen ve oluşturduğu sistemin:

```text
Kod
 ↓
HTTP
 ↓
Backend
 ↓
Container
 ↓
Orchestration
 ↓
Metrics
 ↓
Scaling
```

zincirinin nasıl çalıştığını anlayan sağlam bir backend ve cloud-native geliştirme temeline sahip olmaktır.