# LOOPS

A traditional "for loop" in JavaScript is written like this :

    for (let i = 0; i < 10; i++) {
        console.log(i)
    }

This kind of syntax is common in many C-style languages. Python's `range` syntax is a bit unique.

In English, the code says:

1. Start with `i` equals `0`. (`let i = 0`)
2. if `i` is not less than `10`, exit the loop.
3. Log `i` to the console. (`console.log(i)`)
4. Add `1` to `i`. (`i++`)
5. Go back to step 2.

The result is that the numbers `0-9` are logged toth e console in order.
