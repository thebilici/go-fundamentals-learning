# Aşama 4 — Functions

Bu aşamada Go'da function oluşturma, function'lara veri gönderme ve function'lardan sonuç alma öğrenildi.

## Function

Belirli bir işi yapan tekrar kullanılabilir kod bloğudur.

```go
func greet() {
	fmt.Println("Hello")
}
```

Function çağrılarak çalıştırılır:

```go
greet()
```

## Parameters ve Arguments

Function'ın aldığı değişkenlere **parameter**, çağrılırken gönderilen gerçek değerlere **argument** denir.

```go
func greet(name string) {
	fmt.Println("Hello", name)
}

greet("Fatih")
```

```text
name    → Parameter
"Fatih" → Argument
```

## Return Value

Function bir sonuç döndürebilir:

```go
func add(a, b int) int {
	return a + b
}
```

Kullanımı:

```go
result := add(5, 10)
```

Burada `int`, function'ın return type'ıdır.

## Multiple Return Values

Go function'ları birden fazla değer döndürebilir:

```go
func getUser() (string, int) {
	return "Fatih", 22
}
```

Değerler ayrı variable'lara alınabilir:

```go
name, age := getUser()
```

Bu yapı özellikle Go'nun error handling yaklaşımında sık kullanılır:

```go
result, err := someFunction()
```

## Named Return Values

Return değerleri isimlendirilebilir:

```go
func add(a, b int) (result int) {
	result = a + b
	return
}
```

Ancak okunabilirlik için çoğu durumda explicit return tercih edilebilir.

## Function Scope

Function içerisinde oluşturulan variable'lar kendi scope'larında bulunur.

```go
func calculate() int {
	result := 10 + 20
	return result
}
```

`result` değişkenine başka bir function içerisinden doğrudan erişilemez.

## Yapılan Pratik

`exercises/basics/functions/` altında:

- `greet`
- `add`
- `checkAge`
- `getUser`

function'ları oluşturuldu.

Bu pratikte parameter, argument, return value, multiple return values ve control flow birlikte kullanıldı.

## Aşama Sonu Zihinsel Model

```text
Arguments
    ↓
Parameters
    ↓
Function
    ↓
İşlem
    ↓
Return
    ↓
Sonuç
```