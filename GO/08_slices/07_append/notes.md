# APPEND

The built-in append function is used to dynamically add to a slice:

    func append(slice []Type, elems ...Type) []Type

If the underlying is not large enough, `append()` will create a new underlying array and point the slice to it.

Notice that `append()` is variadic, the following are all valid:

    slice = append(slice, oneThing)
    slice = append(slice,firstThing, secondThing)
    slice = append(slice, anotherSlice...)
