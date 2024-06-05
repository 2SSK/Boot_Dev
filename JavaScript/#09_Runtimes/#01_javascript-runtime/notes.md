# THE JAVASCRIPT RUNTIME

A runtime environment is, put simply, _where your program runs_.

The runtime determines which global objects your program can access, which in turn affects which APIs are available and how your code is interpreted and executed.

### WHAT ARE EXAMPLES OF JAVASCRIPT RUNTIMES?

In approximate order of usage:

- A browser (and technically, each browser's runtime can be slightly different as well)
- Node.js
- A web worker within a browser
- Deno.js
- Bun

Originally, JavaScript _only_ ran in browsers. Today, it runs in many places. Browsers are still the most common, but Node.js is extremely popular as a runtime, especially for backend development.

### HOW DO THE RUNTIMES DIFFER?

They differ in many ways, but the main things to keep in mind are:

- Performance - Some runtimes run JS code faster or slower under different conditions.
- APIs - The runtime exposes APIs to your code. For examle, in a browser you can access the canvas API for drawing, but that's not available in Node. On the other hand, Node.js has some cryptography libraries available by default.
- Global object - In the browser, the global object is called `window`, in Node it's called `global`. The properties and methods avaiable on that global object are quite different depending on the runtime.
