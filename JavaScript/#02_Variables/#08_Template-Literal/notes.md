# TEMPLATE LITERALS IN JS

In JavaScript, you can use a template literal to interpolate dynamic values into a string template. Template literals are JavaScript's version of Python's f-strings.

For example:

    const shadeOfRed = 101
    console.log(`the shade is ${shadeOfRed}`)
    // prints: "the shade is 101"

Template literals _must start and end with a backtick_, and anything inside of the dollar-sign bracket enclouser is automatically _cast_ to a string.

### WHAT DOES "CAST" MEAN IN PROGRAMMING?

To "cast" means to convert from one type to another, like converting from a `number` to a `string`.

### WHAT DOES "INTERPOLATE" MEAN?

To "interpolate" means to insert something into something else. In this case, template literals allow us to **interpolate** dynamic values into a string template - It's super useful!
