# PACKAGE MANGERS

A package manager is a software tool that helps you install other software. Its primary functions include:

- Downloading software from official sources
- Installing software
- Updating software
- Removing software
- Managing dependencies

As a developer, you'll frequently use paclage mangagers to get access to the software you need to get your work done.

---

### APT (UBUNTU)

APT, or "Advanced Package Tool", is the primary package manager for Ubuntu. To be fair, you can use other package managers on Ubuntu, but APT is the default and most common.

If you're on WSL and Ubuntu, you'll be using APT. If you're on another Linux setup, I pray you already know what package manager you're using. If you're on WSL or Ubuntu, check to make sure you have APT installed by running the following command:

    apt --version

---

### BREW (MAC OS)

There isn't a "default" package manager for Mac OS. The most popular (but unofficial) package manger is <span style="color:violet">Homebrew</span>.

If you're on Mac OS, and you don't have Homebrew installed, you can install it by running the following command:

    /bin/bash -c "$(curl -fsSL httpr://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
