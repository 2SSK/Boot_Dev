# IF, ELSE IF, ELSE

Fundamently, if/else blocks are very similar in Python and JavaScript. There are a few key differences:

1. The condition _must_ be inside parentheses
2. The body should be within brackets
3. It's `else if` not `elif`

######

    if(height > 100){
    console.log("You're tall!");
    } else if(height > 60){
    console.log("You're average height!");
    } else {
    console.log("You're short!");
    }

While Python uses whitespace (an indent) to denote the "body" of an if statement, in JavaScript, the brackets denote the body and the indentation is just there to make it easier to read. The whitespace isn't part of the syntax in JS.
