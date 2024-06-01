# VAR, LET AND CONST

JavaScript is a language that changes rapidly, and because it's been arouond for 20+ years, that means there are a lot of old ways of writing JavaScript. The "old ways" should generally be avoided, but it's important to understand them because you'll come across old code and need to understand it.

The `var` keyword is one such piece of "legacy code". `var` is similar to `let` and `const` in that it allows us to define a new variable. Like `let`, `var` creates "mutable" or "changeable" variable. However, unlike `let`, `var` is _less safe_ because it does some strange things with scope.

Basically, `var` works as you would expect with _function scopes_, but ont within _block scopes_. An "if" statement is an example of a block scope.

    function printX(shouldSet) {
       if (shouldSet) {
            var x = 2
        }
        console.log(x);
        // Prints : 2
    }
    printX(true)

######

    function printX(shouldSet) {
        if (shouldSet) {
            let x = 2
        }
        console.log(x);
        // ReferenceError: x is not defined
    }
    printX(true)

The `let` behavior makes _much more sense_ because it treats block scope the same as function scope, which is what almost every other programming language does as well.

**Always use let and const, never var**.
