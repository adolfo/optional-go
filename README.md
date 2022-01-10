<p align="center">
   <img src="/resources/optional_gopher.png" alt="Optional Gopher"/>
</p>

# Optional Go

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
optString := optional.Of("foo")

// Optional int
optInt := optional.Of(5)

// Optional float64
optFloat := optional.Of(9.99)

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
nilString := optional.OfNil[string]()

// Optional int with nil value
nilInt := optional.OfNil[int]()

// Nil struct
nilThing := optional.OfNil[Thing]()

// Alternative initialization
nilStringAlt := optional.Value[string]{}
```

## Updating optional values

```go
o := optional.Of(42)

// Clearing value and resetting to nil
o.SetNil()

// Updating value
o.Set(5) // ok
o.Set("fred") // compiler error; type is Optional.Value[int]
```

## Getting optional values

```go
o := optional.Of(42)

// Get value with default fallback
fmt.Println(o.ValueOrDefault(99)) // prints 42

// Getting underlying optional value
if ok, val := o.Value(); ok {
    fmt.Println(val) // prints 42
}

// Must getter
fmt.Println(o.MustValue()) // prints 42
o.SetNil()
fmt.Println(o.MustValue()) // panics
```

## Checking for value

```go
if o.HasValue() {
    fmt.Println("has value")
}

// Alternatively

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
