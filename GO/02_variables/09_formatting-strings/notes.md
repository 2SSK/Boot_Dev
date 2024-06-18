# FORMATTING STRINGS IN GO

Go follows the printf tradition from the C language. In my opinion, string formatting/interpolation in Go is _less_ elegant than Python's f-strings, unfortunately.

- fmt.Printf - Prints a formatted string to standard out.
- fmt.Sprintf - Returns a formatted string

These following "formatting verbs" works with the formatting functions above:

### DEFAULT REPRESENTATION

The `%V` variant prints the Go syntax representation of value, it's nice default.

```
s := fmt.Sprintf("I am %V years old", 10)
// I am 10 years old

s := fmt.Sprintf("I am %V years old", "way too many")
// I am way too many years old
```

If you want to print in a more specific way, you can yse the following formatting verbs:

### STRING

    s := fmt.Sprintf("I am %s years old", "way too many")
    // I am way too many years old

### INTEGER

    s := fmt.Sprintf("I am %d years old", 10)
    // I am 10 years old

### FLOAT

```
s := fmt.Sprintf("I am %f years old", 10.523)
// I am 10.523000 years old

// The ".2" rounds the number to 2 decimal places
s := fmt.Sprintf("I am %.2f years old", 10.523)
// I am 10.52 years old
```
