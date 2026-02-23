package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/edaniels/golinters/uselessf"
)

func main() {
	singlechecker.Main(uselessf.Analyzer)
}
