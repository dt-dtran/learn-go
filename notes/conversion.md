# Data Type Conversion

- [Data Type Conversion](#data-type-conversion)
  - [Common Conversions](#common-conversions)
  - [Advance Conversion](#advance-conversion)
    - [Floats and Exponents](#floats-and-exponents)
    - [Base String](#base-string)
  - [Other](#other)

## Common Conversions

| Convert To        | Package | Function                     |
| ----------------- | ------- | ---------------------------- |
| string to int     | strconv | strconv.Atoi(str)            |
| string to float64 | strconv | strconv.ParseFloat(str)      |
| string to bool    | strconv | strconv.ParseBool(str)       |
| int to string     | strconv | strconv.Itoa(int)            |
| float64 to string | strconv | strconv.FormatFloat(float64) |
| bool to string    | strconv | strconv.FormatBool(bool)     |
| float64 to int    | none    | int(floatNum)                |

## Advance Conversion

### Floats and Exponents

| Convert To        | Package | Function                                         |
| ----------------- | ------- | ------------------------------------------------ |
| string to float64 | strconv | strconv.ParseFloat(str, bitSize)                 |
| float64 to string | strconv | strconv.FormatFloat(float64, fmt, prec, bitSize) |

bitSize parameter (int): commonly using 32 or 64

- 0: default int
- 8: int8 range from -128 to 127
- 16: int16 range from -32768 to 32767
- 32: int32 range from -2147483648 to 2147483647
- 64: int64 range from -9223372036854775808 to 9223372036854775807

prec paramter: controls the number of digits (excluding the exponent) printed

- -1: uses the smallest number of digits necessary
- 2: for generic currency etc

fmt parameter (byte): formats

Example using strconv.FormatFloat(3.14, fmt, -1, 64)

| Format | Description          | Example               |
| ------ | -------------------- | --------------------- |
| "f"    | no exponent          | 3.14                  |
| "g"    | e or f               | 3.14                  |
| "G"    | E or f               | 3.14                  |
| "b"    | binary               | 7070651414971679p-51  |
| "e"    | e decimal            | 3.14e+00              |
| "E"    | E decimal            | 3.14E+00              |
| "x"    | hexadecimal + binary | 0x1.91eb851eb851fp+01 |
| "X"    | headecimal + binary  | 0X1.91EB851EB851FP+01 |

Use case for formats:

- "f" for generic w/o exponents
- "g" or "G" to default to generic or expontent if theres large numeric values
- "e" or "E": Scientific reports, large monetary / numeric values,
- "x" or "X": Debug and Log memory address or binary data

### Base String

| Convert To             | Package | Function                              |
| ---------------------- | ------- | ------------------------------------- |
| -/+ int to string      | strconv | strconv.FormatInt(int64, base)        |
| unsigned int to string | strconv | strconv.ParseUint(int64, base)        |
| string to -/+ int      | strconv | strconv.ParseInt(str, base, bitSize)  |
| string to unsigned int | strconv | strconv.ParseUint(str, base, bitSize) |

| Base | Type        | Format Input | Output | Parse Input | Output |
| ---- | ----------- | ------------ | ------ | ----------- | ------ |
| 0    | unsure      | none         | none   | "10"        | 10     |
| 2    | binary      | int64(10)    | "1010" | "1010"      | 10     |
| 8    | octal       | int64(10)    | "12"   | "12"        | 10     |
| 10   | decimal     | int64(10)    | "10"   | "10"        | 10     |
| 16   | hexadecimal | int64(10)    | "a"    | "a"         | 10     |

## Other

| Conversion     | Package | Function                                   |
| -------------- | ------- | ------------------------------------------ |
| []string Joins | strings | strings.Join([]string(value), "separator") |
