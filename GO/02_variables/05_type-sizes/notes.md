# TYPE SIZES

Ints, uints, floats, and complex numbers all have type sizes.

```
int int8 int16 int32 int64  // whole numbers

uint uint8 uint16 uint32 uint64 uintptr  // positive whole numbers

float32 float64  // decimal numbers

complex64 complex128  // imaginary numbers (rare)
```

The size (8, 16, 32, 64, 128, etc) represents how many bits in memory will be used to store the variable. The "default" `int` and `uint` types refer to their respective 32 or 64-bit sizes depending on the environment of the user.

The standard sizes that should be used unless the developer has a specific need are :

- `int`
- `uint`
- `float64`
- `complex128`

Some types can be converted like this:

    temperatureFloat := 88.26
    temperatureInt := int64(temperatureFloat)

Casting a float to an integer in this way truncates the floating point portion.
