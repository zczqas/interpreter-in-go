# Interpreter in Golang

This project is an interpreter for a simple programming language written in Go. The interpreter supports basic arithmetic operations, variable declarations, and function definitions.

## Project Structure


### Directories

- **ast/**: Contains the Abstract Syntax Tree (AST) definitions and related functions.
- **lexer/**: Contains the lexer implementation which is responsible for tokenizing the input source code.
- **parser/**: Contains the parser implementation which converts tokens into an AST.
- **repl/**: Contains the Read-Eval-Print Loop (REPL) implementation for interactive use.
- **token/**: Contains the token definitions used by the lexer and parser.

## Getting Started

### Prerequisites

- Go 1.24.0 or later

### Building the Project

To build the project, run:

```sh
go build