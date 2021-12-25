# guard
[![Go Reference](https://pkg.go.dev/badge/github.com/ammario/guard.svg)](https://pkg.go.dev/github.com/ammario/guard)

A mutex-wrapped value package for Go. This package is perhaps too simple to be worth an external dependency. Also,
I question whether this API is superior to the getter-setter [`atomic.Value`](https://pkg.go.dev/sync/atomic#Value) API.

```
go get github/ammario/guard
```
## Usage
```go
g := New[string]("dog")
// Exclusivity is guaranteed in closure
g.Write(func(v string) string {
    return v + "cat"
})
g.Read(func(v string){
    // Prints "dogcat"... thread safely
    fmt.Println(v)	
})
```
