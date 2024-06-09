# MERGE

Just as we merged branches within a single local repo, we can also merge branches between local and remote repos.

### SYNTAX

    git merge remote/branch

For example, if you wnted to merge the `primeagen` branch of the remote `origin` into your local `main` branch, you would run this inside the local repo while on the `main` branch:

    git merge origin/primeagen
