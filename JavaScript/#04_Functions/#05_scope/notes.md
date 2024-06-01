# SCOPE

Scope deals with where values and functions can be accessed in your code. For example, you'll write code at some point, or see code at some point where the bug is related to a variable being "out of scope".

Some "scopes" include:

- Global scope - Values in the global scope are accessible everywhere
- File or module scope - Values can only be accessed from within the same code file
- Function scope - Values can only be accessed within the same function
- Block scope - Values can only be accessed within that `{ ... }` codeblock

######

    function getCar() {
    const car = 'Volvo';
    // code here can use 'car'
    }
    // code here can NOT use 'car'

######

    const car = 'Volvo'
    // code here can use 'car'

    function getCar() {
    console.log(car);
    // code here CAN use 'car'
    }
