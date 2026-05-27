package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/mna/pigeon/examples/json"
)

func main() {
	in := os.Stdin
	nm := "stdin"
	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			err := f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
		in = f
		nm = os.Args[1]
	}

	b, err := io.ReadAll(in)
	if err != nil {
		log.Fatal(err)
	}

	got, err := json.Parse(nm, b)
	if err != nil {
		fmt.Println(caretError(err, string(b)))
		os.Exit(1)
	}
	fmt.Println(got)
}

func caretError(err error, input string) string { _ = "STUB: not implemented"; return "" }

func extractLine(input string, initPos int) string { _ = "STUB: not implemented"; return "" }
