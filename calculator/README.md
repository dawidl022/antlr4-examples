# Calculator interpreter

A simple calculator language interpreter supporting variable assignments and
expression evaluations implemented using [ANLTR4](https://www.antlr.org/) in Go.

Each statement is terminated by a newline, the results of evaluating statements
that are expressions are printed.

## Usage

Generate the parser:

```bash
make parser
```

Run the interpreter:

```bash
go run main.go
```

Give it some input, terminated with `Ctrl`+`D`:

```
193
a = 5
b = 6
a+b*2
(1+2)*3
```

After terminating the input, the interpreter should print the results. For the
above example, the output is:

```
193
17
9
```
