package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/edaniels/golinters/errresp"
)

func main() {
	singlechecker.Main(errresp.Analyzer)
}
