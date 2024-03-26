# Organizing Data
- [Organizing Data](#organizing-data)
  - [Pointers](#pointers)
  - [Structs](#structs)
    - [Pointers to structs](#pointers-to-structs)
    - [Struct Literals](#struct-literals)
    - [Embedded Structs](#embedded-structs)
    - [Struct with Receiver Functions](#struct-with-receiver-functions)

## Pointers
A pointer holds the memory address of a value.
pointer syntax:
- type "*T": is a pointer to a T value. Its zero value is nil.
- "&" operator: generates a pointer to its operand.
- "*" operator: denotes the pointer's underlying value.
This is known as "dereferencing" or "indirecting".

```go
func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
```

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

### Pointers to structs
Struct fields can be accessed through a struct pointer. 

To access the field of a struct pointer p:   
- (*p).field (long way)
-  p.field (recommended way to use dot notation)
  
```go
func main() {
	diana := person{"Diana", "Tran"}
	diana.lastName = "T"
	fmt.Println(diana.lastName)
}
```

### Struct Literals
A struct literal defines a newly allocated struct value by listing the values of its fields.

```go
var (
	v1 = person{"v1first", "v1last"} // has type person
	v2 = person{firstName: "v2first"} // field Name: syntax is implicit
	v3 = person{} // empty: zero value
	p  = &person{"pfirst", "plast"} // has type *person
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
In the example below, the update to the struct did not occur as expected.
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
 In order to update structs values, you need pointers.