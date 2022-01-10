<p align="center">
   <img src="/resources/optional_gopher.png" alt="Optional Gopher"/>
</p>

# Optional Go
![Tests](https://github.com/adolfo/optional-go/actions/workflows/test.yaml/badge.svg)
[![License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)

Library for conveniently and safely dealing with optional (nullable) values.

<mark>Note</mark>: Requires Go 1.18beta1 with generics support.

# Installation
```bash
$ go get github.com/adolfo/optional-go
```

# Usage

```go
package main

import (
    "github.com/adolfo/optional-go"
)
```

## Initialization with literal values

```go
// Optional string
optStr := optional.Of("foo")

// Optional int
optInt := optional.Of(5)

// Optional float64
optFlt := optional.Of(9.99)

// Optional struct
type Thing struct {
    foo int
    bar string
}
optThing := optional.Of(Thing{42, "baz"})
```

## Initialization with nil values

```go
// Optional string with nil value
nilStr := optional.OfNil[string]()

// Optional int with nil value
nilInt := optional.OfNil[int]()

// Nil struct
nilThg := optional.OfNil[Thing]()
```

## Alternative initialization with nil values

```go
// Optional int with nil value
var nilInt optional.Value[int]

// Optional string with nil value
nilStr := optional.Value[string]{}
```

## Updating optional values

```go
o := optional.OfNil[int]

// Set new value
o.Set(42)

// Clearing value and resetting to nil
o.SetNil()

// Set invalid value
o.Set("fred") // compiler error; type is Optional.Value[int]
```

## Getting optional values

```go
o := optional.Of(42)

// Safely get underlying optional value
if ok, val := o.Value(); ok {
    fmt.Println(val) // prints 42
}

// Unsafely get value
fmt.Println(o.MustValue()) // prints 42
o.SetNil()
fmt.Println(o.MustValue()) // panics
```

## Get value with default
```go
o := optional.Of("foo")

fmt.Println(o.ValueOrDefault("fred")) // prints "foo"

o.SetNil()
fmt.Println(o.ValueOrDefault("fred")) // prints "fred"
```

## Checking for value

```go
if o.HasValue() {
    fmt.Println("has value")
}

if o.IsNil() {
    fmt.Println("value is nil")
}
```

## JSON Marshaling & Unmarshaling

```go
type Person struct {
    Name optional.Value[string] `json:"name"`
    Age  optional.Value[int]    `json:"age"`
}

j := `{ "name": null, "age": 42 }`
p := new(Person)
json.Unmarshal([]byte(j), p)

fmt.Println("name is %s", p.Name.ValueOrDefault("Unknown"))
// prints "name is Unknown" since `name` was null
```

# License

Released under the MIT License