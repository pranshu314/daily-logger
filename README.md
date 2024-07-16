# lg - Daily Logger for Developers

`lg` is a command-line tool that allows developers to quickly log their thoughts and progress while working on projects.

## Installation

```bash
git clone https://github.com/pranshu314/daily-logger.git
go install
make build
```

## Usage

```bash
lg [flags]
lg [command]
```

_Available Commands:_

- `completion` - Generates autocompletion script for the specified shell
- `delete` - Deletes log entry by ID
- `help` - Help about any command
- `list` - List all logs for the specified project
- `projects` - List all projects
- `project` - Add log to the specified project
- `update` - Update log entry by ID
- `where` - Show where your logs are being stored

_Flags:_

- `-h --help` - Displays the help message

## Getting Help

For information about a specific command, use:

```bash
lg [command] --help
```

## Tech Stack

- Golang
- Cobra
- Sqlite
