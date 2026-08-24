# Go Flashcards

## Aşama 1 — Go Ortamı ve Temeller

### 1. Go Module nedir?

Bir Go projesinin üst seviye kimliği ve dependency yönetim sınırıdır.

---

### 2. `go.mod` ne işe yarar?

Module path bilgisini, kullanılan Go sürümünü ve gerekli dependency bilgilerini tutar.

---

### 3. Package nedir?

Birbiriyle ilişkili Go kodlarını gruplamak için kullanılan yapıdır.

Bir package birden fazla `.go` dosyasından oluşabilir.

---

### 4. `package main` ne anlama gelir?

Çalıştırılabilir bir Go programının oluşturulmasında kullanılan özel package'dir.

---

### 5. `func main()` ne işe yarar?

Programın entry point'idir. Programın çalışması buradan başlar.

---

### 6. Standard Library nedir?

Go kurulumu ile birlikte gelen hazır package koleksiyonudur.

Örnek:

```text
fmt
net/http
os
time
context
```

---

### 7. `fmt.Println()` ne yapar?

`fmt` package'indeki `Println` fonksiyonunu kullanarak terminale çıktı yazar.

---

### 8. `go run .` ne yapar?

Mevcut package'i derler ve hemen çalıştırır.

```powershell
go run .
```

---

### 9. `go build .` ne yapar?

Mevcut package'i derleyerek executable binary oluşturur.

```powershell
go build .
```

---

### 10. `go run` ve `go build` arasındaki fark nedir?

```text
go run
→ Derle + Çalıştır

go build
→ Derle + Binary oluştur
```

---

### 11. Source Code ve Binary arasındaki fark nedir?

```text
main.go
→ İnsan tarafından okunabilen Source Code

main.exe
→ Compiler tarafından oluşturulan çalıştırılabilir Binary
```

---

### 12. Exported Identifier nedir?

Büyük harfle başlayan ve başka package'lerden erişilebilen identifier'dır.

```text
Hello   → Exported
hello   → Unexported
```

---

### 13. `go fmt` ne işe yarar?

Go kodunu standart Go formatına göre düzenler.

```powershell
go fmt ./...
```

---

### 14. Module ve Package aynı şey midir?

Hayır.

```text
Module
  ↓
Bir veya daha fazla Package
  ↓
.go Files
```

Module daha üst seviyedeki yapıdır.

---

### 15. Go neden kullanılmayan import'lara izin vermez?

Gereksiz kodu engellemek ve hataların compile aşamasında fark edilmesini sağlamak için kullanılmayan import'ları compiler error olarak değerlendirir.

# Aşama 2 — Variables & Data Types

### 1. Variable nedir?

Program içerisinde bir değeri tutmak için kullanılan isimlendirilmiş yapıdır.

---

### 2. `var` ne işe yarar?

Variable tanımlamak için kullanılır.

```go
var age int = 22
```

---

### 3. `:=` ne işe yarar?

Function içerisinde kısa şekilde yeni variable oluşturur ve type inference yapar.

```go
age := 22
```

---

### 4. `:=` ile `=` farkı nedir?

```text
:= → Yeni variable oluşturur.
=  → Mevcut variable'a değer atar.
```

---

### 5. Go'daki temel data type'lar nelerdir?

```text
string
int
float64
bool
```

---

### 6. Static Typing nedir?

Variable type'larının compile aşamasında belirli olmasıdır.

```go
age := 22
age = "22" // Hata
```

---

### 7. Type Inference nedir?

Compiler'ın verilen değerden type'ı otomatik belirlemesidir.

```go
age := 22 // int
```

---

### 8. Zero Value nedir?

Başlangıç değeri verilmeyen variable'ın otomatik aldığı değerdir.

```text
string  → ""
int     → 0
float64 → 0
bool    → false
```

---

### 9. `const` ne işe yarar?

Değeri sonradan değiştirilemeyen constant tanımlar.

```go
const language = "Go"
```

---

### 10. Type Conversion nedir?

Bir değeri başka bir type olarak kullanmak için yapılan açık dönüşümdür.

```go
ageFloat := float64(age)
```

---

### 11. `strconv.Atoi` ne yapar?

String içerisindeki decimal integer değerini `int` olarak parse eder.

```text
"25" → 25
```

---

### 12. `strconv.Itoa` ne yapar?

`int` değerini decimal string'e dönüştürür.

```text
25 → "25"
```

---

### 13. `byte` hangi type'ın alias'ıdır?

```text
byte → uint8
```

---

### 14. `rune` hangi type'ın alias'ıdır?

```text
rune → int32
```

---

### 15. `int` ve `int64` aynı type mıdır?

Hayır. İkisi farklı type'lardır ve gerektiğinde explicit conversion yapılmalıdır.

```go
var number int = 22
converted := int64(number)
```

# Aşama 3 — Control Flow

### 1. Comparison Operator'lar ne üretir?

`true` veya `false` değerinde bir `bool` sonuç üretir.

---

### 2. `=` ile `==` farkı nedir?

```text
=  → Değer atar
== → İki değeri karşılaştırır
```

---

### 3. Temel Logical Operator'lar nelerdir?

```text
&& → AND
|| → OR
!  → NOT
```

---

### 4. `if` ne işe yarar?

Bir koşul `true` olduğunda belirli bir kod bloğunu çalıştırır.

---

### 5. `else` ne zaman çalışır?

`if` koşulu `false` olduğunda çalışır.

---

### 6. `else if` neden kullanılır?

Birden fazla koşulu sırayla kontrol etmek için kullanılır.

---

### 7. `switch` ne işe yarar?

Bir değeri birden fazla olasılıkla karşılaştırmayı kolaylaştırır.

---

### 8. Go'nun temel döngüsü nedir?

`for` döngüsüdür.

```go
for i := 1; i <= 5; i++ {
	fmt.Println(i)
}
```

---

### 9. Go'da `while` var mı?

Hayır. While benzeri kullanım `for` ile yapılır.

```go
for count <= 5 {
	count++
}
```

---

### 10. `break` ne yapar?

Döngüyü tamamen sonlandırır.

---

### 11. `continue` ne yapar?

Mevcut turu atlar ve sonraki tura geçer.

---

### 12. `break` ve `continue` farkı nedir?

```text
break    → Döngüyü bitirir
continue → Sadece mevcut turu atlar
```

---

### 13. `range` ne işe yarar?

Collection içerisindeki elemanları sırayla dolaşmayı sağlar.

---

### 14. `range` ile hangi değerleri alabiliriz?

Örneğin slice üzerinde:

```go
for index, value := range languages {
}
```

ile index ve value alınabilir.

---

### 15. `_` nedir?

Blank Identifier'dır. Kullanmak istemediğimiz değeri yok saymamızı sağlar.

```go
for _, value := range languages {
	fmt.Println(value)
}
```

# Aşama 4 — Functions

### 1. Function nedir?

Belirli bir işi yapan ve gerektiğinde tekrar çağrılabilen kod bloğudur.

---

### 2. Go'da function hangi keyword ile tanımlanır?

`func` keyword'ü ile.

```go
func greet() {
}
```

---

### 3. Parameter nedir?

Function tanımlanırken belirtilen değişkendir.

```go
func greet(name string) {
}
```

Burada `name` parameter'dır.

---

### 4. Argument nedir?

Function çağrılırken gönderilen gerçek değerdir.

```go
greet("Fatih")
```

Burada `"Fatih"` argument'tır.

---

### 5. Parameter ile argument arasındaki fark nedir?

```text
func greet(name string)
           ↑
       Parameter

greet("Fatih")
       ↑
    Argument
```

---

### 6. Return Value nedir?

Function'ın ürettiği sonucu çağrıldığı yere geri göndermesidir.

```go
func add(a, b int) int {
	return a + b
}
```

---

### 7. Return Type nerede yazılır?

Parameter listesinden sonra yazılır.

```go
func add(a, b int) int
                   ↑
              Return Type
```

---

### 8. Bir function birden fazla değer döndürebilir mi?

Evet.

```go
func getUser() (string, int) {
	return "Fatih", 22
}
```

---

### 9. Multiple Return değerleri nasıl alınır?

Sırayla variable'lara atanır.

```go
name, age := getUser()
```

```text
"Fatih" → name
22      → age
```

---

### 10. Birden fazla Return Type nasıl belirtilir?

Parantez içerisinde yazılır.

```go
func example() (string, int, bool) {
	return "Fatih", 22, true
}
```

---

### 11. `_` Multiple Return ile neden kullanılır?

İhtiyacımız olmayan değeri yok saymak için.

```go
name, _ := getUser()
```

---

### 12. Named Return nedir?

Return değerine function tanımında isim verilmesidir.

```go
func add(a, b int) (result int) {
	result = a + b
	return
}
```

---

### 13. Function Scope nedir?

Function içerisinde oluşturulan local variable'ların yalnızca kendi scope'larında erişilebilir olmasıdır.

---

### 14. `:=` ile oluşturulan şey variable mıdır?

Evet.

```go
name := "Fatih"
```

`name` normal bir variable'dır.

---

### 15. `var`, `:=` ve `const` arasındaki temel fark nedir?

```text
var   → Variable tanımlar
:=    → Function içinde kısa variable tanımlar
const → Constant tanımlar
```

---

### 16. Aynı type'taki parameter'lar kısa nasıl yazılır?

```go
func add(a, b int) int {
	return a + b
}
```

Hem `a` hem `b`, `int` type'ındadır.

---

### 17. Function sonucu variable'a alınmak zorunda mı?

Hayır.

```go
fmt.Println(add(5, 10))
```

şeklinde doğrudan kullanılabilir.

---

### 18. `return` çalıştığında ne olur?

Değer çağıran yere gönderilir ve function'ın çalışması sona erer.

---

### 19. Function'ın temel veri akışı nasıldır?

```text
Argument
   ↓
Parameter
   ↓
Function
   ↓
İşlem
   ↓
Return
   ↓
Sonuç
```

---

### 20. `getUser()` örneğinde aşağıdaki kod ne yapar?

```go
name, age := getUser()
```

`getUser()` tarafından döndürülen değerleri sırasıyla `name` ve `age` variable'larına atar.

# Aşama 5 — Arrays & Slices

### 1. Array nedir?

Sabit sayıda ve aynı type'ta eleman tutan veri yapısıdır.

```go
languages := [3]string{"Go", "Python", "Java"}
```

---

### 2. `[3]string` ne anlama gelir?

3 adet `string` eleman tutan Array anlamına gelir.

---

### 3. Slice nedir?

Eleman sayısı sabit olmayan, esnek bir collection yapısıdır.

```go
languages := []string{"Go", "Python", "Java"}
```

---

### 4. Array ve Slice syntax farkı nedir?

```text
[3]string → Array
[]string  → Slice
```

---

### 5. Index nedir?

Bir elemanın collection içerisindeki konumudur.

```go
languages[0]
```

Go'da index `0`'dan başlar.

---

### 6. `len()` ne işe yarar?

Array veya Slice içerisindeki mevcut eleman sayısını verir.

```go
len(languages)
```

---

### 7. `append()` ne işe yarar?

Slice'a yeni eleman eklemek için kullanılır.

```go
languages = append(languages, "Rust")
```

---

### 8. Neden `languages = append(...)` yazıyoruz?

`append()` güncellenmiş slice'ı geri döndürür. Bu nedenle sonucu tekrar variable'a atarız.

---

### 9. `cap()` ne gösterir?

Slice'ın mevcut backing array üzerinde erişebildiği kapasiteyi gösterir.

```text
len → Mevcut eleman sayısı
cap → Mevcut kapasite
```

---

### 10. Slicing nedir?

Bir Slice'ın belirli bir bölümünü almaktır.

```go
selected := languages[1:4]
```

---

### 11. `[1:4]` hangi index'leri alır?

```text
1 → dahil
2 → dahil
3 → dahil
4 → dahil değil
```

---

### 12. `range` ne işe yarar?

Slice içerisindeki elemanları sırayla dolaşmayı sağlar.

```go
for index, value := range languages {
	fmt.Println(index, value)
}
```

---

### 13. Slice içerisindeki bir eleman nasıl değiştirilir?

Index kullanılarak:

```go
languages[0] = "Golang"
```

---

### 14. Array sonradan büyütülebilir mi?

Hayır. Array'in uzunluğu oluşturulduğu anda sabittir.

---

### 15. Slice büyütülebilir mi?

Evet. `append()` kullanılarak yeni elemanlar eklenebilir.

---

### 16. Array ve Slice için aynı variable ismi aynı scope içinde tekrar `:=` ile kullanılabilir mi?

Hayır.

```go
languages := [3]string{"Go", "Python", "Java"}

languages := []string{"Go", "Python", "Java"} // Hata
```

Daha açıklayıcı farklı isimler kullanılabilir:

```go
languages := [3]string{"Go", "Python", "Java"}
languageSlice := []string{"Go", "Python", "Java"}
```

# Aşama 6 — Maps

### 1. Map nedir?

Key-value mantığıyla veri tutan collection yapısıdır.

```go
scores := map[string]int{
	"Fatih": 90,
}
```

---

### 2. `map[string]int` ne anlama gelir?

```text
string → Key type
int    → Value type
```

---

### 3. Map'ten bir değer nasıl okunur?

```go
score := scores["Fatih"]
```

---

### 4. Map'e yeni eleman nasıl eklenir?

```go
scores["Ayşe"] = 85
```

Key yoksa yeni eleman oluşturulur.

---

### 5. Map'teki bir değer nasıl güncellenir?

```go
scores["Fatih"] = 95
```

Key zaten varsa value güncellenir.

---

### 6. `value, ok` ne işe yarar?

Key'in Map içerisinde gerçekten bulunup bulunmadığını kontrol eder.

```go
score, ok := scores["Fatih"]
```

```text
score → Value
ok    → Key bulundu mu?
```

---

### 7. Key varsa `ok` değeri ne olur?

```text
true
```

---

### 8. Key yoksa `ok` değeri ne olur?

```text
false
```

Value ise kendi type'ının zero value'sunu alır.

---

### 9. Olmayan bir `int` value'lu key okunursa ne döner?

`int` type'ının zero value'su olan `0` döner.

---

### 10. Map'ten eleman nasıl silinir?

```go
delete(scores, "Fatih")
```

---

### 11. `len(map)` ne verir?

Map içerisindeki key-value çifti sayısını verir.

---

### 12. Map nasıl dolaşılır?

`range` ile:

```go
for name, score := range scores {
	fmt.Println(name, score)
}
```

---

### 13. Map üzerinde `range` kullanırken sıraya güvenebilir miyiz?

Hayır. Map iteration sırasına güvenilmemelidir.

---

### 14. Slice ve Map arasındaki temel erişim farkı nedir?

```text
Slice
Index → Value

languages[0]


Map
Key → Value

scores["Fatih"]
```

---

### 15. `:=` ile `=` farkı nedir?

```text
:= → Yeni variable tanımlamak için
=  → Mevcut variable'a yeni değer atamak için
```

Örneğin:

```go
score, ok := scores["Ahmet"]

score, ok = scores["Ali"]
```
# Aşama 7 — Structs

### 1. Struct nedir?

Birbiriyle ilişkili verileri tek bir type altında toplamamızı sağlayan yapıdır.

```go
type User struct {
	Name string
	Age  int
}
```

---

### 2. `type User struct` ne anlama gelir?

`User` adında yeni bir struct type tanımlar.

---

### 3. Field nedir?

Struct içerisinde tanımlanan verilerdir.

```go
type User struct {
	Name string
	Age  int
}
```

`Name` ve `Age` field'dır.

---

### 4. Struct'tan bir değer nasıl oluşturulur?

```go
user := User{
	Name: "Fatih",
	Age:  22,
}
```

---

### 5. Struct field'ına nasıl erişilir?

`.` kullanılır.

```go
user.Name
user.Age
```

---

### 6. Struct field'ı nasıl güncellenir?

```go
user.Age = 23
```

---

### 7. Struct function parameter'ı olabilir mi?

Evet.

```go
func printUser(user User) {
	fmt.Println(user.Name)
}
```

---

### 8. Function bir struct döndürebilir mi?

Evet.

```go
func createUser() User {
	return User{
		Name: "Fatih",
		Age:  22,
	}
}
```

---

### 9. `[]User` ne anlama gelir?

`User` değerlerinden oluşan bir Slice anlamına gelir.

```go
users := []User{
	user,
	user2,
}
```

---

### 10. `range` ile `[]User` dolaşılırken `user` nedir?

Her döngüde Slice içerisinden alınan tek bir `User` değeridir.

```go
for _, user := range users {
	fmt.Println(user.Name)
}
```

---

### 11. `range` içindeki `user` ismini Go mu belirler?

Hayır. Variable ismini programcı belirler.

```go
for _, person := range users {
	fmt.Println(person.Name)
}
```

---

### 12. Slice üzerinde `range` ne döndürür?

```go
for index, value := range slice
```

```text
index → Elemanın index'i
value → O index'teki değer
```

---

### 13. Slice'taki `index, value` Map'teki `key, value` ile aynı mıdır?

Hayır.

```text
Slice → index, value
Map   → key, value
```

Index key'e benzer bir rol oynasa da teknik olarak key değildir.

---

### 14. `Age: age` ne anlama gelir?

```go
return User{
	Age: age,
}
```

```text
Sol Age  → Struct field
Sağ age  → Variable / function parameter
```

---

### 15. Struct literal içinde olmayan bir field adı kullanılabilir mi?

Hayır.

Struct:

```go
type User struct {
	IsActive bool
}
```

ise:

```go
User{
	IsActive: true,
}
```

kullanılmalıdır. Struct'ta tanımlanmayan bir field adı compiler error oluşturur.

---

### 16. Büyük harfle başlayan struct field ne anlama gelir?

Exported field'dır.

```go
Age int
```

Başka package'lerden erişilebilir.

---

### 17. Küçük harfle başlayan struct field ne anlama gelir?

Unexported field'dır.

```go
age int
```

Başka package'lerden doğrudan erişilemez.

---

### 18. Function parameter'ının adı ile struct field'ının adı aynı olmak zorunda mı?

Hayır.

```go
func createUser(userAge int) User {
	return User{
		Age: userAge,
	}
}
```

Field ve parameter birbirinden bağımsızdır.

# Aşama 8 — Methods

### 1. Method nedir?

Belirli bir type ile ilişkilendirilmiş function'dır.

```go
func (u User) getName() string {
	return u.Name
}
```

---

### 2. Function ve Method arasındaki temel fark nedir?

```text
Function → getName(user)
Method   → user.getName()
```

Method belirli bir type ile ilişkilidir.

---

### 3. Receiver nedir?

Method'un hangi type üzerinde çalışacağını belirtir.

```go
func (u User) getName() string
```

---

### 4. `(u User)` içerisindeki `u` nedir?

Receiver variable'dır.

İsmini programcı belirler.

---

### 5. `(u User)` içerisindeki `User` nedir?

Receiver type'dır.

Method'un `User` type'ıyla ilişkili olduğunu belirtir.

---

### 6. Neden `user.getName()` çağırabiliriz?

Çünkü `user` değişkeninin type'ı `User`, `getName()` method'unun receiver type'ı da `User`'dır.

---

### 7. Method parameter alabilir mi?

Evet.

```go
func (u User) canAccess(minAge int) bool {
	return u.Age >= minAge
}
```

---

### 8. Receiver ile parameter aynı şey midir?

Hayır.

```go
func (u User) canAccess(minAge int)
```

```text
u      → Receiver
minAge → Parameter
```

---

### 9. Method değer return edebilir mi?

Evet.

```go
func (u User) isAdult() bool {
	return u.Age >= 18
}
```

---

### 10. Value Receiver nedir?

Receiver'ın değer type'ı ile tanımlanmasıdır.

```go
func (u User) getName()
```

Buradaki `(u User)` value receiver'dır.

---

### 11. Value Receiver ile struct değiştirilirse orijinal değer değişir mi?

Genellikle hayır. Method receiver'ın aldığı değer üzerinde çalışır.

Orijinal struct'ı değiştirmek için pointer receiver kullanılabilir.

---

### 12. Method içerisinde struct field'ına nasıl erişilir?

Receiver üzerinden:

```go
func (u User) getName() string {
	return u.Name
}
```

---

### 13. `user.greet("Hello")` içerisindeki `"Hello"` nedir?

Method'a gönderilen argument'tır.

```go
func (u User) greet(message string)
```

Buradaki `message` ise parameter'dır.

---

### 14. Method neden kullanılır?

Bir type'a ait davranışları o type ile ilişkilendirmek için.

Örneğin:

```go
user.isAdult()
user.getName()
user.canAccess(18)
```

---

### 15. `User` type'ına ait method başka bir type üzerinden çağrılabilir mi?

Doğrudan hayır.

```text
Receiver type → User
```

ise method `User` type'ıyla ilişkilidir.

# Aşama 9 — Pointers

### 1. Pointer nedir?

Başka bir variable'ın memory address'ini tutan değerdir.

```go
age := 22
pointer := &age
```

---

### 2. `&` operator ne işe yarar?

Variable'ın memory address'ini alır.

```go
&age
```

---

### 3. `*` operator pointer ile kullanıldığında ne işe yarar?

Pointer'ın gösterdiği address'teki değere erişir.

```go
*pointer
```

Bu işleme dereferencing denir.

---

### 4. `pointer` ile `*pointer` arasındaki fark nedir?

```text
pointer  → Memory address
*pointer → O address'teki value
```

---

### 5. `*pointer = 25` ne yapar?

Pointer'ın gösterdiği address'teki değeri değiştirir.

```go
age := 22
pointer := &age

*pointer = 25
```

Sonuç:

```text
age → 25
```

---

### 6. `*int` ne anlama gelir?

Bir `int` değerine işaret eden pointer type'dır.

```text
int  → normal int value
*int → int pointer
```

---

### 7. `*User` ne anlama gelir?

Bir `User` değerine işaret eden pointer type'dır.

---

### 8. Aşağıdaki kod ne yapar?

```go
value := *pointer
```

Pointer'ın o anda gösterdiği değeri `value` variable'ına kopyalar.

---

### 9. Neden `value := *pointer` sonrasında pointer değişince `value` otomatik değişmez?

Çünkü `value`, pointer değildir. Değerin ayrı bir kopyasını tutar.

---

### 10. Normal parameter ile pointer parameter arasındaki fark nedir?

```go
func changeAge(age int)
```

değer alır.

```go
func changeAge(age *int)
```

int pointer alır ve orijinal değere erişebilir.

---

### 11. Pointer bekleyen function nasıl çağrılır?

```go
changeAge(&age)
```

`&age` ile `age` variable'ının address'i gönderilir.

---

### 12. Aşağıdaki function neden orijinal `age` değerini değiştirebilir?

```go
func changeAge(age *int) {
	*age = 30
}
```

Çünkü pointer üzerinden `age` variable'ının bulunduğu address'teki değer değiştirilir.

---

### 13. Value Receiver nedir?

```go
func (u User) getName()
```

`User` value receiver'dır.

---

### 14. Pointer Receiver nedir?

```go
func (u *User) changeName()
```

Receiver bir `User` pointer'ıdır.

---

### 15. Pointer Receiver neden kullanılır?

Method'un orijinal struct üzerinde değişiklik yapabilmesini sağlar.

```go
func (u *User) changeName(name string) {
	u.Name = name
}
```

---

### 16. `u.Name` yazarken neden `(*u).Name` yazmak zorunda değiliz?

Go, struct pointer'larında field erişimi sırasında gerekli dereference işlemini otomatik yapabilir.

---

### 17. Value Receiver ve Pointer Receiver arasındaki temel fark nedir?

```text
(u User)
→ Value Receiver

(u *User)
→ Pointer Receiver
→ Orijinal struct üzerinde değişiklik yapabilir
```

# Aşama 10 — Interfaces

### 1. Interface nedir?

Bir type'ın sahip olması gereken method'ları tanımlayan yapıdır.

```go
type Notifier interface {
	send() string
}
```

---

### 2. Interface method'un nasıl çalışacağını belirler mi?

Hayır. Sadece hangi method'ların bulunması gerektiğini belirtir.

---

### 3. Bir type interface'i nasıl implement eder?

Interface'in istediği bütün method'lara doğru signature ile sahip olarak.

```go
func (e EmailNotifier) send() string {
	return "Mail gönderildi"
}
```

---

### 4. Implicit Implementation nedir?

Go'da `implements` gibi bir keyword kullanmadan, gerekli method'lara sahip olan type'ın interface'i otomatik olarak implement etmesidir.

---

### 5. Interface birden fazla method içerebilir mi?

Evet.

```go
type Greeter interface {
	greet() string
	getName() string
}
```

Type'ın interface'i implement etmesi için bütün method'lara sahip olması gerekir.

---

### 6. Method Signature neden önemlidir?

Interface'in istediği method ile concrete type'ın method signature'ı uyumlu olmalıdır.

```text
Interface: send() string
Type:      send() string

→ Uyumlu ✅
```

---

### 7. Interface function parameter olarak kullanılabilir mi?

Evet.

```go
func sendNotification(n Notifier) {
	fmt.Println(n.send())
}
```

---

### 8. Interface parameter kullanmanın avantajı nedir?

Function'ın belirli bir concrete type yerine aynı davranışı sağlayan farklı type'larla çalışabilmesini sağlar.

```go
sendNotification(email)
sendNotification(sms)
```

---

### 9. Interface variable olarak kullanılabilir mi?

Evet.

```go
var notifier Notifier
```

---

### 10. Interface variable'a ne atanabilir?

Interface'i implement eden değerler atanabilir.

```go
notifier = email
notifier = sms
```

---

### 11. Concrete Type nedir?

Interface içerisinde bulunan gerçek type'dır.

```go
var notifier Notifier
notifier = email
```

```text
Interface type → Notifier
Concrete type  → EmailNotifier
```

---

### 12. Polymorphism nedir?

Farklı concrete type'ların ortak bir interface üzerinden kullanılabilmesidir.

```text
EmailNotifier ──┐
                ├── Notifier
SMSNotifier ────┘
```

---

### 13. `n.send()` çağrısında hangi `send()` çalışır?

`n` içerisinde bulunan concrete type'ın `send()` implementation'ı çalışır.

```text
n = EmailNotifier
→ EmailNotifier.send()

n = SMSNotifier
→ SMSNotifier.send()
```

---

### 14. Interface kullanırken concrete type önemli midir?

Interface açısından asıl önemli olan concrete type'ın gerekli method'ları sağlayıp sağlamadığıdır.

```text
"Hangi type'sın?"
yerine
"Gerekli davranışlara sahip misin?"
```

---

### 15. Interface neden farklı type'lar için ayrı function yazma ihtiyacını azaltabilir?

Çünkü aynı davranışı sağlayan farklı type'lar ortak interface üzerinden tek function'a gönderilebilir.

```go
func sendNotification(n Notifier)
```

yerine ayrı ayrı:

```go
func sendEmailNotification(e EmailNotifier)
func sendSMSNotification(s SMSNotifier)
```

yazmak gerekmeyebilir.

---

### 16. `EmailNotifier` ve `SMSNotifier` neden aynı `Notifier` interface'ini implement edebilir?

İkisinde de interface'in istediği:

```go
send() string
```

method'u bulunduğu için.

---

### 17. Interface ile polymorphism arasındaki bağlantı nedir?

Interface ortak davranışı tanımlar; farklı type'lar bu davranışı farklı şekilde implement eder ve aynı interface üzerinden kullanılabilir.

# Aşama 11 — Error Handling

### 1. Go'da `error` nedir?

Bir işlem sırasında oluşan hatayı temsil eden type'dır.

---

### 2. `(int, error)` ne anlama gelir?

Function iki değer döndürür:

```text
int   → işlem sonucu
error → hata bilgisi
```

---

### 3. `nil` error açısından ne anlama gelir?

Hata olmadığını belirtir.

```text
err == nil → hata yok
```

---

### 4. `err != nil` ne anlama gelir?

Bir hata oluştuğunu belirtir.

```go
if err != nil {
	// hata var
}
```

---

### 5. Go'da temel error kontrol pattern'i nedir?

```go
result, err := someFunction()

if err != nil {
	return
}
```

---

### 6. `errors.New()` ne işe yarar?

Yeni ve sabit mesajlı bir error oluşturur.

```go
errors.New("cannot divide by zero")
```

---

### 7. `fmt.Errorf()` ne işe yarar?

Formatlı error mesajı oluşturmaya yarar.

```go
fmt.Errorf("invalid age: %d", age)
```

---

### 8. `errors.New()` ile `fmt.Errorf()` arasındaki temel fark nedir?

```text
errors.New()
→ Sabit mesaj

fmt.Errorf()
→ Formatlanabilir/dinamik mesaj
```

---

### 9. Başarılı bir `(value, error)` dönüşünde genellikle ne döndürülür?

```go
return value, nil
```

---

### 10. Hata durumunda neden `User{}` döndürülebilir?

Function `User` döndürmek zorunda olduğu için hata durumunda `User` type'ının zero value hali kullanılabilir.

```go
return User{}, err
```

---

### 11. `User{}` ne anlama gelir?

Field'ları zero value olan bir `User` değeridir.

```text
string → ""
int    → 0
bool   → false
```

---

### 12. Error Propagation nedir?

Bir function'ın aldığı error'ı çağıran üst function'a iletmesidir.

```go
if err != nil {
	return 0, err
}
```

---

### 13. Aşağıdaki kod ne yapar?

```go
result, err := divide(a, b)

if err != nil {
	return 0, err
}
```

`divide()` tarafından döndürülen error varsa bunu üst function'a taşır.

---

### 14. Error Wrapping nedir?

Bir error'a ek context ekleyerek orijinal error'ı koruyup yukarı taşımaktır.

```go
fmt.Errorf("calculate failed: %w", err)
```

---

### 15. `%w` ne işe yarar?

Mevcut error'ı wrap eder.

```go
fmt.Errorf("operation failed: %w", err)
```

---

### 16. `return 0, err` ile `return 0, fmt.Errorf("failed: %w", err)` arasındaki fark nedir?

```text
return 0, err
→ Error'ı olduğu gibi taşır.

fmt.Errorf("failed: %w", err)
→ Error'a context ekleyerek taşır.
```

---

### 17. Error handling ile error propagation arasındaki fark nedir?

```text
Handling
→ Hatayı bulunduğun yerde işle.

Propagation
→ Hatayı üst function'a gönder.
```

---

### 18. Aşağıdaki error akışı neyi gösterir?

```text
divide()
   ↓
calculate()
   ↓
main()
```

Alt function'da oluşan error'ın üst function'lara taşınmasını, yani error propagation'ı gösterir.

---

### 19. `validateAge()` başarılı olduğunda neden `nil` döndürür?

Çünkü herhangi bir hata oluşmamıştır.

```go
return nil
```

---

### 20. Go'da neden error kontrolü sık sık açık şekilde yazılır?

Go'da error'lar genellikle return value olarak taşınır ve çağıran kod tarafından açıkça kontrol edilir.

```go
if err != nil {
	// error işle
}
```
# Aşama 12 — Packages & Modules

### 1. Package nedir?

Birbiriyle ilişkili Go kodlarını gruplamak için kullanılan yapıdır.

---

### 2. `package main` ne anlama gelir?

Çalıştırılabilir Go programının ana package'ıdır.

---

### 3. Go programının entry point'i nedir?

```go
func main()
```

---

### 4. Dosyanın adının `main.go` olması zorunlu mudur?

Hayır. Önemli olan executable program için `package main` içerisinde `func main()` bulunmasıdır.

---

### 5. Bir package birden fazla `.go` dosyasından oluşabilir mi?

Evet.

```text
user/
├── user.go
└── validation.go
```

İki dosyada da:

```go
package user
```

varsa aynı package'ın parçalarıdır.

---

### 6. Exported identifier nedir?

Başka package'lardan erişilebilen identifier'dır.

Büyük harfle başlar.

```go
CreateUser
User
Add
```

---

### 7. Unexported identifier nedir?

Package dışından doğrudan erişilemeyen identifier'dır.

Küçük harfle başlar.

```go
createUser
user
add
```

---

### 8. Go'da `public` ve `private` keyword'leri var mı?

Hayır.

Export durumu identifier'ın ilk harfinin büyük veya küçük olmasıyla belirlenir.

---

### 9. Module nedir?

Go projesinin üst seviyedeki kimliği ve dependency yönetim sınırıdır.

---

### 10. Module path nerede tanımlanır?

`go.mod` dosyasında.

```go
module github.com/thebilici/go-backend-learning
```

---

### 11. Package ile Module arasındaki fark nedir?

```text
Module
→ Projenin üst seviyedeki kimliği

Package
→ Module içerisindeki kod grubu
```

---

### 12. Module, Package ve File ilişkisi nasıldır?

```text
Module
└── Package
    ├── file.go
    └── file.go
```

---

### 13. Package path nedir?

Package'ın module içerisindeki klasör yoludur.

```text
exercises/basics/packages/user
```

---

### 14. Import path nasıl oluşur?

```text
Module Path + Package Path = Import Path
```

---

### 15. Aşağıdakilerden hangisi module path'tir?

```text
github.com/thebilici/go-backend-learning
```

`go.mod` içerisindeki `module` değeridir.

---

### 16. Aşağıdakinin tamamı nedir?

```go
"github.com/thebilici/go-backend-learning/exercises/basics/packages/user"
```

Package'ın import path'idir.

---

### 17. `user.CreateUser()` ifadesinde `user` nedir?

Package adıdır.

```text
user.CreateUser()
 ↑       ↑
package  exported function
```

---

### 18. `fmt.Println()` ile `user.CreateUser()` arasında package kullanımı açısından fark var mı?

Temel mantık aynıdır.

```text
fmt.Println()
↓
fmt package

user.CreateUser()
↓
user package
```

---

### 19. Import alias nedir?

Import edilen package'a farklı bir isim vermektir.

```go
import mathpkg "example.com/project/mathutil"
```

Kullanım:

```go
mathpkg.Add(5, 6)
```

---

### 20. `go mod init` ne yapar?

Yeni bir Go module başlatır ve `go.mod` oluşturur.

---

### 21. `go mod tidy` ne yapar?

Kodun dependency kullanımını analiz ederek module dependency bilgilerini düzenler.

---

### 22. `go.sum` ne işe yarar?

External dependency'lerin checksum bilgilerini tutar ve dependency bütünlüğünün doğrulanmasına yardımcı olur.

---

### 23. Her package için ayrı `go.mod` gerekir mi?

Hayır.

Aynı module içerisindeki package'lar module kökündeki `go.mod` dosyasını kullanabilir.

---

### 24. Aynı package içerisindeki farklı `.go` dosyaları birbirlerinin type ve function'larını kullanabilir mi?

Evet.

Örneğin:

```text
mathutil/
├── math.go
│   └── Add()
│
└── validation.go
    └── IsPositive()
```

İkisi de:

```go
package mathutil
```

ise aynı package'ın parçalarıdır.

---

### 25. Packages & Modules için temel zihinsel model nedir?

```text
go.mod
   ↓
Module
   ↓
Packages
   ↓
.go Files
   ↓
Functions / Types
```