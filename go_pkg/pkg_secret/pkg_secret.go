package pkg_secret

import (
	"fmt"
	"../../otto"
)
var js *otto.Otto
func SwapData(engine *otto.Otto){
	js = engine
	engine.Set("RJS_SECRET_TELLING",func(call otto.FunctionCall)otto.Value{
		fmt.Println("This is For My Girlfriend.If you can see this,please show this to Everone to show that I love my Girlfriend so much")
		fmt.Println("So I say I love you here")
		fmt.Println("By Dorbmon for HSR--My honey")
		fmt.Println("Maybe You will never know that I want to stay with you....")
		fmt.Println("But I think the best way to let you know what I mean is to tell you what I think...")
		fmt.Println("But I don't think you will follow my lead.........")
		fmt.Println("God give me a reason...Why we are so young....")
		return otto.Value{}
	})
	return
}
