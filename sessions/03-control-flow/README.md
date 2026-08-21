# Aşama 3 — Control Flow

Bu aşamada Go programının koşullara göre karar vermesi ve kod bloklarını tekrar çalıştırması öğrenildi.

## Comparison Operators

Karşılaştırmalar sonucunda `true` veya `false` üretilir.

```text
==   Eşit
!=   Eşit değil
>    Büyük
<    Küçük
>=   Büyük veya eşit
<=   Küçük veya eşit
```

Örnek:

```go
age := 22
result := age >= 18
```

## Logical Operators

Birden fazla koşulu birleştirmek için kullanılır.

```text
&&   AND
||   OR
!    NOT
```

Örnek:

```go
if age >= 18 && isActive {
	fmt.Println("Giriş yapılabilir")
}
```

## if / else

Koşula göre farklı kodların çalışmasını sağlar.

```go
if age >= 18 {
	fmt.Println("Reşit")
} else {
	fmt.Println("Reşit değil")
}
```

Birden fazla koşul için `else if` kullanılabilir.

## switch

Bir değeri farklı seçeneklerle karşılaştırmak için kullanılır.

```go
switch role {
case "admin":
	fmt.Println("Admin")
case "user":
	fmt.Println("User")
default:
	fmt.Println("Bilinmeyen rol")
}
```

## for

Go'nun temel döngü yapısıdır.

```go
for i := 1; i <= 5; i++ {
	fmt.Println(i)
}
```

Go'da ayrı bir `while` keyword'ü yoktur. `for` bu amaçla da kullanılabilir.

```go
count := 1

for count <= 5 {
	count++
}
```

## break ve continue

`break` döngüyü tamamen sonlandırır.

```go
if i == 8 {
	break
}
```

`continue` sadece mevcut turu atlar.

```go
if i == 4 {
	continue
}
```

## range

Collection içerisindeki elemanları dolaşmak için kullanılır.

```go
languages := []string{"Go", "Python", "JavaScript"}

for index, value := range languages {
	fmt.Println(index, value)
}
```

`range` konusu Arrays & Slices aşamasında daha detaylı incelenecektir.

## Yapılan Pratik

`exercises/basics/control-flow/` altında:

- Kullanıcı yaş ve aktiflik kontrolü
- Role göre `switch`
- `for` döngüsü
- `break`
- `continue`
- `range`

kullanılarak küçük bir uygulama geliştirildi.

## Aşama Sonu Zihinsel Model

```text
Comparison / Logical Operators
            ↓
        true / false
            ↓
        if / switch
            ↓
      Program kararı


for / range
     ↓
Tekrarlanan işlemler
     ↓
break / continue
     ↓
Döngü akışını kontrol et
```