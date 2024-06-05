# THROWING YOUR OWN ERRORS

Errors are _not_ something to be scared of. Every program that runs in production handles errors on a constant basis. Our job as developers is to handle the erros gracefully and in a way that aligns with out expectations.

When something in our own code happens that we don't expect, we should throw an error. For example, if someone passes bad inputs to a function, we shouldn't be afraid to throw an error to let them know they did something wrong.

### SYNTAX

    throw new Error('something went wrong')

It's worth mentioning that in JavaScript you can throw primitive values as well, like a string:

    throw 'something went wrong'

The problem with doing it this way is that now when developers go to _catch the error_ they can't assume it's an error object and will need to do some nasty shape-checking:

    catch (err) {
        if (err.message) {
            console.log(err.message)
        } else {
            console.log(err)
        }
    }

**I recommend always throwing error objects for consistency's sake**
