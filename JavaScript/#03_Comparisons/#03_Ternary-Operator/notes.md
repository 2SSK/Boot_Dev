# TERNARY OPERATORS IN JS

Sometimes using 3-5 lines of code to write an if/else block is overkill. The ternary operator provides a simpler alternative in some scenarios.

    const price = isMember ? '$2.00' : '$10.00'

I like to read it in English as:

> If `isMember` is true, then evaluate to `$2.00`, otherwise evaluate to `$10.00`.

The smae logic using if/else would be:

    let price
    if(isMember){
        price = '$2.00'
    } else {
        price = '$10.00'
    }

### WHY IS IT CALLED "TERNARY"

ternary's latin root means "3", and it's the only JavaScript operator that takes _three_ operands.

- A condition followed by a question mark (`?`)
- An expression to execute if the condition is truthy, followed by a colon (`:`)
- The expression to execute if the condition is falsy.

### WHEN TO USE A TERNARY

You will use if/else statements more often than ternaries, generally speaking.

Ternary operations are great for _small_, single-line conditionals. I've seen developers in the professional world get "clever" with ternaries, and start creating nested monstrosities like this:

    const vehicleName = isTruck ? 'truck' : isCar ? 'car' : isScooter ? 'scooter' : 'vehicle'

I mean it works, and its all on one line, but it's _much_ harder to understand. I'd rather see this in an if/else block.

Remember: code is written for _humans_ not machines. We write JavaScript instead of binary so that we can read and uderstand it easily. If your code is needlessly hard to understand it's not good code.
