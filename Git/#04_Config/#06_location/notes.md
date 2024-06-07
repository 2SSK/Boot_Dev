# LOCATIONS

There are several locations where Git can be configured. From more general to more specific, they are:

- **system**: `/etc/gitconfig`, a file that configures Git for all users on the system
- **global**: `~/.gitconfig`, a file that configures Git for all projects of a user
- **local**: `.git/config`, a file that configures Git for a specific project
- **worktree**: `.git/config.worktree`, a file that configures Git for part of a project

In my experience, 90% of the time you will be using `--global` to set things like your username and email. The other 9% of the time you will be using `--local` to set project-specific configurations. The last 1% of the time you _might_ need to futz with system and worktree configurations, but it's extremely rare.

### OVERRIDING

If you set a configuration in a more specific location, it will override the same configuration in a more general location. For example, if you set `user.name` in the configuration, it will override the `user.name` set in the global configuration.
