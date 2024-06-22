# TYPE SWITCHES

A type switch is similar to a regular switch statement, but the cases specfiy _types_ instead of _values_.

```
func printNumbericValue(num interface{}){
    switch v := num.(type){
    case int:
        fmt.Println("%T\n", v)
    case string:
        fmt.Println("%T\n", v)
    default:
        fmt.Println("%T\n", v)
    }
}

func main(){
    printNumbericValue(1)
    // prints "int"

    printNumbericValue("1")
    // prints "string"

    printNumbericValue(struct {}{})
    // prints "struct {}"
}
```

`fmt.Printf("%T\n", v)` prints the type of a variable.
