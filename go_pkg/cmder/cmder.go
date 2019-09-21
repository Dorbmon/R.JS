package cmder

import (
	"fmt"
	"github.com/Dorbmon/otto"
	"io"
	"log"
	"os/exec"
)

/*
控制台相关操作
 */


type Pkg_Cmder struct{
	js *otto.Otto
}
func (this Pkg_Cmder) SwapJs (js *otto.Otto) {
	this.js = js
	//开始初始化
	js.Set("NewCmder",func(call otto.FunctionCall)otto.Value {
		//新建一个cmder
		name := call.Argument(0).String()
		var ArgumentList []string
		for _,v := range call.ArgumentList [1:] {
			ArgumentList = append(ArgumentList, v.String ())
		}
		//开始创建
		cmd := exec.Command(name,ArgumentList...)
		var stdout io.ReadCloser
		var err error
		if stdout, err = cmd.StdoutPipe(); err != nil { //获取输出对象，可以从该对象中读取输出结果
			log.Fatal(err)
			return otto.NullValue()
		}
		obj,_ := js.Object("({})")
		obj.Set("Read",func(call otto.FunctionCall)otto.Value {
			var buffer []byte
			stdout.Read(buffer)
			ret,_ := otto.ToValue(string (buffer))
			return ret
		})
		obj.Set("Close",func(call otto.FunctionCall)otto.Value {
			stdout.Close()
			return otto.NullValue()
		})
		obj.Set("Start",func(call otto.FunctionCall)otto.Value {
			//非阻塞
			cmd.Start ()
			return otto.NullValue()
		})
		var stdin io.WriteCloser
		stdin,err = cmd.StdinPipe()
		if err != nil {
			fmt.Println (err)
		return otto.NullValue()
	}
	obj.Set("Write",func(call otto.FunctionCall)otto.Value {
		val := call.Argument(0).String ()
		stdin.Write([]byte(val))
		return otto.NullValue()
	})
	ret,_ := otto.ToValue(obj)
	return ret
	})
}
