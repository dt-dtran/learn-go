# Intro to Maps

- [Intro to Maps](#intro-to-maps)
  - [Map vs Struct](#map-vs-struct)
  - [Initialization](#initialization)
  - [Manipulating Maps](#manipulating-maps)
    - [Insert / Update Value](#insert--update-value)
    - [Retrieve](#retrieve)
    - [Delete](#delete)
    - [Test if key is present](#test-if-key-is-present)
  - [Iterate Over Map](#iterate-over-map)

Folder Reference: [maps](../maps/main.go)

## Map vs Struct

Map is a collection of key-value pairs.

Map consist of:

- all keys must be the same type
  - keys are indexed and can be iterated
- all values must be the same type
- **use case:** used to represent a collection of related properties
- **compile time**: does not need to know all the keys at compile time
- keys are indexed and can be iterated
- referenced type

Vs Struct:

- values can be of different value
- keys don't support indexing
- **use case**: used to represent a thing with diferent properities
- **compile time**: all fields are needed at compile time
- value type

## Initialization

There are a few ways to initialize a map. The preferred method is using go built in make function.

```go
func main() {
    // 1. literal
	colors := map[string]string{
		"red": "#red",
		"green": "#green",
	}

    // 2. var
    var colors map[string]string

    // 3. make function (preferred)
    colors := make(map[string]string)
	fmt.Println(colors)
}
```

## Manipulating Maps

To access value in map you need to use square backet [] syntax.

### Insert / Update Value

```go
m[key] = element
colors["white"] = "#white"
```

### Retrieve

```go
element = m[key]
x = colors["white"]
```

### Delete

Use go built in delete function.

```go
delete(m, key)
delete(colors,"white")
```

### Test if key is present

```go
element, ok := m[key]
v, ok := colors["white"]
fmt.Println("The value:", v, "Present?", ok)
```

## Iterate Over Map

map iterate function syntax:

- keyword: func
- name: name
- argument: (...)
  - argument name:
  - type of the map
- function body: loop
  - key, value:

```go
printMap(colors)

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
```
