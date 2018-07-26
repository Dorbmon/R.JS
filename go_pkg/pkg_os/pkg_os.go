package pkg_os
import (
	"os/exec"
	"github.com/robertkrimen/otto"
	"../pkg_stack"
)
var js *otto.Otto
var exec_stack = pkg_stack.New()
func Swap_data(js_engine *otto.Otto){
	js = js_engine
	js.Set("OS_COMMAND",func(call otto.FunctionCall)otto.Value{
		//
		obj,_ := js.Object("")
		obj.Set("Run",func(call otto.FunctionCall)otto.Value{
			//获取参数
			n := 0
			var datas []string
			datas = make([]string,1)
			for {
				value := call.Argument(n)
				if !value.IsDefined(){
					break
				}
				datas[n] = value.String()
				n++

			}
			cmd := exec.Command(datas[0],datas[0:]...)
			err := cmd.Run()
			value,_ := cmd.Output()
			if err != nil{
				return otto.FalseValue()
			}
			rvalue,_ := otto.ToValue(value)
			return rvalue
		})
		return otto.Value{}
	})
}
//关于操作系统的一些操作

func Run_exec(order []string){
	exec_stack.Push(exec.Command(order[0],order[1:]...))
}
