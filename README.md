Here's a simple README for your Todo CLI app:

```markdown
# Todo CLI - Command Line Task Manager

A simple command-line todo application written in Go that stores tasks in a JSON file.

## Installation

```bash
# Clone the repository
git clone <git remote add origin https://github.com/KernelH132/todo-cli.git
>

# Navigate to the directory
cd todo-cli

# Build the executable
go build -o todo

# (Optional) Move to your PATH
sudo mv todo /usr/local/bin/
```

## Usage

### Add a new task
```bash
./todo add "Buy groceries"
# Output: Added task: Buy groceries
```

### List all tasks
```bash
./todo list
```

Example output:
```
Todo List:
----------
1. [ ] Buy milk
2. [✓] Walk dog
3. [ ] Read book
```

### Delete a task
```bash
./todo delete 2
# Output: Task 2 deleted
```

### Get help
```bash
./todo help
```

## Commands

| Command | Description | Example |
|---------|-------------|---------|
| `add <title>` | Add a new task | `todo add "Write code"` |
| `list` | Show all tasks | `todo list` |
| `delete <id>` | Delete a task by ID | `todo delete 3` |
| `help` | Show help message | `todo help` |

## Data Storage

Tasks are stored in `task.json` in the same directory as the executable. The file is created automatically when you add your first task.

Example `task.json`:
```json
[
  {
    "id": 1,
    "title": "Buy milk",
    "completed": false
  },
  {
    "id": 2,
    "title": "Walk dog",
    "completed": true
  }
]
```

## Features

- ✅ Add new tasks
- 📋 List all tasks with status
- 🗑️ Delete tasks by ID
- 💾 Persistent storage using JSON
- 🎯 Auto-incrementing task IDs

## Building from Source

```bash
# Build for current platform
go build -o todo

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o todo-linux

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o todo.exe

# Build for macOS
GOOS=darwin GOARCH=amd64 go build -o todo-mac
```

## Requirements

- Go 1.16 or higher