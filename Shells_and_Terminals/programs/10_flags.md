# FLAGS

Some commands can take _flags_. Flags are options that you can pass to a command to change its behaviour.

For example, the **ls** command can take a **-l** flag to show a "long" listing of files:

    ls -l

The **ls** command can also take a **-a** flag to show "all" files, including hidden files:

    ls -a

You can combine flags:

    ls -al

---

### CONVENTIONS

Whether or not a command takes flags, and what those flags are, is up to the developer of the command.
That said there are some common conventions:

    - Single -character flags are prefixed with a single dash (.eg. -a)
    - Multi-character flags are prefixed with two dashes (e.g. --help)
    - Sometimes the same flag can be used with a single dash or two dashes (e.g. -h or --help)
