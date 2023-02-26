# Learn golang

## Getting Started

- create a file with `go` extension like `file-name.go`
- create a mod file for project
```bash
go mod init project-name
```

<h2>Go program structure</h2>

- package declaration
- import packages
- main function

```go
//main.go
package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}

outpur: Hello world!
```

run main file
```bash
go run file-name.go
```

- comments
```go
//single line comment
/*
multi
line 
comment
*/
```

## variables

syntax of variable declaration
```bash
var variable dataType
var variable1, variable2, .... varibaleN DataType
```

example
```go
package main

import "fmt"

func main() {
    const PI = 3.1416
	//variable declaration
    var country, city string
    var year int
    var rate float32
    
    //variable initialization
    country = "Bangladesh"
    city = "Dhaka"
    year = 2023
    rate = 90.23

    //print
    fmt.Println(country)
    fmt.Println(city, "is the capital of ", country)
    fmt.Println("Happy new year", year)
    fmt.Println("Rate is", rate)
}
```
variable initialization using shortcuts
```go
package main

import "fmt"

func main() {
	//variable initialization shortcuts
    country := "Bangladesh"
    city := "Dhaka"
    year := 2023
    rate := 90.23

    //print
    fmt.Println(country)
    fmt.Println(city, "is the capital of ", country)
    fmt.Println("Happy new year", year)
    fmt.Println("Rate is", rate)
}
```
## Formating output

```go
package main

import "fmt"

func main() {
	//variable initialization shortcuts
    country := "Bangladesh"
    city := "Dhaka"
    year := 2023
    rate := 90.23

    //print
    fmt.Printf("%v is our motherland\n", country)
    fmt.Println("%v is the capital of %v\n", city, country)
    fmt.Println("Happy new year %v\n", year)
    fmt.Println("Rate is %v\n", rate)
}
```

## getting input from user

```go
package main

import "fmt"

func main() {
    var name string
    var age, num1, num2 int

	fmt.Printf("Enter your name: ")
    fmt.Scanf("%v", &name)

    fmt.Printf("Enter your age: ")
    fmt.Scan(&age)

    fmt.Printf("Enter 2 numbers: ")
    fmt.Scan(&num1, &num2)
}
```

## Number coversion
```go
package main

import "fmt"

func main() {
   var decNum int

   fmt.Printf("Decimal number: ")
   fmt.Scan(&decNum)
   
   fmt.Printf("%v", decNum) //decimal format
   fmt.Printf("%b", decNum) //binary format
   fmt.Printf("%o", decNum) //octal format
   fmt.Printf("%x", decNum) //hexadecimal format
}
```

## Operators

- Arithmetic operators: +, -, *, /, %, ++, --
- Assignment operators: =, +=, -=, *=, /=, %=, &=, |=, ^=, >>=, <<=
- comparison operators: ==, !=, >, <, >=, <=
- Logical operators: &&, ||, !
- Bitwise operators: &, |, ^, <<, >>

## Conditions

-if statements
```bash
if codition {
    //statement
}
```

example
```go
package main

import "fmt"

func main() {
    x := 5
    y := 2

   if x > y {
    fmt.Println("5 is greater than 2")
   }
}
```

-if else statements
```bash
if codition {
    //statement
} else {
    //statement
}
```

example
```go
package main

import "fmt"

func main() {
    x := 5
    y := 2

   if x > y {
    fmt.Println("right")
   } else {
    fmt.Println("wrong")
   }
}
```

- else if statements
```bash
if codition1 {
    //statement
} else if condition2 {
    //statement
} else {
    //statement
}
```