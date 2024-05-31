# UNDEFINED VARIABLES

Not all variables have a value, some don't even have a `null` value. In JavaScript we can declare an empty or undefined variables:

    let emptyVar

The value of `emptyVar` in this instance is `undefined` until we use the assignment operator, `=`, to give it a value.

### UNDEFINED IS NOT A SPECIFIC STRING

Note that the `undefined` type is _not_ the same as a string with a value of "undefined":

    let myUndefinedVar // this is an undefined variable
    let myUndefinedVar = "undefined" // this is a defined string

### UNDEFINED IS ALSO FALSY

Jut like `null`, `undefined` is _also_ falsy.

### A NOTE ON 'UNDEFINED' VS 'NULL'

In my opinion, the way JavaScript deals with `null` and `undefined` is _really bad language design._ Most programming languages, like Python, just have one `null`, `None`, or `nill` type of value. JavaScript Unfortunately has 2: `null` and `undefined`.

In a nutshell, `undefined` literally means that the variable was never assigned a value, whereas the `null` value must be assigned explicitly.

My recommendation to you is to **keep it that way**. If you need a "nonetype" value, the `null` value. Try to avoid explicitly assigning `undefined` to value. In short, **prefer null!**
