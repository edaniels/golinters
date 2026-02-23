package printf_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/edaniels/golinters/printf"
)

func TestAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), printf.Analyzer, "a")
}
