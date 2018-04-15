// +build windows

//Only Windows
package pkg_load
//RJS动态加载链接库

import(
	"plugin"
	"github.com/robertkrimen/otto"
	"fmt"
	"strings"
	"os"
	"syscall"
	"unsafe"
)
type LoadedFileFunc struct{
	func_map map[string]*plugin.Symbol
	FilePointer *plugin.Plugin
	has bool
}
var LoadedFile map[string]*LoadedFileFunc
func (this *JSLoader)SwapJS(engine *otto.Otto){
	this.js = engine
	LoadedFile = make(map[string]*LoadedFileFunc,1)
}
type JSLoader struct{
	LoadedFile map[string]*LoadedFileFunc
	js *otto.Otto
}
func (this *JSLoader)Load(file_name string)bool{
	if LoadedFile[file_name].has{
		return true
	}
	temp,err := plugin.Open(file_name)
	if err != nil{
		return false
	}
	this.LoadedFile[file_name].has = true
	this.LoadedFile[file_name].FilePointer = temp
	FunctionList,err := temp.Lookup("RjsFunctionList")
	if err != nil{
		fmt.Print("不规范的RJS库")
		return false
	}
	FuncList := strings.SplitAfter(FunctionList.(func()string)()," ")
	this.LoadedFile[file_name].func_map = make(map[string]*plugin.Symbol)
	var row string
	//加载函数
	for _,row = range FuncList{
		*this.LoadedFile[file_name].func_map[row],err = temp.Lookup(row)
		if err != nil{	//与函数表中函数不符
			fmt.Print("ERROR When load function:\" " + row +"\" from: \" " + file_name +"\"")
			OnStrictMode(this.js)
		}
	}
	//提供回调接口给对方库
	RjsGetVariableFunction,err := temp.Lookup("RjsSetGetVariableFunction")
	if err != nil{
		fmt.Print("不规范的RJS库")
		return false
	}
	//RjsGetVariableFunction.(func(func()))(GetVariableValue)
	GetVariableValueCallBack := syscall.NewCallback(this.GetVariableValue)
	RjsGetVariableFunction.(func(unsafe.Pointer))(unsafe.Pointer(GetVariableValueCallBack))
	RjsSetVariableFunction,err := temp.Lookup("RjsSetSetVariableFunction")
	if err != nil{
		fmt.Print("不规范的RJS库")
		return false
	}
	SetVariableFunctionCallBack := syscall.NewCallback(this.SetVariableValue)
	RjsSetVariableFunction.(func(unsafe.Pointer))(unsafe.Pointer(SetVariableFunctionCallBack))
	return true
}
/*		RJS动态库规范		*/
/*		调用约定为extern "C" __declspec(dllexport)		*/
/*		必须有一个名为RjsPluginInformation函数为获取动态链接库信息	*/
/*		必须有一个名为RjsFunctionList函数返回所有挂载到RJS的函数列表，使用字符串，用空格分开		*/
/*		必须有一个名为RjsGetVariableFunction的函数让RJS来设置回调函数的地址。这个函数提供给库用来获取变量的值	，该函数只有一个参数，为函数地址。RJS提供的函数原型为
GetVariableValue(VariableName string)string
必须有一个名为RjsSetVariableValue的函数来让RJS设置回调函数地址。这个函数提供给库来设置变量的值。
RjsSetVariableValue只有一个参数，为函数地址。
RJS函数原型：
SetVariableValue(VaribleName string,Value string)bool
返回字节类型
*/
func OnStrictMode(vm *otto.Otto){
	IfStrictMode,err := vm.Get("RJS_CONFIG_STRIT_MODE")
	if err != nil{
		fmt.Print("ERRO CONFIG.RJS_CONFIG_STRIT_MODE")
		os.Exit(0)
	}
	confident,err :=  IfStrictMode.ToBoolean()
	if err != nil{
		fmt.Print("ERRO CONFIG.RJS_CONFIG_STRIT_MODE")
		os.Exit(0)
	}
	if confident{
		os.Exit(0)
	}
	return
}
func (this JSLoader)GetVariableValue(VariableName string)string{
	value,err := this.js.Get(VariableName)
	if err != nil{
		fmt.Print("ERROR When Call GetVariableValue.ERROR VariableName")
		os.Exit(0)
	}
	return value.String()
}
func (this *JSLoader)SetVariableValue(VaribleName string,Value string)bool{
	err := this.js.Set(VaribleName,Value)
	if err != nil{
		fmt.Print(err)
		OnStrictMode(this.js)
		return false
	}
	return true
}