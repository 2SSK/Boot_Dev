# FIRST CLASS AND HIGHER ORDER FUNCTIONS

A programming language is said to have "first-class functions" when functions in that language are treated like any other variable.For example, in such a language, a function can be apssed as an argument to toher functions, can be returned by another function and can be assigned as a value to a variable.

A function that returns a function or accepts a functionas input is called a Higher-Order Function.

Go supports first-class and higher-order functions. Another way to think of this is that a function is just another type -- just like `int`s and `string`s and `bool`s.

For example, to accept a function as a a parameter:

```
func add(x, y int) int {
    return x + y
}

func mul(x,y int) int {
    return x * y
}

// aggregate applies the given math function to the first 3 inputs
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
return arithmetic(arithmetic(a, b), c)
}

func main(){
    fmt.Println(aggregate(2,3,4, add))
    // Prints 9
    fmt.Println(aggregate(2,3,4, mul))
    // Prints 24
}
```
