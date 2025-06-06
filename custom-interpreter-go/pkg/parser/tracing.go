package parser

import (
	"custom-interpreter-go/pkg/feature"
	"fmt"
	"strings"
)

var traceLevel int = 0

const traceIdentPlaceholder string = "\t"

func identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

func tracePrint(fs string) {
	fmt.Printf("%s%s\n", identLevel(), fs)
}

func incIdent() { traceLevel = traceLevel + 1 }
func decIdent() { traceLevel = traceLevel - 1 }

func trace(msg string) string {
	if feature.EnableTracing {
		incIdent()
		tracePrint("BEGIN " + msg)
	}

	return msg
}

func untrace(msg string) {
	if feature.EnableTracing {
		tracePrint("END " + msg)
		decIdent()
	}
}
