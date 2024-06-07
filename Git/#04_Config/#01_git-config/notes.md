# GIT CONFIG

Git stores author information so that when you're making a commit it can track _who_ made the change. Here's how you might update your global <span style="color:blue">Git configuration</span> :

    git config --add --global user.name "2SSK"
    git config --add --global user.email "100ravsinghkarmwar@gmail.com"

Let's take the command apart:

- `git config`: The command to interact with your Git configuration.
- `--add`: Flag stating you want to _add_ a configuration.
- `--global`: Flag stating you want this configuration to be stored globally in your `~/.gitconfig`. The opposite is "local", which stores the configuaration in the current repository only.
- `user.name`: The configuration key you want to set.
- `2SSK`: The value you want to set for the key.
