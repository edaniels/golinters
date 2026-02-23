package uselessf_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/edaniels/golinters/uselessf"
)

func TestAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), uselessf.Analyzer, "a")
}
