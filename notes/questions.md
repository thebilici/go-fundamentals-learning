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