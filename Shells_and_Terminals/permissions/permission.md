# Permissions

the permissions of an individual file or directory are visually represented as a 10-character strig: drwxrwxrwx

## Files/Directories

    * -: Regular file(e.g. -rwxrwxrwx)
    * d: Dirctory (e.g. drwxrwxrwx)

The next 9 characters are broken up into 3 sets of 'rwx' and represent the permissions themselves fro the "owner", "group", and "others", in order.

Each group of 3 represents the permissions for reading, writing, and executing, in order. So, for example:

    - rwx: All permissions
    - rw-: Read and write, but not execute
    - r-x: Read and execute, but not write

- The first 3 characters are "owner" permissions. The "owner" is usually just the user who created the file or directory, but it can be manually changed.

- the next 4 characters are "group" permissions. unix-like systems support groups of users and a file or directory can be assigned to a single group. to be honest, unless you're a system administrator, you won't often worry about groups.

- The last 3 characters are "others" permissions. This is everyone else.

### Examples

    1. -rwxrwxrws: A file where everyone can do everything
    2. -rwxr-x-r-x: A file where everyone can read and execute, but only the owner can write
    3. drwdr-xr-x: A directory where everyone can read (ls the contents) and execute (cd into it), but only the owner can write (modify the contents)
    4. drwx------: A directory where only the owner can read, write and execute
