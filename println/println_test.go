package println_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/edaniels/golinters/println"
)

func TestAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), println.Analyzer, "a")
}
