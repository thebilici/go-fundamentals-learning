# Aşama 12 — Packages & Modules

Bu aşamada Go projelerinin package ve module yapısıyla nasıl organize edildiği öğrenildi.

## Package

Package, birbiriyle ilişkili Go kodlarını aynı yapı altında toplar.

```go
package user
```

Aynı package içerisinde birden fazla `.go` dosyası bulunabilir.

```text
user/
├── user.go
└── validation.go
```

İki dosyada da:

```go
package user
```

yazıyorsa aynı package'ın parçalarıdır.

## `package main`

Çalıştırılabilir Go uygulamasının ana package'ıdır.

```go
package main

func main() {
}
```

Programın entry point'i `func main()` function'ıdır.

## Custom Package

Kendi package'larımızı oluşturabiliriz.

```go
package mathutil

func Add(a, b int) int {
	return a + b
}
```

Başka bir package'tan kullanmak için import edilir.

## Exported / Unexported

Büyük harfle başlayan identifier'lar exported'dır.

```text
Add        → Exported
CreateUser → Exported
User       → Exported
```

Küçük harfle başlayanlar unexported'dır.

```text
add
createUser
user
```

Başka package'lardan yalnızca exported identifier'lara erişilebilir.

## Module

Module, Go projesinin üst seviyedeki kimliğidir.

`go.mod` içerisinde tanımlanır:

```go
module github.com/thebilici/go-backend-learning
```

Bir module birden fazla package içerebilir.

```text
Module
├── Package
├── Package
└── Package
```

## Package ve Module Farkı

```text
Module
→ Projenin ana kimliği ve dependency sınırı

Package
→ Module içerisindeki kod grubu

.go File
→ Package içerisindeki kaynak kod dosyası
```

Temel model:

```text
Module
└── Package
    └── Go Files
```

## Import Path

Bir package'ın import path'i:

```text
Module Path
+
Package'ın module içindeki yolu
=
Import Path
```

Örneğin:

```text
Module:
github.com/thebilici/go-backend-learning

Package:
exercises/basics/packages/user
```

Import:

```go
import "github.com/thebilici/go-backend-learning/exercises/basics/packages/user"
```

## Package Kullanımı

Import edilen package içerisindeki exported function'lar package adı üzerinden çağrılır:

```go
user.CreateUser(...)
```

veya:

```go
mathutil.Add(...)
```

Bu mantık standard library'de de aynıdır:

```go
fmt.Println()
```

## Import Alias

Package'a farklı bir isim verilebilir:

```go
import mathpkg "github.com/thebilici/go-backend-learning/exercises/basics/packages-final/mathutil"
```

Sonra:

```go
mathpkg.Add(5, 6)
```

şeklinde kullanılır.

## `go.mod`

Module bilgisini ve dependency'leri yönetir.

Temel yapı:

```go
module github.com/thebilici/go-backend-learning

go 1.26.5
```

## `go mod init`

Yeni bir Go module oluşturur ve `go.mod` dosyasını başlatır.

```bash
go mod init github.com/example/project
```

## `go mod tidy`

Kullanılan dependency'leri analiz eder.

Genel olarak:

```text
Eksik dependency → eklenir
Gereksiz dependency → kaldırılır
go.mod / go.sum → güncellenebilir
```

## `go.sum`

External dependency'lerin checksum bilgilerini tutar ve dependency bütünlüğünün doğrulanmasına yardımcı olur.

## Yapılan Pratik

Bu aşamada:

- `user` custom package'ı oluşturuldu.
- `User`, `CreateUser()` ve `GetName()` exported olarak kullanıldı.
- Aynı package içerisinde `user.go` ve `validation.go` kullanıldı.
- Module path ile package path birleştirilerek import path oluşturuldu.
- `mathutil` package'ı oluşturuldu.
- `math.go` ve `validation.go` aynı package altında kullanıldı.
- Import alias kullanıldı.

## Aşama Sonu Özeti

```text
go.mod
  ↓
Module
  ↓
Packages
  ↓
.go Files
  ↓
Exported Functions / Types
  ↓
import
  ↓
package.Function()
```