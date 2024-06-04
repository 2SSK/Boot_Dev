# UPDATING AND ADDING KEYS TO OBJECTS

You can update and even add new keys using the `.` operator as well:

    const apple = {
        name : 'Apple',
        radius : 2,
        color : 'red',
    }

    apple.ranking = 3
    apple.color = 'green'
    // {"name":"Apple", "radius":2, "color":"green", "ranking":3}
