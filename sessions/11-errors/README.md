# Aşama 11 — Error Handling

Bu aşamada Go'da hata üretme, hata kontrolü, hata iletme ve hata sarmalama mantığı öğrenildi.

## `error` Nedir?

Go'da function'lar işlem sonucu ile birlikte hata da döndürebilir.

```go
func divide(a int, b int) (int, error)
```

Burada:

```text
int   → işlem sonucu
error → hata bilgisi
```

## `nil`

`nil`, error tarafında hata olmadığını ifade eder.

```text
err == nil  → hata yok
err != nil  → hata var
```

Başarılı durumda:

```go
return result, nil
```

## Error Kontrolü

Go'da sık kullanılan kalıp:

```go
result, err := someFunction()

if err != nil {
	return
}
```

## `errors.New()`

Sabit bir error oluşturmak için kullanılır.

```go
return errors.New("cannot divide by zero")
```

## `fmt.Errorf()`

Değişken içeren veya daha açıklayıcı error mesajları üretmek için kullanılır.

```go
return fmt.Errorf("cannot divide %d by %d", a, b)
```

## `(value, error)` Yapısı

Function hem sonuç hem de hata döndürebilir.

```go
func createUser(name string, age int) (User, error)
```

Başarılı durumda:

```go
return user, nil
```

Hata durumunda:

```go
return User{}, err
```

`User{}` struct'ın zero value'larıyla oluşturulmuş boş halidir.

## Error Propagation

Bir function aldığı hatayı üst function'a iletebilir.

```go
result, err := divide(a, b)

if err != nil {
	return 0, err
}
```

Akış:

```text
divide()
   ↓
calculate()
   ↓
main()
```

Hata alt katmandan üst katmana taşınabilir.

## Error Wrapping

Hatanın üzerine context eklemek için `%w` kullanılabilir.

```go
return 0, fmt.Errorf("calculate failed: %w", err)
```

Bu kullanım orijinal error bilgisini koruyarak daha açıklayıcı bir hata üretir.

## Yapılan Pratik

`exercises/basics/error-handling/` altında:

- `validateAge()` oluşturuldu.
- `error` ve `nil` kullanıldı.
- `errors.New()` kullanıldı.
- `(User, error)` dönüşü uygulandı.
- Error propagation yapıldı.
- `fmt.Errorf("%w")` ile error wrapping uygulandı.
- Hata durumunda `User{}` döndürüldü.

## Aşama Sonu Özeti

```text
Function
   ↓
(value, error)
   ↓
err == nil
→ Başarılı

err != nil
→ Hata var
   ↓
Handle veya propagate et
   ↓
Gerekirse wrap et
```

Temel Go error pattern:

```go
result, err := someFunction()

if err != nil {
	return err
}

// result güvenle kullanılabilir
```