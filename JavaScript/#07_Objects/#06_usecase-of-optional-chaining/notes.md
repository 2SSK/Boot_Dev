# WHEN TO USE OPTIONAL CHAINING

### DON'T OVERUSE OPTIONAL CHAINING

Use `?.` only where it is `expected` that an object may not exist.

For example, if according to our business logic, users must exist but addresses are optional, we would write:

    user.address?.street

**not:**

    user?.address?.street

We don't want to overuse it because if we _expect_ that all users have objects, and we come across one that doesn't we _want_ an error thrown so we can see it and go fix the problem. _Errors make debugging easier._
