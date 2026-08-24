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