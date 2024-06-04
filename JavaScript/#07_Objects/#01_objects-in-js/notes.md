# OBJECTS IN JS

JavaScript has an Object type. Objects are single variable that can hold more complex information than the basic types kike `String`, `Number` and `Boolean`.
though they contain _properties_ that are often those basic types.

For example, we may want to store data about a fruit. We can do so in a single variable:

    const apple = {
        name : 'Apple',
        radius : 2,
        color : 'red',
    }

Then we can access properties stored in an object using the `.` operator :

    console.log(apple.name) // Prints "Apple"
    console.log(apple.radius) // Prints "2"
    console.log(apple.color) // Prints "red"

In a way, JavaScript's objects are used in a similar way to Python's dictionary: as a key/value store.

**Obejcts are how we get around JavaScript's limitation of only being able to return 1 value from a function.**
