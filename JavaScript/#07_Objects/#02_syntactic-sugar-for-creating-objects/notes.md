# SYNTACTIC SUGAR FOR CREATING OBJECTS IN JS

The following pieces of JS code are equivalent:

    const name = 'Apple'
    const radius = 2
    const color = 'red'
    const apple = {
        name : name,
        radius : radius,
        color : color,
    }

######

    const name = 'Apple'
    const radius = 2
    const color = 'red'
    const apple = {
        name,
        radius,
        color,
    }

Prefer the `2nd` example - it's more concise, but just as easy to understand. The 2nd syntax basically says "the existing variable name is what I want the key in the object ot be, so just use that"
