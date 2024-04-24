# CHANGING PERMISSIONS

- The 'chmod' command let you change the permissions of a file or directory. It's short for "change mode" (I wish it was called "change permissions", but alas).

- The ls command has a -l option (lowercase 'L') that will print out the permissions of each file and directory in long format.
  **_Example:_**
  _ls -l permission.md_
  -rw-r--r-- 1 saurav saurav 1461 Apr 24 13:11 permission.md

  > For permission.md file the user has permissions of read, write but not execute, and groups and others are only allowed to read.

  _chmod -R u=rwx,g=r--,o=r-- permission.md_
  -rwxr--r-- 1 saurav saurav 1461 Apr 24 13:11 permission.md

  > Now the permissions have changed, now the user can read, write and also execute but groups and others can only read.

* The -R means "recursively", which means "do this to all of the contents of the directory as well".
* Remember, . is a special alias for the current directory.
