# FUNCTIONS

As you might have guessed, JavaScript supports the ability to define your own functions. To create a new function, you'll ust the `function` keyword. Similar to other blocks like `if` and `else`, the body of the function is surrounded by brackets:

    function main() {
    const sum = getSum(1, 2);
    console.log(`The sum is ${sum}`);
    }

    function getSum(a, b) {
    return a + b;
    }

    main()
