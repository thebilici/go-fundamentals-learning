# Go Backend Learning — Progress

**Son Güncelleme:** 2026-08-24  
**Güncel Aşama:** Aşama 11 — Error Handling

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
Error Handling              ← GÜNCEL
        ↓
Packages & Modules
        ↓
Concurrency
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

## Tamamlanan Aşamalar

### Aşama 8 — Methods ✅

- Method
- Receiver
- Value Receiver
- Method Parameters
- Method Return Values

### Aşama 9 — Pointers ✅

- Memory Address
- `&` ve `*`
- Dereferencing
- Pointer Types
- Pointer + Function
- Pointer Receiver
- Value Receiver vs Pointer Receiver

### Aşama 10 — Interfaces ✅

- Interface tanımlama
- Method Signature
- Implicit Implementation
- Birden fazla method içeren interface
- Interface Parameter
- Interface Variable
- Concrete Type
- Polymorphism

### Yapılan Pratik

`exercises/basics/interfaces/` altında:

- `Notifier` interface oluşturuldu.
- `EmailNotifier` ve `SMSNotifier` oluşturuldu.
- `send()` method'ları implement edildi.
- Aynı interface üzerinden farklı concrete type'lar kullanıldı.
- Interface function parameter olarak kullanıldı.
- Interface variable oluşturuldu.
- Polymorphism uygulandı.

## Güncel Aşama

### Aşama 11 — Error Handling 🔄

Sıradaki konular:

- `error` nedir?
- Go'da hata yönetimi mantığı
- `error` return etmek
- `nil` kontrolü
- `if err != nil`
- `errors.New()`
- `fmt.Errorf()`
- Multiple return ile `(value, error)`
- Error propagation