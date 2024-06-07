# LOG FLAGS

As you know, `git log` shows you the history of commits in your repo. There are a few flags i like to use from time to time to make the output easier to read.

The firs is <span style="color:blue">--decorate</span>. It can be one of:

- `short` (the default)
- `full` (shows the full ref name)
- `no` (no decoration)

Run `git log --decorate=full`. You should see that instead of just using your branch name, it will show the full ref name. A <span style="color:blue">ref</span> is justa pointer to a commit. All branches are refs, but not all refs are branches.

Run `git log --decorate=no`. You should see that the branch names are no longer shown at all.

The second is <span style="color:blue">--oneline</span>. This flag will show you a more compact view of the log. I use this one all the time, it just makes it so much easier to see what's going on.

    git log --oneline
