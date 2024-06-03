# CONTINUE IN LOOPS IN JS

The `continue` keyword can be used to break out of a _single iteration_ of a loop early.

    for (let i = 0; i < 10; i++) {
        if (i % 2 === 0) {
            continue
        }
        console.log(i)
    }
    // Prints:
    // 1
    // 3
    // 5
    // 7
    // 9
