# RUN REBASE

To use rebase to bring changes from a branch called `tomsagenius` onto the current branch (let's call it `jds1`), we would run this while on the `jds1` branch:

    git rebase tomsagenius

This will do the following:

1. Checkout hte latest commit on `tomsagenius`
2. Replay one commit at a time from `jds1` onto `tomsagenius`
3. Update the `jds1` branch to point to the last replayed commit
