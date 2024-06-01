# COMPARISON OPERATORS

You're already familiar with these inequality operators, and they work as you would expect in JavaScript:

    5 < 6 // evaluates to true
    5 > 6 // evaluates to false
    5 >= 6 // evaluates to false
    5 <= 6 // evaluates to true

The equality operators however are a bit... _strange_. To compare two values to see if they are same or different, you use an extra `=` sign:

    5 === 6 // evaluates to false
    5 !== 6 // evaluates to true

The more common equality operators have unintuitive behavior and frankly shoud be avoided:

    5 == 6 // evaluates to false
    5 != '5' // evaluates to true

    5 != 6 // evaluates to true
    5 != '5' // evaluates to false

The "strict equals" (`===`) and "strict not equals" (`!==`) compare both the value AND the type. The "loose equals" (`==`) and "loose not equals" (`!=`) just compare the _value_. That means that with loose equals, the string `'5'` and the number `5` are considered "equal" - which is basically just something that you shouldn't be comparing in good code.

This is a fairly unique quirk of the JavaScript language.

---

# COMPARISON OPERATOR EVALUATIONS

When a comparison happens, the result of the comparison is just a boolean value, it's either `true` or `false`.

Take the following two examples:

    const isBigger = 5 > 4

######

    const isBigger = true

In both of the above cases, we're creating a `Boolean` variable called `isBigger` with a value of `true`.

Since `5` is greater than `4`, `isBigger` is assigned teh value of `true` once the `5 > 4` expression is evaluated.
