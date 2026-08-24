# Go Backend Learning — Progress

**Son Güncelleme:** 2026-08-24  
**Güncel Aşama:** Aşama 10 — Interfaces

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
Interfaces                  ← GÜNCEL
        ↓
Error Handling
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

### Aşama 1 — Go Ortamı ve Temeller ✅

- Go Module ve `go.mod`
- Package yapısı
- `go run`, `go build`, `go fmt`

### Aşama 2 — Variables & Data Types ✅

- Variables ve Constants
- Temel Data Type'lar
- Type Inference
- Zero Values
- Type Conversion

### Aşama 3 — Control Flow ✅

- `if / else`
- `switch`
- `for`
- `break / continue`
- `range`

### Aşama 4 — Functions ✅

- Parameters ve Arguments
- Return Values
- Multiple Return Values
- Function Scope

### Aşama 5 — Arrays & Slices ✅

- Arrays ve Slices
- `append`
- `len` / `cap`
- Slicing
- `range`

### Aşama 6 — Maps ✅

- Key / Value
- Ekleme ve güncelleme
- `value, ok`
- `delete`
- `range`

### Aşama 7 — Structs ✅

- Struct ve Fields
- Struct + Function
- Slice of Structs
- Exported / Unexported Fields

### Aşama 8 — Methods ✅

- Method
- Receiver
- Value Receiver
- Method Parameters
- Method Return Values

### Aşama 9 — Pointers ✅

- Memory Address
- `&` Operator
- `*` Operator
- Dereferencing
- Pointer Types
- Pointer + Function
- Pointer Receiver
- Value Receiver vs Pointer Receiver

### Yapılan Pratik

`exercises/basics/pointers/` altında:

- Variable address'i alındı.
- Pointer oluşturuldu.
- Dereferencing yapıldı.
- Pointer üzerinden orijinal değer değiştirildi.
- Pointer function parameter olarak kullanıldı.
- `User` üzerinde pointer receiver ile field değiştirildi.

## Güncel Aşama

### Aşama 10 — Interfaces 🔄

Sıradaki konular:

- Interface nedir?
- Interface tanımlama
- Method Set
- Interface'i implement etmek
- Implicit Implementation
- Interface'i function parameter olarak kullanmak
- Birden fazla struct + aynı interface