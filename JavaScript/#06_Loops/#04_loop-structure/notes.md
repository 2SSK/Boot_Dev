# LOOP STRUCTURE

We can generalize the structure of a for-loop as follows:

    for (INIT; CONDITION; POST){
        // do stuff
    }

- `INIT` runs once at the beginning of the loop. Here we typically initialize a variable for use within the loop. Conventionally this variable is simply called `i` for simplicity.

- `CONDITION` is run before execution of each iteration of the loop. If the condition evaluates to `true` then the loop body is executed, otherwise the loop is broken and ends.

- `POST` is evaluated at the end of each iteration. This is typically used to modify the value of a variable.

Another example that prints 0-99 ascending:

    for (let i = 0; i < 100; i++) {
        console.log(i)
    }
