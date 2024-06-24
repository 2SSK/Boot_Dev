# THERE IS NO WHILE LOOP IN GO

Most programming languages have a concept of a `while` loop. Becuase Go allows for the omission of sections of a `for` loop, a `while` loop is just a `for` loop that only has a CONDITION.

    for CONDITION {
        // do some stuff while CONDITION is true
    }

For example:

```
planHeight := 1
for planHeight < 5 {
    fmt.println("still growing! current height :", planHeight)
    planHeight++
}
fmt.println("plant has grown to ", planHeight, "inches")
```

which prints:

```
still growing! current height : 1
still growing! current height : 2
still growing! current height : 3
still growing! current height : 4
plant has grow to 5 inches
```
