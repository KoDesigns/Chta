# Git Cheat Sheet 🔧

*Essential Git commands for daily development*

## Basic Operations

```bash
# Initialize repository
git init

# Clone repository
git clone <url>
git clone <url> <directory>

# Check status
git status
git status -s              # Short format

# Add files
git add <file>
git add .                  # Add all files
git add -A                 # Add all changes
git add -p                 # Interactive staging
```

## Committing

```bash
# Commit with message
git commit -m "message"
git commit -am "message"   # Add and commit

# Amend last commit
git commit --amend
git commit --amend -m "new message"

# Empty commit (useful for CI triggers)
git commit --allow-empty -m "trigger build"
```

## Branching

```bash
# List branches
git branch
git branch -a              # Show all branches
git branch -r              # Show remote branches

# Create branch
git branch <branch-name>
git checkout -b <branch-name>
git switch -c <branch-name>

# Switch branches
git checkout <branch>
git switch <branch>

# Delete branch
git branch -d <branch>     # Safe delete
git branch -D <branch>     # Force delete
```

## Remote Operations

```bash
# Show remotes
git remote -v

# Add remote
git remote add origin <url>

# Fetch and pull
git fetch
git pull
git pull origin <branch>

# Push
git push
git push origin <branch>
git push -u origin <branch>  # Set upstream
git push --force-with-lease  # Safer force push
```

## Viewing History

```bash
# View log
git log
git log --oneline
git log --graph --oneline --all
git log -p                 # Show patches
git log --since="2 weeks ago"

# Show differences
git diff
git diff --staged
git diff HEAD~1
git diff <branch1>..<branch2>
```

## Undoing Changes

```bash
# Unstage files
git reset HEAD <file>
git restore --staged <file>

# Discard working changes
git checkout -- <file>
git restore <file>

# Reset commits
git reset --soft HEAD~1    # Keep changes staged
git reset --mixed HEAD~1   # Keep changes unstaged
git reset --hard HEAD~1    # Discard changes

# Revert commit
git revert <commit-hash>
```

## Stashing

```bash
# Stash changes
git stash
git stash save "message"
git stash -u               # Include untracked files

# List stashes
git stash list

# Apply stash
git stash apply
git stash apply stash@{0}
git stash pop              # Apply and delete

# Drop stash
git stash drop stash@{0}
git stash clear            # Clear all stashes
```

## Merging & Rebasing

```bash
# Merge branch
git merge <branch>
git merge --no-ff <branch>

# Rebase
git rebase <branch>
git rebase -i HEAD~3       # Interactive rebase

# During conflicts
git add <file>             # Mark as resolved
git rebase --continue
git rebase --abort
git merge --abort
```

## Tags

```bash
# List tags
git tag

# Create tag
git tag <tag-name>
git tag -a <tag-name> -m "message"

# Push tags
git push origin <tag-name>
git push origin --tags

# Delete tag
git tag -d <tag-name>
git push origin --delete <tag-name>
```

## Configuration

```bash
# Global config
git config --global user.name "Your Name"
git config --global user.email "your@email.com"
git config --global init.defaultBranch main

# Show config
git config --list
git config user.name

# Aliases
git config --global alias.st status
git config --global alias.co checkout
git config --global alias.br branch
git config --global alias.cm commit
```

## Pro Tips

- Use `git status` frequently to understand current state
- Write clear, descriptive commit messages
- Use branches for features and experiments
- Pull before pushing to avoid conflicts
- Learn interactive rebase for clean history
- Use `.gitignore` to exclude unnecessary files

---

*Happy coding! 🚀* 