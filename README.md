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
