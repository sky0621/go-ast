package skel_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
	"skel"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, skel.Analyzer, "a")
}
