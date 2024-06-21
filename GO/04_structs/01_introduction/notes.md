# STRUCTS IN GO

We use structs in Go to represent structured data. It's often convenient to group different types of variables together. For example, if we want to represent a car we could do the following:

```
type car struct {
    make string
    model string
    year int
}
```

This creates a new struct type called `car`. All cars have a `make`, `model`, `doors` and `mileage`.

Structs in Go are often used to represent data that you might use a dictionary or object for in other languages.
