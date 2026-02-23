package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/edaniels/golinters/fatal"
)

func main() {
	singlechecker.Main(fatal.Analyzer)
}
