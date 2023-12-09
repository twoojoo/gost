# GoST - Go Safe Types

Gost is a Golang package that provides two safe generic types, `Option` and `Result`, inspired by Rust's Option and Result types. These types are designed to simplify working with optional and error-handling scenarios, reducing the need for excessive nil and error checks.

## Option Type

The `Option` type represents an optional value and includes various methods for handling optional values.

### Usage Examples

#### Creating Options

```go
someValue := Some(42)
noneValue := None[int]()
```

#### Checking Option State

```go
if someValue.IsSome() {
    // Option has a value
} else {
    // Option is empty
}

if noneValue.IsNone() {
    // Option is empty
} else {
    // Option has a value
}
```

#### Extracting Values 

```go
value := someValue.Unwrap()
valueOrDefault := noneValue.UnwrapOr(10)
valueOrFunction := noneValue.UnwrapOrElse(func() int { return calculateDefaultValue() })
```

#### Panic, Fatal, Exit

```go
valueOrPanic := noneValue.UnwrapOrPanic("Value is required!")
valueOrLogFatal := noneValue.UnwrapOrLogFatal("Fatal: Value is required!")
valueOrExit := noneValue.UnwrapOrExit(1)
```

## Result Type

The `Result` type represents a value or an error and includes methods for handling successful results and errors.

### Usage Examples

#### Creating Results

```go
okResult := Ok("Success")
errorResult := Error[string](errors.New("Something went wrong"))
```

#### Checking Result State

```go
if okResult.IsOk() {
    // Result is successful
} else {
    // Result is an error
}

if errorResult.IsError() {
    // Result is an error
} else {
    // Result is successful
}
```

#### Extracting Values And Errors 

```go
value := okResult.Unwrap()
err := errorResult.UnwrapError()
valueOrDefault := errorResult.UnwrapOr("Default")
valueOrFunction := errorResult.UnwrapOrElse(func() string { return calculateDefaultString() })
```

#### Panic, Fatal, and Exit

```go
valueOrPanic := errorResult.UnwrapOrPanic()
valueOrLogFatal := errorResult.UnwrapOrLogFatal()
valueOrExit := errorResult.UnwrapOrExit(1)
valueOrDynamicExit := errorResult.UnwrapOrDynamicExit(func(err error) int { return calculateExitCode(err) })
```
