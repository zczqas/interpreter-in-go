package repl

import (
	"bufio"
	"fmt"
	"go_inter/evaluator"
	"go_inter/lexer"
	"go_inter/parser"
	"io"
)

const (
	PROMPT = ">>"
	KOMI   = `
	⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⣰⠿⣿⣿⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⣴⣦⠀⠀
⠀⠀⠀⠀⠀⠀⢰⡟⠀⠸⣿⣿⣿⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣴⣿⣿⣿⠋⣿⠃⠀
⠀⠀⠀⠀⠀⠀⣼⠁⣿⣟⣿⣿⣿⣿⣷⡄⠀⣀⣀⣀⣀⣀⣀⣀⣀⠀⣠⣾⣿⣿⣿⣿⡅⢸⣿⠀⠀
⠀⠀⠀⠀⠀⠀⣿⠀⢯⣉⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠟⣿⢛⣿⣟⣛⠛⣿⣟⠸⣿⣿⡇⠀⠀
⠀⠀⠀⠀⠀⠀⣿⣤⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠏⣸⣿⣿⣿⣿⣏⣸⣶⣿⣿⣿⣿⠀⠀⠀
⠀⠀⠀⠀⣠⢴⣿⣿⣿⣿⣿⣿⣿⠿⣻⣿⡿⣿⣿⣿⣿⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀
⠀⠀⣠⢾⣅⣾⣿⣿⣿⣿⣿⠟⠁⠀⣿⣿⣆⣿⣿⣿⣿⣛⣿⠿⣿⢿⣿⣿⣿⣿⣿⣿⣿⣅⡀⠀⠀
⠀⣼⢫⣿⣿⣿⣿⣿⣱⡟⠁⠀⠀⠀⢻⣿⣿⣿⣿⣿⣿⣿⣟⢿⣿⣿⣿⣿⣿⣿⡛⠛⢿⣿⣿⣷⡄
⢸⣿⣿⣿⣿⣿⣿⣿⠏⠀⠀⢀⣀⣀⣀⡙⠪⠙⠛⠿⠿⠿⠿⢆⣹⣿⣿⣿⣿⣿⠻⣷⣼⣿⡟⠛⣿
⠀⠻⣿⣿⣿⣿⣿⡿⠀⣠⠞⠋⣉⣀⣌⡉⠛⠂⠀⠀⠀⠐⠚⠉⣉⣤⣉⠉⠓⢎⡁⢻⣿⣿⣷⣾⠏
⠀⠀⠸⡿⣿⣿⣿⣧⠘⠃⠀⡾⠟⠁⠙⢿⡄⠀⠀⠀⠀⠀⠀⡾⠟⠀⠻⢷⡄⠈⠓⠈⣿⣿⣿⠋⠀
⠀⠀⠀⠈⣨⣿⣿⡟⠛⠀⠀⠻⣷⣄⣼⠿⠃⠀⠀⠀⠀⠀⠀⠻⣷⣀⣴⠟⠁⠀⠀⠀⠸⠟⠓⢦⡀
⠀⠀⠀⢸⣿⣿⣿⡇⠀⠈⠲⠤⣤⣤⣤⠴⠊⠀⠀⠀⠀⠀⠈⠢⢤⣭⣥⡤⠖⠃⠀⠀⣠⣺⡟⢸⠇
⠀⠀⠀⠀⢻⣿⣿⡇⠀⠀⠀⠀⡖⣄⢀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣦⣶⠀⠄⢠⣿⠗⢀⠞⠀
⠀⠀⠀⠀⢸⣿⣿⣷⡀⠀⠀⢠⠺⠟⠛⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣸⡵⠷⠃⢀⣾⣴⠖⠁⠀⠀
⠀⠀⠀⠀⢸⣿⣿⣿⣷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠁⢀⣴⣿⣿⣿⠀⠀⠀⠀
⠀⠀⠀⠀⢸⣿⣿⢻⣿⣿⣿⣶⣤⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣤⣶⣿⣿⣿⡇⣿⠀⠀⠀⠀
⠀⠀⠀⠀⣸⠹⣿⠘⣿⣿⣿⣿⣿⣿⣿⣷⣶⣦⠀⠀⠀⣶⣶⣿⣿⣿⣿⣿⣿⣿⡏⣇⣿⠀⠀⠀⠀
⠀⠀⠀⠀⣿⣦⣼⢀⣿⣃⣹⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⣿⣿⣿⣿⣿⣿⣭⡅⣶⣆⣿⢹⡀⠀⠀⠀
⠀⠀⠀⣸⣿⣿⣏⣾⣿⣿⣿⣿⣿⣿⣿⡿⠁⣯⠀⠀⠀⣿⠈⢿⣿⣿⣿⣿⣧⣿⣿⣿⡜⡇⠀⠀⠀
⠀⠀⢀⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⢟⡟⠀⠀⠸⣆⠀⣸⠃⠀⠈⢿⡛⠻⠿⣿⣿⣿⣿⣧⢻⡀⠀⠀
⠀⠀⣸⣗⣿⣿⣿⣿⠟⠋⠁⠀⢀⡞⠀⢀⣠⣄⣹⣶⣋⡤⢤⣀⠀⢳⡀⠀⠀⠉⠙⢿⣿⡼⡇⠀⠀
⠀⢠⣿⣧⣿⣿⣿⠘⣆⠀⠀⠀⢾⣶⣿⣿⣿⣷⣽⣿⣽⣿⣿⣿⣿⣶⡗⠀⠀⠀⣼⠁⢿⡇⣧⠀⠀
⠀⠘⣿⣿⣿⣿⣿⣆⠘⣆⠀⠀⠘⣟⣟⣿⣿⣿⣿⡿⣿⣿⡿⡿⣷⣿⠃⠀⠀⣰⠃⠀⠘⣿⡿⠀⠀
⠀⠀⠉⠉⠉⠉⠉⠈⠁⠈⠀⠀⠀⠙⠛⠙⠙⠉⠁⠁⠉⠉⠋⠉⠉⠉⠀⠀⠘⠁⠀⠀⠀⠛⠁⠀⠀
	`
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, KOMI)
	io.WriteString(out, "Woops! We ran into an error!\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
