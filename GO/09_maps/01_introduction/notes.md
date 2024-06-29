# MAPS

Maps are similar to JavaScripts, Python dictionaries, and Ruby hashes. Maps are a data structure that provides key->value mapping.

The zero value ofa map is `nil`.

We can create a map by using a literal or by using the `make()` function:

```
ages := make(map[string]int)
ages["John"] = 37
ages["mary"] = 24
ages["mary"] = 27   // overwrites 24
```

```
ages = map[string]int{
    "John": 37,
    "Mary": 24,
}
```

The `len()` function works on a map, it returns the total number of key/value pairs.

```
ages = map[string]int{
    "John": 37,
    "Mary": 21,
}
fmt.Println(len(ages))  // 2
```
