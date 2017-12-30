package main

import "./repl"
import "os"

func main()  {
	repl.Start(os.Stdin, os.Stdout)
	
}




