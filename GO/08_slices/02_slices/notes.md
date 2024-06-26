# SLICES IN GO

_99 times out of 100 you will use a slice_ instead of an array when working with ordered lists.

Arrays are fixed in size. Once you make an array like `[10]int` you can't add an 11th element.

A slice is a _dynamically-sized,flexible _ view of the elements of an array.

<span stye="color:blue">The zero value of slice is `nil`.</span>

Slice **always** have an underlying array, though it isn't always specified explicitly. To explicitly create a slice on top of an array we can do:

    primes := [6]int {2, 3, 5, 7, 11, 13}
    mySlice := primes[1:4]
    // mySlice = {3, 5, 7}

The syntax is :

    arrayname[lowIndex : highIndex]
    arrayname[lowIndex : ]
    arrayname[ : highIndex]
    arrayname[ : ]

where `lowIndex` is inclusive and `highIndex` is exclusive.

`lowIndex`, `highIndex`, or _both_ can be omitted to use the entire array on that side of the colon.
