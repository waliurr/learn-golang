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