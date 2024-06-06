# DIFFERENT HASHES

My latest Git commit has was:

    59cd5ac986f1b37d2f90642767866e1c8e1a1938

You may have noticed that even though we (you an i) both have the **same content** in our repositories, we have **different commit hashes**. While commit hases _are_ derived from their content changes, there's also some other stuff that affects the end hash. For example:

- The commit message
- The author's name and email
- The date and time
- Parent (previous) commit hashes

All this to say that hashes are (almost) always unique, and because they're generated automatically for you, you don't need to worry too much about what goes into them right now.

#### NOTE: HASH = SHA

Git uses a cryptographic hash function called SHA-1 to generate commit hashes. We won't go into the details of how SHA-1 works in this course, but it's important to know because you might also hear commit mashes referred to as "SHAs".
