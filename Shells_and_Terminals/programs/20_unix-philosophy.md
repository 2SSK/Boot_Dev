# UNIX PHILOSOPHY

The <span style="color:violet">Unix Philosophy</span> is a simple set of principles that have guided the development of Unix-like operating systems for decades. It can be summarized as:

1. Write programs that do one thing and do it well.
2. Write programs to work together.
3. Write programs to handle text streams, because that is a universal interface.

---

### 1. WRITE PROGRAMS THAT DO ONE THING AND DO IT WELL

This is why programs like **ls**, **grep**, and **less** exist. They do one thing, and they do it well. They don't try to do too much.

- **ls** lists files and directories
- **grep** searches for text
- **less** displays text

---

### 2. WRITE PROGRAMS TO WORK TOGETHER

Because, at least according to the Unix Philosophy, programs should do one thing and do it well, it's easy to write programs that work together. For example, you can use **grep** to search for text in a file, and then pipe the output of **grep** into **less** to display the results interactively:

    grep "hello" some_file.txt | less

---

### 3. WRITE PROGRAMS TO HANDLE TEXTX STREAMS, BECAUSE THAT IS A UNIVERSAL INTERFACE

This point is more the "how" of the previous point. Programs work together easily when they all use the same interface: text streams. A text stream is just a sequence of characters that can be read or written sequentially. In other words, a text stream is just text.

This hearkens back to the point we talked about at the beginning of this course: the shell is a command-line(text) interface. Text-based interfaces are much more powerful and extensible than graphical interfaces. That's why developers have been using them for decades, and why what we can do with them looks like magic to the uninitiated.
