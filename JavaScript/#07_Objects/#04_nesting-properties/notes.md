# NESTING PROPERTIES

Objects can contain other objects! This means we can have properties that are potentially nested quite deep.

    const tournament = {
        referee : {
            name : 'Sally',
            age : 25,
        },
        prize : {
            units : 'dollars',
            value : 100,
        }
    }

We access nested properties the same way: `tournament.referee.name`
