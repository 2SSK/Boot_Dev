# STRINGS AS KEYS

just like Python dictionaries, we can use a bracket notation to access the properties of objects, as opposed to the dot operator we have been using so far.

For example:

    const desk = {
        wood: 'maple',
        width: 100
    }

    console.log(desk.wood)
    // prints "maple"

    console.log(desk['wood'])
    // also prints "maple"

Bracket notation is powerful, it allows us to use _dynamically_ created strings as keys rather than hard-coding them as we have been doing.

For example:

    const animals = ['cat', 'cat', 'cat', 'dog', 'dog', 'rat']

    const animalCount = {}

    for (let animal of animals) {
        // here we initialize the count to 0 if the key doesn't exist yet
        if (!animalCount[animal]) {
            animalCount[animal] = 0
        }

        // here we increment the count
        animalCount[animal]++
    }

    /*
    animalCounts = {
        cat: 3,
        dog: 2,
        rat: 1
    }
    */

Using strings keys is a great way to ensure uniqueness in a set.
