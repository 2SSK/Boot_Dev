# JAVASCRIPTS STRINGS

In JavaScript, a string can be written using either single quotes, for example:
`Hello`, or double quotes, such as `"Hello"`.

I prefer single quotes for JavaScript. It's important to have styling conventions so that all the code in a project looks consistent, making it easier to read and contribute to.

# INDEXING INTO STRINGS

Square brackets are used to access individual characters inside a string. The chars are numbered starting with `0`, and running up to length-1. You'll notice this is similar to how strings and lists work in Python.

    const greeting = 'Hello'
    greeting[0] // 'H'
    greeting[1] // 'e'
    greeting[2] // 'l'
    greeting[3] // 'l'
    greeting[4] // 'o'
    // you can also get the last char at length-1
    greeting[greeting.length-1] // 'o'

the `.length` _property_ on a JavaScript string returns its length.
