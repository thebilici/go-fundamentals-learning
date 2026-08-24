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

# Aşama 4 — Functions

## 1. Multiple Return Values Mantığı

Bir function birden fazla değer döndürebilir:

```go
func getUser() (string, int) {
	return "Fatih", 22
}
```

Dönen değerler sırayla variable'lara alınır:

```go
name, age := getUser()
```

```text
"Fatih" → name
22      → age
```

---

## 2. Multiple Return Type Syntax'ı

Birden fazla değer döndürüldüğünde dönüş tipleri zaten `()` içerisinde belirtilir:

```go
func example() (string, int, bool) {
	return "Fatih", 22, true
}
```

Burada:

```text
(string, int, bool)
```

function'ın return type'larıdır.

---

## 3. `:=` ile Oluşturulan Değer Variable'dır

```go
name := "Fatih"
```

Burada `name`, normal bir variable'dır.

```text
var   → Variable tanımlar
:=    → Function içinde kısa şekilde variable tanımlar
const → Constant tanımlar
```

# Aşama 5 — Arrays & Slices

## 1. `len` ve `cap` Farkı

Bu aşamada `cap` kavramını anlamakta zorlanıldı.

```text
len → Slice'ta şu anda bulunan eleman sayısı
cap → Slice'ın mevcut backing array üzerinde erişebildiği kapasite
```

Capacity dolduktan sonra `append()` yapılırsa Go daha büyük bir backing array oluşturabilir ve `cap` değeri artabilir.

---

## 2. Array ve Slice İçin Aynı Variable İsmini Kullanmak

Aynı scope içerisinde:

```go
languages := [3]string{"Go", "Python", "Java"}
```

tanımlandıktan sonra tekrar:

```go
languages := []string{"Go", "Python", "Java"}
```

şeklinde `:=` kullanılamaz.

Bunun yerine:

```go
languages := [3]string{"Go", "Python", "Java"}
languageSlice := []string{"Go", "Python", "Java"}
```

gibi farklı ve açıklayıcı isimler kullanılabilir.

# Aşama 6 — Maps

## 1. `else` Syntax'ı

Go'da `else`, kapanan `}` ile aynı satırda olmalıdır.

Yanlış:

```go
if ok {
	fmt.Println(score)
}
else {
	fmt.Println("Bulunamadı")
}
```

Doğru:

```go
if ok {
	fmt.Println(score)
} else {
	fmt.Println("Bulunamadı")
}
```

---

## 2. `:=` ve `=` Kullanımı

İlk kontrolde variable'lar oluşturuldu:

```go
score, ok := scores["Ahmet"]
```

Aynı variable'ları tekrar kullanırken:

```go
score, ok = scores["Ali"]
```

kullanılabilir.

```text
:= → Yeni variable tanımlama
=  → Mevcut variable'a değer atama
```

# Aşama 7 — Structs

## 1. `users` ve `user` Farkı

```go
for _, user := range users {
}
```

Burada:

```text
users → Tüm []User slice'ı
user  → Her döngüde alınan tek bir User
```

`user` ismini programcı belirler.

---

## 2. Slice'ta `index, value` Mantığı

```go
for index, user := range users {
}
```

Burada:

```text
index → User'ın slice içerisindeki konumu
user  → O index'teki User değeri
```

Map'teki `key, value` yapısına benzer ancak Slice'ta `index` bir key değildir.

---

## 3. Struct Field ve Function Parameter Farkı

```go
func createUser(name string, age int) User {
	return User{
		Name: name,
		Age:  age,
	}
}
```

Burada:

```text
Age → Struct field
age → Function parameter
```

İsimlerinin aynı olması zorunlu değildir.

Örneğin:

```go
Age: userAge
```

şeklinde de kullanılabilir.

---

## 4. Struct'ta Tanımlanmayan Field Kullanmak

Struct literal içerisindeki sol taraf, struct'ta gerçekten tanımlanmış bir field olmalıdır.

```go
type User struct {
	IsActive bool
}
```

Bu nedenle:

```go
User{
	IsActive: true,
}
```

geçerlidir.

Struct'ta bulunmayan bir field adı kullanılırsa compiler error oluşur.

# Aşama 8 — Methods

## 1. Type, Variable ve Method Bağlantısı

```go
user := User{
	Name: "Fatih",
	Age:  22,
}
```

Burada `User` type'ı variable'a atanmaz.

`User` type'ından bir değer oluşturulur ve bu değer `user` variable'ında tutulur.

```text
User → Type
user → User type'ındaki variable
```

Method:

```go
func (u User) greet() {
	fmt.Println(u.Name)
}
```

Receiver type `User` olduğu için:

```go
user.greet()
```

şeklinde çağrılabilir.

```text
user'ın type'ı → User
greet receiver → User
                 ↓
              Uyumlu
```