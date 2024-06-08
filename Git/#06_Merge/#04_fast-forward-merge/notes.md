# FAST FORWARD MERGE

The simples type of merge is a fast-forward merge. Let's say we start with this:

          C   delete_vscode
         /
    A - B     main

And we run this while on `main`:

    git merge delete_vscode

Because `delete_vscode` has _all the commits_ that `main` hs, Git automatically does a fast-forward merge. It just moves the pointer of the "base" branch to the tip of the "feature" branch:

                other_branch
    A - B - C   main

Notice that with a fast-forward merge, no merge commit is created.

This is common workflow when working with Git on a team of developers.

1. Create a branch for a new change
2. Make the change
3. Merge the branch back into `main`(or whatever branch your team dubs the "deafault" branch)
4. Remove the branch
5. Repeat
