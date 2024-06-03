# FOR-OF LOOPS

With newer version of JavaScript we can use a clean syntax to loop over an array without keeping track of the index manually.

### INDEX-BASED ITERATION:

    let woods = ['oak', 'pine', 'maple']
    for(let i = 0; i < woods.length; i++) {
        console.log(woods[i])
    }
    // Prints:
    // oak
    // pine
    // maple

### NEW SYNTAX

    let woods = ['oak', 'pine', 'maple']
    for(let wood of woods) {
        console.log(wood)
    }
    // Prints:
    // oak
    // pine
    // maple

In the example above, the `wood` variable is declared using the `of` keyword, and it directly accesses the _value_ in the array rather than the _index_ of the value. If we don't need to update the item, and only need to access its value then this is a more clean way to write the code. You can use either `const` or `let` in a loop header, depending on if the loop changes the variable.

It's similar to Python's `for wood in woods :` syntax. Both examples effectively do the same thing! (Not to confused with JavaScript's `for-in` syntax.)
