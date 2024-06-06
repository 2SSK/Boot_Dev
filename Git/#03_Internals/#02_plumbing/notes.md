# THE PLUMBING

So far we've been using Git in a "procelain" manner. But to sate or instatiable curiosity, let's take a look at some of the "plumbing".

### IT'S JUST FILES ALL THE WAY DOWN

All the data in a Git repository is stored directly in the (hidden) `.git` directory. That includes all the commits, branches, tags and other bojects we'll learn about later.

Git is made up of objects that are stored in the `.git/objects` directory. A commit is just a type of object.
