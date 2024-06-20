# MULTIPLE PARAMETERS

When multiple arguments are of the same type, and are next to each other in the function signature, the type only needs to be declared after the last argument.

Here are some examples:

    func addToDatabse(hp, damage int){
        // ...
    }

######

    func addToDatabse(hp, damage int, name string){
    // ?
    }

######

    func addToDatabse(hp, damage int, name string, level int) {
    // ?
    }
