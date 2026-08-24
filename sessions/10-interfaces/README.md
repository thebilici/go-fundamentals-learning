# Aşama 10 — Interfaces

Bu aşamada Go'da interface, implicit implementation ve polymorphism mantığı öğrenildi.

## Interface Nedir?

Interface, bir type'ın sahip olması gereken davranışları yani method'ları tanımlar.

```go
type Notifier interface {
	send() string
}
```

Bu interface:

```text
send() string
```

method'una sahip type'ların `Notifier` olarak kullanılabileceğini belirtir.

## Implicit Implementation

Go'da bir type'ın interface'i implement ettiğini ayrıca belirtmeyiz.

Örneğin:

```go
type EmailNotifier struct {
	Address string
}

func (e EmailNotifier) send() string {
	return e.Address + " adresine mail gönderildi"
}
```

`EmailNotifier`, `send() string` method'una sahip olduğu için otomatik olarak `Notifier` interface'ini implement eder.

```text
Notifier istiyor:
send() string

EmailNotifier sahip:
send() string

↓
EmailNotifier → Notifier ✅
```

## Bir Interface Birden Fazla Method İsteyebilir

```go
type Greeter interface {
	greet() string
	getName() string
}
```

Bir type'ın `Greeter` olabilmesi için iki method'a da sahip olması gerekir.

```text
greet() string   ✅
getName() string ✅

→ Greeter implement edilir
```

Bir method eksikse interface implement edilmez.

## Interface Function Parameter Olarak Kullanılabilir

```go
func sendNotification(n Notifier) {
	fmt.Println(n.send())
}
```

Burada function belirli bir struct istemez.

`Notifier` interface'ini implement eden herhangi bir type kabul edilebilir.

```go
sendNotification(email)
sendNotification(sms)
```

## Interface Variable Olarak Kullanılabilir

```go
var notifier Notifier
```

Interface'i implement eden farklı değerler bu variable'a atanabilir:

```go
notifier = email
fmt.Println(notifier.send())

notifier = sms
fmt.Println(notifier.send())
```

Aynı interface variable'ı farklı concrete type'larla çalışabilir.

## Polymorphism

Farklı concrete type'ların aynı interface üzerinden kendi method implementation'larını çalıştırabilmesine polymorphism denir.

```text
EmailNotifier ──┐
                │
                ├── send() string
                │
SMSNotifier ────┘
        │
        ▼
     Notifier
        │
        ▼
sendNotification()
```

Aynı çağrı:

```go
n.send()
```

gelen concrete type'a göre farklı method'u çalıştırabilir.

```text
n = EmailNotifier
→ EmailNotifier.send()

n = SMSNotifier
→ SMSNotifier.send()
```

## Interface'in Temel Mantığı

Interface:

```text
"Sen hangi struct'sın?"
```

sorusundan çok:

```text
"İhtiyacım olan method'lara sahip misin?"
```

mantığıyla çalışır.

Bu sayede kod belirli concrete type'lara daha az bağımlı hale getirilebilir.

## Yapılan Pratik

`exercises/basics/interfaces/` altında:

- `Notifier` interface oluşturuldu.
- `EmailNotifier` oluşturuldu.
- `SMSNotifier` oluşturuldu.
- Her iki type için `send()` method'u yazıldı.
- Implicit implementation uygulandı.
- Interface function parameter olarak kullanıldı.
- Interface variable oluşturuldu.
- Aynı interface variable'ına farklı concrete type'lar atandı.
- Polymorphism uygulandı.

## Aşama Sonu Özeti

```text
Interface
    ↓
Gerekli method'ları tanımlar
    ↓
Concrete type bu method'lara sahipse
    ↓
Implicit olarak interface'i implement eder
    ↓
Farklı type'lar aynı interface üzerinden kullanılabilir
    ↓
Polymorphism
```