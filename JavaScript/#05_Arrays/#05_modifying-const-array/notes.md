# MODIFYING CONST ARRAYS

It's important to note that `const` arrays can still be _modified_, they just can't be _reassigned_. That means we can add and remove elements, but we can't set a new array value with the assignment operator : `=`.

    const drinks = []

    // this works
    drinks.push('lemonade')

    // this breaks
    drinks = ['root beer']
