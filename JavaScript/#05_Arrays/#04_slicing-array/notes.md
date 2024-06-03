# SLICING ARRAYS

JavaScript makes it easy to slice and idce arrays to work only with the section you care about.

The syntax is as follows: `slice(fromIndex, untilIndex)`.

Like in Python, the first argument of a slice is inclusive in inclusive and the second argument is exclusive.

    const animals = ['ant', 'bison', 'camel', 'duck', 'elephant']

    console.log(animals.slice(2))
    // ['camel', 'duck', 'elephant']

    console.log(animals.slice(2, 4))
    // ['camel', 'duck']

    console.log(animals.slice(1, 5))
    // ['bison', 'camel', 'duck', 'elephant']

    console.log(animals.slice(-2))
    // ['duck', 'elephant']

    console.log(animals.slice(2, -1))
    // ['camel', 'duck']

    console.log(animals.slice())
    // ['ant', 'bison', 'camel', 'duck', 'elephant']

While JavaScript doesn't support negative indexing inthe built-in array index syntax, the `slice()` method `does` support negative indexes.
