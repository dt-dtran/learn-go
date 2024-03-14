# Go Basic Syntax

- [Go Basic Syntax](#go-basic-syntax)
	- [Variable Declarations](#variable-declarations)
		- [Reassigning Variables](#reassigning-variables)
	- [Basic Data types](#basic-data-types)
	- [Function and Return Types](#function-and-return-types)
	- [Slices](#slices)
		- [Splitting Range of a Slice](#splitting-range-of-a-slice)
	- [For Loops](#for-loops)
	- [Object Oriented Programming with Go](#object-oriented-programming-with-go)
		- [Receiver](#receiver)

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

### Splitting Range of a Slice

Splitting an existing slice consist of a starting index to include and up to index that is not included:

- element[startIndexIncluding : upToNotIncluding]
- this create a copy of the slice and not modify the original slice.

```go
cards := []string{"0", "1", "2", "3"}
cards[0:2] = cards[:2] = ["0", "1"]
cards[2:] = ["2", "3"]
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

Class does not exist in Go. The Go approach to OOP is to extend the base type and add additional functionality.

### Receiver

A function receiver is like a method on types to call on the instance of the type. This allows any variable of the type to get access to the method.

A function with receiver consist of:

- keyword: func
- receiver: (variableShortHand Type) ex. (d deck)
  - variable: copy of the instance of type
    - variable is similar to this (js) or self (python).
    - convention to use the first letter of the type (ex. d)
  - type: every variable of type can call this method on itself
- name: ex. print
- parameter: (..)
- function post statement: {...}

```go
type deck []string

// function with a receiver is like a method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
```

Receiver is not needed for:

- constructor function to create and return a new instance
- utility function for generic operations/calculations
  - ex. math addition operations
- package function: i.e. go builtin function like fmt.Println()
