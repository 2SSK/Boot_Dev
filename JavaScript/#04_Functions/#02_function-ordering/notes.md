# FUNCTION ORDERING

In Python, a function had to defined _before_ it could be used. That's not the case in JavaScript! Like most other programming languages, as long as the function is defined _somewhere_ in the code, it can be called even _before_ the definition.

    console.log(getLabel(3));
    // prints 'awful'

    function getLabel(numStars) {
        if (numStars > 7) {
            return 'great'
        } else if (numStars > 3) {
            return 'okay'
        } else {
            return 'awful'
        }
    }

The way to think about how it works is that the JavaScript interpreter _reads all the code first_(including function definitions), then goes back and starts executing the code from the beginning.
