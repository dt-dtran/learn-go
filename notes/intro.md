# Intro to Go

## Go CLI Commands

| Commands   | Description                                   |
| ---------- | --------------------------------------------- |
| go build   | Compiles executable file                      |
| go run     | Compiles and execute files                    |
| go fmt     | Format code                                   |
| go install | Compiles and install package                  |
| go get     | Downloads raw source code of external package |
| go test    | Run test                                      |

## Go Basics

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
```

### Packages

**Package** is collection of related files like a project or workspace. All related files delcares package main at the start of the source file.

There are 2 types of packages:

- Executable: Generates a file that can run
  - keyword "main" denotes an executable file at build/run: package main
- Reusable: Code used as "helpers" such as code dependencies and libraries
  - any other keyword will default to a reuseable package and do not a build .exe file.

### Imports

Import statements forms a link of standard library or resuable packages with main executable package. This allows packages to access code within another package.

[Standard Library packages](https://pkg.go.dev/std)

### Function

Function consist of:

- keyword: func
- name of function: main
- list of arguments to pass into function: ()
- function body: {...}

### File Organization

1. Package declaration: package main / package xyz
2. Import packages: import "fmt"
3. Declare function: func main() {fmt.Println("Hello, world!")}
