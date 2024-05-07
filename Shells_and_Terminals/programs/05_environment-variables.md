# ENVIRONMENT VARIABLES

    To create and use local variables in your shell:
        name "Lane"
        echo $name
        #Lane

There are another type of variable called **environment variables**. They are available to all programs that you run in your shell.

You can view all of the environment variables that are currently set in your shell with the command **env**.

    If you want to set a variable in your shell, you can use the "export" command:
        export NAME="Lane"

> _The interesting part is that programs and scripts you run in your shell can also use that variable._
