# Aşama 5 — Arrays & Slices

Bu aşamada Go'da birden fazla değeri birlikte tutmak için kullanılan Array ve Slice yapıları öğrenildi.

## Array

Array sabit uzunlukta bir veri yapısıdır.

```go
languages := [3]string{"Go", "Python", "Java"}
```

Burada:

```text
[3]    → Eleman sayısı
string → Elemanların type'ı
```

Array elemanlarına index ile erişilir:

```go
fmt.Println(languages[0])
```

Go'da index `0`'dan başlar.

## Slice

Slice, Array'e göre daha esnek bir yapıdır.

```go
languages := []string{"Go", "Python", "Java"}
```

Temel fark:

```text
Array → [3]string
Slice → []string
```

Array boyutu sabittir. Slice ise `append()` ile büyütülebilir.

## append

Slice'a yeni eleman eklemek için kullanılır.

```go
languages = append(languages, "JavaScript")
```

Birden fazla eleman da eklenebilir:

```go
languages = append(languages, "JavaScript", "Rust")
```

## len ve cap

`len()` mevcut eleman sayısını verir.

```go
len(languages)
```

`cap()` ise slice'ın mevcut backing array üzerinde erişebildiği kapasiteyi gösterir.

```go
cap(languages)
```

Kısaca:

```text
len → Kullanılan eleman sayısı
cap → Mevcut kapasite
```

## Slicing

Bir slice'ın belirli bir bölümünü almak için kullanılır.

```go
selected := languages[1:4]
```

Kural:

```text
[start:end]

start → dahil
end   → dahil değil
```

## range

Slice elemanlarını dolaşmak için kullanılır.

```go
for index, value := range languages {
	fmt.Println(index, value)
}
```

Burada:

```text
index → Elemanın konumu
value → Elemanın değeri
```

## Eleman Güncelleme

Slice içindeki bir değer index üzerinden değiştirilebilir:

```go
languages[0] = "Golang"
```

## Yapılan Pratik

`exercises/basics/arrays-slices/` altında:

- Array oluşturma
- Index kullanımı
- `len`
- Slice oluşturma
- `append`
- `cap`
- Slicing
- `range`
- Eleman güncelleme

uygulandı.

## Aşama Sonu Özeti

```text
Array
→ Sabit boyut

Slice
→ Esnek yapı
→ append ile büyüyebilir
→ len / cap
→ slicing
→ range ile dolaşılabilir
```