# RANGE

Go provides syntactic sugar to iterate easily over elements of a slice:

    for INDEX< ELEMENT := range SLICE {
    }

For example:

```
fruits := []string{"apple", "banana", "grape"}
for i, fruit := range fruits {
    fmt.Println(i, fruit)
}
// 0 apple
// 1 banana
// 2 grape
```
