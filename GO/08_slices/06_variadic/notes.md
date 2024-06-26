# VARIADIC

Many functions, especiallythose in the standard library, can take an arbitrary nubmer of _final_ arguments. This is accomplished by using the "..." syntax in the funcction signature.

A variadic function receives the variadic arguments as a slice.

```
func concat(strs ...string) string {
    final := ""
    // strs is just a slice of strings
    for i := 0; i < len(strs); i++ {
        final += strs[i]
    }
}

func main(){
    final := concat("Hello", "there", "friend!")
    fmt.Pritln(final)
    // Output : Hello there friend!
}
```

The familiar fmt.Pritln() and fmt.Sprintf() are variadic! `fmt.Pritln()` prints each element with space delimiters and a newline at the end.

    func Println(a ...interface{}) (n int, err error)

### SPREAD OPERATOR

The spread operator allows us to pass a sliec into a variadic function. The spread operator consists of three dots following the slice in the function call.

```
func PrintStrings(strings ...string){
    for i := 0; i < len(strings); i++ {
        fmt.Println(strings[i])
    }
}

func main(){
    names := []string{"bob", "sue", "alice"}
    PrintStrings(names...)
}
```
