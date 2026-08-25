    # Aşama 13 — Concurrency

Bu aşamada Go'da concurrent çalışma, goroutine, WaitGroup, channel ve temel synchronization mantığı öğrenildi.

## Sequential vs Concurrent

Sequential çalışma:

```text
Task 1
  ↓
Task 2
  ↓
Task 3
```

Concurrent çalışma:

```text
Task 1 ────────→
Task 2   ────────→
Task 3      ────────→
```

Birden fazla iş aynı zaman aralığında ilerleyebilir.

## Goroutine

Bir function çağrısını goroutine olarak çalıştırmak için `go` keyword'ü kullanılır.

```go
go sendMessage(ch)
```

Normal çağrı:

```go
sendMessage(ch)
```

Goroutine çağrısı:

```go
go sendMessage(ch)
```

## WaitGroup

Birden fazla goroutine'in tamamlanmasını beklemek için kullanılır.

```go
var wg sync.WaitGroup

wg.Add(2)
```

Goroutine tamamlandığında:

```go
defer wg.Done()
```

Main goroutine bekler:

```go
wg.Wait()
```

Temel mantık:

```text
Add()  → Beklenen iş sayısını artır
Done() → Bir iş tamamlandı
Wait() → Sayaç 0 olana kadar bekle
```

## Channel

Goroutine'ler arasında veri taşımak için kullanılır.

```go
ch := make(chan string)
```

Veri gönderme:

```go
ch <- "Hello"
```

Veri alma:

```go
message := <-ch
```

## Unbuffered Channel

```go
ch := make(chan string)
```

Sender ve receiver birbirini bekler.

```text
Sender ←→ Receiver
```

## Buffered Channel

```go
ch := make(chan string, 3)
```

Buffer içerisinde belirli sayıda değer tutulabilir.

```text
len(ch) → Buffer'daki mevcut değer sayısı
cap(ch) → Buffer kapasitesi
```

## Channel ile Sonuç Döndürme

Bir goroutine ürettiği sonucu channel üzerinden başka bir goroutine'e gönderebilir.

```go
func calculate(a, b int, ch chan int) {
	result := a + b
	ch <- result
}
```

## close ve range

Channel'a artık yeni veri gönderilmeyeceği belirtilir:

```go
close(ch)
```

Channel içerisindeki değerler `range` ile okunabilir:

```go
for value := range ch {
	fmt.Println(value)
}
```

Channel kapandığında ve veri kalmadığında döngü sona erer.

## WaitGroup + Channel

Birden fazla worker'ın sonuç göndermesi:

```text
Worker 1 ──┐
Worker 2 ──┼──→ Channel ──→ Receiver
Worker 3 ──┘
```

Tüm worker'lar bittikten sonra channel kapatılabilir:

```go
go func() {
	wg.Wait()
	close(ch)
}()
```

## Race Condition

Birden fazla goroutine aynı shared data'yı aynı anda değiştirdiğinde beklenmedik sonuçlar oluşabilir.

```text
Goroutine A ──→ counter
Goroutine B ──→ counter
```

Bu duruma race condition denir.

Go race detector:

```bash
go run -race .
```

## Mutex

Shared data'ya aynı anda yalnızca bir goroutine'in erişmesini sağlamak için kullanılabilir.

```go
mu.Lock()
counter++
mu.Unlock()
```

## Yapılan Pratik

`exercises/basics/concurrency/` altında:

- 3 goroutine oluşturuldu.
- `sync.WaitGroup` kullanıldı.
- `chan int` oluşturuldu.
- Goroutine'ler sonuçlarını channel'a gönderdi.
- Tüm worker'lar tamamlandıktan sonra channel kapatıldı.
- `range` ile sonuçlar okundu.
- Buffered channel kullanıldı.

## Aşama Sonu Özeti

```text
Goroutine
    ↓
Concurrent iş

WaitGroup
    ↓
İşlerin tamamlanmasını bekle

Channel
    ↓
Goroutine'ler arasında veri taşı

close + range
    ↓
Channel yaşam döngüsünü yönet

Mutex
    ↓
Shared data'yı koru
```