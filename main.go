package main

import "molescript/repl"
import "os"

func main() {
	repl.Start(os.Stdin, os.Stdout)

}
