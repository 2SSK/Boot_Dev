# MAN

the **MAN** command is short for "manual". It's a program that displays the manual for other programs.

---

### USING MAN

The **MAN** command will only work for programs that it has a manual for, but most built-in commands and Unix programs are supported. You just pass the naem of the command you want to read the manual of as an argument. The most intuitive place to start, of course, is reading the manual-manual:

    # open the man pages for the 'man' command
    man man

---

### SEARCHING

Most people don't just _read man pages for fun_. More often, the manual is used as a reference to quickly look up usage or special flags.

You can search for text in the manual by pressing the **/** key and typing your search, then pressing enter. Let's try to look up what the **-r** flag does for the **ls** command:

    man ls
    # type '/-r' to start searching
    # press 'n' to jump to the next result
    # press 'N' to go back if you went too far
