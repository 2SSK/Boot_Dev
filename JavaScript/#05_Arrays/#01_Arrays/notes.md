# ARRAYS

JavaScript's arrays are the equivalent of Python's lists. Arrays have nearly identical syntax in JavaScript and Python.

One important thing about an array is that items in an array are _not_ required to be of the same type.

    const numbers = [1, 2, 3, 4, 5]
    const strings = ['banana', 'apple', 'pear']

### INDEX INTO AN ARRAY

Just like we already saw with strings, you can get items from an array using square brackets.

    const strings = ['banana', 'apple', 'pear']
    console.log(strings[0])
    // Prints: 'banana'

### PUSH TO AN ARRAY

In JavaScript, we `.push()` instead of `.append()`

    const drinks = []
    drinks.push('lemonade')
    console.log(drinks)
    // Prints: ['lemonade']
    drinks.push('root beer')
    console.log(drinks)
    // Prints: ['lemonade', 'root beer']
