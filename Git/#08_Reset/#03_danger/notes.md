# DANGER

I wnat to stress how **dangerous** this command can be. When you deleted the file, because it was tracked in Git, it was trivially easy to recover. However, if you have some changes that you _do_ want to keep, running `git reset --hard` will delete them for good.

Always be careful when using `git reset --hard`. It's a powerful tool, but it's also a dangerous one.

### RESET TO A SPECIFIC COMMIT

If you want to reset back to a specific commit, you can use the <span style="color:blue">git reset --hard</span> command and provide a commit hash. For example:

    git reset --hard 1a2b3c4d

This will reset your working direcotry and index to the state of that commit, and all the changes made after that commit are lost forever.
