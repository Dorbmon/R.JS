package pkg_load

import(
	"github.com/Dorbmon/otto"
)
type JSLoader struct{
	js *otto.Otto
}
func (this *JSLoader)SwapJS(js *otto.Otto){
	this.js = js

	//初始化函数
	//初始化map
}