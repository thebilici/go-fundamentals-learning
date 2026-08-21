# Aşama 6 — Maps

Bu aşamada Go'da key-value mantığıyla çalışan `map` veri yapısı öğrenildi.

## Map Nedir?

Map, verilere numeric index yerine bir key üzerinden erişmemizi sağlar.

```go
scores := map[string]int{
	"Fatih": 90,
	"Ahmet": 75,
}
```

Burada:

```text
string → key type
int    → value type
```

## Değer Okuma

```go
score := scores["Fatih"]
```

## Eleman Ekleme ve Güncelleme

Yeni key eklemek:

```go
scores["Ayşe"] = 85
```

Mevcut key'i güncellemek:

```go
scores["Fatih"] = 95
```

## Key Kontrolü

Bir key'in gerçekten var olup olmadığını kontrol etmek için:

```go
score, ok := scores["Fatih"]
```

Burada:

```text
score → value
ok    → key bulundu mu?
```

Örnek:

```go
if ok {
	fmt.Println(score)
}
```

## delete

Map'ten eleman silmek için:

```go
delete(scores, "Fatih")
```

## len

Map içerisindeki eleman sayısını verir:

```go
len(scores)
```

## range

Map içerisindeki key-value çiftlerini dolaşmak için:

```go
for name, score := range scores {
	fmt.Println(name, score)
}
```

Map dolaşılırken belirli bir sıralamaya güvenilmemelidir.

## Yapılan Pratik

`exercises/basics/maps/` altında:

- Map oluşturma
- Eleman ekleme
- Eleman güncelleme
- `value, ok`
- `delete`
- `len`
- `range`

kullanılarak basit bir öğrenci not sistemi oluşturuldu.

## Aşama Sonu Özeti

```text
Map
 ↓
Key → Value

Okuma       → map[key]
Ekleme      → map[key] = value
Güncelleme  → map[key] = value
Kontrol     → value, ok := map[key]
Silme       → delete(map, key)
Dolaşma     → range
```