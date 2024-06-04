# OPTIONAL CHAINING IN JS

When working with nested data, it can quickly become confusing when you're trying to access the correct properties. When using the normal `.` operator to access nested properties, if you attempt to access a field on an object that doesn't exist an error will be thrown. Thankfully, JavaScript has recently added a new operator to make dealing with this headache easier, the optional chaining operator: `?.`

    const tournament = {
        prize : {
            units : 'dollars',
            value : 100,
        }
    }

    const h = tournament.referee.height
    // TypeError: Cannot read properties of undefined (reading 'height')

Instead, we should use the optional chaining oeprator if we're unsure whether the `referee` object exists or not:

    const tournament = {
        prize: {
            units: 'dollars',
            value: 100,
        }
    }

    const h = tournament.referee?.height
    // h is simply undefined, no error
