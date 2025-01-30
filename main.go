package main

import (
	"github.com/sean9999/wgvet/pkg"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {

	//	use standard checkers
	checkers := otherCheckers()

	//	use my checker(s)
	checkers = append(checkers, pkg.JsonPerf, pkg.Cyclops)

	//	run them all
	multichecker.Main(checkers...)

}
