# THE INITIAL STATEMENT OF AN IF BLOCK

An `if` conditional can have an "initial" statement. The variables(s) created in the initial statement are _only_ defined within the scope of the `if` body.

    if INITIAL_STATEMENT; CONDITION {
        // Do something
    }

### WHY WOULD I USE THIS?

This is just some syntactc sugar that Go offers to shorten up code in some cases. For example, instead of writing:

```
length := getLength(email)
if length < 1 {
    fmt.Println("Email is invalid")
}
```

We can do:

    if length := getLength(email); length < 1 {
        fmt.Println("Email is invalid")
    }

Not only is this code a bit shorter, but is also removes `length` from the parent scope, which is convenient because we don't need it there - we only need access to it while checking a condition.
