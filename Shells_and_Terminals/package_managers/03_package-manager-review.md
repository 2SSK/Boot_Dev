# PACKAGE MANAGER REVIEW

Apt and Brew aren't the only package managers out there. There are many, many more, though they do happen to be two of the most popular, especially on Linux and Mac OS.

---

### HOW DOES A PACKAGE MANAGER WORK?

When you type a command like **apt install neovim**, the package manager will:

1. Check to see if the package is already installed.
2. If it's not installed, it will download the package from a repository.
3. It will install the package on your computer.
4. It will install any dependencies that the package needs to run.
5. It will (hopefully) add the package to your PATH if it should be there.

Good package mangers keep track of what packages you have installed, and what versions of those packages you have installed. They keep your filesystem nice nad tidy, making sure you haven't installed 10 different instances of the same package or application.

For your edification, take a look at where your package manager installed **nvim** on your filesystem. The **which** command will help:

    which nvim
