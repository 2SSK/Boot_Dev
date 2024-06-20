# THE BENEFITS OF NAMED RETURNS

#### GOOD FOR DOCUMENTATION

#### (UNDERSTANDING)

Named return parameters are great for documenting a function. We know what the function is returning directly from its signature, no need for a comment.

Named return parameters are particularly important in longer functions with many return values.

```
func calculator(a, b int) (mul, div int, err error){
    if b == 0 {
        return 0, 0, errors.New("Can't divide by zero")
    }
    mul = a * b
    div = a / b
    return mul, div, nil
}
```

Which is easier to understand than:

```
func calculator(a, b int) (int, int, error) {
    if b == 0 {
        return 0, 0, errors.New("Can't divide by zero")
    }
    mul := a * b
    div := a / b
    return mul, div, nil
}
```

We know _the meaning_ of each return value just by looking at the function signature: `func calculator(a, b int) (mul, div int, err error)`

_Note: `nil` is the zero value of an error. More on this later._

### LESS CODE (SOMETIMES)

If there are mulitple return statements in a function, you don't need to write all the return values each time, though you _probably_ should.

When you choose to omit return values, it's called a _naked_ return. Naked returns should only be used in short and simple functions.
