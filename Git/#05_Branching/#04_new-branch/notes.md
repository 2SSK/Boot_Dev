# NEW BRANCH

You should already be on the `main` branch: your "default" branch. You can always check with `git branch`.

### TWO WAYS TO CREATE A BRANCH

    git branch my_new_branch

This creates a new branch called `my_new_branch`. The thing is, I rarely use this command because ususally i want to create a branch and switch to it immediately. So i use this command instead:

    git switch -c my_new_branch

The switch command allows you to switch branches, and the `-c` flag tells Git to create a new branch if it doesn't already exist.

When you create a new branch, it uses the _current commit_ you are on as the branch base. For example, if you're on your `main` branch with 3 commits, `A`, `B`, and `C`, and then you run `git switch -c my_new_branch`, your new branch will also will look like this:

![branch image](./new-branch.png)
