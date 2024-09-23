# TASK TRACKER

A simple command-line application for tracking tasks, written in Go. This program allows you to add, update, delete, and list tasks, as well as mark them as "in-progress" or "done". Tasks are stored in a JSON file for persistence.

## Features

- **Add Task**: Create new tasks with a status of "todo"
- **List Tasks**:  List all tasks or based on their status (todo, in-progress or done)
- **Update Task**:  Update the name of an existing task by its ID.
- **Mark Tasks as In Progress**: Mark a task as "in-progress".
- **Mark Tasks as Done**: Mark a task as "done".
- **Delete Task**: Remove tasks by their ID.

## Usage

### Add Task

Adds a new task with the given name and assigns it a "todo" status.

```bash
./task-tracker add "<task_name>"
```

### List Tasks

Lists tasks filtered by their status (todo, in-progress, or done). Leave <status> empty to list all tasks.

```bash
./task-tracker list <status>
```

### Update Task

Updates the name of the task with the specified ID.

```bash
./task-tracker update <task_id> "<new_name>"
```

### Mark Task as In Progress

Marks the task with the given ID as "in-progress".

```bash
./task-tracker mark-in-progress <task_id>
```

### Mark Task as Done

Marks the task with the given ID as "done".

```bash
./task-tracker mark-done <task_id>
```

###  Delete Task

Deletes the task with the given ID.

```bash
./task-tracker delete <task_id>
```

## Notes

- Ensure that the `task_tracker.json` file exists in the same directory as the script for it to function correctly.
- The application will create the `task_tracker.json` file if it does not already exist when adding a new task.

## Project Link

For more details about this project, visit the [Task Tracker Project Roadmap](https://roadmap.sh/projects/task-tracker).