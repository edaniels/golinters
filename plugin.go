package golinters

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"

	"github.com/edaniels/golinters/deferfor"
	"github.com/edaniels/golinters/errresp"
	"github.com/edaniels/golinters/fatal"
	"github.com/edaniels/golinters/mustcheck"
	"github.com/edaniels/golinters/printf"
	"github.com/edaniels/golinters/println"
	"github.com/edaniels/golinters/uselessf"
)

func init() {
	register.Plugin("println", newSyntaxPlugin(println.Analyzer))
	register.Plugin("printf", newSyntaxPlugin(printf.Analyzer))
	register.Plugin("uselessf", newSyntaxPlugin(uselessf.Analyzer))
	register.Plugin("deferfor", newSyntaxPlugin(deferfor.Analyzer))
	register.Plugin("errresp", newSyntaxPlugin(errresp.Analyzer))
	register.Plugin("mustcheck", newSyntaxPlugin(mustcheck.Analyzer))
	register.Plugin("fatal", newTypesPlugin(fatal.Analyzer))
}

func newSyntaxPlugin(a *analysis.Analyzer) register.NewPlugin {
	return func(any) (register.LinterPlugin, error) {
		return &linterPlugin{a: a, mode: register.LoadModeSyntax}, nil
	}
}

func newTypesPlugin(a *analysis.Analyzer) register.NewPlugin {
	return func(any) (register.LinterPlugin, error) {
		return &linterPlugin{a: a, mode: register.LoadModeTypesInfo}, nil
	}
}

type linterPlugin struct {
	a    *analysis.Analyzer
	mode string
}

func (p *linterPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{p.a}, nil
}

func (p *linterPlugin) GetLoadMode() string {
	return p.mode
}
