# ERRORS IN JAVASCRIPT

The fundamentals of error handling in JavaScript and Python are similar. However, instead of being called `Exceptions`, they're jsut called `Errors`.

JavaScript uses a try-catch pattern for handling errors.

    try {
        const car = {}
        console.log(car.make.name)
    } catch (err) {
        console.log(err.message)
        // Cannot read properties of undefined (reading 'name')
    } finally {
        console.log('This will always run')
    }

- The `try` block is executed untill an error is thrown or it completes, whichever happens first.
- The `catch` block is only executed if an error is raised in the `try` block. It then exposes the error as an error object (`err` in out case) so that the program can handle the error gracefully without crashing.
- The `finally` block will always run, regardless of whether or not an error was thrown

## FINALLY IS NOT AS COMMON

The `finally` block is optional, and personally, I rarely use it. If you want something to run regardless of a thrown error in the `try` block, you can usually just put it _after_ the `try/catch` block.

I've used `finally` when I want to do something dangerous with the error object in the `catch` block, but still want to run some code when the error is thrown.
`finally` is a place to put code that you want to run regardless of whether or not an error is thrown in the `try` _or_ `catch` block.

### THE ERROR OBJECT

The error object has two useful properties:

- name
- message

Most of the time, you're going to want to do something with the _message_, like log it to the console.
