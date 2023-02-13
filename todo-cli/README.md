# Learning Go - REST API

A small [Go](https://go.dev) project that fetches data from an external API and retrieves it in your own endpoint.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What things you need to install the software and how to install them:

- [Go](https://go.dev/learn)

### Installing

1. Clone this repository to your local machine:

```bash
git clone https://github.com/isneru/learning-go.git
```

2. Navigate to the project directory:

```bash
cd learning-go/todo-cli
```

3. Build the project:

```go
go build
```

4. Now, depending on your GO environment variables, this is will build an executable accordingly. Take windows as example:

```bash
./todo-cli.exe -flag
```

## Flags

The following flags are available for you:

#### Adding a task

`-add <task name example>` - Adds a task.

#### Completing a task

`-complete <task number>` - Marks a task as done.

#### Deleting a task

`-delete <task number>` - Deletes a task.

#### Listing the tasks

`-list` - Returns a list of all tasks.
