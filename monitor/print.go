package main

import (
	"fmt"
)

/*
This go file shoud be put in a directory outside of the directory of  main package.
We can build a plugin custom the pluginpath with  'go build -ldflags "-pluginpath=plugin/hot-$(date +%s)" -buildmode=plugin -o gplugin.so gplugin.go'.

'go build -buildmode=plugin -o print.so print.go'


*/
func PrintByPlugin(strIn string) {
	fmt.Println("string in plugin print.so: ", strIn)
	//go build --buildmode=plugin -o print.so monitor/print.go
}
