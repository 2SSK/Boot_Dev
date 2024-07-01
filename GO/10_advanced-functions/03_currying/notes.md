# CURRYING

Function currying is a concept from functional programming and involves partial application of functions. It allows a function with multiple arguments to be transformed into a sequencea of functions, each taking a single argument.

Although Go does not support full currying, it is possible to simulate this behavior. By designing functions that take other functions as inputs and return new functions, we can achieve a similar effect.

For example:

```
func main() {
    squareFun := selfMath(multiply)
    doubleFunc := selfMath(add)

    fmt.Println(squareFun(5))
    // Prints 25

    fmt.Println(doubleFunc(5))
    // Prints 10
}

func multiply(x, y int) int {
    return x * y
}

func add(x, y int) int {
    return x + y
}

func selfMath(mathFunc func(int, int) int) func(int) int {
    return func(x int) int {
        return mathFunc(x, x)
    }
}
```

In the example above, the `selfMath` function takes in a function takes in a function as its parameter, and returns a functions that itself returns the value of running that input function on its parameter.
