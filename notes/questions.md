# Go Tekrar Soruları

## Aşama 1 — Go Ortamı ve Temeller

1. Go Module nedir ve ne işe yarar?

2. `go.mod` dosyasının görevi nedir?

3. Yeni bir Go Module oluşturmak için hangi komut kullanılır?

4. Module ile Package arasındaki fark nedir?

5. Bir Package birden fazla `.go` dosyasından oluşabilir mi?

6. `package main` neden özeldir?

7. `func main()` ne işe yarar?

8. Entry Point ne demektir?

9. Go Standard Library nedir? Üç örnek package ver.

10. `import "fmt"` neden gereklidir?

11. `fmt.Println()` ifadesindeki `fmt` ve `Println` neyi temsil eder?

12. Exported ve Unexported Identifier arasındaki fark nedir?

13. `go run main.go` ile `go run .` arasındaki fark nedir?

14. `go build` ne yapar?

15. `go run` ile `go build` arasındaki temel fark nedir?

16. Source Code ile Binary arasındaki fark nedir?

17. `go build .` sonucunda neden executable dosyası oluşur?

18. `go fmt` ne işe yarar?

19. Go'da kullanılmayan bir package import edilirse ne olur?

20. Aşağıdaki süreci açıklayabilir misin?

```text
Go Source Code
      ↓
Package
      ↓
Go Compiler
      ↓
Binary
      ↓
Operating System
      ↓
CPU
```
# Aşama 2 — Variables & Data Types

1. Variable nedir?

2. `var` keyword'ü ne işe yarar?

3. Aşağıdaki üç tanımlama arasındaki fark nedir?

```go
var age int = 22
var age = 22
age := 22
```

4. `:=` ile `=` arasındaki fark nedir?

5. Declaration, initialization ve assignment kavramları arasındaki fark nedir?

6. `string`, `int`, `float64` ve `bool` hangi tür değerleri tutar?

7. Go'nun statically typed olması ne anlama gelir?

8. Type inference nedir?

9. Aşağıdaki değişkenlerin type'ları nedir?

```go
name := "Fatih"
age := 22
height := 1.73
isActive := true
```

10. Zero Value nedir? `string`, `int` ve `bool` için zero value değerleri nelerdir?

11. `const` ile `var` arasındaki temel fark nedir?

12. Aşağıdaki kod neden çalışmaz?

```go
age := 22
age = "22"
```

13. `int` bir değer nasıl `float64` type'ına dönüştürülür?

14. `strconv.Atoi()` ne işe yarar?

15. `strconv.Itoa()` ne işe yarar?

16. `"25"` ile `25` arasındaki fark nedir?

17. `int` ile `int64` aynı type mıdır?

18. `byte` hangi type'ın alias'ıdır?

19. `rune` hangi type'ın alias'ıdır?

20. Aşağıdaki dönüşüm akışını açıklayabilir misin?

```text
"25"
  ↓
strconv.Atoi
  ↓
25
  ↓
float64(...)
  ↓
25.0
```

# Aşama 3 — Control Flow

1. Control Flow nedir?

2. Comparison operator'ları hangi type'ta sonuç üretir?

3. `=` ile `==` arasındaki fark nedir?

4. `&&`, `||` ve `!` operator'ları ne işe yarar?

5. Aşağıdaki koşul ne zaman `true` olur?

```go
age >= 18 && isActive
```

6. `if`, `else if` ve `else` arasındaki fark nedir?

7. `switch` hangi durumlarda `if / else if` yerine tercih edilebilir?

8. `switch` içerisindeki `default` ne zaman çalışır?

9. Go'da temel döngü yapısı nedir?

10. Aşağıdaki döngü kaç kez çalışır?

```go
for i := 1; i <= 5; i++ {
	fmt.Println(i)
}
```

11. Go'da neden ayrı bir `while` keyword'ü yoktur?

12. `break` ne işe yarar?

13. `continue` ne işe yarar?

14. `break` ile `continue` arasındaki fark nedir?

15. Aşağıdaki kodun çıktısı nedir?

```go
for i := 1; i <= 5; i++ {
	if i == 3 {
		continue
	}

	fmt.Println(i)
}
```

16. `range` ne işe yarar?

17. `range` ile bir slice dolaşılırken `index` ve `value` neyi temsil eder?

18. Blank Identifier `_` neden kullanılır?

19. Aşağıdaki kodda neden `4` yazdırılmaz?

```go
for i := 1; i <= 5; i++ {
	if i == 4 {
		continue
	}

	fmt.Println(i)
}
```

20. `break` kontrolünün `fmt.Println()` öncesinde veya sonrasında olması programın çıktısını nasıl etkiler?

# Aşama 4 — Functions

1. Function nedir ve neden kullanılır?

2. Go'da function hangi keyword ile tanımlanır?

3. Parameter ile argument arasındaki fark nedir?

4. Aşağıdaki kodda parameter ve argument hangisidir?

```go
func greet(name string) {
	fmt.Println(name)
}

greet("Fatih")
```

5. Return value nedir?

6. Aşağıdaki function'daki son `int` ne anlama gelir?

```go
func add(a, b int) int {
	return a + b
}
```

7. `return` keyword'ü ne işe yarar?

8. Multiple Return Values nedir?

9. Aşağıdaki function kaç değer döndürür ve type'ları nelerdir?

```go
func getUser() (string, int) {
	return "Fatih", 22
}
```

10. Aşağıdaki satırda `name` ve `age` hangi değerleri alır?

```go
name, age := getUser()
```

11. Bir function üç farklı değer döndürebilir mi?

12. Multiple return kullanırken `_` ne işe yarar?

13. Named Return Value nedir?

14. Function Scope ne anlama gelir?

15. Bir function içerisinde oluşturulan local variable'a başka bir function doğrudan erişebilir mi?

# Aşama 5 — Arrays & Slices

1. Array nedir?

2. Slice nedir?

3. `[3]string` ile `[]string` arasındaki fark nedir?

4. Go'da index kaçtan başlar?

5. Bir Array veya Slice'ın ilk elemanına nasıl erişilir?

6. `len()` ne işe yarar?

7. `append()` ne işe yarar?

8. Neden aşağıdaki şekilde assignment yapıyoruz?

```go
languages = append(languages, "Rust")
```

9. `len` ile `cap` arasındaki fark nedir?

10. Slice'ın capacity değeri dolduğunda yeni bir eleman `append` edilirse ne olabilir?
Slice'ın capacity değeri dolduğunda `append()` yapılırsa Go daha büyük bir backing array oluşturabilir, mevcut elemanları yeni alana taşıyabilir ve slice'ın capacity değeri artabilir.

11. Slicing nedir?

12. Aşağıdaki kod hangi elemanları alır?

```go
selected := languages[1:4]
```

13. Slicing işleminde `start` ve `end` değerlerinden hangisi dahil değildir?

14. `languages[:3]` ne anlama gelir?

15. `languages[2:]` ne anlama gelir?

16. `range` ne işe yarar?

17. `range` kullanırken `index` ve `value` neyi temsil eder?

18. Slice içerisindeki mevcut bir eleman nasıl değiştirilir?

19. Array neden `append()` ile büyütülemez?

20. Aynı scope içerisinde neden iki kez aşağıdaki gibi `:=` kullanamayız?

```go
languages := [3]string{"Go", "Python", "Java"}
languages := []string{"Go", "Python", "Java"}
```

# Aşama 6 — Maps

1. Map nedir?

2. Map ile Slice arasındaki temel erişim farkı nedir?

3. `map[string]int` ifadesindeki `string` ve `int` neyi temsil eder?

4. Aşağıdaki Map'ten `"Fatih"` değerine nasıl erişilir?

```go
scores := map[string]int{
	"Fatih": 90,
}
```

5. Map'e yeni bir key-value çifti nasıl eklenir?

6. Map'teki mevcut bir value nasıl güncellenir?

7. Olmayan bir key doğrudan okunursa ne olur?

8. `map[string]int` içinde olmayan bir key okunursa neden `0` döner?

9. Aşağıdaki kodda `score` ve `ok` neyi temsil eder?

```go
score, ok := scores["Fatih"]
```

10. `ok == false` ne anlama gelir?

11. `value, ok` kontrolüne neden ihtiyaç duyabiliriz?

12. Map'ten bir eleman nasıl silinir?

13. Olmayan bir key için `delete()` çağrılırsa ne olur?

14. `len(scores)` neyi verir?

15. Map içerisindeki tüm key-value çiftleri nasıl dolaşılır?

16. Map üzerinde `range` kullanırken çıktı sırasına güvenebilir miyiz?

17. Aşağıdaki kod ne yapar?

```go
scores["Mehmet"] = 70
```

18. `:=` ile `=` arasındaki fark nedir?

19. Neden ikinci kullanımda `=` tercih edilebilir?

```go
score, ok := scores["Ahmet"]
score, ok = scores["Ali"]
```

20. Backend uygulamasında `"değer 0"` ile `"kayıt bulunamadı"` arasındaki farkı nasıl tespit edebiliriz?

# Aşama 7 — Structs

1. Struct nedir ve neden kullanılır?

2. `type User struct` ifadesi ne anlama gelir?

3. Struct içerisindeki `Name`, `Age` gibi değerlere ne ad verilir?

4. Aşağıdaki kodda `User` ve `user` arasındaki fark nedir?

```go
user := User{
	Name: "Fatih",
	Age:  22,
}
```

5. Struct field'larına nasıl erişilir?

6. Bir struct field'ının değeri nasıl güncellenir?

7. Struct function'a parameter olarak gönderilebilir mi?

8. Function bir struct return edebilir mi?

9. Aşağıdaki `User` neyi ifade eder?

```go
func createUser(name string, age int) User
```

10. Aşağıdaki kodda soldaki `Age` ile sağdaki `age` arasındaki fark nedir?

```go
User{
	Age: age,
}
```

11. Function parameter'ının adı ile struct field'ının adı aynı olmak zorunda mıdır?

12. Struct içerisinde tanımlanmayan bir field, `User{...}` içerisinde kullanılabilir mi?

13. `[]User` ne anlama gelir?

14. Aşağıdaki kodda `users` ve `user` arasındaki fark nedir?

```go
for _, user := range users {
	fmt.Println(user.Name)
}
```

15. `range` içerisindeki `user` variable ismini kim belirler?

16. Slice üzerinde `range` kullanıldığında `index` ve `value` neyi temsil eder?

17. Slice'taki `index, value` ile Map'teki `key, value` arasındaki fark nedir?

18. `append()` ile `[]User` içerisine yeni bir User nasıl eklenir?

19. Büyük harfle başlayan struct field ne anlama gelir?

20. Aşağıdaki iki field arasındaki temel fark nedir?

```go
Age int
age int
```

# Aşama 8 — Methods

1. Method nedir?

2. Function ile Method arasındaki temel fark nedir?

3. Receiver nedir?

4. Aşağıdaki kodda `u` ve `User` neyi temsil eder?

```go
func (u User) getName() string
```

5. Receiver variable'ın ismini kim belirler?

6. Neden aşağıdaki şekilde method çağırabiliriz?

```go
user.getName()
```

7. Method parameter alabilir mi?

8. Aşağıdaki kodda receiver ve parameter hangileridir?

```go
func (u User) canAccess(minAge int) bool
```

9. Method bir değer return edebilir mi?

10. Aşağıdaki method ne döndürür?

```go
func (u User) isAdult() bool {
	return u.Age >= 18
}
```

11. `user.greet("Hello")` içerisindeki `"Hello"` nedir?

12. Value Receiver nedir?

13. Aşağıdaki receiver neden Value Receiver'dır?

```go
(u User)
```

14. Value Receiver içerisinde `u.Name` değiştirilirse orijinal `user.Name` neden değişmez?

15. Orijinal struct üzerinde kalıcı değişiklik yapmak istediğimizde hangi receiver türünü kullanabiliriz?

16. Method içerisinde struct field'larına nasıl erişilir?

17. `User` type'ına ait bir method neden `Product` type'ındaki bir variable üzerinden doğrudan çağrılamaz?

18. Aşağıdaki iki kullanım arasındaki fark nedir?

```go
printUser(user)

user.printUser()
```

19. Receiver ile normal function parameter'ı aynı şey midir?

20. Method kullanmanın struct + function kullanımına göre sağladığı temel avantaj nedir?

# Aşama 9 — Pointers

1. Pointer nedir?

2. Memory address nedir?

3. `&` operator ne işe yarar?

4. Aşağıdaki kodda `agePointer` ne tutar?

```go
age := 22
agePointer := &age
```

5. `agePointer` ile `*agePointer` arasındaki fark nedir?

6. Dereferencing nedir?

7. Aşağıdaki kod ne yapar?

```go
*agePointer = 25
```

8. `*int` ne anlama gelir?

9. `*User` ne anlama gelir?

10. Aşağıdaki kodda `agePointerValue` pointer mıdır?

```go
agePointerValue := *agePointer
```

11. `agePointerValue := *agePointer` yapıldıktan sonra `*agePointer` değiştirilirse neden `agePointerValue` değişmez?

12. Aşağıdaki iki function arasındaki fark nedir?

```go
func changeAge(age int)

func changeAge(age *int)
```

13. Pointer parameter alan function'a neden `&age` göndeririz?

14. Aşağıdaki function neden orijinal değeri değiştirebilir?

```go
func changeAge(age *int) {
	*age = 30
}
```

15. Value Receiver nedir?

16. Pointer Receiver nedir?

17. Aşağıdaki iki receiver arasındaki fark nedir?

```go
(u User)

(u *User)
```

18. Struct üzerinde kalıcı değişiklik yapmak için hangisi kullanılabilir?

19. Aşağıdaki method neden orijinal `user.Name` değerini değiştirebilir?

```go
func (u *User) changeName(newName string) {
	u.Name = newName
}
```

20. `u` bir `*User` olmasına rağmen neden field'a şu şekilde erişebiliriz?

```go
u.Name
```

# Aşama 10 — Interfaces

1. Interface nedir?

2. Interface içerisinde method'un gövdesi neden yazılmaz?

3. Aşağıdaki interface bir type'dan ne ister?

```go
type Notifier interface {
	send() string
}
```

4. Implicit Implementation nedir?

5. Go'da neden `implements Notifier` gibi bir ifade yazmamıza gerek yoktur?

6. Bir type'ın `Notifier` interface'ini implement etmesi için ne gerekir?

7. Interface birden fazla method içeriyorsa concrete type kaç tanesini implement etmek zorundadır?

8. Method signature'larının interface ile uyumlu olması neden önemlidir?

9. Aşağıdaki function'daki `n` ve `Notifier` neyi temsil eder?

```go
func sendNotification(n Notifier) {
	fmt.Println(n.send())
}
```

10. Neden aşağıdaki iki çağrı aynı function ile çalışabilir?

```go
sendNotification(email)
sendNotification(sms)
```

11. Interface'i function parameter type'ı olarak kullanmanın avantajı nedir?

12. Aşağıdaki kod ne oluşturur?

```go
var notifier Notifier
```

13. Neden aynı variable'a aşağıdaki iki değer de atanabilir?

```go
notifier = email
notifier = sms
```

14. Interface variable ile concrete type arasındaki fark nedir?

15. Aşağıdaki durumda interface type ve concrete type nedir?

```go
var notifier Notifier
notifier = email
```

16. Polymorphism nedir?

17. `notifier.send()` çağrıldığında hangi `send()` method'unun çalışacağı nasıl belirlenir?

18. `EmailNotifier` gerekli bütün method'lara sahipken `SMSNotifier` bir method'u eksik bırakırsa ne olur?

19. Interface kullanmak neden `EmailNotifier`, `SMSNotifier` gibi concrete type'lara olan bağımlılığı azaltabilir?

20. Aşağıdaki yapının interface + polymorphism açısından çalışma mantığını açıklayın:

```go
type Notifier interface {
	send() string
}

func sendNotification(n Notifier) {
	fmt.Println(n.send())
}

sendNotification(email)
sendNotification(sms)
```
# Aşama 11 — Error Handling

1. Go'da `error` nedir?

2. Aşağıdaki function signature ne anlama gelir?

```go
func divide(a int, b int) (int, error)
```

3. Error açısından `nil` ne anlama gelir?

4. `err == nil` ile `err != nil` arasındaki fark nedir?

5. Aşağıdaki kod neyi kontrol eder?

```go
if err != nil {
	return
}
```

6. `errors.New()` ne işe yarar?

7. `fmt.Errorf()` ne işe yarar?

8. `errors.New()` ile `fmt.Errorf()` arasındaki temel fark nedir?

9. Başarılı bir işlemde neden genellikle aşağıdaki gibi `nil` döndürülür?

```go
return result, nil
```

10. Aşağıdaki kodda `result` ve `err` neyi temsil eder?

```go
result, err := divide(10, 2)
```

11. Error propagation nedir?

12. Aşağıdaki kod ne yapar?

```go
if err != nil {
	return 0, err
}
```

13. `divide()` → `calculate()` → `main()` zincirinde `divide()` içinde oluşan bir error nasıl `main()`e ulaşabilir?

14. Error handling ile error propagation arasındaki fark nedir?

15. Error wrapping nedir?

16. `%w` ne işe yarar?

```go
fmt.Errorf("calculate failed: %w", err)
```

17. Aşağıdaki iki kullanım arasındaki fark nedir?

```go
return 0, err
```

```go
return 0, fmt.Errorf("calculate failed: %w", err)
```

18. `(User, error)` döndüren bir function hata durumunda neden aşağıdaki gibi dönebilir?

```go
return User{}, err
```

19. `User{}` ne anlama gelir?

20. Aşağıdaki kodun error akışını baştan sona açıklayın:

```go
func createUser(name string, age int) (User, error) {
	err := validateAge(age)

	if err != nil {
		return User{}, fmt.Errorf("create user failed: %w", err)
	}

	return User{
		Name: name,
		Age:  age,
	}, nil
}
```
# Aşama 12 — Packages & Modules

1. Package nedir?

2. `package main` ne anlama gelir?

3. Go programının gerçek entry point'i nedir?

4. Dosyanın adının mutlaka `main.go` olması gerekir mi?

5. Bir package birden fazla `.go` dosyasından oluşabilir mi?

6. Aşağıdaki iki dosya neden aynı package'ın parçasıdır?

```text
user/
├── user.go        → package user
└── validation.go  → package user
```

7. Exported identifier nedir?

8. Unexported identifier nedir?

9. `CreateUser()` başka bir package'dan kullanılabilirken `createUser()` neden kullanılamaz?

10. Go'da `public` / `private` yerine hangi mekanizma kullanılır?

11. Module nedir?

12. Package ile Module arasındaki fark nedir?

13. `go.mod` dosyasının görevi nedir?

14. Aşağıdaki satır neyi tanımlar?

```go
module github.com/thebilici/go-backend-learning
```

15. Module path nedir?

16. Package path nedir?

17. Import path nasıl oluşturulur?

18. Aşağıdaki yapıda module path, package path ve import path'i ayrı ayrı belirtin:

```text
Module:
github.com/thebilici/go-backend-learning

Package klasörü:
exercises/basics/packages/user
```

19. Aşağıdaki import'un tamamı module path midir? Değilse nedir?

```go
import "github.com/thebilici/go-backend-learning/exercises/basics/packages/user"
```

20. `user.CreateUser()` ifadesindeki `user` ve `CreateUser` neyi temsil eder?

21. `fmt.Println()` ile kendi oluşturduğumuz `mathutil.Add()` arasında package kullanımı açısından nasıl bir benzerlik vardır?

22. Import alias nedir?

23. Aşağıdaki kodda `mathpkg` neyi temsil eder?

```go
import mathpkg "github.com/example/project/mathutil"
```

24. `go mod init` ne işe yarar?

25. `go mod tidy` ne işe yarar?

26. `go.sum` ne işe yarar?

27. Aynı module içerisindeki her package için ayrı `go.mod` oluşturmak gerekir mi?

28. Aşağıdaki yapıyı Module → Package → File açısından açıklayın:

```text
go-backend-learning/
│
├── go.mod
│
└── mathutil/
    ├── math.go
    └── validation.go
```

29. `math.go` içerisindeki bir type neden `validation.go` içerisinde ayrıca import edilmeden kullanılabilir?

30. Aşağıdaki zinciri kendi cümlelerinizle açıklayın:

```text
Module Path
      +
Package Path
      ↓
Import Path
      ↓
import
      ↓
package.Function()
```
# Aşama 13 — Concurrency

1. Concurrency nedir?

2. Sequential çalışma ile concurrent çalışma arasındaki fark nedir?

3. Concurrency ile parallelism arasındaki fark nedir?

4. Goroutine nedir?

5. Normal function çağrısı ile aşağıdaki çağrı arasındaki fark nedir?

```go
go task()
```

6. `main()` function'ı da bir goroutine içerisinde mi çalışır?

7. Main goroutine diğer goroutine'ler tamamlanmadan biterse ne olur?

8. `sync.WaitGroup` ne işe yarar?

9. Aşağıdaki kod ne anlama gelir?

```go
wg.Add(3)
```

10. `wg.Done()` ne yapar?

11. Neden aşağıdaki kullanım tercih edilir?

```go
defer wg.Done()
```

12. `wg.Wait()` ne yapar ve nereye yazıldığı neden önemlidir?

13. Channel nedir?

14. Aşağıdaki kod ne oluşturur?

```go
ch := make(chan string)
```

15. Channel'a veri nasıl gönderilir?

16. Channel'dan veri nasıl alınır?

17. Aşağıdaki iki işlemi açıklayın:

```go
ch <- "Hello"

message := <-ch
```

18. Unbuffered channel nedir?

19. Unbuffered channel'da sender neden receiver'ı bekleyebilir?

20. Aşağıdaki kod neden deadlock oluşturabilir?

```go
ch := make(chan string)

ch <- "Hello"

message := <-ch
```

21. Buffered channel nedir?

22. Aşağıdaki `3` ne anlama gelir?

```go
ch := make(chan int, 3)
```

23. `len(ch)` ile `cap(ch)` arasındaki fark nedir?

24. Buffered channel tamamen dolduğunda yeni bir send yapılırsa ne olabilir?

25. Bir goroutine'in hesapladığı sonucu `main` goroutine'e nasıl gönderebiliriz?

26. Neden aşağıdaki kullanım yapılamaz?

```go
result := go calculate()
```

27. Birden fazla goroutine aynı channel'a veri gönderebilir mi?

28. Birden fazla goroutine aynı channel'a sonuç gönderdiğinde sonuçların geliş sırasına güvenebilir miyiz?

29. `close(ch)` ne anlama gelir?

30. Channel neden genellikle tüm sender'lar tamamlandıktan sonra kapatılır?

31. Kapalı channel'a send yapılırsa ne olur?

32. `for value := range ch` ne işe yarar?

33. `range` ile okunan channel hiç kapatılmazsa ne tür bir problem oluşabilir?

34. Aşağıdaki pattern'i açıklayın:

```go
go func() {
	wg.Wait()
	close(ch)
}()

for result := range ch {
	fmt.Println(result)
}
```

35. Race condition nedir?

36. `counter++` işlemini iki goroutine'in aynı anda yapması neden problem oluşturabilir?

37. Go race detector nasıl çalıştırılır?

```bash
go run -race .
```

38. `sync.Mutex` ne işe yarar?

39. `Lock()` ve `Unlock()` ne işe yarar?

40. WaitGroup, Channel ve Mutex arasındaki temel görev farklarını açıklayın.

```text
WaitGroup → ?
Channel   → ?
Mutex     → ?
```

41. Aşağıdaki kodu baştan sona açıklayın:

```go
func calculateSquare(a int, results chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	results <- a * a
}
```

42. Aşağıdaki concurrency akışını kendi cümlelerinizle açıklayın:

```text
main
 ↓
wg.Add(3)
 ↓
3 goroutine başlat
 ↓
hesaplamalar
 ↓
results channel
 ↓
wg.Done()
 ↓
wg.Wait()
 ↓
close(results)
 ↓
range tamamlanır
```
# Aşama 15 — Generics Review Questions

## Temel Kavramlar

1. Generics nedir ve hangi problemi çözmek için kullanılır?

2. Type Parameter nedir?

3. Type Argument nedir?

4. Type Parameter ile Type Argument arasındaki fark nedir?

5. Type Constraint nedir?

6. `[T any]` ne anlama gelir?

7. `[T comparable]` ne anlama gelir?

8. `any` ile `comparable` arasındaki fark nedir?

---

## Generic Functions

9. Aşağıdaki kodda `T` neyi temsil eder?

```go
func GetFirst[T any](arr []T) T
```

10. `[]T` ne anlama gelir?

11. Aşağıdaki çağrıda `T` hangi tip olur?

```go
GetFirst([]string{"Go", "Java"})
```

12. Aşağıdaki generic function neden `comparable` kullanmalıdır?

```go
func Contains[T comparable](arr []T, target T) bool
```

13. Bir generic function birden fazla Type Parameter kullanabilir mi?

14. Type Inference nedir?

15. Aşağıdaki çağrıda Go `T` tipini nasıl belirler?

```go
Sum(10, 20)
```

---

## K ve V Mantığı

16. Aşağıdaki yapıda `K` ve `V` neyi temsil eder?

```go
map[K]V
```

17. Generic map yapılarında `K` neden genellikle `comparable` constraint'ine sahiptir?

18. Aşağıdaki function'ı kendi cümlelerinle açıkla:

```go
func GetValue[K comparable, V any](data map[K]V, key K) (V, bool)
```

19. Yukarıdaki function'daki `(V, bool)` neden iki değer döndürür?

20. Aşağıdaki map `GetValue()` function'ına gönderildiğinde `K` ve `V` hangi tip olur?

```go
students := map[string]int{
	"Fatih": 100,
}
```

21. Aşağıdaki map gönderildiğinde `K` ve `V` hangi tip olur?

```go
users := map[int]string{
	1: "Fatih",
}
```

22. `K`, `V` ve `T` Go'nun özel keyword'leri midir?

---

## Generic Struct

23. Generic Struct nedir?

24. Aşağıdaki struct'ta `T` neyi temsil eder?

```go
type Response[T any] struct {
	Data    T
	Success bool
}
```

25. Aşağıdaki iki kullanım arasındaki fark nedir?

```go
Response[string]
Response[int]
```

26. `Response[string]` kullanıldığında `Data` field'ının tipi ne olur?

27. `Response[int]` kullanıldığında `Data` field'ının tipi ne olur?

---

## Custom Constraints

28. Custom Constraint nedir?

29. Aşağıdaki constraint ne anlama gelir?

```go
type Number interface {
	int | float64
}
```

30. Aşağıdaki function neden hem `int` hem de `float64` ile çalışabilir?

```go
func Sum[T Number](a T, b T) T {
	return a + b
}
```

---

## Kod Okuma

31. Aşağıdaki function'ı satır satır açıkla:

```go
func GetValue[K comparable, V any](data map[K]V, key K) (V, bool) {
	value, exists := data[key]
	return value, exists
}
```

32. Aşağıdaki çağrıda:

```go
score, exists := GetValue(students, "Fatih")
```

`score` neyi temsil eder?

33. Aynı çağrıda `exists` neyi temsil eder?

34. Aranan key map içerisinde bulunmazsa `V` için ne döner?

35. Aranan key map içerisinde bulunmazsa `bool` değeri ne olur?

---

## Genel Anlama

36. Normal:

```go
map[string]int
```

yapısı generic olarak nasıl ifade edilebilir?

37. `any` kullanmak yerine neden bazen daha dar bir constraint kullanmak gerekir?

38. `==` kullanan generic bir function'da neden doğrudan `any` kullanamayabiliriz?

39. Generics'in type safety açısından avantajı nedir?

40. Generics hangi durumda gereksiz olabilir?

41. Aynı işlemi yalnızca tek bir veri tipi üzerinde yapıyorsak Generics kullanmak mantıklı mıdır?

42. Generics'in temel amacını kendi cümlelerinle açıkla.

# Aşama 16 — Context Review Questions

## Temel Kavramlar

1. Context nedir ve hangi amaçlarla kullanılır?

2. Cancellation nedir?

3. Timeout nedir?

4. Deadline nedir?

5. Timeout ile Deadline arasındaki fark nedir?

---

## context.Context

6. `context.Context` nedir?

7. Aşağıdaki function'da `ctx` neyi temsil eder?

```go
func worker(ctx context.Context)
```

8. Aynı function'da `context.Context` neyi temsil eder?

9. `context` burada neden yazılmıştır?

10. `Context` neden büyük harfle başlamaktadır?

---

## Background

11. `context.Background()` ne işe yarar?

12. Aşağıdaki kodda `ctx` hangi tiptedir?

```go
ctx := context.Background()
```

13. `context.Background()` başlangıçta timeout veya cancellation içerir mi?

---

## Cancellation

14. `context.WithCancel()` ne işe yarar?

15. Aşağıdaki kod hangi iki değeri döndürür?

```go
ctx, cancel := context.WithCancel(ctx)
```

16. `cancel()` ne yapar?

17. `cancel()` bir goroutine'i doğrudan öldürür mü?

18. Goroutine cancellation sinyalini nasıl öğrenir?

---

## ctx.Done()

19. `ctx.Done()` nedir?

20. `ctx.Done()` neden `select` içerisinde kullanılabilir?

21. Aşağıdaki kod neyi kontrol eder?

```go
case <-ctx.Done():
	return
```

22. Context iptal edildiğinde neden `return` kullanıyoruz?

---

## ctx.Err()

23. `ctx.Err()` ne işe yarar?

24. Manuel cancellation sonrasında `ctx.Err()` genellikle ne döndürür?

25. Timeout sonrasında `ctx.Err()` genellikle ne döndürür?

---

## Timeout

26. `context.WithTimeout()` ne işe yarar?

27. Aşağıdaki context ne kadar süre sonra timeout olur?

```go
ctx, cancel := context.WithTimeout(
	context.Background(),
	3*time.Second,
)
```

28. `WithTimeout()` kullanırken neden genellikle `defer cancel()` yazılır?

---

## Deadline

29. `context.WithDeadline()` ne işe yarar?

30. Aşağıdaki kod ne anlama gelir?

```go
deadline := time.Now().Add(5 * time.Second)
```

31. `WithTimeout()` ile `WithDeadline()` arasındaki temel fark nedir?

---

## Context + Goroutine

32. Aşağıdaki iki çağrı arasındaki fark nedir?

```go
worker(ctx)
```

ve:

```go
go worker(ctx)
```

33. `go worker(ctx)` çağrıldığında main goroutine worker'ın bitmesini bekler mi?

34. Context kullanarak sonsuz döngüde çalışan bir goroutine nasıl kontrollü şekilde durdurulabilir?

35. `cancel()` çağrıldıktan sonra worker neden kendiliğinden anında yok olmaz?

---

## Context Propagation

36. Context Propagation nedir?

37. Aşağıdaki akışta neden aynı `ctx` aşağı doğru gönderilir?

```text
main
 ↓ ctx
service
 ↓ ctx
repository
```

38. Backend uygulamasında context hangi katmanlar arasında taşınabilir?

39. HTTP request iptal edilirse context propagation neden faydalıdır?

---

## Kod Okuma

40. Aşağıdaki kodu kendi cümlelerinle açıkla:

```go
func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return

		default:
			fmt.Println("Processing...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
```

41. Yukarıdaki kodda `for` neden sonsuz döngüdür?

42. Sonsuz döngüden nasıl çıkılmaktadır?

43. Context henüz sona ermediyse `select` içerisindeki hangi bölüm çalışır?

44. Context sona erdiğinde hangi bölüm çalışır?

45. Context'in temel amacını kendi cümlelerinle açıkla.