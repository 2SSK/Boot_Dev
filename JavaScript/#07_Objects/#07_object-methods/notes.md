# OBJECT METHODS

JavaScript objects can have `methods`, just like in Python. This is interesting to note because objects in JavaScript are essentially used to replace both dictionaries
_and_ objects in Python.

Methods are functions that are defined on an object. They can access and change the properties of the object in question.

In the context of a method, the `this` keyword referes to the object as a whole.

Let's look at an example:

    const person = {
        firstName: 'Lane',
        lastName: 'Wagner',
        getFullName: function() {
            return this.firstName + ' ' + this.lastName;
        }
    }

    console.log(person.getFullName());

As you can see, within the scope of the `getFullName()` method, `this` refers to the object itself and we can access the object's properties.
