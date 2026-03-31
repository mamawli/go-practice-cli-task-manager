# Go CLI Task Manager

A clean Command Line Interface (CLI) Task Manager built with Go.

## Features

- Add new tasks
- List all tasks
- Get task details by ID
- Mark task as done
- Delete task
- Persistent storage using JSON file (`tasks.json`)


## Available Commands

| Command              | Description                        | Example |
|----------------------|------------------------------------|--------|
| `add`                | Add a new task                     | `go run cmd/cli/main.go add "Learn Go"` |
| `list`               | List all tasks                     | `go run cmd/cli/main.go list` |
| `get <id>`           | Get task by ID                     | `go run cmd/cli/main.go get 1` |
| `done <id>`          | Mark task as completed             | `go run cmd/cli/main.go done 1` |
| `delete <id>`        | Delete a task                      | `go run cmd/cli/main.go delete 1` |

## How to Run

```bash
# Clone the repo
git clone https://github.com/mamawli/go-practice-cli-task-manager.git
cd go-practice-cli-task-manager

# Run examples:
go run cmd/cli/main.go add "Complete Project 1"
go run cmd/cli/main.go list
go run cmd/cli/main.go get 1
go run cmd/cli/main.go done 1
