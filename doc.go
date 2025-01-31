/*

wgvet does what `go vet` does (more or less), plus some additional checks:

1. Cyclomatic complexity

2. Usage of [encoding/json] from stdlib. Some shops prefer the performance trade-off that certain drop-in replacements offer.

3. Usage of [errors] from stdlib. Some shops will want to ensure that you're using errors with stack traces.

It can be easily extended. See package documention: [pkg]

*/

package main
