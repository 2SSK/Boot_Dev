# IGNORING RETURN VALUES

A function can return a value that the caller doesn't care about. We can explicitly ignore variables by using an underscore: `_`

For example:

```
func getPoint() {x int, y int}{
    return 3, 4
}

// ignore y value
x, _ := getPoint()
```

Even though `getPoint()` returns two values, we can capture the first one and ignore the second.

### WHY MIGHT YOU IGNORE A RETURN VALUE?

Maybe a function called `getCircle` returns the center point and the radius, but you only need the radius for your calculation. In that case, you can ignore the center pint variable.

The Go compiler will **throw and error** if you have any unused variable declarations in your code, so you _need_ to ignore anything you don't intend to use.
