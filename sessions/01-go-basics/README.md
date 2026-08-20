# Aşama 1 — Go Ortamı ve Temeller

Bu aşamada Go'nun temel çalışma mantığı öğrenildi. İlk Go programı oluşturuldu, Go Module yapısı incelendi ve yazılan kaynak kodun nasıl derlenerek çalıştırılabilir bir programa dönüştüğü gözlemlendi.

---

## 1. Go Ortamı

Kurulu Go sürümü aşağıdaki komutla kontrol edildi:

```powershell
go version
```

Kullanılan ortam:

```text
go version go1.26.5 windows/amd64
```

Burada:

- `go1.26.5` → kullanılan Go sürümü
- `windows` → hedef işletim sistemi
- `amd64` → x86-64 işlemci mimarisi

---

## 2. Go Module

Repo aşağıdaki komutla bir Go Module haline getirildi:

```powershell
go mod init github.com/thebilici/go-backend-learning
```

Bu işlem sonucunda repo kökünde:

```text
go.mod
```

dosyası oluşturuldu.

Module, Go projesinin üst seviye kimliği ve dependency yönetim sınırıdır.

Bizim module path'imiz:

```text
github.com/thebilici/go-backend-learning
```

Genel yapı:

```text
Module
  ↓
Packages
  ↓
.go Files
  ↓
Functions / Types / Variables
```

Bir Module içerisinde birden fazla Package bulunabilir.

---

## 3. İlk Go Programı

İlk Go programı:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
```

Bu küçük program Go'nun temel program yapısını göstermektedir.

---

## 4. `package main`

Her Go dosyası bir package'e ait olmak zorundadır.

```go
package main
```

bu dosyanın `main` package'ine ait olduğunu belirtir.

`main` package'i çalıştırılabilir Go uygulamaları için özeldir.

Çalıştırılabilir bir Go programının temel yapısı:

```text
package main
      +
func main()
      ↓
Executable Program
```

---

## 5. `import`

Başka bir package içerisindeki kodları kullanabilmek için `import` kullanılır.

```go
import "fmt"
```

Burada Go Standard Library içerisinde bulunan `fmt` package'i programa dahil edilmiştir.

Go ile birlikte gelen bazı Standard Library package'leri:

```text
fmt
os
time
errors
net/http
context
encoding/json
```

---

## 6. `func main()`

```go
func main() {

}
```

programın başlangıç noktasıdır.

Go programı çalıştırıldığında execution `main()` fonksiyonundan başlar.

```text
Program başlatılır
      ↓
package main
      ↓
func main()
      ↓
Program kodları çalışır
```

---

## 7. `fmt.Println`

```go
fmt.Println("Hello, Go!")
```

terminal üzerine çıktı yazdırmak için kullanılır.

Buradaki yapı:

```text
fmt.Println
 │      │
 │      └── Println fonksiyonu
 │
 └── fmt package
```

Genel olarak:

```text
package.Identifier
```

mantığı kullanılmaktadır.

`Println` isminin büyük harfle başlaması da önemlidir.

Go'da büyük harfle başlayan identifier'lar başka package'lerden erişilebilir, yani **exported** kabul edilir.

---

## 8. `go run`

Programı çalıştırmak için:

```powershell
go run main.go
```

veya package bazlı olarak:

```powershell
go run .
```

kullanıldı.

`go run` kaynak kodu derler ve oluşan programı hemen çalıştırır.

```text
main.go
   ↓
Compile
   ↓
Temporary Binary
   ↓
Run
   ↓
Program Output
```

Development sırasında hızlı şekilde kodu çalıştırmak için kullanışlıdır.

---

## 9. `go build`

Programı executable haline getirmek için:

```powershell
go build main.go
```

kullanıldı.

Bu işlem sonucunda Windows üzerinde:

```text
main.exe
```

oluştu.

Package bazlı:

```powershell
go build .
```

kullanıldığında ise package build edildi ve:

```text
basics.exe
```

oluştu.

Temel build süreci:

```text
Go Source Code
      ↓
Go Compiler
      ↓
Machine Code
      ↓
Executable Binary
```

`go run` ile `go build` arasındaki temel fark:

```text
go run
→ Derle + Çalıştır

go build
→ Derle + Binary oluştur
```

---

## 10. Dosya Bazlı ve Package Bazlı Çalışma

Dosya bazlı:

```powershell
go run main.go
go build main.go
```

Package bazlı:

```powershell
go run .
go build .
```

Go projeleri büyüdükçe package bazlı çalışma daha önemli hale gelir.

Bir package içerisinde:

```text
main.go
server.go
config.go
```

gibi birden fazla `.go` dosyası bulunabilir.

Bu dosyalar aynı package'e aitse birlikte build edilebilir.

---

## 11. `go fmt`

Go kodunu standart biçimde formatlamak için:

```powershell
go fmt main.go
```

kullanıldı.

Örneğin kötü formatlanmış:

```go
package main
import "fmt"
func main(){
fmt.Println("Hello")
}
```

kodunu standart Go formatına getirir:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello")
}
```

Go ekosisteminde ortak ve standart bir kod formatı kullanılması teşvik edilir.

---

## 12. Go Compiler Davranışı

Go compiler kullanılmayan import'lara izin vermez.

Örneğin:

```go
import "fmt"

func main() {

}
```

kodunda `fmt` kullanılmadığı için compilation başarısız olur.

Aynı şekilde:

```go
func main() {
	fmt.Println("Hello")
}
```

yazıp `fmt` import edilmezse de hata alınır.

Bu davranış gereksiz ve hatalı kodların compile aşamasında fark edilmesini sağlar.

---

## 13. Source Code ve Binary

Bu aşamada source code ile executable arasındaki fark gözlemlendi.

```text
main.go
   │
   │ go build
   ▼
main.exe
```

`main.go`:

- İnsan tarafından okunabilir kaynak koddur.
- Git repository içerisinde tutulur.

`main.exe` / `basics.exe`:

- Compiler tarafından üretilir.
- İşletim sistemi tarafından çalıştırılabilir.
- Bir build artifact'tır.
- Normalde Git repository'ye eklenmez.

---

## Aşama Sonu Zihinsel Model

Bu aşamanın sonunda Go uygulamasının temel çalışma modeli:

```text
Go Module
    ↓
Package
    ↓
.go Source Files
    ↓
Go Compiler
    ↓
Executable Binary
    ↓
Operating System
    ↓
CPU
```

olarak öğrenildi.

## Aşama Sonunda Öğrenilen Temel Kavramlar

- Go Module
- `go.mod`
- Package
- `package main`
- `import`
- Standard Library
- `func main()`
- Entry Point
- Exported Identifier
- `fmt.Println`
- `go run`
- `go build`
- `go fmt`
- Compiler
- Source Code
- Binary / Executable
- Build Artifact
- Package bazlı çalışma