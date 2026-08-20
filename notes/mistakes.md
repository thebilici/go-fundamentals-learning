# Hatalar ve Çıkarılan Dersler

## Aşama 1 — Go Ortamı ve Temeller

### 1. Import Etmeden Package Kullanmak

`fmt.Println()` kullanılmasına rağmen `fmt` import edilmezse program derlenmez.

Yanlış:

```go
package main

func main() {
	fmt.Println("Hello, Go!")
}
```

Doğru:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
```

**Ders:** Başka bir package içerisindeki bir identifier kullanılacaksa ilgili package import edilmelidir.

---

### 2. Kullanılmayan Package Import Etmek

Go kullanılmayan import'lara izin vermez.

```go
package main

import "fmt"

func main() {
}
```

Bu kod compiler error oluşturur çünkü `fmt` import edilmiş ancak kullanılmamıştır.

**Ders:** Go compiler gereksiz import'ların kod içerisinde kalmasını engeller.

---

### 3. Binary Dosyasını Source Code Sanmak

`go build` sonrasında `main.exe` veya `basics.exe` oluşturuldu.

```text
main.go
→ Source Code

main.exe / basics.exe
→ Build Artifact
```

**Ders:** `.exe` dosyaları kaynak kod değildir. Compiler tarafından üretilen çalıştırılabilir binary dosyalarıdır ve normalde Git repository'ye eklenmemelidir.

---

### 4. Dosya ve Package Bazlı Build Farkı

İlk olarak:

```powershell
go build main.go
```

kullanıldığında `main.exe` oluştu.

Daha sonra:

```powershell
go build .
```

kullanıldığında package bazlı build yapıldı ve `basics.exe` oluştu.

**Ders:** Go projelerinde sadece dosya bazlı değil, package bazlı düşünmek önemlidir.