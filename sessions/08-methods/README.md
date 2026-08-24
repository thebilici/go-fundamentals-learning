# Aşama 8 — Methods

Bu aşamada Go'da method kavramı, receiver yapısı ve struct'larla method kullanımının mantığı öğrenildi.

## Method Nedir?

Method, belirli bir type ile ilişkilendirilmiş function'dır.

```go
func (u User) getName() string {
	return u.Name
}
```

Burada `getName`, `User` type'ına ait bir method'dur.

## Receiver

Method ile type arasındaki bağlantıyı receiver kurar.

```go
func (u User) getName() string
```

```text
u    → Receiver variable
User → Receiver type
```

Method şu şekilde çağrılır:

```go
user.getName()
```

## Function ve Method Farkı

Function:

```go
printUser(user)
```

Method:

```go
user.getName()
```

Method, davranışı doğrudan ilgili type ile ilişkilendirir.

## Method Parameter

Method kendi parameter'larını da alabilir:

```go
func (u User) canAccess(minAge int) bool {
	return u.Age >= minAge
}
```

Çağırma:

```go
user.canAccess(18)
```

## Method Return Value

Method bir değer return edebilir:

```go
func (u User) isAdult() bool {
	return u.Age >= 18
}
```

## Value Receiver

```go
func (u User) getName() string
```

şeklindeki receiver bir value receiver'dır.

Value receiver genellikle struct'ın bilgisini okumak veya hesaplama yapmak için uygundur.

Struct üzerinde kalıcı değişiklik yapmak istediğimizde pointer receiver konusu devreye girer.

## Yapılan Pratik

`exercises/basics/methods/` altında:

- `getName()`
- `isAdult()`
- `canAccess()`
- `greet()`

method'ları oluşturuldu.

Method'larda:

- Receiver
- Parameter
- Return Value
- Struct field erişimi

uygulandı.

## Aşama Sonu Özeti

```text
User
 ↓
user
 ↓
user.method()
 ↓
Receiver
 ↓
Struct bilgileri üzerinde işlem
```

Method'lar, davranışları ilgili type ile ilişkilendirmemizi sağlar. 