# COMMIT

After staging a file, we can commit it.

A commit is a snapshot of the reopsitory at a give point in time It's a way to save the state of the repository, and it's how Git keeps track of changes to the project. A commit comes with a message that describes the changes made in the commit.

Here's how to commit all of your staged files:

    git commit -m "your message here"

### TIP

If you screw up a commit message, you can change it with the `--amend` flag. For example:

    # Change the last commit message
    git commit --amend -m "A: add contents.md"
