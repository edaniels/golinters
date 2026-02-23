package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/edaniels/golinters/printf"
)

func main() {
	singlechecker.Main(printf.Analyzer)
}
