# Aşama 9 — Pointers

Bu aşamada Go'da pointer, memory address ve pointer kullanarak orijinal değerleri değiştirme mantığı öğrenildi.

## Pointer Nedir?

Pointer, başka bir variable'ın memory address'ini tutar.

```go
age := 22
agePointer := &age
```

```text
age        → 22
agePointer → age'in memory address'i
```

## `&` Operator

Bir variable'ın memory address'ini almak için kullanılır.

```go
&age
```

## `*` Operator

Pointer'ın gösterdiği adresteki değere erişmek için kullanılır.

```go
*agePointer
```

```text
agePointer  → Address
*agePointer → Address'teki value
```

Bu işleme dereferencing denir.

## Pointer ile Değer Değiştirme

```go
*agePointer = 25
```

Pointer üzerinden address'teki değer değiştirildiği için orijinal `age` de değişir.

```text
age → 22

*agePointer = 25

age → 25
```

## Value Kopyası ve Pointer Farkı

```go
agePointerValue := *agePointer
```

Burada değer ayrı bir variable'a kopyalanır.

```text
agePointerValue → Değerin kopyası
*agePointer     → Orijinal değere erişim
```

Bu nedenle daha sonra:

```go
*agePointer = 25
```

yapılması `agePointerValue` değerini değiştirmez.

## Pointer Type

```go
agePointer := &age
```

Eğer `age` bir `int` ise:

```text
age        → int
agePointer → *int
```

`*int`, bir `int` değerine işaret eden pointer type'dır.

## Pointer + Function

Normal parameter:

```go
func changeAge(age int) {
	age = 30
}
```

orijinal değeri değiştirmez.

Pointer parameter:

```go
func changeAge(age *int) {
	*age = 30
}
```

Çağırırken:

```go
changeAge(&age)
```

orijinal `age` değiştirilebilir.

## Pointer Receiver

Struct üzerinde kalıcı değişiklik yapmak için pointer receiver kullanılabilir.

```go
func (u *User) changeName(newName string) {
	u.Name = newName
}
```

Çağırma:

```go
user.changeName("Ahmet")
```

Bu işlem orijinal `user.Name` değerini değiştirir.

## Value Receiver vs Pointer Receiver

```text
(u User)
↓
Value Receiver
Orijinal struct'ı değiştirmek için uygun değildir.


(u *User)
↓
Pointer Receiver
Orijinal struct üzerinde değişiklik yapabilir.
```

## Yapılan Pratik

`exercises/basics/pointers/` altında:

- Memory address görüntülendi.
- `&` kullanıldı.
- `*` ile dereferencing yapıldı.
- Pointer üzerinden değer değiştirildi.
- Value kopyası ile pointer arasındaki fark incelendi.
- Pointer function parameter olarak kullanıldı.
- Struct üzerinde pointer receiver kullanıldı.

## Aşama Sonu Özeti

```text
age := 22
   │
   │ &age
   ▼
Memory Address
   │
   │ *pointer
   ▼
22


& → Address'i al
* → Address'teki değere eriş


*pointer = 25
→ Orijinal değeri değiştir


func changeAge(age *int)
→ Pointer parameter


func (u *User) changeName()
→ Pointer receiver
```