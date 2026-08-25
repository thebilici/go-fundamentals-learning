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

# Aşama 9 — Pointers

## 1. Pointer'dan Alınan Değer ile Pointer'ı Karıştırmak

Başta:

```go
agePointerValue := *agePointer
agePointerValue = 25
```

ile orijinal `age` değerinin değişeceği düşünüldü.

Ancak:

```go
agePointerValue := *agePointer
```

pointer'ın gösterdiği değeri ayrı bir variable'a kopyalar.

```text
agePointer      → Address
*agePointer     → Address'teki değer
agePointerValue → Değerin ayrı kopyası
```

Orijinal değeri değiştirmek için:

```go
*agePointer = 25
```

kullanılmalıdır.

---

## 2. Struct Field İsimlerinde Büyük/Küçük Harf

Struct başlangıçta:

```go
type User struct {
	name string
	age  int
}
```

şeklinde tanımlanıp oluşturulurken:

```go
User{
	Name: "Fatih",
	Age:  22,
}
```

kullanıldı.

Go case-sensitive olduğu için:

```text
name != Name
age  != Age
```

Field isimleri struct tanımıyla aynı olmalıdır.

Örneğin:

```go
type User struct {
	Name string
	Age  int
}
```

ve:

```go
user := User{
	Name: "Fatih",
	Age:  22,
}
```

---

## 3. Struct Variable Oluşturma Syntax'ı

Başlangıçta:

```go
user:User{
```

yazıldı.

Yeni variable oluştururken:

```go
user := User{
	Name: "Fatih",
	Age:  22,
}
```

kullanılmalıdır.

# Aşama 10 — Interfaces

## 1. Interface Parameter Mantığını Karıştırmak

Başta:

```go
func printGreeting(g Greeter) {
	fmt.Println(g.greet())
}
```

yapısındaki:

```go
g Greeter
```

kısmı karıştırıldı.

Burada:

```text
g       → Parameter adı
Greeter → Parameter type'ı
```

Function belirli bir `User` veya `Admin` istemez.

`Greeter` interface'ini implement eden herhangi bir değer kabul edebilir.

```go
printGreeting(user)
printGreeting(admin)
```

---

## 2. Interface ile Tek Function Kullanma Mantığı

Başta farklı type'lar için ayrı function gerektiği düşünüldü:

```go
func printUserGreeting(user User)
func printAdminGreeting(admin Admin)
```

Interface sayesinde ortak davranış üzerinden tek function kullanılabilir:

```go
func printGreeting(g Greeter)
```

```text
User  → greet() string ──┐
                         ├── Greeter → printGreeting()
Admin → greet() string ──┘
```

Bu kullanım polymorphism örneğidir.

---

## 3. Interface Variable Mantığını Karıştırmak

Şu kullanım başlangıçta net değildi:

```go
var greeter Greeter
```

Interface normal bir variable type'ı olarak kullanılabilir.

```go
greeter = user
greeter = admin
```

Atanan concrete type'ın `Greeter` interface'ini implement etmesi gerekir.

```text
Variable type → Greeter

greeter = user
Concrete type → User

greeter = admin
Concrete type → Admin
```

Aynı interface variable farklı zamanlarda farklı concrete type'ları tutabilir.

# Aşama 11 — Error Handling

## 1. `nil` Mantığını Karıştırmak

Başta `nil` kavramı net değildi.

Error açısından:

```text
err == nil
→ hata yok

err != nil
→ hata var
```

Örneğin:

```go
return a / b, nil
```

işlemin başarılı olduğunu belirtir.

---

## 2. `err != nil` Hatayı Döndürmez

Başta:

```go
if err != nil
```

ifadesinin hatayı döndürdüğü düşünüldü.

Aslında sadece hata olup olmadığını kontrol eder.

```go
if err != nil {
	fmt.Println(err)
}
```

Burada:

```text
err != nil       → hata var mı kontrolü
fmt.Println(err) → hatayı yazdırır
```

---

## 3. Error Propagation Mantığı

Başta hatanın function'lar arasında nasıl ilerlediği karıştırıldı.

```text
main()
  ↓
calculate()
  ↓
divide()
```

`divide()` hata oluşturur:

```go
return 0, errors.New("cannot divide by zero")
```

`calculate()` hatayı alır ve yukarı taşır:

```go
if err != nil {
	return 0, err
}
```

`main()` hatayı işler.

```text
divide()    → error oluşturur
calculate() → error'ı propagate eder
main()      → error'ı handle eder
```

---

## 4. `:=` ve `=` Scope Farkı

`calculate()` içinde:

```go
result, err = divide(a, b)
```

kullanıldı.

Ancak `result` ve `err`, `calculate()` scope'unda henüz oluşturulmamıştı.

Doğrusu:

```go
result, err := divide(a, b)
```

Temel kural:

```text
Variable bu scope'ta ilk kez oluşturuluyor → :=

Variable aynı scope'ta zaten mevcut → =
```

Aynı variable isimlerinin başka bir function'da bulunması önemli değildir çünkü function'ların scope'ları ayrıdır.

---

## 5. `User{}` Mantığını Karıştırmak

Hata durumunda:

```go
return User{}, err
```

kullanımındaki `User{}` başlangıçta net değildi.

`User{}`, struct'ın zero value'larla oluşturulmuş halidir.

```go
type User struct {
	Name string
	Age  int
}
```

için:

```go
User{}
```

yaklaşık olarak:

```go
User{
	Name: "",
	Age:  0,
}
```

anlamına gelir.

Hata bulunduğu için gerçek bir `User` oluşturmak yerine boş/default `User` değeri ile error birlikte döndürülür.

# Aşama 12 — Packages & Modules

## 1. `main.go` ile `func main()` Farkı

Başta programın `main.go` dosyasından başladığı düşünüldü.

Asıl entry point:

```go
func main()
```

function'ıdır.

Executable Go programında:

```go
package main

func main() {
}
```

yapısı önemlidir.

Dosyanın adının mutlaka `main.go` olması gerekmez.

---

## 2. Import Path'i Module Sanmak

Şu yolun tamamının module olduğu düşünüldü:

```text
github.com/thebilici/go-backend-learning/exercises/basics/packages/user
```

Aslında:

```text
github.com/thebilici/go-backend-learning
→ Module Path

exercises/basics/packages/user
→ Package Path
```

Birleşince:

```text
github.com/thebilici/go-backend-learning/exercises/basics/packages/user
→ Import Path
```

Temel kural:

```text
Module Path + Package Path = Import Path
```

---

## 3. Package Adı ile Variable Adını Çakıştırmak

İlk exercise'da:

```go
user := user.CreateUser("Fatih", 22)
```

kullanıldı.

Burada:

```text
user → package
user → variable
```

aynı isim kullanıldığı için sonraki kullanımlarda package adı shadow edilir ve karışıklık oluşur.

Daha açık kullanım:

```go
createdUser := user.CreateUser("Fatih", 22)
```

veya import alias:

```go
import userpkg "example.com/project/user"

user := userpkg.CreateUser("Fatih", 22)
```

---

## 4. Function Parameter Yazım Sırasını Karıştırmak

İlk olarak:

```go
func CreateUser(string Name, int Age) User
```

şeklinde yazıldı.

Go'da parameter sırası:

```text
parameterName type
```

şeklindedir.

Doğrusu:

```go
func CreateUser(name string, age int) User
```

---

## 5. Package ile Dosyayı Aynı Şey Sanmamak

Bir package yalnızca tek `.go` dosyasından oluşmak zorunda değildir.

```text
mathutil/
├── math.go
└── validation.go
```

İkisinde de:

```go
package mathutil
```

varsa aynı package'ın parçalarıdır.

Bu nedenle aynı package içerisindeki type ve function'lar farklı dosyalarda olsalar bile birlikte kullanılabilir.
# Aşama 13 — Concurrency

## 1. Goroutine'i Ayrı Bir Function Türü Sanmak

Başta goroutine'in ayrı bir function türü olup olmadığı karıştırıldı.

Aslında normal bir function:

```go
func sendMessage(ch chan string) {
	ch <- "Hello"
}
```

normal şekilde çağrılabilir:

```go
sendMessage(ch)
```

veya `go` keyword'ü ile yeni bir goroutine içerisinde başlatılabilir:

```go
go sendMessage(ch)
```

Temel fark:

```text
sendMessage(ch)
→ Normal function çağrısı

go sendMessage(ch)
→ Function çağrısını yeni goroutine olarak başlat
```

---

## 2. `wg.Wait()` Konumunu Karıştırmak

`wg.Wait()` her yere yazılamaz.

Yanlış:

```go
wg.Add(2)

wg.Wait()

go task1(&wg)
go task2(&wg)
```

Burada counter `2` olur ancak `Done()` yapacak goroutine'ler henüz başlatılmamıştır.

Doğru:

```go
wg.Add(2)

go task1(&wg)
go task2(&wg)

wg.Wait()
```

Temel sıra:

```text
Add
↓
Goroutine'leri başlat
↓
Wait
```

---

## 3. `Done()` Konumunun Önemsiz Olduğunu Düşünmek

`Done()` goroutine'in işi tamamlandığında çalışmalıdır.

Yanlış mantık:

```go
func task(wg *sync.WaitGroup) {
	wg.Done()

	// işlemler...
}
```

Burada WaitGroup'a iş tamamlanmadan "tamamlandı" bilgisi verilmiş olur.

Tercih edilen kullanım:

```go
func task(wg *sync.WaitGroup) {
	defer wg.Done()

	// işlemler...
}
```

---

## 4. Unbuffered Channel'ı Veri Saklayan Kutu Gibi Düşünmek

Başta:

```go
ch := make(chan string)

ch <- "Hello"
```

işleminin `"Hello"` değerini channel içerisine bırakıp doğrudan devam ettiği düşünülebilir.

Unbuffered channel'da send ve receive birbirleriyle senkronize olur.

```text
Sender
   │
   │ "Hello"
   ▼
Channel
   │
   ▼
Receiver
```

Gönderici ve alıcının eşleşmesi gerekir.

Bu yüzden:

```go
go sendMessage(ch)

message := <-ch
```

çalışır.

Bir goroutine gönderirken main goroutine alabilir.

---

## 5. Aynı Goroutine'de Unbuffered Send ve Receive Yapmak

Şu kullanım problem oluşturabilir:

```go
ch := make(chan string)

ch <- "Hello"

message := <-ch
```

`main`, send sırasında receiver bekler.

Ancak receiver bir sonraki satırda olduğu için oraya ulaşamaz.

```text
Send
 ↓
Receiver bekleniyor
 ↓
Aynı goroutine ilerleyemiyor
 ↓
Receive satırına ulaşılamıyor
 ↓
Deadlock
```

Çözüm olarak send farklı bir goroutine'den yapılabilir:

```go
go func() {
	ch <- "Hello"
}()

message := <-ch
```

---

## 6. `close()` İşlemini Çok Erken Yapmamak

Channel, sender'lar hâlâ veri gönderecekken kapatılmamalıdır.

Yanlış mantık:

```go
go worker(ch)

close(ch)
```

Worker daha sonra:

```go
ch <- result
```

yaparsa kapalı channel'a send yapılmış olur ve runtime panic oluşabilir.

Birden fazla worker olduğunda güvenli pattern:

```go
go func() {
	wg.Wait()
	close(ch)
}()
```

Yani:

```text
Tüm sender'lar tamamlandı
↓
Artık yeni veri gelmeyecek
↓
Channel kapatılabilir
```