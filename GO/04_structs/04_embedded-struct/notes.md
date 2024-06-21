# EMBEDDED STRUCTS

Go is not an object-oriented language. However, embedded structs provide a kind of _data-only_ inheritance that can be useful at times. Keep in mind, Go doesn't support classes or inheritance in the _complete_ sense, but embedded structs are a way to elevate and **share fields between struct definitions.**

```
type car struct {
    make string
    model string
}

type truck struct {
    // "Car" is embedded, so the definition of a
    // "truck" now also additionally contains all
    // of the fields of the car struct
    car
    bedSize int
}
```

### EMBEDDED VS NESTED

- Unlike nested structs, an embedded struct's fields are accessed at the top level like normal fields.
- Like nested structs, you assign the prompted fields with the embedded struct in a composite literal.

```
lanesTruck := truck {
    bedSize: 10,
    car: car {
        make: "toyota",
        model: "camry",
    },
}

fmt.Println(lanesTruck.bedSize)

// embedded fields promoted to the top-level
// instead of lanesTruck.car.make
fmt.Println(lanesTruck.make)
fmt.Println(lanesTruck.model)
```
