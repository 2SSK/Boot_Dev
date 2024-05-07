# PATH

There are environment variable that are sort of "built-in" to your shell. By "built-in" I just mean that different programs and parts of your system know about them and use them. The **PATH** variable is one of those.

### WHY DO WE CARE ABOUT THE PATH?

If it weren't for the **PATH**, you'd have to rememnber the filesystem path of every execuatble you wanted to run. Instead of just running **ls**, you'd have to run **/bin/ls** (or whatever the location of the **ls** executable is on your system). That's not very convenient.

The **PATH** variable is a list of directories that your shell will look into when you try to run a command. If you type **ls**, your shell will look in each directory listed in your **PATH** variable for an executable called **ls**. If it finds one, it will just run it. If it dosen't, it will give you an error like: "command not found".

### WHAT'S IN THE PATH VARIABLE?

Take a look at your current **PATH** variable:

    echo $PATH

you should see a giant list of directories separated by colons':':', Each of those directories is a placce where your shell will look for executables. For example, with a **PATH** like this:

    /usr/local/bin:/usr/bin:/bin:/urs/sbin:/sbin

Your shell will look for executables in the following directories:

    * /usr/local/bin
    * /usr/bin
    * /bin
    * /usr/sbin
    * /sbin
