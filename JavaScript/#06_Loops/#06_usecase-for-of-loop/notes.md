# WHEN SHOULD YOU USE FOR-OF LOOPS?

### TRADITIONAL SYNTAX

    for (let i = 0; i < woods.length; i++) {
        console.log(woods[i])
    }

### FOR-OF SYNTAX

    for(let wood of woods) {
        console.log(wood)
    }

Just like Python's `for wood in woods` syntax, you should use JavaScript `for wood of woods` syntax if:

- You need to iterate over an entire array
- You need to go in ascending order
- You don't need access to the index value
- You don't need to make updates to the array's items

As it turns out, all of those conditions are _usually_ true, so you'll use the `for-of` syntax often. That said, there are absolutely times when one of the conditions isn't true, so the traditional `i` - based for loop is the necessary alternative.
