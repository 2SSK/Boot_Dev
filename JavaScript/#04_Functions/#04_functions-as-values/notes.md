# FUNCTIONS AS VALUES

"Fuctions as values", or "first-class" functions are supported by JavaScript. That means that functions themselves can be passed around as _data_.

    function sayHello() {
        console.log("Hello there!");
    }

    function doTwice(f){
        f();
        f();
    }

    doTwice(sayHello);
    // prints:
    // Hello there!
    // Hello there!

In the example above, the `doTwice` function accepts a _function_ as a parameter, then calls that function twice.

### CALLBACKS

Another name for this sort of thing is a "callback". In this case, `sayHello` function is a `callback` function that's passed into the `doTwice` function for it to execute later. Callbacks are _very_ common in front-end applications. You'll often see code that essentially says

> "When the 'sign up' button is clicked, do X"

What's `X`? Well, it's a function! So you essentially connect functions to UI components (like buttons) so hey can be called later when the element is clicked.

---

## WHEN TO USE CALLBACKS

The #1 rule of programming in my book is "KISS":

> Keep It Simple, Stupid!

What that means in regards to fancy programming tactics like a callbacks is **only use them if there isn't a simpler alternative**. 95% of the time you won't need a callback to accomplish what you're trying to do.

When callbacks are needed in application code, they are usually being passed to functions imported from a library or framework. Applications rarely define their own functions which accepts callbacks, while libraries and frameworks commonly do because they need to stay _application agnostic_.

For example, the requestAnimationFrame() is part of every web browser's JavaScript API. It _needs_ to accept a callback because it needs to be able to run an arbitrary piece of code the next time the browser redraws the webpage. It can't hard-code the behavior because it needs to support every web application imaginable.

However, when you are writing your _very specific app_, you can usually "hard code" what happens next - you don't need to write functions that callback to any other fucntion... **usually**.
