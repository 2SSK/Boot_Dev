# MERGE COMMITS

A merge commit is the result of merging two branches together.

Let's say we stat with this:

    A - B - C   main
     \
      D - E     vimchadsonly

And we merge `vimchadsonly` into `main` by runnig this while on `main`:

    git merge vimchadsonly

The merge will:

1. Find the "merge base" commit, or "best common ancestor" of the two branches. In this case, "A".
2. Replays the changes from `vimchadsonly` onto `main` starting from the best common ancestor.
3. Records the result as a new commit, in our case "F".
4. "F" is special because it hase _two parents_, "C" and "E".

**After**:

    A - B - C - F   main
     \      /
      D - E         vimchadsonly
