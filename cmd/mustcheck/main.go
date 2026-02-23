package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/edaniels/golinters/mustcheck"
)

func main() {
	singlechecker.Main(mustcheck.Analyzer)
}
