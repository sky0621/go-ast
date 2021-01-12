package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"
	"skel"
)

func main() { unitchecker.Main(skel.Analyzer) }
