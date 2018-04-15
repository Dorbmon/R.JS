package pkg_secret

import (
	"fmt"
	"github.com/robertkrimen/otto"
)
var js *otto.Otto
func SwapData(engine *otto.Otto){
	fmt.Print("set...")
	js = engine
	engine.Set("RJS_SECRET_TELLING",func(call otto.FunctionCall)otto.Value{
		fmt.Println("This is For My Girlfriend.If you can see this,please show this to Everone to show that I love my Girlfriend so much")
		fmt.Println("So I say I love you here")
		fmt.Print("By Dorbmon for HSR--My honey")
		return otto.Value{}
	})
	return
}
