# Aşama 7 — Structs

Bu aşamada Go'da kendi veri tiplerimizi oluşturmak için kullanılan `struct` yapısı öğrenildi.

## Struct Nedir?

Struct, birbiriyle ilişkili farklı type'lardaki verileri tek bir yapı altında toplar.

```go
type User struct {
	Name     string
	Age      int
	IsActive bool
}
```

Burada `User` bizim oluşturduğumuz yeni bir type'tır.

## Field

Struct içerisindeki değerlere field denir.

```text
Name     → string field
Age      → int field
IsActive → bool field
```

## Struct Oluşturma

```go
user := User{
	Name:     "Fatih",
	Age:      22,
	IsActive: true,
}
```

Field'lara `.` ile erişilir:

```go
user.Name
user.Age
```

## Struct + Function

Struct, function parameter'ı veya return type olarak kullanılabilir.

```go
func createUser(name string, age int, isActive bool) User {
	return User{
		Name:     name,
		Age:      age,
		IsActive: isActive,
	}
}
```

Burada:

```text
Name: name

Sol  → Struct field
Sağ  → Function parameter'ı
```

## Slice of Structs

Birden fazla `User`, slice içerisinde tutulabilir:

```go
users := []User{
	user,
	user2,
	user3,
}
```

`range` ile tek tek dolaşılabilir:

```go
for index, user := range users {
	fmt.Println(index, user.Name)
}
```

```text
users → Tüm User slice'ı
user  → Her döngüde alınan tek User
index → O User'ın slice içerisindeki index'i
```

## Exported Field

Büyük harfle başlayan field başka package'lerden erişilebilir:

```go
Age int
```

Küçük harfle başlayan field unexported'dır:

```go
age int
```

## Yapılan Pratik

`exercises/basics/structs/` altında:

- `User` struct oluşturuldu.
- `createUser()` function'ı yazıldı.
- Birden fazla User oluşturuldu.
- `[]User` kullanıldı.
- `range` ile kullanıcılar dolaşıldı.
- Aktif kullanıcılar filtrelendi.
- `append()` ile yeni kullanıcı eklendi.
- `len()` ile kullanıcı sayısı bulundu.

## Aşama Sonu Özeti

```text
Struct → Kendi type'ımız
Field  → Struct içerisindeki değerler

User
 ↓
User{...}
 ↓
[]User
 ↓
range
 ↓
Tek tek User'lar
```