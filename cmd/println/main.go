package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/edaniels/golinters/println"
)

func main() {
	singlechecker.Main(println.Analyzer)
}
