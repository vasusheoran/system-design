package repl

import (
	"bufio"
	"custom-interpreter-go/pkg/lexer"
	"custom-interpreter-go/pkg/token"
	"fmt"
	"io"
	"strings"
)

const (
	PROMPT = ">> "
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		if strings.EqualFold(strings.Trim(line, "\n"), "quit") {
			return
		}

		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
