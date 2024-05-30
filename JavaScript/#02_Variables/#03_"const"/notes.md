# DECLARING VARIABLES WITH "CONST"

In JavaScript, the `const` keyword can _also_ be used to create variables. There is a very important difference between `let` and `const`. The variables declared with `let` _can_ be reassigned. Variables declared with `const` _cannot_ be reassigned later.

    let name = 'John'
    name = 'Jane'
    console.log(name)
    // prints: Jane

#####

    const name = 'John'
    name = 'Jane'
    // TypeError: Assignment to constant variable.

### CONST IS "BETTER" THAN LET

In general, if you don't plan to change the value of a variable later, use `const`. This will save you from unexpected bugs and confusion! If you need to mutate a variable for some reason though, it's okay to use `let`. Generally, it's good practice to create new variables rather than overwrite them, which is why I say prefer `const`.
