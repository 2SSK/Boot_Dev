# NESTED STRUCTS IN GO

Structs can be nested to represent more complex entities:

```
type car struct {
    make string
    model string
    doors int
    mileage int
    frontWheel wheel
    backWheel wheel
}

type wheel struct {
    radius int
    material string
}
```

The fields of a struct can be accessed using the dot `.` operator.

    myCar := car{}
    myCar.frontWheel.radius = 5
