# SHORT VARIABLE DECLARATION

Inside a function (like the main function) the `:=` short assignment statement can be used in place of a var declaration. The `:=` operator infers the type of the new variable based on the value. It's colloquially called the walrus operator because it looks like a walrus... sort of.

These two lines of code are equivalent:

    var empty string

######

    empty := ""

```
numCars := 10   // inferred as an integer
temperature := 0.0  // temperature is inferred as a float becuase it has a decimal point
var isFunny = true  // inferred as a boolean
```

Outside of a function (in the global/package scope), every statement begins with a keyword(`var`, `func`, and so on) and so the `:=` construct is not available.
