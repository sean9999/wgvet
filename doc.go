/*
wgvet is a binary go vet utility that does what go vet does (pretty much), plus some additional checks:

1. cyclomatic complexity

2. Usage of encoding/json from stdlib

3. Useage of "errors" from stdlib

It can be easily extended. See package documention: [pkg]
*/
package main
