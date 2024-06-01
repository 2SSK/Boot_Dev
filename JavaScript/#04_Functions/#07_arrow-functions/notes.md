# FAT ARROW FUNCTIONS

Fat arrow functions, or "arrow functions" are another way to define fucntions in JavaScript. Arrow functions are _newer_ than the `function` keyword, however, unlike the `let/var` issue, arrow functions are _sometimes_ more appropriate, not _always_.

### 'FUNCTION' KEYWORD

    const add = function(x, y) {
        return x + y
    }

### FAT ARROW

    const add = (x, y) => {
        return x + y
    }

---

#### WHAT ARE THE DIFFERENCES BETWEEN THE TWO?

- Fat arrow functions are always declared as variables - usually `const` , while the `function` keyword may or may not be a variable.
- Fat arrow functions handle object scoping in a more intuitive way (we'll talk about this in later chapters)
- A few other minor differneces
