# METHODS MUTATE

Methods can mutate (change) the properties of their objects:

    const tree = {
        height: 256,
        color: 'green',
        cut () {
            this.height /= 2;
        }
    }

    tree.cut()
    console.log(tree.height)
    // Prints 128

Just like the contents of a `const` array can change, the properties of a `const` object can change as well. Remember, the rule with `const` is simply that the name itself can't be _reassigned_.
