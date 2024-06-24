# CONTINUE & BREAK

### `CONTINUE`

The `continue` keyword stops the current iteration of a loop and continues to the next iteration. `continue` is a powerful way to use the guard clause pattern within loops.

```
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue
    }
    fmt.Println(i)
}

// Prints 1, 3, 5, 7, 9
```

### `BREAK`

The `break` keyword stops the current iteration of a loop and exits the loop.

```
for i := 0; i < 10; i++ {
    if i == 5 {
        break
    }
    fmt.Println(i)
}
// Prints 0, 1, 2, 3, 4
```
