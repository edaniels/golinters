package errresp_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/edaniels/golinters/errresp"
)

func TestAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), errresp.Analyzer, "a")
}
