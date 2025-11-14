# Git Guardian

**Git Guardian** is a simple yet powerful CLI tool to automatically back up your project folders to GitHub. Designed for Mac users, Guardian watches your directories for changes or backs them up on a regular interval, committing and pushing changes automatically. It’s perfect for developers who want effortless version control without thinking about it.

The CLI tool is called `guardian`.

---

## 🚀 Core Features

* **Auto-Backup:** Automatically commit and push changes to a GitHub repository.
* **Watch Mode:** Monitor a folder for file changes and push updates with a configurable debounce time (e.g., every 30 seconds).
* **Interval Mode:** Push all changes at a set interval (e.g., every 5 minutes).
* **Folder Management:** Link folders to Guardian and manage them easily.
* **Listeners Dashboard:** View all active watchers, pause, resume, or remove them.
* **Autostart:** Guardian can start automatically when your Mac boots.
* **Safe and Reliable:** Stops watchers if a folder is deleted and prevents duplicate or conflicting commits.

---

## 💻 Installation

### From Source

1. Clone the repository:

```bash
git clone https://github.com/itzcodex24/git-guardian.git
cd git-guardian
```

2. Build the CLI tool:

```bash
go build -o guardian ./...
```

3. Move it to your `PATH` for global access:

```bash
sudo mv guardian /usr/local/bin/
```

### Homebrew (Recommended for users)

Once the Homebrew formula is published:

```bash
brew install itzcodex24/git-guardian/guardian
```

---

## ⚡ Usage

### Initialize Guardian in a folder:

```bash
guardian init
```

### Link a folder to Guardian:

```bash
guardian link ~/Projects/myproject
```

### Start automatic backup:

* **Watch mode (on change):**

```bash
guardian start ~/Projects/myproject --watch --debounce 30s
```

* **Interval mode (every 5 minutes):**

```bash
guardian start ~/Projects/myproject --interval 5m
```

### Manage watchers:

```bash
guardian listeners          # List active watchers
guardian pause <id>         # Pause a watcher
guardian resume <id>        # Resume a watcher
guardian remove <id>        # Remove a watcher
```

### Autostart Guardian on Mac login:

```bash
guardian autostart enable
guardian autostart disable
```

---

## 📝 Collaboration Guidelines

Git Guardian is open source! Your feedback, bug reports, and contributions are highly appreciated. Let’s make automatic Git backups easier for everyone!

We welcome contributions! Here’s how you can collaborate:

1. **Fork the repository** and create a feature branch:

```bash
git checkout -b feature/my-awesome-feature
```

2. **Write code and tests** for your feature.
3. **Commit changes** with clear messages:

```bash
git commit -m "Add feature X with Y improvements"
```

4. **Push to your fork** and open a pull request.
5. **Code review:** PRs will be reviewed, and feedback may be requested before merging.
6. **Issues:** Report bugs or suggest features via GitHub Issues.

**Code Style:**

* Follow Go conventions (`gofmt`, `golint`)
* Keep commits focused and atomic
* Document public functions and methods clearly

