package fatal_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/edaniels/golinters/fatal"
)

func TestAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), fatal.Analyzer, "a")
}
