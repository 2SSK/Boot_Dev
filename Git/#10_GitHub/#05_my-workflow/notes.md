# MY WORKFLOW

I wnat to leave you with a note on my simple workflow. 90% of the time, you're only using a handful of git commands to get your coding work done.

### KEEP STUFF ON GITHUB

I keep all my serious projects on GitHub. That way if my computer explodes, I have a backup, and if I'm ever on another computer, I can just clone ther repo and get back to work.

### REBASE VS MERGE

I've configured Git to rebase by default on pull, rather than merge so I keep a linear history. If you want todo the same, you can add this to your global Git config:

    git config --global pull.rebase true

### MY SOLO WORKFLOW

When I'm working by myself, I usually stick to a single branch, `main`. I mostly use Git on soo projects to keep a backup remotely and to keep a history of my changes. I only rarely use separate branches.

1. Make changes to files
2. `git add .` (or `git add <files>` if i only wnat to add specific files)
3. `git commit -m "a message describing the changes"`
4. `git push origin main`

It really is that simple for most solo work. `git log`, `git reset`, and some others are ofcourse useful from time to time, but the above is the core of waht i do day-to-day.

### MY TEAM WORKFLOW

When you're working with a _team_ Git gets a bit more involved. Here's what i do:

- Update my local `main` branch with `git pulll origin main`
- Checkout a new branch for the changes i want to make with `git switch -c <branch_name>`
- Make changes to files
- `git add .`
- `git commit -m "a message describing the changes"`
- `git push origin <branch_name>` (I push to the _new_ branch name, not `main`)
- Open a <span style='color:blue'>pull request</span> on GitHub to merge my changes into `main`
- Ask a team member to review my pull request
- Once approved, click the "Merge" button on GitHub to merge my changes into `main`
- Delete my feature branch, and repeat with a new branch for the next set of changes
