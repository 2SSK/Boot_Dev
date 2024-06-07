# GIT FILES

KungFu? No, _Gitfu_.

Remember, Git stores all its information in files in the `.git` subdirectory at the root of your project. even information about branches. The "heads"(or "tips") of branches are stored in the `.git/refs/heads` directory. If you `cat` one of the files in that direcotry,k you should be able to see the commit hash taht the branch points to.
