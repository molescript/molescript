package repl

import (
	"bufio"
	"fmt"
	"io"
	"runtime"

	"molescript/lexer"
)

const PROMPT = ">>> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	fmt.Printf("MoleScript REPL on %s\n", runtime.GOOS)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lex := lexer.New(line)
		for {
			token := lex.ReadToken()
			if token.Type == lexer.EOF {
				break
			}

			fmt.Printf("%+v\n", token)
		}
	}
}
