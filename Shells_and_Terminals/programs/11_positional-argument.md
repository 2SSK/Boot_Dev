# POSITIONAL ARGUMENTS

Programming languages have functions, and functions take arguments. For example, this python function takes two arguments: **xp** and **level**:

    def print_player(xp, level):
        print("player has", xp, "xp and is level", level)

It can then be called with two arguments:

    print_player(100, 2)
    # Player has 100 xp and is level 2

In a shell, command (programs) can also take arguments. For example, the **cd** command takes a single argument (the directory to change to):

    cd /home/wagslange

Other commands might take mulitple arguments. For example, the **mv** command takes two arguments: the file to move, and the destination to move it to:

    mv file.txt newfile.txt
