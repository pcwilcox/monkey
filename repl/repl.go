// repl/repl.go
//
// defines the Read-Eval-Print Loop functionality

package repl

import (
	"bufio"
	"fmt"
	"io"

	"../lexer"
	"../token"
)

// PROMPT is the shell prompt
const PROMPT = ">> "

// Start runs the loop
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
