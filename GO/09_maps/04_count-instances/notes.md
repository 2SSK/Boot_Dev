# COUNT INSTANCES

Remember that you can check if a key is already present in a map by using the second return value from the index operation.

```
names := map[string]int()
missingNames := []string{}

if _, ok := names["Denna"]; !ok {
    // if the key doesn't exist yet,
    // append the name to the missingNames slice
    missingNames = append(missingNames, "Denna")
}
```
