# PIPING

One of the most beautiful things about the shell is that you can **pipe** the output of one program into the input of another program. With this one simple concept, you can run incredibly powerful automation tasks.

## PIPE

The pipe operator is **|**. It's the character that looks like a vertical line. It's usually on the same key as the backslash(\*\*\*\*) above the enter key. The pipe operator takes the stdout of the program on the left and "pipes" it inot the stdin of the program on the right.

    echo "Have you heared the tragedy of Darth Plagueis the Wise?" | wc -W
    # 10

In the example above, the **echo** command sends "Have you heared the tragedy of Darht Plagueis the Wise?" to stdout.

However, instead of that text being sent to your terminal, the pipe operator pipes it into the **wc** (word count) command. The **wc** command counts the number of words in the input it receives. The **-w** flag tells **wc** to only count words.

This only works because the **wc** command, like most shell commands, can optionally read from stdin instead of a filepath.
