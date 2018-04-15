// +build !windows

package pkg_load
import(
	"github.com/robertkrimen/otto"
)
type LoadedFileFunc struct{
	//func_map map[string]*plugin.Symbol
	//FilePointer *plugin.Plugin
	//has bool
}
type JSLoader struct{
	LoadedFile map[string]*LoadedFileFunc
	js *otto.Otto
}
func (this *JSLoader)SwapJS(engine *otto.Otto){
	//Do nothing......We will suport this soon...
	this.js = engine
}