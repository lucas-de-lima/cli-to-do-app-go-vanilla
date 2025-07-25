# Simple Task Manager (Study Purpose Only)

This is a minimal Go program that demonstrates how to create a basic task manager.

## Overview

- The program prompts the user to enter a task description.
- It creates a task with the description, a "not done" status, and a start timestamp.
- It stores the task in memory (no data is saved permanently).
- Finally, it prints the list of tasks registered during that run.

## Important Notice

**This program is for educational and learning purposes only.**  
It does **not** include features such as:

- Data persistence (no saving/loading from files or databases)
- Multiple tasks input (only one task per run)
- Task completion or editing
- Error handling beyond basic input reading

## How to Run

1. Make sure you have [Go installed](https://golang.org/doc/install).
2. Save the source code in a file, for example, `main.go`.
3. Run the program using:
```bash
   go run main.go
````

4. When prompted, type the task description and press Enter.
5. The program will display the task list with timestamps.

## Code Summary

The program defines a `Todo` struct with the following fields:

* `Description` (string): the task text entered by the user
* `Done` (bool): completion status (currently always false)
* `StartData` (time.Time): timestamp when the task was created
* `EndData` (time.Time): timestamp when the task is completed (not used yet)

The main function handles user input via `bufio.Scanner`, trims whitespace, and appends the new task to a slice of `Todo`.
