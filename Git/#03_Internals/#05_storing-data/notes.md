# STORING DATA

Git stores an entire _snapshot_ of files on a _per-commit_ level. This was a surprise to me! I always assumed each commit only stored the _changes_ made in that commit.

### OPTIMIZATION

While it's true that Git stores entire snapshots, it _does_ have some performance optimizations so that your `.git` directory doesn't get too unbearably large.

- Git compresses and packs files to store them more efficiently.
- Git deduplicates files that are the same across different commits. If a file doesn't chage between commits, Git will only store it once.
