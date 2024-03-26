# Intro to Interface

- [Intro to Interface](#intro-to-interface)
  - [Concrete Type](#concrete-type)
  - [Interface Type](#interface-type)
  - [Interface Usage](#interface-usage)
- [Standard Library](#standard-library)
  - [http](#http)
    - [Reader / Closer Interface](#reader--closer-interface)
    - [Writer Interface](#writer-interface)

Folder Reference:

- [bot](../interfaces/bot/main.go)
- [http](../interfaces/http/main.go)

## Concrete Type

A concrete type has complete definition and can be instantiated directly to create actual values.

Concrete type includes:

- basic types: int, float64, string
- composite types: structs, array, slices, maps, channels
- user defined / custom types: created using struct / type to exend a basic or composite

## Interface Type

Interface is a type that defines a set of methods:

- It defines behavior but does not implement it.
- Interface is implicit: don't have to manually state custom type satisfies interface
- a contract to help manage types.

## Interface Usage

Interface syntax:

- type: interface
- name: interface of name
- body: all methods or all embedded interfaces
  - set of methods: method includes
    - method name
    - (optional) params and types: () or (param1, type1)
    - return type
  - set of embedded interfaces:
    - interface name

Any type that implements all the methods defined by the Interface is an honorary "of type" of the defined Interface.

- for interface with embedded interfaces: once methods in embedded interface is satisfied it implicity is the parent interface type

```go
type bot interface {
	getGreeting() string
}
type englishBot struct{}

func main() {
	eb := englishBot{}

	printGreeting(eb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hello! - English Bot"
}
```

# Standard Library

[net](https://pkg.go.dev/net@go1.22.1): Package net provides a portable interface for network I/O, including TCP/IP, UDP, domain name resolution, and Unix domain sockets.

## http

[http](https://pkg.go.dev/net/http@go1.22.1): Package http provides HTTP client and server implementations.

```go
func main() {
	resp, err := http.Get("http://google.com/")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(resp)
}

type Response struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"
	ProtoMajor int    // e.g. 1
	ProtoMinor int    // e.g. 0
    Header Header
    Body io.ReadCloser
    ContentLength int64
    TransferEncoding []string
    Close bool
    Uncompressed bool
    Trailer Header
    Request *Request
    TLS *tls.ConnectionState
}
```

### Reader / Closer Interface

```go
type ReadCloser interface {
	Reader
	Closer
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}
```

Reader interface can be reused for various source of data and return []byte data. Ex:

- HTTP reqest body
- Text / Image file on hard drive
- User entering text into command line
- Data from analog sensor plugged into machine

Read function (from Reader Interface) will read data into byte slice until the byte slice is full. So at initlization, a capacity is added.

```go
bs := make([]byte, 99999)
resp.Body.Read(bs)
// To output the response in string
fmt.Println(string(bs))
```

### Writer Interface

To achieve the same output of above, a Writer interface can be used to convert []byte to a source of output such as:

- Outgoing HTTP request
- Text / Image file on hard drive
- Terminal

```go
io.Copy(os.Stdout, resp.Body)

func Copy(dst Writer, src Reader) (written int64, err error)
// reader: resp.Body
// writer: os.Stdout
type Writer interface {
	Write(p []byte) (n int, err error)
}
// os.Stdout is type file which has write function
func (f *File) Write(b []byte) (n int, err error)
```

[io](https://pkg.go.dev/io)
