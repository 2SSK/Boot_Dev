# ANONYMOUS STRUCTS IN GO

An anonymous struct is just like a normal struct, but it is defined without a name and therefore cannot be referenced elsewhere in the code.

To create an anonymous struct, just instantiate the instance immediately usig a second pair of brackets after declaring the type:

```
myCar := struct {
    make string
    model string
}{
    make: "tesla",
    model: "model 3",
}
```

You can even nest anonymous structs as fields within other structs:

```
type car struct {
    make string
    model string
    doors int
    mileage int
    // wheel is a field containing an anonymous struct
    wheel struct {
        radius int
        material string
    }
}
```

### WHEN SHOULD YOU USE AN ANONYMOUS STRUCT?

In general, _prefer named structs._ Named structs make it easier to read and understand your code, and they have the nice side-effect of being reusable. I sometimes use anonymous structs when i _know_ I won't ever need to use a struct again. For example, sometimes i'll use one to create te shape of some JSON data in HTTP handlers.

If a struct is only meant to be used once, then it makes sense to declare it in such a way that developers down the road won;t be tempted to accidentally use it again.
