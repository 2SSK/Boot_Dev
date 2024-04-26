# chown

    the chown command allows you to change the owner of a file or directory, and it requires root privileges.

    So when do you need to use `sudo`? `chmod` allows you to change the permissions of any file or directory that you own.

    Example:  [sudo chown -R root contacts]

    Let's break down this command:
        * sudo:     Run as the root user
        * chown:    Command to change the owner
        * -R:       "Recursive", meaning also apply the changes to everything inside the directory
        * root:     The name of the new owner
        * contacts: The directory to change the owner of
