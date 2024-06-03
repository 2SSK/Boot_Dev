# ARRAYS OPERATIONS - CONCATENATE

There are a couple of ways to concatenate two arrays (smushing them together) in JavaScript.

## THE .CONCAT() METHOD

    let nums = [1, 2, 3]
    nums = nums.concat([4, 5, 6])
    console.log(nums)
    // Prints : [1, 2, 3, 4, 5, 6]

## THE SPREAD SYNTAX

    const nums = [1, 2, 3]
    const newNums = [...nums, 4, 5, 6]
    console.log(newNums)
    // Prints : [1, 2, 3, 4, 5, 6]

This comes down to personal preference, but i prefer the spread syntax. It's newer, but i alos like it because, in my opinion, the `.concat` method can be confusing. I would expect the `.concat` method to mutate the array in place, but instead, it returns a new array.

    const nums = [1, 2, 3]
    nums.concat([4, 5, 6])
    console.log(nums)
    // Prints : [1, 2, 3]
