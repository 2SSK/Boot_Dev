# EFFECTIVE GO

### LIKE SLICES, MAPS HOLD REFERENCES

Likes slices, maps hold references to an underlying data structure. If you pass a map to a function that changes the contents of the map, the changes will be visible in the caller.

### MAP LITERALS

Maps can be constructed using the usual composite literal syntax. with colon-separated key-value pairs, so it's easy to build them during initialization.

```
var timeZone = map[string]int{
    "UTC":  0*60*60,
    "EST": -5*60*60,
    "CST": -6*60*60,
    "MST": -7*60*60,
    "PST": -8*60*60,
}
```

### MISSING KEYS

An attempt to fetch a map value with a key that is not present in the map will return the zero value for the type of the entries in the map. For instance, if the map contains integers, looking up a non-existent key will return 0. A set can be implemented as a map with value type bool. Set the map entry to ture to put the value in the set, and then test it by simple indexing.

```
attended := map[string]bool{
    "Ann": true,
    "Joe": true,
    ...
}

if attended[person]{    // will be false if person is not in the map
    fmt.Println(person, "was at the meeting")
}
```

Sometimes you need to distinguish a missing entry from a zero value. Is there an entry for "UTC" or is that 0 becuase it's not in the map at all? You can discriminate with a form of multiple assignment.

```
var seconds int
var ok bool
seconds, ok = timeZone[tz]
```

For obvious reasons, this is called the "comma ok" idiom. In this example, if tz is present, seconds will be set appropriately and ok will be true; if not, seconds will be set to zero and ok will be false. Here's a function that puts it together with a nice error report:

```
func offset(tz string) int {
    if seconds, ok := timezone[tz]; ok {
        return seconds
    }
    log.Println("unknown time zone:", tz)
    return 0
}
```

### DELETING MAP ENTRIES

To delete a map entry, use the delete built-in function, whose arguments are the map and the key to be deleted. It's safe to this even if the key is already absent from the map.

    delete(timeZone, "PDT") // Now on Standard Time
