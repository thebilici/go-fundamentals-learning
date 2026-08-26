# Session 16 — Context

## Context Nedir?

Context, bir işlemin yaşam döngüsünü kontrol etmek için kullanılan Go yapısıdır.

Özellikle:

- Cancellation
- Timeout
- Deadline
- İptal bilgisini alt işlemlere taşıma

için kullanılır.

Backend uygulamalarında bir request'in artık devam etmemesi gerektiğini alt katmanlara bildirmek için sık kullanılır.

---

## context.Context

`context.Context`, `context` package içerisinde bulunan bir interface'tir.

```go
func worker(ctx context.Context)
```

Burada:

```text
ctx              → Parameter adı
context          → Package
Context          → Package içerisindeki type/interface
context.Context  → Parameter'ın tipi
```

---

## context.Background()

Başlangıç context'i oluşturmak için kullanılır.

```go
ctx := context.Background()
```

Bu context başlangıçta:

```text
Cancellation yok
Timeout yok
Deadline yok
```

Yeni context'ler bunun üzerinden oluşturulabilir.

---

## context.WithCancel()

Manuel olarak iptal edilebilen bir context oluşturur.

```go
ctx, cancel := context.WithCancel(ctx)
```

İki değer döndürür:

```text
ctx
→ İptal edilebilir context

cancel
→ Context'i iptal eden function
```

İptal etmek için:

```go
cancel()
```

kullanılır.

---

## ctx.Done()

Context'in sona erdiğini bildiren channel'dır.

```go
<-ctx.Done()
```

Genellikle `select` içerisinde kullanılır:

```go
select {
case <-ctx.Done():
	fmt.Println("Context stopped")
	return

default:
	fmt.Println("Working...")
}
```

Mantık:

```text
Context devam ediyor
        ↓
İşe devam et

Context iptal edildi
        ↓
ctx.Done()
        ↓
return
        ↓
İşlem biter
```

---

## ctx.Err()

Context'in neden sona erdiğini öğrenmek için kullanılır.

```go
ctx.Err()
```

Manuel iptal durumunda:

```text
context canceled
```

Timeout veya deadline durumunda:

```text
context deadline exceeded
```

sonucu alınabilir.

---

## context.WithTimeout()

Bir context'in belirli bir süre sonra otomatik olarak sona ermesini sağlar.

```go
ctx, cancel := context.WithTimeout(
	context.Background(),
	3*time.Second,
)

defer cancel()
```

Mantık:

```text
Context oluştur
      ↓
3 saniye çalış
      ↓
Süre doldu
      ↓
ctx.Done()
      ↓
İşlem durdurulabilir
```

---

## context.WithDeadline()

Context'in belirli bir zamana kadar çalışmasını sağlar.

```go
deadline := time.Now().Add(5 * time.Second)

ctx, cancel := context.WithDeadline(
	context.Background(),
	deadline,
)

defer cancel()
```

Temel fark:

```text
WithTimeout
→ "3 saniye çalış"

WithDeadline
→ "şu zamana kadar çalış"
```

---

## defer cancel()

Context oluşturulurken alınan `cancel` function genellikle:

```go
defer cancel()
```

şeklinde kullanılır.

Bu sayede function sona ererken context ile ilişkili kaynakların serbest bırakılması garanti edilir.

---

## Context + Goroutine

Context bir goroutine'e gönderilebilir:

```go
go worker(ctx)
```

Worker context'i dinleyebilir:

```go
func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		default:
			fmt.Println("Working...")
		}
	}
}
```

`cancel()` çağrıldığında:

```text
main
 ↓
cancel()
 ↓
ctx.Done()
 ↓
worker sinyali görür
 ↓
return
 ↓
goroutine biter
```

---

## Context Propagation

Context bir function'dan diğerine aktarılabilir.

```go
func service(ctx context.Context) {
	repository(ctx)
}

func repository(ctx context.Context) {
	// ctx kullanılabilir
}
```

Akış:

```text
main
 ↓ ctx
service
 ↓ ctx
repository
```

Buna Context Propagation denir.

Backend uygulamalarında ileride:

```text
HTTP Request
     ↓
Handler
     ↓ ctx
Service
     ↓ ctx
Repository
     ↓ ctx
Database
```

şeklinde kullanılabilir.

---

## Yapılan Pratik

3 saniyelik timeout ile çalışan bir worker oluşturuldu.

```go
func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped:", ctx.Err())
			return

		default:
			fmt.Println("Processing...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
```

Akış:

```text
Worker başladı
      ↓
Processing...
      ↓
Processing...
      ↓
3 saniye doldu
      ↓
ctx.Done()
      ↓
context deadline exceeded
      ↓
return
```

---

## Temel Özet

```text
context.Background()
→ Başlangıç context'i

context.WithCancel()
→ Manuel iptal edilebilir context

cancel()
→ İptal et

ctx.Done()
→ İptal / timeout sinyalini dinle

ctx.Err()
→ Context neden sona erdi?

context.WithTimeout()
→ Belirli süre sonra sona er

context.WithDeadline()
→ Belirli zamanda sona er

Context Propagation
→ Context'i alt işlemlere taşı
```

Context'in temel amacı:

> Bir işlemin ne kadar süre yaşayacağını ve ne zaman durması gerektiğini kontrollü şekilde yönetmek ve bu bilgiyi alt işlemlere aktarabilmektir.