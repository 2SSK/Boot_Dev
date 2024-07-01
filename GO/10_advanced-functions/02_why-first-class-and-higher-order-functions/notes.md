# WHY FIRST-CLASS AND HIGHER-ORDER FUNCTIONS

Ar first, it may seem like dynamically creating functions and passing them around as variables adds unnecessary complexity. Most of the time you would be right There are cases however when functions as values make a lot of sense. Some of these include:

- <span style="color:blue">HTTP API</span> handlers
- <span style="color:blue">Pbu/Sub</span> handlers
- Onclick callbacks

Any time you need to run custom code at a _time in the future_, functions as values might make sense.

### DEFINITION: FIRST-CLASS FUNCTIONS

A first-class function is a function that can be treated like any other value. Go supports first-class functions. A function's type is dependent on the types of its parameters and returns values. For example, these are different funciton types:

    func() int

######

    func(string) int

### DEFINITION: HIGHER-ORDER FUNCTIONS

A higher-order function is a function as an argument or returns a function as a return value. Go supports higher-order functions. For example, this function takes a functions as an argument:

    func aggregate(a, b, c int, arithmetic func(int, int) int) int
