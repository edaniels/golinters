package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/edaniels/golinters/deferfor"
)

func main() {
	singlechecker.Main(deferfor.Analyzer)
}
