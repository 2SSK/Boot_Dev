# INTERFACES IN GO

Interfaces are just collections of method signatures. A type "implements" an interface if it has methods that match the interface's method signatures.

In the following example, a "shape" must be able to return its area and perimeter. Both `rect` and `circle` fulfill the interface.

```
type shape interface {
    area() float64
    perimeter() float64
}

type rect struct {
    width, height float64
}

func (r rect) area() float64 {
    return r.width * r.height
}

type circle struct {
    radius float64
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}
```

When a type implements an interface, it can then be used as that interface type. As an example, becuase the empty interface doesn't require a type to have any methods at all, every type automatically implements the empty interface, written as:

    interface{}

### TIP

The length of a string can be obtained using the `len` function, which returns the number of bytes.

    s := "Hello, World!"
    fmt.Println(len(s))
    // 13
