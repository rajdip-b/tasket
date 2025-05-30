# tasket

**A blazing fast, developer-friendly CLI for managing your plain-text todos (todo.txt) in style, built with Go.**

![build](https://img.shields.io/badge/build-passing-brightgreen)
![license](https://img.shields.io/badge/license-MIT-blue.svg)

---

## 🚀 Overview

`tasket` is a feature-rich command-line interface to help you manage tasks the [todo.txt way](https://github.com/todotxt/todo.txt), with a modern Go twist. Quickly add, complete, filter, search, and organize your todos—right from your terminal.

- **Portable:** Works anywhere you can run Go binaries.
- **Efficient:** Designed for speed, even with huge todo lists.
- **Customizable:** Extensible and hackable for your workflow.

---

## 📝 Example Usage

```sh
# Add a new task
tasket add "Read Go documentation +golang @high"

# List tasks
tasket list

# Complete a task
tasket done 2

# Filter by context/project/priority
tasket list --project golang
tasket list --priority high
```

---

## 📦 Installation

#### With Go:

```sh
go install github.com/rajdip-b/tasket@latest
```
(Or download binaries from Releases when available.)

---

## ✨ Features
See FEATURES.md for the full, up-to-date list of features.

---

## 👨‍💻 Contributing
Contributions, bug reports, and feature requests are warmly welcomed!
Check out CONTRIBUTING.md (to be created) or open an issue to get started.

Made with ❤️ and Go by Raj