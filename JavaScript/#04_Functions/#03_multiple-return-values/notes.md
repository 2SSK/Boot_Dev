# MULTIPLE RETURN VALUES

In Python, you'll recall that we could return multiple values from a function.

    def get_user():
        return "name@domain.com", 21, "active"

    email, age, status = get_user()

However, in JavaScript, _that's not allowed!_

    function getUser() {
        return "name@domain.com", 21, "active"
        // DON'T DO THIS
        // only returns 'active'
    }

The above code won't actually throw any sort of error, it will silently return the `active` string, which as you can probably guess is strange, unintuitive behavior.
_You can only return one value from a function in JavaScript!_

### THERE IS A WORKAROUND

There is a common workaround used by JavaScript developers, and that's to return an `object` that _contains_ multiple values. You're still just returning a single `object`, but that way you can return more than one "thing". We'll cove that in later chapters.
