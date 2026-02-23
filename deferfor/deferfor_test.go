package deferfor_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/edaniels/golinters/deferfor"
)

func TestAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), deferfor.Analyzer, "a")
}
