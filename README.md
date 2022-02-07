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
