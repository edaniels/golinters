package main

import (
	"github.com/edaniels/golinters/fatal"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(fatal.Analyzer)
}
