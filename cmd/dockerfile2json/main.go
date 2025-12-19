package main

import (
	"fmt"
	"os"

	"github.com/takanoriyanagitani/go-dockerfile2json"
)

func main() {
	if e := dockerfile2json.StdinToDockerfileToParsedToJsonToStdout(); nil != e {
		fmt.Fprintf(os.Stderr, "Error: %v\n", e)
		os.Exit(1)
	}
}
