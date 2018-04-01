package pkg_os
import (
	"os/exec"
	"github.com/robertkrimen/otto"
	"../pkg_stack"
)
var js *otto.Otto
var exec_stack = pkg_stack.New()
func Swap_data(js_engine *otto.Otto){
	*js = *js_engine
	js.Set("OS_COMMAND",func(call otto.FunctionCall)otto.Value{
		//
		return otto.Value{}
	})
}
//关于操作系统的一些操作

func Run_exec(order []string){
	exec_stack.Push(exec.Command(order[0],order[1:]...))
}