package pkg_secret

import (
	"github.com/robertkrimen/otto"
)
var js *otto.Otto
func SwapData(engine *otto.Otto){
	js = engine
	js.Set("RJS_SECRET_TELLING",func(call otto.FunctionCall){
		
	})
}
