# Go Backend Learning — Progress

**Son Güncelleme:** 2026-08-24  
**Güncel Aşama:** Aşama 12 — Packages & Modules

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
Packages & Modules          ← GÜNCEL
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

### Aşama 9 — Pointers ✅

- Memory Address
- `&` ve `*`
- Dereferencing
- Pointer Types
- Pointer + Function
- Pointer Receiver
- Value Receiver vs Pointer Receiver

### Aşama 10 — Interfaces ✅

- Interface
- Method Signature
- Implicit Implementation
- Interface Parameter
- Interface Variable
- Concrete Type
- Polymorphism

### Aşama 11 — Error Handling ✅

- `error` type
- `(value, error)` dönüş yapısı
- `nil`
- `err != nil`
- `errors.New()`
- `fmt.Errorf()`
- Error Propagation
- Error Wrapping
- `%w`
- Error Handling vs Propagation
- Struct Zero Value (`User{}`)

### Yapılan Pratik

`exercises/basics/error-handling/` altında:

- `validateAge()` oluşturuldu.
- Yaş doğrulama kuralları yazıldı.
- `errors.New()` ile error üretildi.
- `createUser()` içinde `(User, error)` kullanıldı.
- Error üst function'a propagate edildi.
- `fmt.Errorf("%w")` ile error wrapping yapıldı.
- Hata durumunda `User{}` döndürüldü.

## Güncel Aşama

### Aşama 12 — Packages & Modules 🔄

Sıradaki konular:

- Package nedir?
- `package main`
- Custom package oluşturma
- `import`
- Exported / Unexported identifiers
- Package path
- Module nedir?
- `go.mod`
- Module path
- Package ve Module farkı
- Bir package'i başka package'dan kullanma