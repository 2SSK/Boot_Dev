# FAT ARROW FUNCTIONS AND 'THIS'

I said we would get back to the differences between the `function` keyword and fat arrow functions and the time is now!

Now that we've talked how `this` is bound to the instance of the current object, we can understand a bit more about fat arrow functions. First, let's acknowledge something interesting: In JavaScript, the `this` keyword is practically _always_ accessible.

### WHEN YOU ARE _NOT_ USING FAT ARROWS:

- If you're in an object, `this` is bound to the object
- If you're in the global, or "top-level" scope, `this` is bound to the "global" object

### FAT ARROW FUNCTIONS PRESERVE 'THIS'

    const author = {
        firstName: 'Lane',
        lastName: 'Wagner',
        getName(){
            return `${this.firstName} ${this.lastName}`
        }
    }

    console.log(author.getName())
    // Prints: Lane Wagner

######

    const author = {
        firstName: 'Lane',
        lastName: 'Wagner',
        getName: () => {
            return `${this.firstName} ${this.lastName}`
        }
    }

    console.log(author.getName())
    // Prints: undefined undefined
    // because the parent scope (the scope outside of the author object)
    // never defined .firstName and .lastName properties

So, the big takeaway is that when you use the non-fat-arrow function syntax, you _sometimes_ get a different `this` object, depending on how the function (or method) was defined and called.

With a fat-arrow function, the `this` keyword will **always refer to the same object** that its parent scope would. In essence, **fat arrow functions "preserve" the `this` context**. For the reason, some developers and frameworks prefer fat arrow functions in many situations. It can alleviate some confusion around the `this` object.
