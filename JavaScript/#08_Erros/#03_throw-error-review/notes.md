# THORWING ERRORS REVIEW

Software applications aren't perfect, and user input and network connectivity are far from predictable. Despite intensive debugging and unit testing, applications will still have failure cases when running in the real world.

Loss of network connectivity, missing database rows, out of memory issues, and unexpected user inputs can all prevent an application from performing "normally".
It is **your job** to catch and handle any and all errors gracefully so that your application performs as expected. When you are able to detect that something is amiss, you should be throwing the errors yourself, in addition to the "default" errors that JavaScript interpreter will throw.

    throw new Error('something bad happened')
