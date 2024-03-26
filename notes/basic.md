# Go Basic Syntax

- [Go Basic Syntax](#go-basic-syntax)
	- [Variable Declarations](#variable-declarations)
		- [Reassigning Variables](#reassigning-variables)
	- [Basic Data types](#basic-data-types)
	- [Arrays](#arrays)
		- [Slices](#slices)
			- [Splitting Range of a Slice](#splitting-range-of-a-slice)
	- [Control Flow:](#control-flow)
		- [Loops](#loops)
			- [For each](#for-each)
			- [While Loop](#while-loop)
			- [Range](#range)
		- [Conditionals](#conditionals)
			- [Switch](#switch)
		- [Other Branching](#other-branching)
		- [Defer](#defer)
			- [Stacking Defers](#stacking-defers)
		- [Function and Return Types](#function-and-return-types)
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

## Arrays

### Slices

Array is a fixed length list of things. Slice is an array that can grow or shrink. A slice does not store any data, it just describes a section of an underlying array.

Slice syntax consist of:

- square brackets: []
- type of values in slice: ex string
- specify elements that exist inside the slice: {}

```go
cards := []string{"1", "newCard()"}
```

Changing the elements of a slice modifies the corresponding elements of its underlying array. Other slices that share the same underlying array will see those changes.

#### Splitting Range of a Slice

Splitting an existing slice consist of a starting index to include and up to index that is not included:

- element[startIndexIncluding : upToNotIncluding]
- this create a copy of the slice and not modify the original slice.

```go
cards := []string{"0", "1", "2", "3"}
cards[0:2] = cards[:2] = ["0", "1"]
cards[2:] = ["2", "3"]
```
dynamically-sized: https://go.dev/tour/moretypes/13
## Control Flow:

### Loops

Loops syntax consist of:

- keyword: for
- (optional) init statement: variable declaration / executed before the first iteration
- condition expression: evaluated before every iteration
- (optional) post statement: executed at end of each iteration
- separation by semicolon: ;

```go
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//optional init and post statements
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
```

#### For each

```go
func main() {
	for i, card := range cards {
		fmt.Println(i, card)
	}
}
```

#### While Loop
Semicolons syntax can be dropped.
```go
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
```

#### Range
You can skip the index or value by assigning to _. If you only want the index, you can omit the second variable.

```go
for i, _ := range pow // omit value
for _, value := range pow // omit index

// index only
for i := range pow
```

### Conditionals

if syntax consist of:
- keyword: if
- (optional) short statement: to execute before the condition
- condition expression: condition to evaluate
- post statement: executed if condition is met
  
Variables declared inside an if short statement can be used inside any of the else blocks.

```go
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		// v:=....; = short statement
		// v < lim = conditional expression
		return v
	} else {
		// can use variable here without declaration.
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use variable outside of the if/else
	return lim
}
```
#### Switch
A switch statement is a shorter way to write a sequence of if - else statements. Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

switch syntax consist of:
- keyword: switch and case
- condition expression: condition to evaluate
  - No condition is the same as switch true {...}
- case: value to compare against expression
- case post statement: code block to execute if case matches the expression
- default case: case to execute if no other cases matches the expression

```go
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd, windows...
		fmt.Printf("%s.\n", os)
	}

	// no condition expression = switch true
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```

### Other Branching
- break: terminates the execution of a loop or switch statement
- continue: skips the current iteration of a loop and proceeds to the next iteration
- goto: transfer control to a labeled statement within the same function
  - Recommended to avoid using goto. goto makes code harder to read.

### Defer
A defer statement defers the execution of a function until the surrounding function returns. The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

defer syntax consist of:
- keyword: defer
- function call: function to defer
- arguments: arguments to pass into the function call

defer use case:
- Resource Cleanup: closing files, releasing locks, or closing database connections
- Logging: recording entry and exit points
- Panic Recovery: recovering from panics by deferring function calls that handle errors

#### Stacking Defers
Multiple defer statements within the same function are executed in the reverse order in which they are declared. This means the function calls are pushed onto a stack and executed in last-in-first-out (LIFO) order.

```go
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```

### Function and Return Types

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

# Object Oriented Programming with Go

Class does not exist in Go. The Go approach to OOP is to extend the base type and add additional functionality.

## Receiver

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
