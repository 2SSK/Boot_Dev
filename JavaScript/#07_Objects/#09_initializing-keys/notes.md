# INITIALIZING KEYS

If a key doesn't exist but we try to access it, we will get a value of `undefined`. One way we can check for this si by using the not `!` operator. Because `undefined`
is considered "falsy" (falsy means it evaluates to `false` in a boolean context). The syntax is fairly simple:

### EXAMPLE 1

    const balances = {
        lane: 100,
        breanna: 150,
        john: 200
    }

    // if bob doesn't have a balance yet
    // create a new key for him with a value of 0
    if (!balances.bob) {
        balances.bob = 0;
    }

### EXAMPLE 2

    let days = {
        monday: 'workday',
        wednesday: 'workday'
        saturday: 'vacation',
    }

    if(!days.monday) {
        console.log('Monday is not a day')
    }

    if(!days.sunday) {
        console.log('Sunday is not a day')
    }
