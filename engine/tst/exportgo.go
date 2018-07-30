package main

import "C"
//import "fmt"
import ".."
//export PrintBye

func main() {
	// Need a main function to make CGO compile package as C shared library
}
func Run(){
	t := ""
	engine.Run(&t)

}