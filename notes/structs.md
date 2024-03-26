# Intro to Structs

- [Intro to Structs](#intro-to-structs)
  - [Structs](#structs)
    - [Struct Literals](#struct-literals)
    - [Embedded Structs](#embedded-structs)
    - [Struct with Receiver Functions](#struct-with-receiver-functions)
  - [Pass By Value](#pass-by-value)
  - [Pointers](#pointers)
    - [Struct with Pointers](#struct-with-pointers)

Folder Reference: [structs](../structs/main.go)

## Structs

Struct is short for structure. A struct is a collection of fields or properties that are related together. Struct is similar to object (javascript), dictionary (python), or a hash (ruby).

struct syntax consist of:

- type and keyword: type struct
- name: structure name
- field:
  - field name and field type: ex. firstName string

```go
type person struct {
    firstName string
    lastName string
}
```

struct field can be accessed via dot notation.

### Struct Literals

A struct literal defines a newly allocated struct value by listing the values of its fields.

```go
var (
	v1 = person{"v1first", "v1last"} // has type person
	v2 = person{firstName: "v2first"} // field Name: syntax is implicit
	v3 = person{} // empty: zero value
	p  = &person{"pfirst", "plast"} // pointer: has type *person
)

func main() {
	fmt.Println(v1, v2, v3, p)
    // {v1first v1last} {v2first } { } &{pfirst plast}
}
```

### Embedded Structs

A struct field could be another struct embedded.

```go
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
    // both lines below is the same
	// contact   contactInfo
    contactInfo
}
func main() {
    jim := person{
        firstName: "Jim",
        lastName:  "Johns",
        // if field = contact contactInfo
        // contact: contactInfo{
        // 	email:   "jim@email.com",
        // 	zipCode: 001122,
        // },

        // if field = contactInfo
        contactInfo: contactInfo{
            email:   "jim@email.com",
            zipCode: 001122,
        },
    }
    // provide entire structure of struct
    fmt.Printf("%+v", jim)
}
// {firstName:Jim lastName:Johns contactInfo:{email:jim@email.com zipCode:594}}%
```

### Struct with Receiver Functions

In the example below, the update to the struct did not occur as expected. Instead the update only occurred at the copy.

```go
{... above main code
	jim.updateName("Jimmy")
	jim.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
```

## Pass By Value

In Go, when you pass an argument to a function, by default it is passed by value. This means that a copy of the argument's value is made and passed to the function. Any modifications made to the parameter inside the function do not affect the original value outside the function.

In order to update structs values, you need pointers.

## Pointers

A pointer holds the memory address of a value.

pointer syntax:

- type: "\*Var" is a pointer to a Variable struct value.
  - Its zero value is nil.
- operator:
  - "&" operator: return a pointer to memeory address
    - turn value into address
  - "\*" operator: return the value of the memory address pointed at
    - turns address into value

This is known as "dereferencing" or "indirecting".

```go
func main() {
	i, j := 42, 2701

	p := &i         // point to i memory address
	fmt.Println(*p) // read i  through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j memory address
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
```

### Struct with Pointers

Struct fields can be accessed through a struct pointer.

To access the field of a struct pointer p:

- (\*p).field (long way)
- p.field (recommended way to use dot notation)

```go
func (pointerToPerson *person) updateName(newFirstName string) {
	// long way
	// (*pointerToPerson).firstName = newFirstName
    // recommended way: able to do this because receiver type *person
	pointerToPerson.firstName = newFirstName
}
```
