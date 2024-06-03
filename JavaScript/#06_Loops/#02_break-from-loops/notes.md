# BREAK FROM LOOPS IN JS

The `break` keyword can be used to break from a loop early.

    for(let i = 0; i < 10; i++) {
        if(i === 3) {
            break
        }
        console.log(i)
    }
    // Prints:
    // 0
    // 1
    // 2
