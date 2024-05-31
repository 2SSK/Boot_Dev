# NULL VARIABLES

Similar to Python's `None` type, JavaScript has the `null` value. Just like you would use `None` in Python to set a "default" value of some kind, you could use `null` for the same thing in JavaScript.

For example, I often use `null` values for varialbes in websites while I'm waiting for the data to load from the server.

    let profilePic = null

    // show a loading spinner while profilePic is null

    const dataFromBackend = waitForDataFromBackend()

    // once the profilePic is loaded, the picture
    // can be displayed in the browser
    prifilePic = dataFromBackend.profilePic

## NULL IS NOT A SPECIFIC STRING

Note that the `null` type is _not_ the same as a string with a value of "null":

    let  profilePic - null // this is a null value
    let profilePic = "null" // this is a string

## NULL VALUES ARE "FALSY"

"Falsy" means that a value evaluates to `false` when cast to a boolean. Here are some examples of "falsy" values:

- `false` (false boolean)
- `" "` (empty string)
- `0` (number zero)

Here are some truthly values:

- `true` (true boolean)
- `1` (a non-zero number)
- `"name"` (a string with more than zero characters)
