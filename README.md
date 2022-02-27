# book-pp-pclag

Working on the Pragmatic Programmers book "Powerful Command Line Applications In Go".

## Chapter 1

Creates a word count program.

Added a flag to count lines only.

Adding the `GOOS` flag will compile for a different OS, for example:

```bash
GOOS=windows go build
file wc
# wc.exe: PE32+ executable (console) x86-64 (stripped to external PDB), for MS Windows
```

## Chapter 2

A ToDo list.

Flag       | Action
-----------|--------
list       | lists all items in the task list
complete n | completes task number n
add text   | adds text to the task list from the command line
del n      | deletes task number n
verbose    | produce verbose output
outstanding| only shows outstanding items

The program will also accept arguments passed via a pipe, eg

```bash
echo "Piped in" | ./todo -add
./todo -add added from command line
./todo -complete 1
todo -list
```

produces:

```text
X 1: Piped in
  2: added from command line
```

Copying the example.env to `.env` and setting the `TODO_FILENAME` will set the name the ToDo lists
will be saved as.

Exercises not completed:

1. Include test cases for the verbose option
1. Update the `getTask()` function allowing it to handle multiline input from STDIN. Each line should be a new task in the list.
   - and associated test

## Chapter 3 Markdown to HTML converter

Converts a Markdown file to a html file. Includes a preview facility.

An alternate template file added that adds a blue h1 style.

Flag | Purpose
-----|---------
file | the name of the file to convert
s    | skip auto preview
t    | use a different template file

Exercised not done.

## Chapter 4: Navigating the File System

Adds a directory walk tool, with a delete flag and a logging facility.

It is possible to install the binary version:

```bash
go install
type ch4-navigating
# walk is home/ns/go/bin/ch4-navigating
```
