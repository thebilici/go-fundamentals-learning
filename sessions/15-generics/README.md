# Session 15 — Generics

## Generics Nedir?

Generics, aynı kodun farklı veri tipleriyle type-safe şekilde çalışmasını sağlar.

Normalde farklı tipler için ayrı function yazmak gerekebilir:

```go
func addInt(a int, b int) int {
	return a + b
}

func addFloat(a float64, b float64) float64 {
	return a + b
}
```

Generics ile ortak bir yapı oluşturulabilir:

```go
func add[T int | float64](a T, b T) T {
	return a + b
}
```

---

## Type Parameter

```go
func printValue[T any](value T) {
	fmt.Println(value)
}
```

Buradaki:

```text
T
```

bir **Type Parameter**'dır.

Function içerisinde kullanılacak tipi temsil eder.

---

## Type Argument

Generic yapı kullanılırken verilen gerçek tipe Type Argument denir.

```go
Response[string]
Response[int]
```

Burada:

```text
string → Type Argument
int    → Type Argument
```

Go generic function çağrılarında çoğu zaman tipi kendisi çıkarabilir:

```go
add(10, 20)
```

Burada:

```text
T = int
```

olduğunu Go otomatik olarak anlar.

---

## Type Constraint

Constraint, `T` için hangi tiplerin kullanılabileceğini belirler.

```go
func add[T int | float64](a T, b T) T {
	return a + b
}
```

Burada:

```text
T → int veya float64 olabilir
```

---

## any

`any`, type parameter'ın herhangi bir tip olabileceğini belirtir.

```go
func getFirst[T any](arr []T) T {
	return arr[0]
}
```

Örneğin:

```go
getFirst([]int{1, 2, 3})
getFirst([]string{"Go", "Java"})
```

aynı function ile çalışabilir.

---

## comparable

`comparable`, `==` ve `!=` ile karşılaştırılabilen tipleri kabul eder.

```go
func contains[T comparable](arr []T, target T) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}

	return false
}
```

Burada `item == target` kullanıldığı için `comparable` gerekir.

---

## Birden Fazla Type Parameter

Bir generic yapı birden fazla type parameter kullanabilir.

```go
func printPair[K any, V any](key K, value V) {
	fmt.Println(key, value)
}
```

Burada:

```text
K → birinci tip
V → ikinci tip
```

Örneğin:

```go
printPair("Age", 22)
```

için:

```text
K = string
V = int
```

---

## Generic Struct

Struct'lar da generic olabilir.

```go
type Response[T any] struct {
	Data    T
	Success bool
}
```

Kullanım:

```go
user := Response[string]{
	Data:    "Fatih",
	Success: true,
}

score := Response[int]{
	Data:    100,
	Success: true,
}
```

Aynı struct farklı tiplerde veri taşıyabilir.

---

## Custom Constraint

Constraint ayrı bir interface olarak tanımlanabilir.

```go
type Number interface {
	int | float64
}
```

Daha sonra:

```go
func sum[T Number](a T, b T) T {
	return a + b
}
```

kullanılabilir.

---

## Generic Map Mantığı

Normal bir map:

```go
map[string]int
```

şu tipleri kullanır:

```text
Key   → string
Value → int
```

Generic hale getirildiğinde:

```go
map[K]V
```

olur.

```text
K → Key tipi
V → Value tipi
```

Örneğin:

```go
func GetValue[K comparable, V any](data map[K]V, key K) (V, bool) {
	value, exists := data[key]
	return value, exists
}
```

`map[string]int` ile kullanılırsa:

```text
K = string
V = int
```

`map[int]string` ile kullanılırsa:

```text
K = int
V = string
```

---

## Yapılan Pratikler

Bu aşamada:

```text
getFirst[T any]
        ↓
Generic Slice

contains[T comparable]
        ↓
Generic karşılaştırma

printPair[K, V]
        ↓
Birden fazla Type Parameter

Response[T any]
        ↓
Generic Struct

Number interface
        ↓
Custom Constraint

GetValue[K, V]
        ↓
Generic Map
```

örnekleri uygulandı.

---

## Temel Özet

```text
T
→ Type Parameter

any
→ Herhangi bir tip

comparable
→ Karşılaştırılabilir tip

K
→ Genellikle Key tipi

V
→ Genellikle Value tipi

[T int | float64]
→ Belirli tiplerle sınırlandırma

Response[T]
→ Generic Struct

map[K]V
→ Generic Map yapısı
```

Generics'in temel amacı:

> Aynı algoritmayı veya veri yapısını farklı tiplerle tekrar kod yazmadan, type-safe şekilde kullanabilmektir.