# Go Kavramları

Bu dosya Go öğrenme sürecinde karşılaşılan önemli kavramların kısa açıklamalarını içerir.

## Go Module

Bir Go projesinin üst seviye kimliği ve dependency yönetim sınırıdır.

Module bilgileri `go.mod` dosyasında tutulur.

```text
Module
  ↓
Packages
  ↓
.go Files
```

## Package

Birbiriyle ilişkili Go kodlarını gruplamak için kullanılan yapıdır.

Bir package birden fazla `.go` dosyasından oluşabilir.

## `package main`

Çalıştırılabilir Go uygulamalarında kullanılan özel package'dir.

`func main()` ile birlikte programın başlangıç yapısını oluşturur.

## Entry Point

Programın çalışmaya başladığı noktadır.

Go uygulamalarında:

```go
func main() {
}
```

entry point olarak kullanılır.

## Standard Library

Go kurulumu ile birlikte gelen hazır package koleksiyonudur.

Örnekler:

- `fmt`
- `net/http`
- `os`
- `time`
- `context`

## Exported Identifier

Go'da büyük harfle başlayan identifier'lar başka package'lerden erişilebilir.

```text
Println → Exported
println → Unexported
```

## Compiler

Go source code'unu işletim sisteminin ve CPU'nun çalıştırabileceği machine code'a dönüştürür.

```text
Source Code
    ↓
Compiler
    ↓
Binary
```

## Binary / Executable

`go build` sonucunda oluşturulan çalıştırılabilir programdır.

Windows üzerinde genellikle `.exe` uzantısına sahiptir.

## `go run`

Go kodunu derler ve hemen çalıştırır.

```powershell
go run .
```

## `go build`

Go kodunu derleyerek executable binary oluşturur.

```powershell
go build .
```

## `go fmt`

Go kodunu standart Go formatına göre düzenler.

```powershell
go fmt ./...
```
# Aşama 2 — Variables & Data Types

## Variable

Program içerisinde bir değeri tutmak için kullanılan isimlendirilmiş yapıdır.

```go
var age int = 22
```

## Static Typing

Go statically typed bir dildir. Variable'ın type'ı compile aşamasında bellidir ve sonradan başka bir type'a dönüşmez.

```go
age := 22
age = "22" // Hata
```

## Type Inference

Compiler'ın verilen değere bakarak variable'ın type'ını belirlemesidir.

```go
age := 22       // int
height := 1.73  // float64
```

## Short Variable Declaration `:=`

Function içerisinde kısa şekilde yeni variable oluşturmak için kullanılır.

```go
age := 22
```

`:=` yeni variable oluştururken `=` mevcut variable'a değer atar.

## Zero Value

Başlangıç değeri verilmeyen variable'lara Go otomatik olarak zero value verir.

```text
string  → ""
int     → 0
float64 → 0
bool    → false
```

## Constant

Program içerisinde değeri değişmemesi gereken değerler `const` ile tanımlanır.

```go
const language = "Go"
```

## Type Conversion

Bir değerin başka bir type olarak kullanılmasını sağlar.

```go
age := 22
ageFloat := float64(age)
```

## Parsing

Text olarak bulunan bir değerin yorumlanarak başka bir type'a çevrilmesidir.

```go
number, err := strconv.Atoi("25")
```

Burada `"25"` string değeri `25` integer değerine dönüştürülür.

## `strconv.Atoi`

String → int dönüşümü için kullanılır.

```go
number, err := strconv.Atoi("25")
```

## `strconv.Itoa`

Int → string dönüşümü için kullanılır.

```go
text := strconv.Itoa(25)
```

## `byte` ve `rune`

Go'daki iki önemli alias:

```text
byte → uint8
rune → int32
```

`byte` binary/veri işlemlerinde, `rune` ise Unicode karakterlerle çalışırken sık kullanılır.
# Aşama 3 — Control Flow

## Comparison Operators

İki değeri karşılaştırır ve `bool` sonuç üretir.

```text
==   Eşit
!=   Eşit değil
>    Büyük
<    Küçük
>=   Büyük veya eşit
<=   Küçük veya eşit
```

## Logical Operators

Birden fazla koşulu birleştirmek için kullanılır.

```text
&& → AND
|| → OR
!  → NOT
```

## if / else

Bir koşula göre hangi kod bloğunun çalışacağını belirler.

```go
if age >= 18 {
	fmt.Println("Reşit")
} else {
	fmt.Println("Reşit değil")
}
```

Birden fazla koşul için `else if` kullanılabilir.

## switch

Bir değerin farklı olasılıklarını kontrol etmek için kullanılır.

```go
switch role {
case "admin":
	fmt.Println("Admin")
case "user":
	fmt.Println("User")
default:
	fmt.Println("Bilinmeyen")
}
```

## for

Go'daki temel döngü yapısıdır.

```go
for i := 1; i <= 5; i++ {
	fmt.Println(i)
}
```

Go'da ayrı bir `while` keyword'ü bulunmaz. `for` farklı döngü biçimleri için kullanılır.

## break

İçinde bulunduğu döngüyü sonlandırır.

```go
if i == 8 {
	break
}
```

## continue

Döngünün mevcut turunun geri kalanını atlayarak sonraki tura geçer.

```go
if i == 4 {
	continue
}
```

## range

Slice gibi collection'ların elemanlarını dolaşmak için kullanılır.

```go
for index, value := range languages {
	fmt.Println(index, value)
}
```

## Blank Identifier `_`

Bir fonksiyondan, `range`'den vb. gelen ancak kullanılmayacak bir değeri yok saymak için kullanılır.

```go
for _, value := range languages {
	fmt.Println(value)
}
```

# Aşama 4 — Functions

## Function

Belirli bir işi yapan ve tekrar kullanılabilen kod bloğudur.

```go
func greet() {
	fmt.Println("Hello")
}
```

## Parameter ve Argument

Parameter, function tanımlanırken belirtilen değişkendir.

Argument ise function çağrılırken gönderilen gerçek değerdir.

```go
func greet(name string) {
	fmt.Println(name)
}

greet("Fatih")
```

```text
name    → Parameter
"Fatih" → Argument
```

## Return Value

Function yaptığı işlemin sonucunu geri döndürebilir.

```go
func add(a, b int) int {
	return a + b
}
```

Buradaki son `int`, function'ın return type'ıdır.

## Multiple Return Values

Go function'ları birden fazla değer döndürebilir.

```go
func getUser() (string, int) {
	return "Fatih", 22
}
```

Dönen değerler sırayla variable'lara atanır:

```go
name, age := getUser()
```

```text
"Fatih" → name
22      → age
```

## Named Return Values

Return değerlerine function tanımında isim verilebilir.

```go
func add(a, b int) (result int) {
	result = a + b
	return
}
```

## Function Scope

Function içerisinde oluşturulan variable'lar kendi scope'ları içerisinde kullanılabilir.

```go
func calculate() int {
	result := 10 + 20
	return result
}
```

Başka bir function `result` değişkenine doğrudan erişemez.

## Temel Function Akışı

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
# Aşama 5 — Arrays & Slices

## Array

Aynı type'taki birden fazla değeri sabit uzunlukta tutan veri yapısıdır.

```go
languages := [3]string{"Go", "Python", "Java"}
```

`[3]string`, 3 elemanlı bir `string` array'idir.

## Index

Array ve Slice elemanlarına konumları üzerinden erişilir.

```go
languages[0]
```

Go'da index `0`'dan başlar.

## Slice

Boyutu sabit olmayan, Go'da collection işlemlerinde sık kullanılan yapıdır.

```go
languages := []string{"Go", "Python", "Java"}
```

Temel syntax farkı:

```text
[3]string → Array
[]string  → Slice
```

## append

Slice'a eleman eklemek için kullanılır.

```go
languages = append(languages, "Rust")
```

`append`, güncellenmiş slice'ı geri döndürdüğü için sonuç tekrar variable'a atanır.

## len

Mevcut eleman sayısını verir.

```go
len(languages)
```

## cap

Slice'ın mevcut backing array üzerinde erişebildiği kapasiteyi verir.

```text
len → Mevcut eleman sayısı
cap → Mevcut kapasite
```

## Slicing

Bir slice'ın belirli bir bölümünü almak için kullanılır.

```go
selected := languages[1:4]
```

Kural:

```text
[start:end]

start → dahil
end   → dahil değil
```

## range

Slice elemanlarını sırayla dolaşmayı sağlar.

```go
for index, value := range languages {
	fmt.Println(index, value)
}
```

## Array ve Slice Farkı

```text
Array
→ Sabit uzunluk
→ [3]string

Slice
→ Esnek uzunluk
→ []string
→ append ile büyüyebilir
```
# Aşama 6 — Maps

## Map

Key-value mantığıyla veri tutan collection yapısıdır.

```go
scores := map[string]int{
	"Fatih": 90,
	"Ahmet": 75,
}
```

```text
map[intstring]
     ↓      ↓
    key    value
```

## Değer Okuma

Bir value'ya key üzerinden erişilir.

```go
score := scores["Fatih"]
```

## Eleman Ekleme ve Güncelleme

Aynı syntax kullanılır:

```go
scores["Ayşe"] = 85
```

Key yoksa yeni eleman eklenir, varsa mevcut value güncellenir.

## Key Kontrolü

```go
score, ok := scores["Fatih"]
```

```text
score → Value
ok    → Key mevcut mu?
```

Key varsa `ok = true`, yoksa `ok = false` olur.

Bu kullanım **comma ok idiom** olarak bilinir.

## Zero Value

Olmayan bir key doğrudan okunursa value type'ın zero value'su döner.

Örneğin `map[string]int` için:

```go
scores["Unknown"]
```

key yoksa `0` döner.

Bu nedenle key'in gerçekten var olup olmadığını anlamak için `value, ok` kullanılabilir.

## delete

Map'ten key-value çiftini siler.

```go
delete(scores, "Fatih")
```

## len

Map'teki eleman sayısını verir.

```go
len(scores)
```

## range

Map'teki key-value çiftlerini dolaşır.

```go
for name, score := range scores {
	fmt.Println(name, score)
}
```

Map üzerinde `range` kullanırken belirli bir sıralamaya güvenilmemelidir.

## Temel Map Modeli

```text
Map
 ↓
Key → Value

scores["Fatih"]
       ↓
       90
```

# Aşama 7 — Structs

## Struct

Birbiriyle ilişkili verileri tek bir type altında toplamak için kullanılır.

```go
type User struct {
	Name     string
	Age      int
	IsActive bool
}
```

`User`, bizim oluşturduğumuz yeni bir type'tır.

## Field

Struct içerisindeki değerlere field denir.

```text
Name     → string
Age      → int
IsActive → bool
```

Field'lara `.` ile erişilir:

```go
user.Name
user.Age
```

## Struct Oluşturma

```go
user := User{
	Name:     "Fatih",
	Age:      22,
	IsActive: true,
}
```

## Struct ve Function

Struct, parameter veya return type olarak kullanılabilir.

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
Age: age
 ↑    ↑
Field Parameter
```

Field ve parameter birbirinden farklı kavramlardır.

## Slice of Structs

Birden fazla struct değeri Slice içerisinde tutulabilir:

```go
users := []User{
	user,
	user2,
	user3,
}
```

## Struct + range

```go
for index, user := range users {
	fmt.Println(index, user.Name)
}
```

```text
users → Tüm Slice
index → Elemanın index'i
user  → O turdaki User
```

`user` ismini programcı belirler.

## Exported / Unexported Field

```go
Age int
```

Büyük harfle başladığı için exported'dır.

```go
age int
```

Küçük harfle başladığı için unexported'dır.

Bu ayrım özellikle farklı package'lerle çalışırken önemlidir.

# Aşama 8 — Methods

## Method

Belirli bir type ile ilişkilendirilmiş function'dır.

```go
func (u User) getName() string {
	return u.Name
}
```

Method, ilgili type'tan oluşturulan değer üzerinden çağrılır:

```go
user.getName()
```

## Receiver

Method'un hangi type ile ilişkili olduğunu belirtir.

```go
func (u User) getName() string
```

```text
u    → Receiver variable
User → Receiver type
```

`u` ismini programcı belirler.

## Function ve Method Farkı

```text
Function:
getName(user)

Method:
user.getName()
```

Method bir davranışı belirli bir type ile ilişkilendirir.

## Method Parameter

Method receiver dışında parameter da alabilir:

```go
func (u User) canAccess(minAge int) bool {
	return u.Age >= minAge
}
```

Burada:

```text
u      → Receiver
minAge → Parameter
```

## Method Return Value

Normal function gibi değer döndürebilir:

```go
func (u User) isAdult() bool {
	return u.Age >= 18
}
```

## Value Receiver

```go
func (u User) getName()
```

`(u User)` bir value receiver'dır.

Value receiver ile method, `User` değerinin bilgilerini okuyabilir ve işlemler gerçekleştirebilir.

Orijinal struct üzerinde kalıcı değişiklik yapmak için pointer receiver kullanılabilir.

## Method Çağırma

Bir variable'ın type'ı receiver type ile uyumluysa ilgili method çağrılabilir:

```go
user := User{
	Name: "Fatih",
}

user.getName()
```

```text
user type → User
receiver  → User

Bu nedenle user.getName() kullanılabilir.
```

# Aşama 9 — Pointers

## Pointer

Pointer, başka bir variable'ın memory address'ini tutar.

```go
age := 22
agePointer := &age
```

```text
age        → 22
agePointer → age'in address'i
```

## Memory Address

Bir variable'ın bellekte bulunduğu konumu ifade eder.

```go
fmt.Println(&age)
```

## `&` Operator

Variable'ın memory address'ini verir.

```go
&age
```

```text
age  → Value
&age → Address
```

## `*` Operator ve Dereferencing

Pointer'ın gösterdiği address'teki value'ya erişir.

```go
*agePointer
```

Bu işleme **dereferencing** denir.

```text
agePointer  → Address
*agePointer → Value
```

## Pointer ile Değer Değiştirme

```go
*agePointer = 25
```

Pointer'ın gösterdiği address'teki değer değiştirilir. Böylece orijinal variable da değişir.

## Value Kopyası

```go
agePointerValue := *agePointer
```

Burada pointer'ın gösterdiği değer normal bir variable'a kopyalanır.

Daha sonra:

```go
*agePointer = 25
```

yapılması `agePointerValue` değerini değiştirmez.

## Pointer Type

```text
int   → int value
*int  → int değerine işaret eden pointer

User  → User value
*User → User değerine işaret eden pointer
```

## Pointer Parameter

Function'a address göndererek orijinal değer üzerinde işlem yapılabilir.

```go
func changeAge(age *int) {
	*age = 30
}
```

Çağırma:

```go
changeAge(&age)
```

## Pointer Receiver

Struct üzerinde kalıcı değişiklik yapmak için kullanılabilir.

```go
func (u *User) changeName(newName string) {
	u.Name = newName
}
```

Burada:

```text
u     → Receiver variable
*User → Pointer receiver type
```

## Value Receiver vs Pointer Receiver

```go
func (u User) method()
```

Value receiver'dır.

```go
func (u *User) method()
```

Pointer receiver'dır.

Temel fark:

```text
Value Receiver
→ Değer üzerinde çalışır

Pointer Receiver
→ Orijinal değere pointer üzerinden erişebilir
→ Struct'ı değiştirebilir
```
# Aşama 10 — Interfaces

## Interface

Interface, bir type'ın sahip olması gereken method'ları tanımlar.

```go
type Notifier interface {
	send() string
}
```

Interface method'un nasıl çalışacağını değil, hangi method'un bulunması gerektiğini belirtir.

## Interface Implementation

Bir type, interface'in istediği bütün method'lara sahipse o interface'i implement eder.

```go
type Notifier interface {
	send() string
}

func (e EmailNotifier) send() string {
	return e.Address + " adresine mail gönderildi"
}
```

Burada `EmailNotifier`, `Notifier` interface'ini implement eder.

## Implicit Implementation

Go'da:

```text
implements Notifier
```

gibi bir ifade yazılmaz.

Method signature'ları eşleştiğinde implementation otomatik gerçekleşir.

## Method Signature

Interface'in istediği method ile type'ın method'u uyumlu olmalıdır.

```go
send() string
```

Interface birden fazla method içeriyorsa type bütün method'ları karşılamalıdır.

```go
type Greeter interface {
	greet() string
	getName() string
}
```

## Interface Parameter

Interface function parameter type'ı olarak kullanılabilir.

```go
func sendNotification(n Notifier) {
	fmt.Println(n.send())
}
```

Bu function `Notifier` interface'ini implement eden farklı type'ları kabul edebilir.

```go
sendNotification(email)
sendNotification(sms)
```

## Interface Variable

Interface normal variable type'ı olarak kullanılabilir.

```go
var notifier Notifier
```

Interface'i implement eden değerler atanabilir:

```go
notifier = email
notifier = sms
```

## Concrete Type

Interface variable'ın içerisinde bulunan gerçek type'tır.

Örneğin:

```go
var notifier Notifier
notifier = email
```

Burada:

```text
Variable type → Notifier
Concrete type → EmailNotifier
```

## Polymorphism

Farklı concrete type'ların ortak bir interface üzerinden kullanılabilmesidir.

```text
EmailNotifier ──┐
                ├── Notifier
SMSNotifier ────┘
```

Aynı çağrı:

```go
notifier.send()
```

concrete type'a göre farklı implementation çalıştırabilir.

## Temel Mantık

```text
Interface
↓
Davranışı tanımlar

Concrete Type
↓
Gerekli method'ları sağlar

Function
↓
Concrete type yerine interface'e bağımlı olabilir
```

# Aşama 11 — Error Handling

## Error

Go'da hatalar `error` type'ı ile temsil edilir.

Function hem değer hem error döndürebilir:

```go
func divide(a int, b int) (int, error)
```

```text
int   → başarılı işlem sonucu
error → hata bilgisi
```

## nil

Error için `nil`, hata olmadığını belirtir.

```text
err == nil  → hata yok
err != nil  → hata var
```

Başarılı işlem:

```go
return result, nil
```

## Error Kontrolü

Go'da temel error kontrol pattern'i:

```go
result, err := someFunction()

if err != nil {
	// hata var
	return
}

// hata yok
```

## errors.New()

Sabit bir error oluşturur.

```go
errors.New("cannot divide by zero")
```

Örnek:

```go
if b == 0 {
	return 0, errors.New("cannot divide by zero")
}
```

## fmt.Errorf()

Formatlı error oluşturmak için kullanılır.

```go
fmt.Errorf("cannot multiply %d by %d", a, b)
```

Değişkenleri error mesajına ekleyebilir.

## Value + Error

Go'da yaygın function dönüşlerinden biri:

```go
(value, error)
```

Örneğin:

```go
func createUser(name string, age int) (User, error)
```

Başarılı:

```go
return user, nil
```

Hatalı:

```go
return User{}, err
```

## Zero Value Struct

```go
User{}
```

struct'ın field'larının zero value aldığı halidir.

Örneğin:

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

## Error Propagation

Bir function'ın aldığı error'ı çağıran üst function'a iletmesidir.

```go
result, err := divide(a, b)

if err != nil {
	return 0, err
}
```

Akış:

```text
divide()
   ↓ error
calculate()
   ↓ error
main()
```

## Error Wrapping

Mevcut error'a ek context ekleyerek yukarı taşımaktır.

```go
return 0, fmt.Errorf("calculate failed: %w", err)
```

Buradaki:

```text
%w
```

orijinal error'ı wrap eder.

Örnek:

```text
Orijinal:
cannot divide by zero

Wrapped:
calculate failed: cannot divide by zero
```

## Handle vs Propagate

Error oluştuğunda iki temel seçenek vardır.

### Handle

Hatayı bulunduğun yerde işle:

```go
if err != nil {
	fmt.Println("Error:", err)
	return
}
```

### Propagate

Hatayı çağıran function'a ilet:

```go
if err != nil {
	return 0, err
}
```

veya context ekleyerek:

```go
if err != nil {
	return 0, fmt.Errorf("operation failed: %w", err)
}
```

## Temel Akış

```text
Function çağır
      ↓
value, err
      ↓
err != nil?
   /       \
 Evet      Hayır
  ↓          ↓
Handle /    value
Propagate   kullan
```
# Aşama 12 — Packages & Modules

## Package

Package, birbiriyle ilişkili Go kodlarını gruplamak için kullanılır.

```go
package user
```

Bir package birden fazla `.go` dosyasından oluşabilir.

```text
user/
├── user.go
└── validation.go
```

Her iki dosyada da:

```go
package user
```

bulunuyorsa aynı package'ın parçalarıdır.

## package main

Çalıştırılabilir Go programının ana package'ıdır.

```go
package main

func main() {
}
```

Programın entry point'i:

```go
func main()
```

function'ıdır.

Dosyanın adının `main.go` olması zorunlu değildir.

## Custom Package

Kendi package'larımız oluşturulabilir.

```go
package mathutil

func Add(a, b int) int {
	return a + b
}
```

Başka package'lardan kullanılmak için import edilir.

## Exported Identifier

Büyük harfle başlayan identifier başka package'lardan erişilebilir.

```go
Add
User
CreateUser
IsPositive
```

## Unexported Identifier

Küçük harfle başlayan identifier package dışından erişilemez.

```go
add
user
createUser
isPositive
```

Temel kural:

```text
Aynı package
→ Exported + Unexported kullanılabilir

Başka package
→ Exported kullanılabilir
```

## Module

Module, Go projesinin üst seviyedeki kimliğidir.

`go.mod` içerisinde tanımlanır:

```go
module github.com/thebilici/go-backend-learning
```

Bir module birden fazla package içerebilir.

## Module vs Package

```text
Module
→ Projenin ana kimliği / dependency sınırı

Package
→ Kodları organize eden birim

.go file
→ Package içerisindeki kaynak kod dosyası
```

İlişki:

```text
Module
└── Package
    ├── file.go
    └── file.go
```

## Module Path

`go.mod` içerisindeki:

```go
module github.com/thebilici/go-backend-learning
```

değeri module path'tir.

```text
github.com/thebilici/go-backend-learning
```

## Package Path

Package'ın module içerisindeki klasör yoludur.

Örneğin:

```text
exercises/basics/packages/user
```

## Import Path

Bir package'ın tam import adresidir.

```text
Module Path
+
Package Path
=
Import Path
```

Örnek:

```text
github.com/thebilici/go-backend-learning
+
exercises/basics/packages/user

=

github.com/thebilici/go-backend-learning/exercises/basics/packages/user
```

## Package Kullanımı

Import edilen package'ın exported identifier'ları package adı üzerinden kullanılır:

```go
user.CreateUser()
```

```go
mathutil.Add()
```

Standard library'de de aynı mantık vardır:

```go
fmt.Println()
```

```text
fmt     → package
Println → exported function
```

## Import Alias

Import edilen package'a farklı bir isim verilebilir.

```go
import mathpkg "github.com/example/project/mathutil"
```

Sonra:

```go
mathpkg.Add(5, 6)
```

şeklinde kullanılır.

## go.mod

Module bilgisini ve dependency tanımlarını tutar.

Örneğin:

```go
module github.com/thebilici/go-backend-learning

go 1.26.5
```

## go mod init

Yeni module başlatır.

```bash
go mod init github.com/example/project
```

Temel olarak `go.mod` oluşturur ve module path'i tanımlar.

## go mod tidy

Kod tarafından kullanılan dependency'leri analiz ederek module dependency bilgilerini düzenler.

```text
Eksik dependency
→ eklenebilir

Artık kullanılmayan dependency
→ kaldırılabilir
```

## go.sum

External dependency'lere ait checksum bilgilerini tutar.

Dependency bütünlüğünün doğrulanmasına yardımcı olur.

## Temel Mimari

```text
go.mod
   ↓
Module
   ↓
Packages
   ↓
Go Files
   ↓
Types / Functions
   ↓
Exported olanlar
   ↓
Başka package'lardan kullanılabilir
```
# Aşama 13 — Concurrency

## Concurrency

Concurrency, birden fazla işin aynı zaman aralığında ilerleyebilmesidir.

```text
Sequential:

Task A → biter → Task B → biter → Task C


Concurrent:

Task A ─────────→
Task B    ─────────→
Task C       ─────────→
```

Concurrency ile parallelism aynı şey değildir.

```text
Concurrency
→ Birden fazla işin ilerleyişini aynı zaman aralığında yönetmek

Parallelism
→ Birden fazla işi gerçekten aynı anda çalıştırmak
```

## Goroutine

Goroutine, Go runtime tarafından yönetilen hafif bir concurrent execution unit'tir.

Normal function çağrısı:

```go
task()
```

Goroutine olarak:

```go
go task()
```

`go` keyword'ü function çağrısını yeni bir goroutine olarak başlatır.

`main()` de main goroutine içerisinde çalışır.

## Main Goroutine

Programın `main()` function'ını çalıştıran goroutine'dir.

Main goroutine bittiğinde program sona erer.

Bu nedenle diğer goroutine'lerin tamamlanması gerekiyorsa senkronizasyon sağlanmalıdır.

## sync.WaitGroup

Birden fazla goroutine'in tamamlanmasını beklemek için kullanılır.

```go
var wg sync.WaitGroup
```

### Add

Beklenen iş sayısını artırır.

```go
wg.Add(3)
```

```text
counter = 3
```

### Done

Bir işin tamamlandığını bildirir.

```go
wg.Done()
```

Counter'ı bir azaltır.

```text
3 → 2 → 1 → 0
```

Genellikle:

```go
defer wg.Done()
```

şeklinde kullanılır.

### Wait

Counter `0` olana kadar bulunduğu goroutine'i bekletir.

```go
wg.Wait()
```

Temel model:

```text
Add()
 ↓
Goroutine'leri başlat
 ↓
Done()
 ↓
Counter azalır
 ↓
Wait()
 ↓
Counter = 0
 ↓
Devam et
```

## Channel

Channel, goroutine'ler arasında typed veri iletişimi sağlar.

String channel:

```go
ch := make(chan string)
```

Integer channel:

```go
ch := make(chan int)
```

## Channel Send

Channel'a veri göndermek:

```go
ch <- value
```

Örnek:

```go
ch <- "Hello"
```

## Channel Receive

Channel'dan veri almak:

```go
value := <-ch
```

Örnek:

```go
message := <-ch
```

Temel model:

```text
Goroutine A

value
  │
  ▼
Channel
  │
  ▼

Goroutine B
```

## Unbuffered Channel

Buffer kapasitesi olmayan channel'dır.

```go
ch := make(chan string)
```

Send ve receive birbirleriyle senkronize olur.

```text
Sender ←→ Receiver
```

Receiver hazır değilse sender bekleyebilir.

Sender hazır değilse receiver bekleyebilir.

## Buffered Channel

Belirli kapasitede buffer içeren channel'dır.

```go
ch := make(chan string, 3)
```

Örneğin:

```text
Capacity = 3

[ A ][ B ][ boş ]
```

Buffer dolu olmadığı sürece send işlemi doğrudan bir receiver beklemeden ilerleyebilir.

Buffer dolduğunda yeni send işlemi yer açılana kadar bekleyebilir.

## Channel len ve cap

```go
len(ch)
```

Buffer içerisinde şu anda bulunan değer sayısını verir.

```go
cap(ch)
```

Channel buffer kapasitesini verir.

Örneğin:

```text
[ A ][ B ][ boş ]

len = 2
cap = 3
```

## Channel ile Goroutine Sonucu Alma

Goroutine'in hesapladığı sonuç channel üzerinden başka bir goroutine'e gönderilebilir.

```go
func calculate(a, b int, ch chan int) {
	result := a + b
	ch <- result
}
```

Main:

```go
go calculate(10, 20, ch)

result := <-ch
```

Akış:

```text
calculate goroutine
       ↓
      30
       ↓
    Channel
       ↓
main goroutine
       ↓
    result
```

## close

Channel'a artık yeni değer gönderilmeyeceğini belirtir.

```go
close(ch)
```

Kapalı channel'a tekrar veri göndermek runtime panic oluşturur.

Channel genellikle veri gönderen tarafın işi tamamlandığında kapatılır.

## range ile Channel Okuma

Channel'dan değerleri okumak için:

```go
for value := range ch {
	fmt.Println(value)
}
```

`range`, channel kapanana ve mevcut değerler tüketilene kadar okumaya devam eder.

## WaitGroup + Channel + close

Birden fazla sender olduğunda:

```go
wg.Add(3)

go worker(...)
go worker(...)
go worker(...)
```

Worker'ların tamamlanması beklenebilir:

```go
go func() {
	wg.Wait()
	close(ch)
}()
```

Receiver:

```go
for result := range ch {
	fmt.Println(result)
}
```

Temel akış:

```text
Workers
   ↓
Channel
   ↓
Receiver

Workers tamamlandı
   ↓
WaitGroup = 0
   ↓
close(channel)
   ↓
range tamamlanır
```

## Deadlock

Goroutine'lerin birbirlerini beklediği ve hiçbirinin ilerleyemediği durumdur.

Örneğin unbuffered channel:

```go
ch := make(chan string)

ch <- "Hello"

message := <-ch
```

Aynı goroutine send sırasında receiver beklediği için sonraki receive satırına ulaşamaz.

Bu deadlock oluşturabilir.

## Race Condition

Birden fazla goroutine aynı shared data üzerinde eşzamanlı ve uygun senkronizasyon olmadan işlem yaptığında oluşabilir.

Örneğin:

```go
counter++
```

iki goroutine tarafından aynı anda çalıştırılırsa beklenmeyen sonuçlar oluşabilir.

Race detector:

```bash
go run -race .
```

## Mutex

Shared data'ya erişimi senkronize etmek için kullanılır.

```go
var mu sync.Mutex
```

Kritik bölgeyi kilitlemek:

```go
mu.Lock()

counter++

mu.Unlock()
```

Bu sayede aynı kritik bölgeye aynı anda bir goroutine'in girmesi sağlanabilir.

## WaitGroup vs Channel vs Mutex

```text
WaitGroup
→ İşlerin tamamlanmasını beklemek

Channel
→ Goroutine'ler arasında veri/işaret iletmek

Mutex
→ Shared data'ya erişimi korumak
```

## Temel Concurrency Modeli

```text
                Goroutines
               /    |     \
              /     |      \
             ↓      ↓       ↓
           Worker Worker  Worker
              \     |      /
               \    |     /
                    ↓
                 Channel
                    ↓
                 Receiver

WaitGroup
→ Worker yaşam döngüsünü takip edebilir

Mutex
→ Shared state varsa erişimi koruyabilir
```

# Aşama 15 — Generics

## Generics

Generics, aynı kodun farklı veri tipleriyle type-safe şekilde çalışmasını sağlar.

Örneğin farklı tipler için ayrı function yazmak yerine:

```go
func Add[T int | float64](a T, b T) T {
	return a + b
}
```

kullanılabilir.

---

## Type Parameter

Generic yapı içerisinde kullanılacak tipi temsil eder.

```go
func PrintValue[T any](value T)
```

Burada:

```text
T → Type Parameter
```

---

## Type Argument

Generic yapı kullanılırken verilen gerçek tiptir.

```go
Response[string]
Response[int]
```

Burada:

```text
string → Type Argument
int    → Type Argument
```

---

## Type Constraint

Type Parameter'ın hangi tipleri kabul edebileceğini belirler.

```go
[T int | float64]
```

Burada `T` yalnızca:

```text
int
veya
float64
```

olabilir.

---

## any

`any`, herhangi bir Go tipini kabul eder.

```go
func GetFirst[T any](arr []T) T {
	return arr[0]
}
```

Burada `T`:

```text
int
string
bool
float64
struct
...
```

gibi farklı tipler olabilir.

---

## comparable

`comparable`, `==` ve `!=` ile karşılaştırılabilen tipleri kabul eder.

```go
func IsEqual[T comparable](a T, b T) bool {
	return a == b
}
```

`==` karşılaştırması yapıldığı için `comparable` kullanılır.

---

## Generic Slice

```go
[]T
```

`T` tipinde elemanlardan oluşan slice anlamına gelir.

```go
func GetFirst[T any](arr []T) T {
	return arr[0]
}
```

Örneğin:

```text
[]int    → T = int
[]string → T = string
```

---

## Birden Fazla Type Parameter

Generic yapılarda birden fazla type parameter kullanılabilir.

```go
func PrintPair[K any, V any](key K, value V) {
	fmt.Println(key, value)
}
```

Burada:

```text
K → Birinci tip
V → İkinci tip
```

---

## Generic Struct

Struct'lar generic olarak tanımlanabilir.

```go
type Response[T any] struct {
	Data    T
	Success bool
}
```

Kullanım:

```go
Response[string]{
	Data: "Fatih",
}

Response[int]{
	Data: 100,
}
```

Aynı struct farklı veri tipleriyle kullanılabilir.

---

## Custom Constraint

Constraint'ler interface kullanılarak ayrı tanımlanabilir.

```go
type Number interface {
	int | float64
}
```

Daha sonra:

```go
func Sum[T Number](a T, b T) T {
	return a + b
}
```

kullanılabilir.

---

## Generic Map

Normal map:

```go
map[string]int
```

şu anlama gelir:

```text
string → Key tipi
int    → Value tipi
```

Generic karşılığı:

```go
map[K]V
```

Burada:

```text
K → Key tipi
V → Value tipi
```

Örneğin:

```go
func GetValue[K comparable, V any](data map[K]V, key K) (V, bool) {
	value, exists := data[key]
	return value, exists
}
```

`map[string]int` gönderildiğinde:

```text
K = string
V = int
```

olur.

`map[int]string` gönderildiğinde:

```text
K = int
V = string
```

olur.

---

## `(V, bool)` Mantığı

```go
func GetValue[K comparable, V any](data map[K]V, key K) (V, bool)
```

şöyle okunur:

```text
key K
→ K tipinde bir key al

V
→ Bulunan değeri V tipinde döndür

bool
→ Değer bulundu mu bilgisini döndür
```

Örneğin:

```go
score, exists := GetValue(students, "Fatih")
```

sonucu:

```text
score  → 100
exists → true
```

olabilir.

---

## Temel Generics Mantığı

```text
Normal:

map[string]int

string → Key
int    → Value


Generic:

map[K]V

K → Key tipi
V → Value tipi
```

Generics'in temel amacı:

> Aynı algoritmayı veya veri yapısını farklı tipler için tekrar kod yazmadan, type-safe şekilde kullanabilmektir.