# GoST - Go Safe Types

Gost is a Golang package that provides two safe generic types, `Option` and `Result`, inspired by Rust's Option and Result types. These types are designed to simplify working with optional and error-handling scenarios, reducing the need for excessive nil and error checks.

## Option Type

The `Option` type represents an optional value and includes various methods for handling optional values.

### Usage Examples

#### Returning Options

```go
func maybe() gost.Option[int] {
    if /* some condition here */ {
        return gost.Some(10)
    } 

    return gost.None[int]()
}

val := maybe()
```

#### Wrapping Existing Functions

```go
slice := []int{1, 2, 3, 4, 5}
target := 4

opt := AsOption(slices.BinarySearch(slice, target))
```

#### Checking Option State

```go
if val.IsSome() {
    // Option has a value
} else {
    // Option is empty
}

if val.IsNone() {
    // Option is empty
} else {
    // Option has a value
}
```

#### Extracting Values 

```go
value := val.Unwrap()
valueOrDefault := val.UnwrapOr(10)
valueOrFunction := val.UnwrapOrElse(func() int { return calculateDefaultValue() })
valueOrZeroValue := val.UnwrapOrZero()
```

#### Panic, Fatal, Exit

```go
valueOrPanic := val.UnwrapOrPanic("Value is required!")
valueOrLogFatal := val.UnwrapOrLogFatal("Fatal: Value is required!")
valueOrExit := val.UnwrapOrExit(1)
```

## Result Type

The `Result` type represents a value or an error and includes methods for handling successful results and errors.

### Usage Examples

#### Creating Results

```go
func risky() gost.Result[int] {
    v, err := /* some risky function call */
    if err != nil {
        return gost.Ok(v)
    } 

    return gost.Error[int](err)
}

val := risky()
```

#### Wrapping Existing Functions

```go
slice := []int{1, 2, 3, 4, 5}
target := 4

res := AsResult(os.ReadFile("non-existent-file.txt"))
```


#### Checking Result State

```go
if val.IsOk() {
    // Result is successful
} else {
    // Result is an error
}

if val.IsError() {
    // Result is an error
} else {
    // Result is successful
}
```

#### Extracting Values And Errors 

```go
value := val.Unwrap()
err := val.UnwrapError()
valueOrDefault := val.UnwrapOr("Default")
valueOrFunction := val.UnwrapOrElse(func() string { return calculateDefaultString() })
valueOrZeroValue := val.UnwrapOrZero()
```

#### Panic, Fatal, and Exit

```go
valueOrPanic := val.UnwrapOrPanic()
valueOrLogFatal := val.UnwrapOrLogFatal()
valueOrExit := val.UnwrapOrExit(1)
valueOrDynamicExit := val.UnwrapOrDynamicExit(func(err error) int { return calculateExitCode(err) })
```
