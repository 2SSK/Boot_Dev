# MERGE LOG

Your output from `git log --oneline --decorate --graph --parents` (aside from the hashes) should look _something_ like:

    *   89629a9 d234104 b8dfd64 (HEAD -> main) F: Merge branch 'add_classics'
    |\
    | * b8dfd64 fba0999 (tag: 5.8, add_classics) D: add classics
    * | d234104 fba0999 (tag: 6.1) E: update contents
    |/
    * fba0999 1381199 (tag: 3.8, origin/master, origin/main, master) C: add quotes
    * 1381199 a21228f (tag: 3.7) B: add titles.md
    * a21228f A: add contents.md

Each asterisk `*` represents a commit in the repository. There are mulitple commit hashes on each line because the `--parent` flag logs the parent has(es) as well.

- The first line, with these three hashes: `89629a9 d234104 b8dfd64` is our recent merge commit. The first hash,, `89629a9` is the merge commit's hash, and the other two are the parent commits.
- The next section is visual representation of the branch structure. It shows the commits on the `add_classics` branch and the `main` branch before the merge Notice that they both share a common parent.
- The next three lines are just "normal" commits, each poiting to their parent.
- The last line is the initial commit and terefore has no parent.
