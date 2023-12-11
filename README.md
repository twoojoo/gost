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

v := AsOption(slices.BinarySearch(slice, target))
```

#### Checking Option State

```go
if val.IsSome() {
    // Option has a value
}

if val.IsNone() {
    // Option is empty
}
```

#### Extracting Values 

```go
v := val.Unwrap()
v := val.UnwrapOr(10)
v := val.UnwrapOrElse(func() int { return calculateDefaultValue() })
v := val.UnwrapOrZero()
```

#### Panic, Fatal, Exit

```go
v := val.UnwrapOrPanic("Value is required!")
v := val.UnwrapOrLogFatal("Fatal: Value is required!")
v := val.UnwrapOrExit(1)
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
val := AsResult(os.ReadFile("non-existent-file.txt"))
```


#### Checking Result State

```go
if val.IsOk() {
    // Result is successful
}

if val.IsError() {
    // Result is an error
}
```

#### Extracting Values And Errors 

```go
v := val.Unwrap()
err := val.UnwrapError()
v := val.UnwrapOr("Default")
v := val.UnwrapOrElse(func() string { return calculateDefaultString() })
v := val.UnwrapOrZero()
```

#### Panic, Fatal, and Exit

```go
valueOrPanic := val.UnwrapOrPanic()
valueOrLogFatal := val.UnwrapOrLogFatal()
valueOrExit := val.UnwrapOrExit(1)
valueOrDynamicExit := val.UnwrapOrDynamicExit(func(err error) int { return calculateExitCode(err) })
```

### Pattern Matching

You can also handle `Result` and `Options` using a sort of pattern matching syntax:

```go
val := AsOption(slices.BinarySearch([]int{1, 2, 3, 4, 5}, 10)).
        OnSome(func(v *int) *int { return v }).
        OnNone(func() int { return -1 })

val := AsResult(os.ReadFile("non-existent-file.txt")).
        OnOk(func(v *[]byte) *[]byte { return v }).
        OnError(func(e error) []byte { return []byte{} })
```

This is syntax is very limited in his extensibility due to the lack of generics on golang struct methods, thus it only comes handy when you need some side effects while extracting the option or result inner value.

```go
func respondSuccess[T any](w http.ResponseWriter) func (v T) {
    return func (value T) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(v)
    }
}

func respondNotFound[T any](w http.ResponseWriter) T {
    w.WriteHeader(http.StatusNotFound)
    W.Write("")
}

func userHandler(w http.ResponseWriter, req *http.Request) {
    AsOption(db.FetchUser(4)).
        OnSome(func(v *int) *int { return respondSuccess(w, v) }).
        OnNone(func() int { return respondNotFound(w) })
}
```