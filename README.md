# Task CLI

My workly task-cli  (command-line interface) in terminal !

## Features

- **Add tasks**: Quickly add new tasks with descriptions.
- **List tasks**: View all tasks or filter them by status (e.g., "todo", "in-progress", "done").
- **Update tasks**: Modify task descriptions or change their status.
- **Delete tasks**: Remove completed or unnecessary tasks.
- **Persistent storage**: Tasks are saved to a JSON file for future use.

## Installation

1. Clone the repository:
   ```terminal
   git clone https://github.com/yourusername/task-cli.git 
   cd task-cli
2. Build the project:
   ```terminal
   go build -o task-cli
3. (Optional) Move the binary to a directory in your PATH:
   ```terminal
   sudo mv task-cli /usr/local/bin/
