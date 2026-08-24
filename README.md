# Go Backend Learning

Go programlama dilini temelden başlayarak backend ve cloud-native geliştirme seviyesine kadar öğrenmek amacıyla oluşturulmuş uygulamalı bir öğrenme reposudur.

Bu reponun amacı yalnızca Go syntax'ını öğrenmek değil; Go'nun veri yapıları, type sistemi, error handling yaklaşımı, concurrency modeli ve backend geliştirme mantığını uygulayarak öğrenmektir.

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
├── exercises/
├── project/
│   └── service-a/
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

Her aşamada konu önce burada öğrenilir ve test edilir.

### `exercises/`

Öğrenilen konuların bağımsız olarak uygulanması için hazırlanan alıştırmaları içerir.

Amaç konuyu örnekten kopyalamadan tekrar uygulayabilmektir.

### `project/`

Öğrenilen Go kavramlarının gerçek bir backend uygulamasında birleştirileceği ana proje alanıdır.

İlerleyen aşamalarda HTTP, REST API, database, testing, Docker ve Kubernetes konuları burada uygulanacaktır.

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
Notları Güncelle
      ↓
Progress'i Güncelle
      ↓
Commit Et
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
Packages & Modules          🔄
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

konuları uygulamalı olarak öğrenildi.

---

## Güncel Aşama

**Aşama 12 — Packages & Modules**

Bu aşamada Go projelerinin birden fazla package'a nasıl ayrıldığı ve package'lerin birbirleriyle nasıl çalıştığı öğrenilecektir.

Özellikle:

```text
Module
  ↓
Packages
  ↓
Go Files
```

ilişkisi üzerinde durulacaktır.

Öğrenilecek temel konular:

- Custom package oluşturma
- `package` kullanımı
- `import`
- Exported / Unexported identifiers
- Package path
- Module path
- `go.mod`
- Package ve Module farkı
- Package'ler arası kullanım

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
Kubernetes
```

konularını birlikte kullanabilen ve yazdığı kodun arkasındaki mantığı anlayan bir backend geliştirme temeline sahip olmaktır.