# TREES AND BLOBS

Now that we understand some of our plumbing equipment, let's get into the pipes. Here are some terms to know:

- `tree`: git's way of storing a directory
- `blob`: git's way of storing a file

Here's what i got when i inspected my last commit:

    > git cat-file -p 59cd5ac986f1b37d2f90642767866e1c8e1a1938

    tree bdfc7a4e697a001b46b68670fc4cff2f323fc1c4
    parent 87829bbcbd8216a75bce4ec96aec9167e1efeb29
    author 2SSK <100ravsinghkarmwar@gmail.com> 1717614520 +0530
    committer 2SSK <100ravsinghkarmwar@gmail.com> 1717614520 +0530

    Add Git repository notes and update README

Notice that we can see:

- The `tree` object
- The `author`
- The `committer`
- The commit message

However, we _can not_ see the contents of the `contents.md` file itself! That's because the `tree` object stores it.
