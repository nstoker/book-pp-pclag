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

Can add archive files to the mix.

It is possible to install the binary version:

```bash
go install
type ch4-navigating
# walk is home/ns/go/bin/ch4-navigating
```

## Chapter 5: Improving performance of CLI tools

Up to p126, but getting a panic:

```bash
csv_test.go:70: [{Column2 3 [2056 899 3054 4133 950] <nil> 0xc00008a300} {Column3 3 [236 220 226 218 238] <nil> 0xc00008a330} {FailRead 1 [] 0xc00006a410 0xc00000c150} {FailedNotNumber 1 [] 0xc00006a420 0xc00008a390} {FailedInvalidColumn 4 [] 0xc00006a430 0xc00008a3c0}]
  panic: runtime error: index out of range [0] with length 0
192.168.0.199,2056,236
192.168.0.88,899,220
192.168.0.199,3054,226
  /usr/local/go/src/testing/testing.go:1209 +0x24e
192.168.0.199,950,238
  /usr/local/go/src/testing/testing.go:1212 +0x218
csv_test.go:72: tc.col 3
  /usr/local/go/src/runtime/panic.go:1038 +0x215
panic: runtime error: index out of range [0] with length 0 [recovered]
  /home/neil/Code/github/nstoker/book-pp-pclag/ch5-performance/colStats/csv_test.go:93 +0x454

  /usr/local/go/src/testing/testing.go:1259 +0x102
testing.tRunner.func1.2({0x50d540, 0xc000014390})
  /usr/local/go/src/testing/testing.go:1306 +0x35a
testing.tRunner.func1()
/usr/local/go/src/testing/testing.go:1212 +0x218
panic({0x50d540, 0xc000014390})
/usr/local/go/src/runtime/panic.go:1038 +0x215
github.com/nstoker/book-pp-pclag/performance/colStats.TestCSV2Float.func1(0xc0000ad860)
/home/neil/Code/github/nstoker/book-pp-pclag/ch5-performance/colStats/csv_test.go:93 +0x454
testing.tRunner(0xc0000ad860, 0xc00008a3f0)
/usr/local/go/src/testing/testing.go:1259 +0x102
created by testing.(*T).Run
/usr/local/go/src/testing/testing.go:1306 +0x35a
```
