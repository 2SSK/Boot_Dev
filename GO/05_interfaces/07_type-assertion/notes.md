# TYPE ASSERTIONS IN GO

When working with iterfaces in Go, every once-in-awhile you'll need acccess to the underlying type of an interface value. You can cast an interface to its underlying type using a type assertion.

In the example below:

- we want to check if `s` is a `circle`
- we know `s` is an instance of the `shape` interface, but we do not know if it's also a `circle`
- `c` is a new `circle` struct from `s`
- `ok` is `true` if `s` is indeed a `circle`, or `false` if `s` is NOT a `circle`

```
type shape interface {
    area() float64
}

type circle struct {
    radius float64
}

c, ok := s.(circle)
if !ok {
    // log an error if s isn't a circle
    log.Fatal("s is not a  circle")
}

radius := c.radius
```
