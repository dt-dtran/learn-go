# Go Basic Syntax

- [Go Basic Syntax](#go-basic-syntax)
	- [Variable Declarations](#variable-declarations)
		- [Reassigning Variables](#reassigning-variables)
	- [Basic Data types](#basic-data-types)
	- [Function and Return Types](#function-and-return-types)
	- [Slices](#slices)
	- [For Loops](#for-loops)
	- [Object Oriented Programming with Go](#object-oriented-programming-with-go)

Folder Reference: [Cards](../cards/main.go)

## Variable Declarations

variables consist of:

- keyword: var
- name of variable: name
- type of variable: string
- value assignment: "value"

```go
var card string = "Ace of Spages"
// is the same as
card := "Ace of Spades"
```

### Reassigning Variables

Only use := when declaring new variables
card = "Five of Diamonds"

## Basic Data types

| Type    | Example           |
| ------- | ----------------- |
| bool    | true, false       |
| string  | "string"          |
| int     | -1, 0, 1          |
| float64 | -1.01, 0.1, 1.091 |

## Function and Return Types

function syntax consist of:

- keyword: func
- function name: name
- list of arguments to pass into function: () || (arg1, arg2)
- data type of the return: string, bool, etc
- function body: {...}

```go
func newCard() string {
	return "Five of Diamonds"
}
```

## Slices

Array is a fixed length list of things. Slice is an array that can grow or shrink.

Slice syntax consist of:

- square brackets: []
- type of values in slice: ex string
- specify elements that exist inside the slice: {}

```go
cards := []string{"1", "newCard()"}
```

## For Loops

Loops syntax consist of:

- keyword: for
- init statement: variable declaration
- condition expression: evaluated before every iteration
- post statement: executed at end of each iteration

```go
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

for i, card := range cards {
	fmt.Println(i, card)
}
```

## Object Oriented Programming with Go

Class doest not exist in Go. The Go approach to OOP is to extend the base stype and add additional functionality.
